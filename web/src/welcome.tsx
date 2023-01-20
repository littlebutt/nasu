import React, {useState} from "react";
import {Button, Col, Input, message, Row, Space} from "antd";
import Axios from "./axios";
import {useNavigate} from "react-router-dom";
import {setCookie} from "typescript-cookie";
import axios from "axios";

function Welcome() {
    const [password, setPassword] = useState('');
    const navigate = useNavigate();

    const handleOnChange: React.ChangeEventHandler<HTMLInputElement> = (e) => {
        setPassword(e.target.value)
    }
    const handleOnClick = () => {
        const hash = window.md5(password);
        axios(
            Axios.defaults.baseURL + '/login',
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
                console.log(res.data?.token)
                setCookie('token', res.data?.token);
                navigate('/');
            } else {
                message.error("请求错误");
            }
        }).catch(err => {
            console.warn(err);
        })
    }
    return (
        <>
            <Row style={{margin: '0px', height: '753px', backgroundColor: '#5D2667', textAlign: 'center'}}>
                <Col span={18} style={{backgroundColor: 'white'}}>
                    <embed src='./logo.svg' type="image/svg+xml"/>
                </Col>
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