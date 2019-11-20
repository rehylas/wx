import React, { Component } from 'react';
import { Table } from 'antd';
import $axios from '../../axios/$axios';
import { openNotification_err,openNotification_succ } from '../notification/index'


const Host = process.env.API_HOST || "http://localhost:8888"
  

// {
// 	"insertdt": "2019-10-30 20:56:01",
// 	"symbol": "ru2001",
// 	"type": "sell",
// 	"offset": "close",
// 	"price": 11965,
// 	"vol": 1,
// 	"date": "2019-10-30",
// 	"state": 1,
// 	"execdt": "2019-10-30 20:56:00",
// 	"orderidsys": "",
// 	"exchange": "",
// 	"FrontID": "21",
// 	"SessionID": "-2040445757",
// 	"OrderID": "1"
//   },

class BizTableBasic extends Component {
	state = {
		data: [],
		pagination: {},
		loading: false,
		selectedRowKeys: [],
		columns: [
			{
				title: '时间',
				dataIndex: 'insertdt'
			},
			{
				title: '执行',
				dataIndex: 'execdt'
			},			
			{
				title: '合约',
				dataIndex: 'symbol'
			},
			{
				title: '买卖',
				dataIndex: 'type',
				key: 'type',
				render: val => <a>{val}</a> 

			},
			{
				title: '开平',
				dataIndex: 'offset',
				key: 'offset', 
				render: val => { 
					let colorsMap = {'close': 'red',  'open': 'blue'}
					return <div style={{color: colorsMap[val]}}> {val} </div>  
				}
				 
			},


			{
				title: '价格',
				dataIndex: 'price'
			},	
			{
				title: '数量',
				dataIndex: 'vol'
			},		
			{
				title: '状态',
				dataIndex: 'state',
				render(state){
					return state===0?'未执行':'已执行'   
				}
			},	
			{
				title: 'CTP编号',
				dataIndex: 'OrderID'
			}																		
		 
		]
	};

	componentWillMount() {
		this.fetch();
	}

	componentWillUnmount() {
		// componentWillMount进行异步操作时且在callback中进行了setState操作时，需要在组件卸载时清除state
		this.setState = () => {
			return;
		};
	}

	handleTableChange = (pagination, filters, sorter) => {
		const pager = { ...this.state.pagination };
		pager.current = pagination.current;
		this.setState({
			pagination: pager
		});
		this.fetch({
			results: pagination.pageSize,
			page: pagination.current,
			sortField: sorter.field,
			sortOrder: sorter.order,
			...filters
		});
	};
	fetch = () => {
		this.setState({ loading: true });

		// $axios({
		// 	url: 'https://randomuser.me/api?results=10',
		// 	method: 'get',
		// 	type: 'json'
		// }).then(data => {
		// 	const pagination = { ...this.state.pagination };
		// 	// Read total count from server
		// 	// pagination.total = data.totalCount
		// 	pagination.total = 200;
		// 	this.setState({
		// 		loading: false,
		// 		data: data.data.results,
		// 		pagination
		// 	});
		// });

		$axios({
			url: Host +'/v1/order/listtoday',
			method: 'get',
			type: 'json'
		}).then(result => {
			let data =[]
			console.log("resp:", result )
			if (result.data.code === "0000"){
				data =  result.data.data
				console.log(   (data.length) );
				for (let i=0;i<data.length;i++ ){
					data[i].key = i
				}
							 
			}else{
				console.log( result.data.code );
				openNotification_err("读取失败，请重新刷新", result.data.msg )
			}

			const pagination = { ...this.state.pagination };
			// Read total count from server
			// pagination.total = data.totalCount
			//pagination.total = 200;
			this.setState({
				loading: false,
				data:data ,
				pagination
			});			
	 
		});			




	};
	render() {
		return (
			<div className="shadow-radius">
				<Table
					bordered
					columns={this.state.columns}
					dataSource={this.state.data}
					loading={this.state.loading}
					onChange={this.handleTableChange}
					pagination={this.state.pagination}
					
				/>
			</div>
		);
	}
}
//rowKey={record => record.location.postcode}
export default BizTableBasic;
