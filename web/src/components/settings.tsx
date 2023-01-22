import React, { useState } from 'react'
import { Button, Col, Divider, Input, InputNumber, message, Row, Select, Slider, Space } from 'antd'
import { Option } from 'antd/es/mentions'
import Axios from '../axios'
import { useNavigate } from 'react-router-dom'
import { setCookie } from 'typescript-cookie'

interface ISettings {
  show: boolean
}

const Settings: React.FC<ISettings> = (props) => {
  const [newPassword, setNewPassword] = useState('')
  const [oldPassword, setOldPassword] = useState('')
  const [hashPrefix, setHashPrefix] = useState('')
  const [maxSize, setMaxSize] = useState('')
  const [maxSizeUnit, setMaxSizeUnit] = useState('MB')
  const [tokenTtl, setTokenTtl] = useState('')
  const navigate = useNavigate()

  const handleChangeSizeUnit = (value: string) => {
    setMaxSizeUnit(value)
  }
  const selectAfter = (
        <Select defaultValue="MB" style={{ width: 70 }} onChange={handleChangeSizeUnit}>
            <Option value='B'>B</Option>
            <Option value="KB">KB</Option>
            <Option value="MB">MB</Option>
            <Option value="GB">GB</Option>
        </Select>
  )

  const handleChangePassword = async () => {
    const data = new URLSearchParams()
    data.append('newPassword', window.md5(newPassword))
    data.append('oldPassword', window.md5(oldPassword))
    await Axios({
      method: 'POST',
      url: '/api/changePassword',
      data
    }).then(res => {
      if (res.status === 200 && res.data?.success) {
        message.success('修改成功')
        navigate('/welcome')
        setCookie('token', '')
      } else {
        message.error('未知错误')
      }
    }).catch(err => {
      console.warn(err)
    })
  }

  const handleChangeHashPrefix = () => {
    const data = new URLSearchParams()
    data.append('hashPrefix', hashPrefix)
    Axios({
      method: 'POST',
      url: '/api/changeHashPrefix',
      data
    }).then(res => {
      if (res.status === 200 || res.data?.success) {
        message.success('修改成功')
      } else {
        message.error(res.data?.reason ?? '未知错误')
      }
    }).catch(err => { console.warn(err) })
  }

  const handleChangeMaxSize = async () => {
    let unit = 20
    switch (maxSizeUnit) {
      case 'B': unit = 1; break
      case 'KB': unit = 10; break
      case 'MB': unit = 20; break
      case 'GB': unit = 30; break
    }
    const total = Number(maxSize) << unit
    const data = new URLSearchParams()
    data.append('size', total.toString())
    await Axios({
      method: 'POST',
      url: '/api/changeMaxFileSize',
      data
    }).then(res => {
      if (res.status === 200 || res.data?.success) {
        message.success('修改成功')
      } else {
        message.error(res.data?.reason ?? '未知错误')
      }
    }).catch(err => { console.warn(err) })
  }

  const handleChangeTokenTtl = async () => {
    const data = new URLSearchParams()
    data.append('tokenTtl', tokenTtl)
    await Axios({
      method: 'POST',
      url: '/api/changeTokenTtl',
      data
    }).then(res => {
      if (res.status === 200 || res.data?.success) {
        message.success('修改成功')
      } else {
        message.error(res.data?.reason ?? '未知错误')
      }
    }).catch(err => { console.warn(err) })
  }
  return (
        <div style={{ margin: '20px' }}>
            {props.show &&
                <>
                    <h2>重置密码</h2>
                    <p>用于更改登录密码，第一次登录建议修改默认密码。</p>
                    <Row >
                        <Col span={2}/>
                        <Col>
                        <Space direction='vertical'>
                            <Input.Password placeholder='旧密码' width={250} onChange={(e) => { setOldPassword(e.target.value) }}/>
                            <Input.Password placeholder='新密码' width={250} onChange={(e) => { setNewPassword(e.target.value) }}/>
                            <Button type='primary' style={{ width: '250px' }} disabled={newPassword === '' || oldPassword === ''} onClick={handleChangePassword}>重置</Button>
                        </Space>
                    </Col>
                    </Row>
                    <Divider/>
                    <h2>更改哈希前缀</h2>
                    <p>用于计算文件存储路径，该数值越大存储的越分散。</p>
                    <Row >
                        <Col span={2}/>
                        <Col>
                            <Space>
                                <InputNumber min={1} max={32} style={{ width: '180px' }} onChange={(v) => { setHashPrefix(v?.toString() ?? '') }}/>
                                <Button type='primary' disabled={hashPrefix === ''} onClick={handleChangeHashPrefix}>确定</Button>
                            </Space>
                        </Col>
                    </Row>
                    <Divider/>
                    <h2>更改最大文件限制</h2>
                    <p>用于限制上传文件的大小，更改后需要重新启动服务。</p>
                    <Row >
                        <Col span={2}/>
                        <Col>
                            <Space>
                                <InputNumber addonAfter={selectAfter} min={0} style={{ width: '180px' }} onChange={(v) => { setMaxSize(v?.toString() ?? '') }}/>
                                <Button type='primary' disabled={maxSize === ''} onClick={handleChangeMaxSize}>确定</Button>
                            </Space>
                        </Col>
                    </Row>
                    <Divider/>
                    <h2>更改最长登录时间</h2>
                    <p>用于限制用户登录时间，超过该时间需要重新登录，更改后需要重新启动服务。</p>
                    <Row >
                        <Col span={2}/>
                        <Col>
                            <Space>
                                <Slider min={0} max={24} tooltip={{ formatter: (value: number | undefined) => `${value}小时` }} style={{ width: '180px' }} onChange={(v) => { setTokenTtl(v.toString()) }}/>
                                <Button type='primary' disabled={tokenTtl === ''} onClick={handleChangeTokenTtl}>确定</Button>
                            </Space>
                        </Col>
                    </Row>
                </>
            }
        </div>
  )
}

export default Settings
