import React, { Component } from 'react';
import Particles from 'react-particles-js';
import { Form, Icon, Input, Button, notification } from 'antd';
import { connect } from 'react-redux';
import { setUserInfo } from '@/redux/actions/userInfo';
import '@/assets/css/login';
import $axios from '../axios/$axios';
//import {post,fetch,   } from '../axios/$axios';

import { openNotification_err,openNotification_succ } from './notification/index'

const Host = process.env.API_HOST || "http://localhost:8888"
           
// /v1/admin/login
// console.log("env:",process.env)

const FormItem = Form.Item;
class Login extends Component {
	state = { clientHeight: document.documentElement.clientHeight || document.body.clientHeight };
	constructor(props) {
		super(props);
		this.onResize = this.onResize.bind(this);
		//debug
		console.log("env:",process.env)
		
		 
	}
	login = e => {
		e.preventDefault();
		this.props.form.validateFields((err, values) => {
			console.log("values:", values )
			console.log("url:", process.env.PUBLIC_URL )
			let reqdata =  { "username": values.userName , "loginpwd": values.password }
			

			if (!err) {
				
			} else {
				
				return 
			}
 
			$axios({
				url: Host +'/v1/admin/login',
				method: 'post',
				data : reqdata,
				type: 'json'
			}).then(data => {

				console.log("resp:", data )
				if (data.data.code === "0000"){
					localStorage.setItem('isLogin', '1');
					// 模拟生成一些数据
					this.props.setUserInfo(Object.assign({}, values, { role: { type: 1, name: '超级管理员' } }));
					localStorage.setItem('userInfo', JSON.stringify(Object.assign({}, values, { role: { type: 1, name: '超级管理员' } })));
					this.props.history.push('/dashboard');					
				}else{
					console.log( data.data.code );
					openNotification_err("获取数据失败", data.data.msg )
				}
		 
			});		
 

	 
		});
	};
	componentDidMount() {
		window.addEventListener('resize', this.onResize);
	}
	componentWillUnmount() {
		window.addEventListener('resize', this.onResize);
		// componentWillMount进行异步操作时且在callback中进行了setState操作时，需要在组件卸载时清除state
		this.setState = () => {
			return;
		};
	}
	onResize() {
		this.setState({ clientHeight: document.documentElement.clientHeight || document.body.clientHeight });
	}
	render() {
		const { getFieldDecorator } = this.props.form;
		return (
			<div className="container">
				<Particles
					height={this.state.clientHeight - 5 + 'px'}
					params={{
						number: { value: 50 },
						ize: { value: 3 },
						interactivity: {
							events: {
								onhover: { enable: true, mode: 'repulse' }
							}
						}
					}}
				/>
				<div className="content">
					<div className="title">后台管理系统</div>
					<Form className="login-form">
						<FormItem>
							{getFieldDecorator('userName', {
								rules: [{ required: true, message: '请填写用户名！' }]
							})(<Input prefix={<Icon type="user" style={{ color: 'rgba(0,0,0,.25)' }} />} placeholder="用户名" />)}
						</FormItem>
						<FormItem>
							{getFieldDecorator('password', {
								rules: [{ required: true, message: '请填写密码！' }]
							})(<Input.Password prefix={<Icon type="lock" style={{ color: 'rgba(0,0,0,.25)' }} />} placeholder="密码" />)}
						</FormItem>
						<FormItem>
							<Button type="primary" htmlType="submit" block onClick={this.login}>
								登录
							</Button>
							<div style={{ color: '#999',paddingTop:'10px',textAlign:'center' }}>Tips : 输入任意用户名密码即可</div>
						</FormItem>
					</Form>
				</div>
			</div>
		);
	}
}

const mapStateToProps = state => state;
const mapDispatchToProps = dispatch => ({
	setUserInfo: data => {
		dispatch(setUserInfo(data));
	}
});
export default connect(
	mapStateToProps,
	mapDispatchToProps
)(Form.create()(Login));
