import React, {useState} from "react";
import {Button, Col, Input, message, Row, Space} from "antd";
import Axios from "./axios";
import {useNavigate} from "react-router-dom";

function Welcome() {
    const [password, setPassword] = useState('');
    const [messageApi, contextHolder] = message.useMessage();
    const navigate = useNavigate();

    const handleOnChange: React.ChangeEventHandler<HTMLInputElement> = (e) => {
        setPassword(e.target.value)
    }
    const handleOnClick = () => {
        const hash = window.md5(password);
        Axios(
            '/login',
            {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded'
                },
                data: {
                    password: hash
                }
            }
        ).then(res => {
            if (res.status == 200) {
                window.token = res.data?.token;
                navigate('/');
            } else {
                messageApi.open({
                    type: 'error',
                    content: '请求错误',
                });
            }
        }).catch(err => {
            console.warn(err);
        })
    }
    return (
        <>
            <Row style={{margin: '0px', height: '721px', backgroundColor: '#5D2667', textAlign: 'center'}}>
                <Col span={18} style={{backgroundColor: 'black'}}/>
                <Col span={6} >
                    <Row align="middle" justify='center' style={{margin: '0px', height: '721px', textAlign: 'center'}}>
                        <Space direction='vertical' align='center'>
                            <h1>NASU</h1>
                            <Space direction='horizontal'>
                                <Input.Password placeholder="请输入密钥" onChange={handleOnChange}/>
                                <Button onClick={handleOnClick}>确认</Button>
                            </Space>
                        </Space>
                    </Row>
                </Col>

            </Row>
        </>
    )
}

export default Welcome;