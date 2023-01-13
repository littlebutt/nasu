import React from 'react';
import { Col, Row } from 'antd';

interface IOverview {

}

const Overview: React.FC<IOverview> = (props) => {
    return (
        <>
            <Row>
                <Col span={4}>
                    概况总览
                </Col>
                <Col span={20}/>
            </Row>
            <Row>
                <Col span={8}>col-12</Col>
                <Col span={8}>col-12</Col>
                <Col span={8}></Col>
            </Row>
            <Row>
                <Col span={24}>col-12</Col>
            </Row>
        </>
    )
}