package database

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	mgo "gopkg.in/mgo.v2"
)

// mongo db 操作参考  https://godoc.org/gopkg.in/mgo.v2

type ServDB struct {
	__session     *mgo.Session
	__redisClient *redis.Client
}

var (
	dbserv ServDB
)

var (
	TransactionDB *mongo.Database
	Ctx           context.Context
	Client        *mongo.Client
)

const DB_NAME = "bnkz"

/////////////////////////////////////////////////////////////////////////
//输出

func Initdb(mongoip string, redisip string) error {
	var err error
	if dbserv.__session == nil {

		dbserv.__session, err = mgo.Dial(mongoip)
		if err != nil {

			return err
		} else {
			dbserv.__session.SetMode(mgo.Strong, true)
			//__session.SetMode(mgo.Eventual   )
		}
	}

	// dbserv.__redisClient = redis.NewClient(&redis.Options{
	// 	Addr:     redisip,
	// 	Password: "", // no password set
	// 	DB:       0,  // use default DB
	// })

	// _, err := dbserv.__redisClient.Ping().Result()
	// if err != nil {
	// 	return err
	// }
	return err
}

func Closedb() {
	dbserv.__session.Close()
	dbserv.__redisClient.Close()
}

func GetRedisClient() *redis.Client {
	return dbserv.__redisClient
}

func GetCollect(collectname string) (*mgo.Collection, error) {
	dbname := DB_NAME
	sess := getSession()
	if sess == nil {
		return nil, errors.New("mongo db error")
	} else {
		return sess.DB(dbname).C(collectname), nil
	}

}

//////////////////////////////////////////////////////////////////////////
//内部函数

func getSession() *mgo.Session {
	return dbserv.__session
}

func GetCollectEx(dbname, collectname string) (*mgo.Collection, error) {

	sess := getSession()
	if sess == nil {
		return nil, errors.New("mongo db error")
	} else {
		return sess.DB(dbname).C(collectname), nil
	}

}

func InitTransactionDB(mongoip string) error {
	connectString := "mongodb://" + mongoip + "/" + DB_NAME
	dbUrl, err := url.Parse(connectString)
	if err != nil {
		panic(err)
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(connectString))
	if err != nil {
		logs.Error("获取mongo事务异常:", Client, err)
		panic(err)
	}

	transactionDB := client.Database(dbUrl.Path[1:])
	if transactionDB == nil {
		return errors.New("Client.Database异常,返回为空！")
	}
	Client = client
	TransactionDB = transactionDB
	//Ctx=ctx
	//defer Client.Disconnect(ctx)
	return nil
}

func CallTransaction(DoReleasePartner func(db *mongo.Database, sessionContext mongo.SessionContext, args ...interface{}) error, args2 ...interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	err := Client.Connect(ctx)
	if err != nil {
		return err
	}
	defer TransactionDB.Client().Disconnect(ctx)

	err = TransactionDB.Client().UseSession(ctx, func(sessionContext mongo.SessionContext) error {
		err := sessionContext.StartTransaction()
		if err != nil {
			fmt.Println(err)
			return err
		}
		//做具体业务
		err = DoReleasePartner(TransactionDB, sessionContext, args2)
		if err == nil {
			sessionContext.CommitTransaction(sessionContext)
			return nil
		}
		return err
	})
	return err
}
