/**
 *  閺嶈宓� 鐠囦礁鍩滅紓鏍у娇閼惧嘲褰囨禍宀�娣惍浣蜂繆閹垽绱濋獮鑸垫▔缁�锟�
 */
 
 
 
 var jq = $.noConflict(); 
 var qrcode = null;
 var tempScenceId = null;
 var cust_userid =null;
 var cust_userpwd =null;
 var cust_code=null;
 var cust_date=null;
 var nCount =0;
 var ObjInter = null;
 var fileList = null; 
 
 
 jq(document).ready(function(){
 	
 			//fnSendMsgTest('1001','1001','娑堟伅鍐呭')
	         qrcode = new QRCode(document.getElementById("qrcode"), {
	            width : 196,//鐠佸墽鐤嗙�逛粙鐝�
	            height : 196
	        });
	        fnGetQruInfoRegAPI('5001');
	        
	        	                	     
	        // qrcode.makeCode( '' ); 	
	        //jq("#imgQrID").show(); 
	        
	        		
//	        		jq("#SM_Info_userid").text( '鏇存柊鏄剧ず' );
//					console.log(  '鎴愬姛' );					
//	        		jq("#SM_Info_ID").show(); 	        
	
});

function Interval_TestScan()
{
	 fnGetScanInfoByScenceId( tempScenceId );
	 nCount ++
	 if( nCount >= 20*10 )
	 	clearInterval( ObjInter );
}	


// 涓嶆柇鍒锋柊璇诲彇锛屽垽鏂敤鎴锋槸鍚﹀凡鎵爜锛屽鏋滃凡鎵爜锛屽睍绀鸿处鎴风浉鍏充俊鎭�
fnGetScanInfoByScenceId = function( ScenceId ) 
{
  //tempScenceID =  {'tempScenceID':'50010001'} 
	reqdata = " { 'tempScenceID' : '"+ ScenceId +"' " 
			+"} ";   
			
	console.log('reqdata:');
	console.log( reqdata );	
	jq.ajax({
	        type: "POST",
	        url: "/DXBizGate/server?cmd=getscaninfobyscenceid",
	        contentType: "application/json; charset=utf-8",
	        data:  reqdata ,    //JSON.stringify(GetJsonData())
	        dataType: "json",
	        success: function (message) { 
	        	
				if (typeof( message ) == "undefined")
				{
				  		return ;
				}  	
	        	console.log( message );
	        	console.log( message.total );
	        	var msg = message.msg; //a.length();
	        	var code = msg.substring(0,4);
	        	console.log( code );
	        	
	        	if( code == '0000' )
	        	{
	        	 
	        		var datalist = message;
	        		
	        		var TempState = datalist.data[0].state;
	        		cust_userid = datalist.data[0].userid;
	        		cust_userpwd = datalist.data[0].userpwd;
	        		state = datalist.data[0].state;
	        		console.log(  '0000' );	
	        		if( state == '10' )
	        		{
		        		jq("#SM_Info_userid").text( '账号：' +cust_userid );
		        		jq("#SM_Info_userpwd").text( '密码：' +cust_userpwd );
		        		//jq("#test_userid_ID").val(  cust_userid );
		        		//jq("#test_userpwd_ID").val(   +cust_userpwd );	
		        		
		        		//按钮可用 attr("disabled", true );
	 	                jq("#bigtick_getbtn_ID").removeAttr("disabled"); 
						jq("#text_info_ID").text("关注成功，请输入证券代码和年月提取数据");
	 	                	
										
		        		//不显示没必要  jq("#SM_Info_ID").show(); 
		        		clearInterval( ObjInter );
		        		console.log(  'scan over' );	
	        		}

	        		return ;
	        	}
	        	else
	        	{
	        		alert( msg );
	        	}

	            if (message > 0) {  
			  		console.log(  message );
	            }
	        },
	        error: function (message) { // 
	        	alert( 'GetScanInfoByScenceId ajax  error:' + message.msg );
	            jq("#request-process-patent").html(" ");
	        }	        
	    });	 
	return  ;
 	
}

// 璇锋眰鑾峰彇浜岀淮鐮佷俊鎭紝骞跺睍绀�
fnGetQruInfoRegAPI = function( QrType ) 
{ //QrType = 5001, 4001, 6001  {'QrType':'5001'} 
	reqdata = " { 'qrtype' : '"+ QrType +"' " 
			+"} ";   
			
	console.log('reqdata:');
	console.log( reqdata );	
	jq.ajax({
	        type: "POST",
	        url: "/DXBizGate/server?cmd=getqrinfoex",
	        contentType: "application/json; charset=utf-8",
	        data:  reqdata ,    //JSON.stringify(GetJsonData())
	        dataType: "json",
	        success: function (message) { 
	        	
				if (typeof( message ) == "undefined")
				{
				  		return ;
				}  	
	        	console.log( message );
	        	console.log( message.total );
	        	var msg = message.msg; //a.length();
	        	var code = msg.substring(0,4);
	        	console.log( code );
	        	
	        	if( code == '0000' )
	        	{
	        		//alert( '鎴愬姛' );
	        		var datalist = message;
	        		var qrInfo = datalist.data[0].qrurl;
	        		tempScenceId = datalist.data[0].tempid;
					console.log(  qrInfo );
	        		qrcode.makeCode( qrInfo );
	        		jq("#imgQrID").show();
	        		// 获取二维码成功后再循环
	        		ObjInter = setInterval("Interval_TestScan()",3000); 

	        		return ;
	        	}
	        	else
	        	{
	        		alert( msg );
	        	}

	            if (message > 0) {  
			  		console.log(  message );
	            }
	        },
	        error: function (message) { // 
	        	alert( 'GetQruInfoRegAPI :' + message.msg );
	            jq("#request-process-patent").html(" ");
	        }	        
	    });	 
	return  ;
}

// 测试发送数据
//  {'userid':'10011001','userpwd':'123456','txt':'淇℃伅鍐呭' }
fnSendMsgTest =  function(userid,userpwd, msg_txt)
{
	var userid =jq("#test_userid_ID").val();
	var userpwd =jq("#test_userpwd_ID").val();  
	var msg_txt =jq("#test_msg_ID").val(); 
	
	
	//console.log( jq("#test_msg_ID").text() );
	//console.log( jq("#test_msg_ID").val() );
	
	
		
	reqdata = " { 'userid' : '"+ userid +"', " 
			 +"   'userpwd' : '"+ userpwd +"', " 
			 +"   'txt' : '"+ msg_txt +"', " 
			+"} ";   
			
	console.log('reqdata:');
	console.log( reqdata );	
	jq.ajax({
	        type: "POST",
	        url: "/DXBizGate/server?cmd=sendmsg4client",
	        contentType: "application/json; charset=utf-8",
	        data:  reqdata ,    //JSON.stringify(GetJsonData())
	        dataType: "json",
	        success: function (message) { 
	        	
				if (typeof( message ) == "undefined")
				{
				  		return ;
				}  	
	        	console.log( message );
	        	console.log( message.total );
	        	var msg = message.msg; //a.length();
	        	var code = msg.substring(0,4);
	        	console.log( code );
	        	
	        	if( code == '0000' )
	        	{
	        		alert( 'send success' );
//	        		var datalist = message;
//	        		var qrInfo = datalist.data[0].qrurl;
//	        		tempScenceId = datalist.data[0].tempid;
//					console.log(  qrInfo );
//	        		qrcode.makeCode( qrInfo );
//	        		jq("#imgQrID").show(); 
	        		return ;
	        	}
	        	else
	        	{
	        		alert( msg );
	        	}

	            if (message > 0) {  
			  		console.log(  message );
	            }
	        },
	        error: function (message) { // 
	        	alert( 'send msg error ajax msg:' + message.msg );
	            jq("#request-process-patent").html(" ");
	        }	        
	    });	 
 
}

// 获取文件清单
//  {'userid':'10011001','userpwd':'123456', 'code':'600177', 'yyyymm':'201706'  }
fnBigTickFileList =  function(userid,userpwd, msg_txt)
{
      		
	var userid = cust_userid ;
	var userpwd = cust_userpwd ;  
 
	var code =jq("#bigtick_stock_ID").val();
	var yyyymm =jq("#bigtick_date_ID").val();  
	
	cust_code= code ;
  	cust_date= yyyymm ;	 
	
	//console.log( jq("#test_msg_ID").text() );
	//console.log( jq("#test_msg_ID").val() );
	
	
		
	reqdata = " { 'userid' : '"+ userid +"', " 
			 +"   'userpwd' : '"+ userpwd +"', " 
			 +"   'code' : '"+ code +"', " 
			 +"   'yyyymm' : '"+ yyyymm +"' " 
			+"} ";   
			
	console.log('reqdata:');
	console.log( reqdata );	
	jq.ajax({
	        type: "POST",
	        url: "/DXBizGate/server?cmd=getbigdatabigticklist",
	        contentType: "application/json; charset=utf-8",
	        data:  reqdata ,    //JSON.stringify(GetJsonData())
	        dataType: "json",
	        success: function (message) { 
	        	
				if (typeof( message ) == "undefined")
				{
				  		return ;
				}  	
	        	console.log( message );
	        	console.log( message.total );
	        	var msg = message.msg; //a.length();
	        	var resp_code = msg.substring(0,4);
	        	console.log( resp_code );
	        	
	        	//if( message.total >0 )
	        	fileList = message;	        	
	        	
	        	if( resp_code == '0000' )
	        	{
	        		console.log( 'getbigdatabigticklist success' );
	        		fnShowFileList();
	        		jq("#fileList_ID").show();
	       
	        		return ;
	        	}
	        	else
	        	{
	        		alert( msg );
	        	}
 
	        },
	        error: function (message) { // 
	        	alert( 'send msg error ajax msg:' + message.msg );
	            jq("#request-process-patent").html(" ");
	        }	        
	    });	 
 
}


fnOnRefreshqrClick = function()
{
	 	
	cust_stockcode  =  'www.bangnikanzhe.com';   //jq("#stockuserID").val();
	if( cust_stockcode == '' ){
		alert('杈撳叆浜岀淮鐮佹湁闂');
		return ;
	}		
	
	qrcode.makeCode( cust_stockcode );
	jq("#imgQrID").show(); 	
	
		
	//qrcode.makeCode( "debug" );
}


fnShowFileList = function()
{
	//<label><input name="Fruit" type="radio" value="" />方正证券5  </label>
	var id="";
	var name="";
	var valStr="";
	
	console.log( 'fnShowFileList' );

	/*
	if( yybCount == 1 )
	{
		jq("#companyyybID").val( '总部' );
		jq("#companyyybID").attr("disabled", true);
		console.log( 'fnShowCompanyyybList over yybCount == 1' );
		return ;
	}

	for(var i = 0;i < 5; i++) {
		id = "id_yyb"+i;
		name = "方正证券部"+i;
		valStr = ""+i;
		txt_input =' <label id="lab_' +id+ '" style="font-size:16px;"  ><input id="'+id+'" class="stdRadio" onclick="javascript:fnCheckedCompanyYyb(this.id)" name="Companyyyb" type="radio" value="'+valStr+'"  />'+name+'</label>' ;
		jq("#companyyyblistid").append( txt_input );		
		console.log( txt_input );
	}
	*/
	jq("#data_desp_ID").text( '股票代码：'+ cust_code + '  日期：'+cust_date  + '  文件个数：'+ fileList.total );
	if( fileList.total == 0 )
	{
		
		return ;
					
	}

	for(var i = 0;i < fileList.total; i++) {
		
		file_date = ""+fileList.data[i].datadt;
		file_type = ""+ fileList.data[i].filetype ;
		//file_url = ""+fileList.data[i].fileurl;		
		
		file_url = '  <a href="+fileList.data[i].fileurl+" > 点击下载 </a>  '
		
 
		
		
		txt_input =' <tr> <td>'+ file_date +'</td><td> '+ file_type +' </td><td>'+file_url+'</td>    </tr> ' ;
		jq("#table_filelist_ID").append( txt_input );			
		console.log( txt_input );
	}	

	console.log( 'fnShowFileList over  ' );
}


