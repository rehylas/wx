/**
 * 鏍规嵁 璇佸埜缂栧彿鑾峰彇浜岀淮鐮佷俊鎭紝骞舵樉绀�
 */

var jq = $.noConflict();
var qrcode = null;
var tempScenceId = null;
var cust_userid = null;
var cust_userpwd = null;
var nCount = 0;
var ObjInter = null;

jq(document).ready(function() {

	// fnSendMsgTest('1001','1001','消息内容')
	qrcode = new QRCode(document.getElementById("qrcode"), {
		width : 196,// 璁剧疆瀹介珮
		height : 196
	});
	//fnGetQruInfoRegAPI('4001');

	// ObjInter = setInterval("Interval_TestScan()",3000);//1000为1秒钟

	// qrcode.makeCode( '' );
	// jq("#imgQrID").show();

	// jq("#SM_Info_userid").text( '更新显示' );
	// console.log( '成功' );
	// jq("#SM_Info_ID").show();

});

function Interval_TestScan() {
	// fnGetScanInfoByScenceId(tempScenceId);
	// nCount++
	// if (nCount >= 20 * 10)
	// 	clearInterval(ObjInter);
}

// 不断刷新读取，判断用户是否已扫码，如果已扫码，展示账户相关信息
fnGetScanInfoByScenceId = function(ScenceId) {
	// tempScenceID = {'tempScenceID':'50010001'}
	reqdata = " { 'tempScenceID' : '" + ScenceId + "' " + "} ";

	console.log('reqdata:');
	console.log(reqdata);
	jq.ajax({
		type : "POST",
		url : "/DXBizGate/server?cmd=getscaninfobyscenceid",
		contentType : "application/json; charset=utf-8",
		data : reqdata, // JSON.stringify(GetJsonData())
		dataType : "json",
		success : function(message) {

			if (typeof (message) == "undefined") {
				return;
			}
			console.log(message);
			console.log(message.total);
			var msg = message.msg; // a.length();
			var code = msg.substring(0, 4);
			console.log(code);

			if (code == '0000') {
				// alert( '成功' );
				var datalist = message;

				var TempState = datalist.data[0].state;
				cust_userid = datalist.data[0].userid;
				cust_userpwd = datalist.data[0].userpwd;
				state = datalist.data[0].state;
				console.log('成功');
				if (state == '10') {
					jq("#SM_Info_userid").text('股讯账号：' + cust_userid);
					jq("#SM_Info_userpwd").text('股讯密码：' + cust_userpwd);
					jq("#test_userid_ID").val(cust_userid);
					jq("#test_userpwd_ID").val(+cust_userpwd);

					jq("#SM_Info_ID").show();
					clearInterval(ObjInter);
					console.log('申请成功');
				}

				return;
			} else {
				alert(msg);
			}

			if (message > 0) {
				console.log(message);
			}
		},
		error : function(message) { // 
			alert('获取失败:' + message.msg);
			jq("#request-process-patent").html(" ");
		}
	});
	return;

}

// 请求获取二维码信息，并展示
fnGetQruInfoRegAPI = function(QrType) { // QrType = 5001, 4001, 6001
										// {'QrType':'5001'}
	reqdata = " { 'qrtype' : '" + QrType + "' " + "} ";

	console.log('reqdata:');
	console.log(reqdata);
	jq.ajax({
		type : "POST",
		url : "/DXBizGate/server?cmd=getqrinfoex",
		contentType : "application/json; charset=utf-8",
		data : reqdata, // JSON.stringify(GetJsonData())
		dataType : "json",
		success : function(message) {

			if (typeof (message) == "undefined") {
				return;
			}
			console.log(message);
			console.log(message.total);
			var msg = message.msg; // a.length();
			var code = msg.substring(0, 4);
			console.log(code);

			if (code == '0000') {
				// alert( '成功' );
				var datalist = message;
				var qrInfo = datalist.data[0].qrurl;
				tempScenceId = datalist.data[0].tempid;
				console.log(qrInfo);
				qrcode.makeCode(qrInfo);
				jq("#imgQrID").show();
				// interval at this time
				ObjInter = setInterval("Interval_TestScan()", 3000);// 1000为1秒钟
				return;
			} else {
				alert(msg);
			}

			if (message > 0) {
				console.log(message);
			}
		},
		error : function(message) { // 
			alert('获取失败:' + message.msg);
			jq("#request-process-patent").html(" ");
		}
	});
	return;
}

// 测试一下发送信息
// {'userid':'10011001','userpwd':'123456','txt':'信息内容' }
fnSendMsgTest = function(userid, userpwd, msg_txt) {
	var userid = jq("#test_userid_ID").val();
	var userpwd = jq("#test_userpwd_ID").val();
	var msg_txt = jq("#test_msg_ID").val();

	// console.log( jq("#test_msg_ID").text() );
	// console.log( jq("#test_msg_ID").val() );

	reqdata = ' { "userid" : "' + userid + '", ' + ' "userpwd" : "' + userpwd
			+ '", ' + '   "txt" : "' + msg_txt + '"  ' + '} ';

	 			
	console.log('reqdata:');
	console.log(reqdata);
	jq.ajax({
		type : "POST",
		url : "/DXBizGate/server?cmd=sendmsg4client",
		contentType : "application/json; charset=utf-8",
		data : reqdata, // JSON.stringify(GetJsonData())
		dataType : "json",
		success : function(message) {

			if (typeof (message) == "undefined") {
				return;
			}
			console.log(message);
			console.log(message.total);
			var msg = message.msg; // a.length();
			var code = msg.substring(0, 4);
			console.log(code);

			if (code == '0000') {
				alert('发送成功');
				// var datalist = message;
				// var qrInfo = datalist.data[0].qrurl;
				// tempScenceId = datalist.data[0].tempid;
				// console.log( qrInfo );
				// qrcode.makeCode( qrInfo );
				// jq("#imgQrID").show();
				return;
			} else {
				console.log(message);
				alert(msg);
			}

			if (message > 0) {
				console.log(message);
			}
		},
		error : function(message) { // 
			alert('获取失败??:' + message.msg);
			jq("#request-process-patent").html(" ");
		}
	});

}

fnOnRefreshqrClick = function() {

	cust_stockcode = 'www.bangnikanzhe.com'; // jq("#stockuserID").val();
	if (cust_stockcode == '') {
		alert('输入二维码有问题');
		return;
	}

	qrcode.makeCode(cust_stockcode);
	jq("#imgQrID").show();

	// qrcode.makeCode( "debug" );
}

/*
 * reqdata = " { 'stockcode' : '"+ cust_stockcode +"', " +" 'openid' : '"+
 * "000000" +"' " +"} ";
 * 
 * console.log('reqdata:'); console.log( reqdata ); jq.ajax({ type: "POST", url:
 * "/DXBizGate/server?cmd=getqrinfo", contentType: "application/json;
 * charset=utf-8", data: reqdata , //JSON.stringify(GetJsonData()) dataType:
 * "json", success: function (message) { // 鎴愬姛澶勭悊
 * 
 * if (typeof( message ) == "undefined") { return ; }
 * 
 * console.log( message ); console.log( message.total ); var msg = message.msg;
 * //a.length(); var code = msg.substring(0,4); console.log( code );
 * 
 * if( code == '0000' ) { //alert( '鐢宠鎴愬姛' ); var datalist = message; var
 * qrInfo = datalist.data[0].qrurl; console.log( qrInfo ); qrcode.makeCode(
 * qrInfo ); jq("#imgQrID").show(); return ; } else { alert( msg ); }
 * 
 * if (message > 0) { console.log( message ); } }, error: function (message) {
 * //澶辫触澶勭悊 alert( '鎻愪氦鏁版嵁澶辫触:' + message.msg );
 * jq("#request-process-patent").html("鎻愪氦鏁版嵁澶辫触锛�"); } });
 */

