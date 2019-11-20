import React, { Component } from 'react';
import { Table, Divider, Modal, message } from 'antd';
import ModeEditForm from '../../components/ModeEditForm';
import $axios from '../../axios/$axios';
import { openNotification_err,openNotification_succ } from '../notification/index'
 
const Host = process.env.API_HOST || "http://localhost:8888"
  
// { 
//     "_id" : ObjectId("5d4ec66dc15f077db7f15c44"), 
//     "userid" : "80010121", 
//     "userpwd" : "888889", 
//     "userlevel" : 1.0, 
//     "openid" : "ot_s20be187ldnuWyJhUwUxo6DjA", 
//     "state" : NumberInt(1), 
//     "insertdt" : "2017-06-30 17:55:21", 
//     "mobile" : "139", 
//     "tempid" : "40010121", 
//     "daymsgmax" : 20.0, 
//     "vaildt" : "2019-01-01 00:00:00"
// }

const { confirm } = Modal;
class BizTableEdit extends Component {
	state = {
		data: [],
		pagination: {
			pageSize: 20,
			current: 1
		},
		loading: false,
		selectedRowKeys: [],
		columns: [
			{
				title: '用户编号',
				dataIndex: 'userid'
			},
			{
				title: '用户密码',
				dataIndex: 'userpwd'
			},
			{
				title: '付费用户',
				dataIndex: 'userlevel',
				key: 'userlevel',
				// render(userlevel){
				// 	return userlevel===2?'已付费':'未付费'   
				// }			
				render: userlevel => { 
					let showval = userlevel===2?'已付费':'未付费'  
					let colorsMap = {1: 'red',  2: 'blue'}
					return <div style={{color: colorsMap[userlevel]}}> {showval} </div>  
				}					
			},
			{
				title: '付费时间',
				dataIndex: 'paydt'
			},			
			{
				title: '用户状态',
				dataIndex: 'state',
				key: 'state',
				// render(enable){
				// 	return enable===1?'开启':'关闭'   
				// }
				
				render: state => { 
					let showval = state===1?'开启':'禁用'  
					let colorsMap = {0: 'red',  1: 'blue'}
					return <div style={{color: colorsMap[state]}}> {showval} </div>  
				}
			},
			{
				title: '加入时间',
				dataIndex: 'insertdt'
			},	
			{
				title: '微信ID',
				dataIndex: 'openid'
			},	
			{
				title: '备注信息',
				dataIndex: 'meo'
			},																								
			{
				title: '操作',
				dataIndex: 'Action',
				width: 200,
				align: 'center',
				render: (text, row, index) => (
					<span>
						<button className="link-button" onClick={() => this.handleEdit(row)}>
							编辑/添加
						</button>
						<Divider type="vertical" />
 
					</span>
				)
			}
		],
		currentRow: null
	};

	componentWillMount() {
		const { pageSize, current } = this.state.pagination;
		this.fetch(current, pageSize);
	}

	componentWillUnmount() {
		// componentWillMount进行异步操作时且在callback中进行了setState操作时，需要在组件卸载时清除state
		this.setState = () => {
			return;
		};
	}
	// 分页操作
	handleTableChange = pagination => {
		const pager = { ...this.state.pagination };
		pager.current = pagination.current;
		this.setState({ pagination: pager }, () => {
			this.fetch(pager.current, this.state.pagination.pageSize);
		});
	};
	fetch = (pageCurrent, pageSize) => {
		this.setState({ loading: true });
		setTimeout(() => {



			const pager = { ...this.state.pagination };
			pager.total = 200;
			const data = [];
			
			// for (let i = (pageCurrent - 1) * pageSize; i < pageSize * pageCurrent; i++) {
			// 	data.push({
			// 		key: i,
			// 		id : `52394829${i}`, 
			// 		symbol : "rb2001", 
			// 		"modetype" : "gytk", 
			// 		"bstype" : "buy", 
			// 		"vol" : 1, 
			// 		"voldef" : 1, 
			// 		"enable" : 1, 
			// 		"exec" : 1, 
			// 		"state" : 2, 
			// 		"execdt" : ""	 
			// 	});
			// }

			// this.setState({
			// 	loading: false,
			// 	data,
			// 	pagination: pager
			// });			

			$axios({
				url: Host +'/v1/commuser/list',
				method: 'get',
				type: 'json'
			}).then(data => {
				console.log("resp:", data )
				if (data.data.code === "0000"){
					data =  data.data.data
					console.log(   (data.length) );
					pager.total =  (data.length)  ;
					for (let i=0;i<data.length;i++ ){
						data[i].key = i
					}
				 				
				}else{
					console.log( data.data.code );
					openNotification_err("读取失败，请重新刷新", data.data.msg )
				}

				this.setState({
					loading: false,
					data,
					pagination: pager
				});				
		 
			});	





		}, 100);
	};
	selectRow = record => {
		const selectedRowKeys = [...this.state.selectedRowKeys];
		if (selectedRowKeys.indexOf(record.key) >= 0) {
			selectedRowKeys.splice(selectedRowKeys.indexOf(record.key), 1);
		} else {
			selectedRowKeys.push(record.key);
		}
		this.setState({ selectedRowKeys });
	};
	onSelectedRowKeysChange = selectedRowKeys => {
		console.log('selectedRowKeys changed: ', selectedRowKeys);
		this.setState({ selectedRowKeys });
	};
	// 编辑
	handleEdit(row) {
		this.setState({ currentRow: row, visible: true });
	}
	// 删除
	handleDel(row) {
		let _this = this;
		confirm({
			title: '温馨提示',
			content: '确定要删除当前内容吗？',
			okText: '确定',
			cancelText: '取消',
			onOk() {
				message.info('你点击了确定，当前行key为：' + row.key, 1);
				
				let reqData = {id: row.id }

				$axios({
					url: Host +'/v1/commuser/del',
					method: 'post',
					data:reqData,
					type: 'json'
				}).then(data => {
					console.log("resp:", data )
					if (data.data.code === "0000"){
						openNotification_succ("操作成功", data.data.msg )
					}else{
						console.log( data.data.code );
						openNotification_err("操作失败", data.data.msg )
					}	
					
					// 操作完跳转到第一页
					const pager = { ..._this.state.pagination };
					pager.current = 1;
					_this.setState({ pagination: pager });
					_this.fetch(1, _this.state.pagination.pageSize);
					// console.log(_this.state.selectedRowKeys)								
			 
				});	
			 

			},
			onCancel() {}
		});
	}
	handleOk = () => {
		this.setState({ visible: false });
	};

	handleCancel = () => {
		this.setState({ visible: false });
	};
	// 提交(修改)
	handleSubmit = e => {
		e.preventDefault();
		let _this = this;
		this.formRef.props.form.validateFields((err, values) => {
			if (!err) {
	 
				console.log('Received values of form: ', values);
				this.setState({ visible: false });
				// 此处应该为后台业务逻辑
				setTimeout(() => {
					Modal.info({
						title: '点击了提交修改',
						content: (
							<div>
								<p>当前表单内容为：</p>
								<p>{JSON.stringify(values)}</p>
							</div>
						),
						onOk() {
							let reqData = values

							$axios({
								url: Host +'/v1/commuser/update',
								method: 'post',
								data:reqData,
								type: 'json'
							}).then(data => {
								console.log("resp:", data )
								if (data.data.code === "0000"){
									openNotification_succ("操作成功", data.data.msg )
								}else{
									console.log( data.data.code );
									openNotification_err("操作失败", data.data.msg )
								}	
								
								// 操作完跳转到第一页
								const pager = { ..._this.state.pagination };
								pager.current = 1;
								_this.setState({ pagination: pager });
								_this.fetch(1, _this.state.pagination.pageSize);
								// console.log(_this.state.selectedRowKeys)								
						 
							});	
						}
					});
				}, 500);
			}
		});
	};
	// 增加
	handleAdd = e => {
		e.preventDefault();
		let _this = this;
		this.formRef.props.form.validateFields((err, values) => {
			if (!err) {
				console.log('Received values of form: ', values);
				this.setState({ visible: false });
				// 此处应该为后台业务逻辑
				setTimeout(() => {
					Modal.info({
						title: '点击了增加',
						content: (
							<div>
								<p>当前表单内容为：</p>
								<p>{JSON.stringify(values)}</p>
							</div>
						),
						onOk() {
							let reqData = values

							$axios({
								url: Host +'/v1/commuser/add',
								method: 'post',
								data:reqData,
								type: 'json'
							}).then(data => {
								console.log("resp:", data )
								if (data.data.code === "0000"){
									openNotification_succ("操作成功", data.data.msg )
								}else{
									console.log( data.data.code );
									openNotification_err("操作失败", data.data.msg )
								}	
								
								// 操作完跳转到第一页
								const pager = { ..._this.state.pagination };
								pager.current = 1;
								_this.setState({ pagination: pager });
								_this.fetch(1, _this.state.pagination.pageSize);
								// console.log(_this.state.selectedRowKeys)								
						 
							});	
						}
					});
				}, 500);
			}
		});
	};	

	rowclass = (record, index) => {
		let className = 'light-row';
		if (index % 2 === 1) className = 'dark-row';
		return className;
	}

	
	

	render() {
		const { selectedRowKeys, loading, pagination, columns, data } = this.state;
		const rowSelection = {
			selectedRowKeys,
			onChange: this.onSelectedRowKeysChange
		};
		return (
			<div className="shadow-radius">
				<Table
					bordered
					columns={columns}
					dataSource={data}
					loading={loading}
					onChange={this.handleTableChange}
					pagination={pagination}
					rowSelection={rowSelection}
					rowClassName={this.rowclass}
					onRow={record => ({
						onClick: () => {
							this.selectRow(record);
						}
					})}
				/>
				<Modal title="编辑信息" visible={this.state.visible} onOk={this.handleOk} onCancel={this.handleCancel} footer={null}>
					<ModeEditForm data={this.state.currentRow} visible={this.state.visible} wrappedComponentRef={form => (this.formRef = form)} handleSubmit={this.handleSubmit} handleAdd={this.handleAdd}  />
				</Modal>
			</div>
		);
	}
}

export default BizTableEdit;
