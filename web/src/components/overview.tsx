import React, {useEffect, useState} from 'react';
import {Col, message, Row} from 'antd';
import ChartCard from "./chartcard";
import ListCard from "./listcard";
import Axios from "../axios";

interface IOverview {
    show: boolean
}

const Overview: React.FC<IOverview> = (props) => {
    const [labelLabels, setLabelLabels] = useState<Array<string>>([]);
    const [labelData, setLabelData] = useState<Array<number>>([]);
    const [tagLabels, setTagLabels] = useState<Array<string>>([]);
    const [tagData, setTagData] = useState<Array<number>>([])
    const [messageApi, contextHolder] = message.useMessage();
    const handleOverallLabelInfo = () => Axios({
        method: 'GET',
        url: '/api/overallLabelInfo',
        headers: {
            'Authorization': window.token
        }
        }).then(res => {
            if (res.status == 200) {
                setLabelLabels(Object.keys(res.data));
                setLabelData(Object.values(res.data));
            } else {
                messageApi.open({
                    type: 'error',
                    content: '请求错误',
                });
            }
        }).catch(err => {
            console.warn(err);
    })
    const handleOverallTagInfo = () => Axios({
        method: 'GET',
        url: '/api/overallTagInfo',
        headers: {
            'Authorization': window.token
        }
    }).then(res => {
        if (res.status == 200) {
            setTagLabels(Object.keys(res.data));
            setTagData(Object.values(res.data));
        } else {
            messageApi.open({
                type: 'error',
                content: '请求错误',
            });
        }
    }).catch(err => {
        console.warn(err);
    })

    useEffect(() => {
        handleOverallLabelInfo();
        handleOverallTagInfo();
    }, [])
    return (
        <>
        {props.show &&
        <>
            <Row>
                <Col span={4}>
                    <h1>概况总览</h1>
                </Col>
                <Col span={20}/>
            </Row>
            <Row>
                <Col span={8}>
                    <ChartCard width={400} height={300} title={"标签数据"} labels={labelLabels} data={labelData}/>
                </Col>
                <Col span={8}>
                    <ChartCard width={400} height={300} title={"标记数据"} labels={tagLabels} data={tagData}/>
                </Col>
                <Col span={8}>
                    <ChartCard width={400} height={300} title={"标记数据"} labels={["test1", "test2"]} data={[5, 10]}/>
                </Col>
            </Row>
            <Row>
                <Col span={12}>

                </Col>
            </Row>
        </>}
        </>
    )
}

export default Overview;