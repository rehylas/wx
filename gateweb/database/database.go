package database

import (
	"errors"

	"github.com/go-redis/redis"
	"gopkg.in/mgo.v2"
)

// mongo db 操作参考  https://godoc.org/gopkg.in/mgo.v2

type ServDB struct {
	__session     *mgo.Session
	__redisClient *redis.Client
}

var (
	dbserv ServDB
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
			dbserv.__session.SetMode(mgo.Eventual, true)
			//__session.SetMode(mgo.Monotonic   )
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
