import React, {useState} from 'react';
import SideBar from "./components/sidebar";
import {Col, Row} from "antd";
import Overview from "./components/overview";
import Files from "./components/files";


function App() {
    const [active, setActive] = useState('overview')
  return (
    <>
        <Row>
            <Col span={4}>
                <SideBar setActive={setActive}/>
            </Col>
            <Col span={20}>
                <Overview show={active === 'overview'}/>
                <Files show={active === 'files'}/>
            </Col>
        </Row>

    </>
  );
}

export default App;
