import React, { Component } from 'react';
import { Row,Col,Card } from 'antd';

class Index extends Component {
	// 	#D1EEEE  ECECEC
	// {name:"", "heatbeat":"dt", "bizlog":"", "state":"正常/异常" }

	render() {
		return (
			<div className="shadow-radius">
				<div style={{ background: '#D1EEEE', padding: '30px' }}>  
					<Row gutter={16}>
					<Col span={8}>
						<Card title="Card title" bordered={false}>
						Card content
						</Card>
					</Col>
					<Col span={8}>
						<Card title="Card title" bordered={false}>
						Card content
						</Card>
					</Col>
					<Col span={8}>
						<Card title="Card title" bordered={false}>
						Card content
						</Card>
					</Col>
					</Row>
				</div>
			</div>
		);
	}
}

export default Index;
