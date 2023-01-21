import React, { useState } from 'react'
import SideBar from './components/sidebar'
import { Col, ConfigProvider, Row } from 'antd'
import Overview from './components/overview'
import Files from './components/files'

function App (): any {
  const [active, setActive] = useState('overview')
  return (
    <div>
        <ConfigProvider theme={{
          token: {
            colorPrimary: '#5D2667',
            fontFamily: '-apple-system,BlinkMacSystemFont,Helvetica Neue,PingFang SC,Microsoft YaHei,Source Han Sans SC,Noto Sans CJK SC,WenQuanYi Micro Hei,sans-serif'
          }
        }}>
        <Row>
            <Col span={4}>
                <SideBar setActive={setActive}/>
            </Col>
            <Col span={20}>
                <Overview show={active === 'overview'}/>
                <Files show={active === 'files'}/>
            </Col>
        </Row>
    </ConfigProvider>
    </div>
  )
}

export default App
