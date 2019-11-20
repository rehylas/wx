/**
 *  根据 证券编号获取二维码信息，并显示
 */
 
 
 
 var jq = $.noConflict(); 
 var qrcode = null;
 var cust_stockcode ="";
 
 jq(document).ready(function(){
 	
	         qrcode = new QRCode(document.getElementById("qrcode"), {
	            width : 196,//设置宽高
	            height : 196
	        });
	        qrcode.makeCode("http://www.bangnikanzhe.com"); 	
	
});



fnOnRefreshqrClick = function()
{
	 
	
	cust_stockcode  = jq("#stockuserID").val(); ;
	if( cust_stockcode == '' ){
		alert('请输入证券账号');
		return ;
	}		
	 
	reqdata = " { 'stockcode' : '"+ cust_stockcode +"', " 
			 +" 'openid' : '"+ "000000" +"' " 
			+"} ";
			
	console.log('reqdata:');
	console.log( reqdata );	
	jq.ajax({
	        type: "POST",
	        url: "/DXBizGate/server?cmd=getqrinfo",
	        contentType: "application/json; charset=utf-8",
	        data:  reqdata ,    //JSON.stringify(GetJsonData())
	        dataType: "json",
	        success: function (message) { // 成功处理
	        	
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
	        		//alert( '申请成功' );
	        		var datalist = message;
	        		var qrInfo = datalist.data[0].qrurl;
					console.log(  qrInfo );
	        		qrcode.makeCode( qrInfo );
	        		jq("#imgQrID").show(); 
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
	        error: function (message) { //失败处理
	        	alert( '提交数据失败:' + message.msg );
	            jq("#request-process-patent").html("提交数据失败！");
	        }	        
	    });	 
		
	//qrcode.makeCode( "debug" );
}


