import React from 'react'
import { Menu } from 'antd'

interface ISidebar {
  setActive: (active: string) => void
}

const SideBar: React.FC<ISidebar> = (props) => {
  const { setActive } = props
  const items = [{
    key: 'overview',
    label: '概况总览'
  }, {
    key: 'files',
    label: '文件详情'
  }, {
    key: 'settings',
    label: '平台设置'
  }]
  const onClick = (e: any) => {
    const { key } = e
    setActive(key)
  }
  return (
        <Menu
            onClick={onClick}
            style={{ width: 256, top: 0, bottom: 0, height: 753, backgroundColor: '#5D2667' }}
            defaultSelectedKeys={['overview']}
            mode="vertical"
            items={items}
        />
  )
}

export default SideBar
