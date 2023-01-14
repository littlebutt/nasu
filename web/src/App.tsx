import React, {useState} from 'react';
import './App.css';
import SideBar from "./components/sidebar";
import {Col, Row} from "antd";
import Overview from "./components/overview";


function App() {
    window.host = 'http://' + document.location.host + ':' + document.location.port;
    const [active, setActive] = useState('overview')
  return (
    <>
        <Row>
            <Col span={4}>
                <SideBar setActive={setActive}/>
            </Col>
            <Col span={20}>
                <Overview show={active === 'overview'}/>
            </Col>
        </Row>

    </>
  );
}

export default App;
