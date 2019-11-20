/**
 *  申请页面的交互
 */
 
 
 
var jq = $.noConflict(); 
var yybCount =4; 
var cust_yybid='';
var cust_zqcompanyCode='';
var cust_stockuser='';

var cust_jypwd ='888888';	
var cust_txpwd ='888888'; 	
var cust_phone ='139';
var cust_openid ='';



var companyCount = 0;
var companyList ;

var yybList; 


jq(document).ready(function(){
	setTimeout(function(){console.log("  fnGetRegopenid ");fnGetRegopenid();  }	
		,2000);		
	initUI();
	
});


initUI =  function()
{
	fnGetCompanyList();
	jq("#stockuserID").focus( fnOnfocusStockUserID );
	jq("#companyNameID").focus( fnOnfocusCompanyNameID );
	jq("#companyyybID").focus( fnOnfocusCompanyyybID );
	
	
	jq("#stockuserID").focusout( fnOnfocusoutStockUserID );
	jq("#companyNameID").focusout( fnOnfocusoutCompanyNameID );
	
	jq("#companylistid").focusout( fnOnfocusCompanyListID );
	jq("#companyyyblistid").focusout( fnOnfocusCompanyYybListID );
	


	//clearCompanyList();
	//jq("#companylistid").hide();
	//#("#companylistid").hide();
}

// 焦点进去证券账号输入框
fnOnfocusStockUserID = function()
{
		console.log("fnOnfocusStockUserID");	 
		if( jq("#stockuserID").val()=='请输入证券公司账号' ) 
			jq("#stockuserID").val('');
	 	
		console.log("fnOnfocusStockUserID over ");
}

// 焦点离开证券账号输入框
fnOnfocusoutStockUserID = function()
{

		console.log("fnOnfocusoutCompanyNameID");	 
		if( jq("#stockuserID").val()=='' ) 
			jq("#stockuserID").val('请输入证券公司账号');
		console.log("fnOnfocusoutCompanyNameID over ");
}


// 焦点进入选择证券公司
fnOnfocusCompanyNameID = function()
{

		console.log("fnOnfocusCompanyNameID");	 
		if( jq("#companyNameID").val()=='请选择证券公司' ) 
			jq("#companyNameID").val('');
	    fnRefreshCompanyList();		
		jq("#companylistid").show();  // 证券公司选择框显示出来		
		console.log("selectCompany over ");
}

// 焦点离开选择证券公司
fnOnfocusoutCompanyNameID = function()
{

		console.log("fnOnfocusoutCompanyNameID");	 
		if( jq("#companyNameID").val()=='' ) 
			jq("#companyNameID").val('请选择证券公司');
		//jq("#companylistid").hide();  // 证券公司选择框隐藏 	
	    	
		location.href = "#stockUserInfoID"; 	
		console.log("fnOnfocusoutCompanyNameID over ");
}

//焦点离开证券公司选择清单框
fnOnfocusCompanyListID = function()
{
	jq("#companylistid").hide();  // 证券公司选择框隐藏 
	location.href = "#stockUserInfoID"; 
}


//焦点离开证券公司营业部选择清单框
fnOnfocusCompanyYybListID = function()
{
	jq("#companyyyblistid").hide();  // 证券公司选择框隐藏 
	location.href = "#stockUserInfoID"; 
}




// 焦点进入选择证券公司营业部
fnOnfocusCompanyyybID = function()
{

		console.log("fnOnfocusCompanyyybID");	 
		if( jq("#companyNameID").val()=='' || jq("#companyNameID").val()=='请选择证券公司' ) 
		{ 
			alert('证券公司未选择，请先选择证券公司');
			jq("#companyNameID").focus();
			return ;
		}
		fnRefreshYybInfoList(); // 证券公司营业部显示出来

		jq("#companyyyblistid").show(); 
		console.log("fnOnfocusCompanyyybID over ");
}

// 刷新证券公司清单界面
fnRefreshCompanyList = function()
{
	jq("#companylistid").empty();
	fnAddCompanyList();
}

// 增加证券公司
fnAddCompanyList = function()
{

	var id="";
	var name="";
	var valStr="";
	
	/*
	for(var i = 0;i <5; i++) {
		id = "id_fzzq"+i;
		name = "方正证券"+i;
		valStr = "fzzq"+i;
		txt_input =' <label id="lab_' +id+ '" ><input id="'+id+'" onclick="javascript:fnCheckedCompany(this.id)" name="Company" type="radio" value="'+valStr+'"  />'+name+'</label>' ;
		jq("#companylistid").append( txt_input );		
		console.log( txt_input );
	}
	*/
	
  	if (typeof( companyList ) == "undefined")
  	{	
  		alert('异常: 证券公司列表为空');
  		jq("#stockuserID").focus();
  		return ;
  	}
  	
   
  	console.log(  companyList.total  );
  	console.log(  companyList.data[1].name  );
  	
	for(var i = 0;i <companyList.total ; i++) {
		id = "id_" + companyList.data[i].code ;
		name = companyList.data[i].name ;
		valStr =  companyList.data[i].code ;
		txt_input =' <label id="lab_' +id+ '" style="font-size:48px;" ><input id="'+id+'" class="stdRadio" onclick="javascript:fnCheckedCompany(this.id)" name="Company" type="radio" value="'+valStr+'"  />'+name+'</label>' ;
		jq("#companylistid").append( txt_input );	
		if( i % 5 == 4 )
			jq("#companylistid").append( '<br/>' );	
		
		console.log( txt_input );
	}

}

// 选择证券公司
fnCheckedCompany  = function( id )
{
	console.log("fnCheckedCompany");
	console.log(  id );
 
	// code 保存在  input val 中
	data = jq("#"+id ).val();
	console.log(  data);
	cust_zqcompanyCode = data;

	//名称  保存在 对应lab的 text中
	data = jq("#lab_"+id ).text();
	console.log(  data);
 
	jq("#companyNameID").val( data );
	
	//
	jq("#companyyybID").val( '请选择证券营业厅' );
	jq("#companyyybID").attr("disabled", false);
	
	fnGetCompanyyybList();
 	//隐藏证券公司清单
	jq("#companylistid").hide();
	
		
	console.log("fnCheckedCompany over ");
	
}

// 选择证券营业厅
fnCheckedCompanyYyb  = function( id )
{
	console.log("fnCheckedCompanyYyb");
	console.log(  id );
 
	// code 保存在  input val 中
	data = jq("#"+id ).val();
	console.log(  data);
	cust_yybid = data;

	//名称  保存在 对应lab的 text中
	data = jq("#lab_"+id ).text();
	console.log(  data);
 
	jq("#companyyybID").val( data );
	jq("#companyyyblistid").hide();
 
	console.log("fnCheckedCompanyYyb over ");
	
}

// 刷新显示营业部信息
fnRefreshYybInfoList  = function( CompanyCode )
{
	console.log( 'fnRefreshYybInfoList' );
	jq("#companyyyblistid").empty(); 
	fnShowCompanyyybList();
	
	console.log( 'fnRefreshYybInfoList over' );
}


fnShowCompanyyybList = function()
{
	//<label><input name="Fruit" type="radio" value="" />方正证券5  </label>
	var id="";
	var name="";
	var valStr="";
	
	console.log( 'fnShowCompanyyybList' );

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
	
	if( yybList.total == 1 )
	{
		jq("#companyyybID").val( '总部' );
		jq("#companyyybID").attr("disabled", true);
		cust_yybid = yybList.data[0].yybid;
		console.log( 'fnShowCompanyyybList over yybCount == 1' );
		return ;
	}

	for(var i = 0;i < yybList.total; i++) {
		id = "id_yyb"+yybList.data[i].yybid;
		name = ""+ yybList.data[i].yybname ;
		valStr = ""+yybList.data[i].yybid;
		txt_input =' <label id="lab_' +id+ '" style="font-size:48px;"  ><input id="'+id+'" class="stdRadio" onclick="javascript:fnCheckedCompanyYyb(this.id)" name="Companyyyb" type="radio" value="'+valStr+'"  />'+name+'</label>' ;
		jq("#companyyyblistid").append( txt_input );	
		//if( i%5 == 4 )
			jq("#companyyyblistid").append( '<br/>' );
		console.log( txt_input );
	}	

	console.log( 'fnShowCompanyyybList over  ' );
}


// // http://localhost:8080/DXBizGate/server?cmd=getstockcompany
fnGetCompanyList = function()
{
	//jq("#div_txt_company").load("/DXBizGate/server?cmd=getstockcompany");
 	
    jq.get("/DXBizGate/server?cmd=getstockcompany", function(data,status){
  	console.log( "数据: " + data + "\n状态: " + status  );

  	companyList = JSON.parse(data);
  	
  	if (typeof( companyList ) == "undefined")
  	{
  		return ;
  	}  	
  	console.log(  companyList.total  );
  	console.log(  companyList.data[1].name  );
  	
   
  });
  
}

// // http://localhost:8080/DXBizGate/server?cmd=getcompanyyyb
fnGetCompanyyybList = function()
{
	//jq("#div_txt_company").load("/DXBizGate/server?cmd=getcompanyyyb");

	if( cust_zqcompanyCode == '' ){
		alert('请选择证券公司');
		return ;
	}
	
	reqdata = " { 'companycode' : '"+ cust_zqcompanyCode +"' } ";
	
	jq.ajax({
	        type: "POST",
	        url: "/DXBizGate/server?cmd=getcompanyyyb",
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
	        	if( message.total >0 )
	        		yybList = message;

	            if (message > 0) {  

			  		console.log(  message );
	            }
	        },
	        error: function (message) { //失败处理
	            jq("#request-process-patent").html("提交数据失败！");
	        }
	    });
 
}


 
// {'stockcode':'28001134','companycode':'fzzq','yybid':'1' }
fnRegZNDX = function()
{
	//("/DXBizGate/server?cmd=regzndx4user");
	if( cust_zqcompanyCode == '' ){
		alert('请选择证券公司');
		return ;
	}
	
	cust_stockuser = jq("#stockuserID").val(); ;
	if( cust_stockuser == '' || cust_stockuser == '请输入证券公司账号' ){
		alert('请输入证券账号');
		return ;
	}	
	
	if( isNumeric( cust_stockuser ) == true )
		console.log('数字');
	else
	{
		console.log('非数字');
		alert('证券账号必须为数字');
		return ;		
	}
	
 
	if( cust_yybid == '' ){
		alert('请选择证券营业部');
		return ;
	}		
	 
	reqdata = " { 'stockcode' : '"+ cust_stockuser +"'," 
	            +" 'companycode' : '"+ cust_zqcompanyCode +"'," 
	            +" 'yybid' : '"+ cust_yybid +"' " 	
			+"} ";
			
	console.log('reqdata:');
	console.log( reqdata );
	
	jq.ajax({
	        type: "POST",
	        url: "/DXBizGate/server?cmd=regzndxuser",
	        contentType: "application/json; charset=utf-8",
	        data:  reqdata ,    //JSON.stringify(GetJsonData())
	        dataType: "json",
	        success: function (message) { // 成功处理
	        	console.log( 'jq.ajax post ... ' );
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
	        		alert( '申请成功' );
	        		//self.location='www.bangnikanzhe.com'; 
	        		window.location.href='http://www.baidu.com/';
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
	        	console.log( 'jq.ajax post error ' );
	        	alert( '后台异常:' + message.msg );
	            jq("#request-process-patent").html("提交数据失败！");
	        }
	    });
	    
}

// {'stockcode':'28001134','companycode':'fzzq','yybid':'1','jypwd':'888888','txpwd':'888888' }
fnRegYundx = function()
{
	//("/DXBizGate/server?cmd=regzndx4user");
	if( cust_zqcompanyCode == '' ){
		alert('请选择证券公司');
		return ;
	}
	
	cust_stockuser = jq("#stockuserID").val(); ;
	if( cust_stockuser == '' || cust_stockuser == '请输入证券公司账号' ){
		alert('请输入证券账号');
		return ;
	}	
		
	cust_jypwd = jq("#stockuserJyPwd").val(); ;
	if( cust_jypwd == '' || cust_jypwd == '请输入证券交易密码' ){
		alert('请输入证券交易密码');
		return ;
	}
	
	cust_phone = jq("#phoneID").val(); ;
	 
	
	if( isNumeric( cust_stockuser ) == true )
		console.log('数字');
	else
	{
		console.log('非数字');
		alert('证券账号必须为数字');
		return ;		
	}
	
 
	if( cust_yybid == '' ){
		alert('请选择证券营业部');
		return ;
	}			
	reqdata = " { 'stockcode' : '"+ cust_stockuser +"'," 
	            +" 'companycode' : '"+ cust_zqcompanyCode +"'," 
	            +" 'yybid' : '"+ cust_yybid +"', " 
	            +" 'jypwd' : '"+ cust_jypwd +"', " 	
	            +" 'phone' : '"+ cust_phone +"', "  
	            +" 'openid' : '"+ cust_openid +"', "  	            
	            +" 'txpwd' : '"+ cust_txpwd +"' " 		           
			+"} ";
			
	console.log('reqdata:');
	console.log( reqdata );
	
	jq.ajax({
	        type: "POST",
	        url: "/DXBizGate/server?cmd=regyundxuser",
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
	        		alert( '申请成功' );

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
	    
}

fnGetRegopenid = function()
{
	
	console.log('fnGetRegopenid:');
	var reqdata =" {} ";
	jq.ajax({
	        type: "POST",
	        url: "/DXBizGate/server?cmd=getnewopenidaccessreg_html",
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
	        		console.log( '获取成功' );
	        		var datalist = message;
	        		cust_openid = datalist.data[0].openid;
	        		
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
}

//testClick()
fnTestClick = function()
{
	//fnTestGetCompanyList();
	//fnTestGetCompanyyybList(); 	
	fnRegZNDX();
	console.log( 'fnTestClick over  ' );
}

isNumeric = function( numberStr )
{	 
	var n = Number( numberStr );
	if (!isNaN(n))
	{
		return true
	}
	else
	{
		return false;
	}
}

