import React, { Component } from 'react';
import { Form, Input,InputNumber, Button } from 'antd';

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

class ModeEditForm extends Component {
	componentWillReceiveProps(nextProps) {
		!nextProps.visible && this.props.form.resetFields();
	}
	render() {
		const { getFieldDecorator } = this.props.form;
		const { data } = this.props;
		const formItemLayout = {
			labelCol: { span: 8 },
			wrapperCol: { span: 16 }
		};
		const formTailLayout = {
			labelCol: { span: 4 },
			wrapperCol: { span: 20, offset: 4 ,margin:4}
		};
		return (
			<Form onSubmit={this.handleSubmit} refs="editForm">
				<Form.Item label="用户编号" {...formItemLayout}>
					{getFieldDecorator('userid', {
						initialValue: data.userid,
						rules: [
							{
								required: true,
								message: '用户编号'
							}
						]
					})(<Input />)}
				</Form.Item>

				<Form.Item label="用户密码" {...formItemLayout}>
					{getFieldDecorator('userpwd', {
						initialValue: data.userpwd,
						rules: [
							{
								required: true,
								message: '请输入用户密码'
							}
						]
					})(<Input />)}
				</Form.Item>

				<Form.Item label="支付状态" {...formItemLayout}>
					{getFieldDecorator('userlevel', {
						initialValue: data.userlevel,
						rules: [
							{
								required: true,
								message: '请输入支付状态，1未支付 2已支付'
							}
						]
					})(<InputNumber />)}
				</Form.Item>

				<Form.Item label="支付时间" {...formItemLayout}>
					{getFieldDecorator('paydt', {
						initialValue: data.paydt,
						rules: [
							{
								required: true,
								message: '请输入支付时间'
							}
						]
					})(<Input  />)}
				</Form.Item>				

				<Form.Item label="用户状态" {...formItemLayout}>
					{getFieldDecorator('state', {
						initialValue: data.state,
						rules: [
							{
								required: true,
								message: '请输入状态0 未开启 1已开启'
							}
						]
					})(<InputNumber />)}
				</Form.Item>	

				<Form.Item label="备注信息" {...formItemLayout}>
					{getFieldDecorator('meo', {
						initialValue: data.meo,
						rules: [
							{
								required: true,
								message: '请输入备注信息'
							}
						]
					})(<Input />)}
				</Form.Item>				
 
				<Form.Item {...formTailLayout}>
					<Button type="primary" onClick={this.props.handleSubmit}  >
						保存
					</Button><span></span>
					<Button type="primary" onClick={this.props.handleAdd}  >
						添加
					</Button>					
				</Form.Item>
			</Form>
		);
	}
}
export default Form.create()(ModeEditForm);
