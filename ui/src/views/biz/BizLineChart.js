import React, { Component } from 'react';
import { connect } from 'react-redux';
import echarts from 'echarts';
import Chart from '@/components/chart/Chart';

import $axios from '../../axios/$axios';
import { openNotification_err,openNotification_succ } from '../notification/index'


const Host = process.env.API_HOST || "http://localhost:8888"

// { 
//     "_id" : ObjectId("5dd4acbfa1da650b7c481bee"), 
//     "date" : "2019-11-20", 
//     "totalusers" : 100.0, 
//     "payusers" : 5.0
// }

let dateData = ['13:00', '13:05', '13:10', '13:15', '13:20', '13:25', '13:30', '13:35', '13:40', '13:45', '13:50', '13:55']
let totalsData = [220, 182, 191, 134, 150, 120, 110, 125, 145, 122, 165, 122]
let paysData = [120, 110, 125, 145, 122, 165, 122, 220, 182, 191, 134, 150]
let otherData=[]

let chartData = {
	backgroundColor: '#fff',
	title: {
		top: 30,
		text: '折线图',
		textStyle: {
			fontWeight: 'normal',
			fontSize: 16,
			color: '#57617B'
		},
		left: 'center'
	},
	tooltip: {
		trigger: 'axis',
		axisPointer: {
			type: 'cross'
		},
		padding: [5, 10]
	},
	// tab
	legend: {
		top: 20,
		icon: 'rect',
		itemWidth: 14,
		itemHeight: 5,
		itemGap: 13,
		right: '2%',
		textStyle: {
			fontSize: 12,
			color: '#57617B'
		}
	},

	// 图表
	grid: {
		top: 80,
		left: '2%',
		right: '2%',
		bottom: '6%',
		containLabel: true
	},
	// x轴
	xAxis: [
		{
			type: 'category', //分类
			boundaryGap: false,
			axisLine: {
				lineStyle: {
					color: '#57617B'
				}
			},
			data: dateData
		}
	],
	yAxis: [
		{
			type: 'value',
			name: '(%)',
			axisTick: {
				show: false
			},
			axisLine: {
				lineStyle: {
					color: '#57617B'
				}
			},
			axisLabel: {
				margin: 10,
				textStyle: {
					fontSize: 14
				}
			}
		}
	],
	series: [
		{
			name: '总用户数',
			type: 'line',
			smooth: true,
			symbol: 'circle',
			symbolSize: 5,
			showSymbol: false,
			lineStyle: {
				normal: {
					width: 1
				}
			},
			areaStyle: {
				normal: {
					color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
						{
							offset: 0,
							color: 'rgba(137, 189, 27, 0.3)'
						},
						{
							offset: 0.8,
							color: 'rgba(137, 189, 27, 0)'
						}
					]),
					shadowColor: 'rgba(0, 0, 0, 0.1)',
					shadowBlur: 10
				}
			},
			itemStyle: {
				normal: {
					color: 'rgb(137,189,27)',
					borderColor: 'rgba(137,189,2,0.27)',
					borderWidth: 12
				}
			},
			data: totalsData
		},
		{
			name: '支付用户数',
			type: 'line',
			smooth: true,
			symbol: 'circle',
			symbolSize: 5,
			showSymbol: false,
			lineStyle: {
				normal: {
					width: 1
				}
			},
			areaStyle: {
				normal: {
					color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
						{
							offset: 0,
							color: 'rgba(0, 136, 212, 0.3)'
						},
						{
							offset: 0.8,
							color: 'rgba(0, 136, 212, 0)'
						}
					]),
					shadowColor: 'rgba(0, 0, 0, 0.1)',
					shadowBlur: 10
				}
			},
			itemStyle: {
				normal: {
					color: 'rgb(0,136,212)',
					borderColor: 'rgba(0,136,212,0.2)',
					borderWidth: 12
				}
			},
			data: paysData
		},
		{
			name: '其它',
			type: 'line',
			smooth: true,
			symbol: 'circle',
			symbolSize: 5,
			showSymbol: false,
			lineStyle: {
				normal: {
					width: 1
				}
			},
			areaStyle: {
				normal: {
					color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
						{
							offset: 0,
							color: 'rgba(219, 50, 51, 0.3)'
						},
						{
							offset: 0.8,
							color: 'rgba(219, 50, 51, 0)'
						}
					]),
					shadowColor: 'rgba(0, 0, 0, 0.1)',
					shadowBlur: 10
				}
			},
			itemStyle: {
				normal: {
					color: 'rgb(219,50,51)',
					borderColor: 'rgba(219,50,51,0.2)',
					borderWidth: 12
				}
			},
			data: otherData
		}
	]
};

//const BizLineChart = props => <Chart chartData={chartData} height={'500px'} style={{ padding: 0 }} {...props} />;


class BizLineChart extends Component {

	constructor(props){
        super(props)
        this.state={
			data: {},
			pagination: {},
			loading: false,
		}
	}
 
	componentWillMount() {
		this.state.data = chartData
		// setInterval(this.fetch(), 1000)
		this.fetch() 
	}

	componentWillUnmount() {
		// componentWillMount进行异步操作时且在callback中进行了setState操作时，需要在组件卸载时清除state
		this.setState = () => {
			return;
		};
	}	

	fetch = () => {
		this.setState({ loading: true });

		$axios({
			url: Host +'/v1/summ/list',
			method: 'get',
			type: 'json'
		}).then(result => {
			let data =[]
			console.log("resp:", result )
			if (result.data.code === "0000"){

				data =  result.data.data
				console.log(   (data.length) );
				if( (data.length) <=0 ){
					return 
				}
				dateData = []
				totalsData=[]
				paysData=[]
				for (let i=0;i<data.length;i++ ){
					dateData.push(  data[i].date ) 
					totalsData.push(  data[i].totalusers ) 
					paysData.push(  data[i].payusers ) 
				}	
				// console.log(  "dateData:", dateData );			
				chartData.xAxis[0].data = dateData
				chartData.series[0].data = totalsData
				chartData.series[1].data = paysData

				this.setState({
					loading: false,
					data:chartData 
				});					
							 
			}else{
				console.log( result.data.code );
				openNotification_err("读取失败，请重新刷新", result.data.msg )
			}
 
		
	 
		});	

	 
	}

	render() {
		return (
			<Chart chartData={ this.state.data } height={'500px'} style={{ padding: 0 }}  {...this.props}  />
		)	
	}

}	

const mapStateToProps = state => {
	return {
		collapse: state.collapse
	};
};

export default connect(mapStateToProps)(BizLineChart);
//export default BizLineChart;

