import React, { useEffect, useState } from 'react'
import { Button, Col, DatePicker, Input, message, Row, Select, Space, Table, Tag } from 'antd'
import { EyeOutlined, EditOutlined, DeleteOutlined } from '@ant-design/icons'
import tagRender, { tagOptions, toParam } from './tagrender'
import { ColumnsType } from 'antd/es/table'
import Axios from '../axios'
import FileDrawer from './filedrawer'

interface FileDetail {
  id: number
  filename: string
  extension: string
  labels: string[]
  tags: string[]
  uploadTime: string
  size: string
  location: string
}

interface IOption {
  key: string
  label: string
  value: string
}

interface IFiles {
  show: boolean
}

const Files: React.FC<IFiles> = (props) => {
  const [extensionOptions, setExtensionOptions] = useState<IOption[]>([])
  const [labelOptions, setLabelOptions] = useState<IOption[]>([])
  const [data, setData] = useState<FileDetail[]>([])
  const [filename, setFilename] = useState('')
  const [extension, setExtension] = useState('')
  const [labels, setLabels] = useState('')
  const [tags, setTags] = useState('')
  const [startTime, setStartTime] = useState('')
  const [endTime, setEndTime] = useState('')
  const [show, setShow] = useState(false)
  const [isUpload, setIsUpload] = useState(true)
  const [modifyFile, setModifyFile] = useState<FileDetail | void>()
  const tableColumns: ColumnsType<FileDetail> = [
    {
      title: '文件名',
      dataIndex: 'filename'
    }, {
      title: '大小',
      dataIndex: 'size'
    }, {
      title: '上传日期',
      dataIndex: 'uploadTime'
    }, {
      title: '标签',
      dataIndex: 'labels',
      render: (_, { labels }) => (
                <>
                    {
                        labels.map(label => {
                          return (label &&
                                <Tag key={label}>
                                    {label}
                                </Tag>
                          )
                        })
                    }
                </>

      )
    }, {
      title: '标记',
      dataIndex: 'tags',
      render: (_, { tags }) => (
                <>
                    {
                        tags.map((tag) => {
                          let color = 'null'
                          switch (tag) {
                            case '0': color = '#ffccc7'; break
                            case '1': color = '#fff1b8'; break
                            case '2': color = '#f4ffb8'; break
                            case '3': color = '#b5f5ec'; break
                            case '4': color = '#bae0ff'; break
                          }
                          return (
                                <>
                                    {color !== 'null' &&
                                        <Tag color={color} key={tag}>{'\u3000'}</Tag>
                                    }
                                </>
                          )
                        })
                    }
                </>
      )
    }, {
      title: '操作',
      dataIndex: 'operation',
      render: (_, record) => (
                <Space size='middle'>
                    <Button shape='circle' size='small' icon={<EyeOutlined />} onClick={() => { window.open(Axios.defaults.baseURL + record.location) }}/>
                    <Button shape='circle' size='small' icon={<EditOutlined />} onClick={() => { handleClickModify(record) }}/>
                    <Button shape='circle' size='small' icon={<DeleteOutlined />} onClick={async () => { await handleDeleteFile(record.filename) }}/>
                </Space>
      )
    }
  ]

  const handleDeleteFile = async (filename: string) => {
    const data = new URLSearchParams()
    data.append('filename', filename)
    await Axios({
      method: 'POST',
      url: '/api/deleteFile',
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded'
      },
      data
    }).then(res => {
      if (res.status === 200) {
        message.success('删除成功')
      }
    }).catch(err => {
      console.warn(err)
    })
    handleListFilesByCondition()
  }

  const handleChangeFilename = (e: { target: { value: React.SetStateAction<string> } }) => {
    setFilename(e.target.value)
  }

  const handleChangeExtension = (v: any, event: any) => {
    setExtension(v)
  }

  const handleChangeLabels = (v: any, event: any) => {
    setLabels(v.join(':'))
  }

  const handleChangeTags = (v: string[], event: any) => {
    setTags(v.map(t => toParam(t)).join(':'))
  }

  const handleChangeDate = (d: any, s: string[]) => {
    setStartTime(s[0] ? s[0] + ' 00:00:00' : '')
    setEndTime(s[1] ? s[1] + ' 23:59:59' : '')
  }

  const handleClickUpload = () => {
    setShow(true)
    setIsUpload(true)
  }

  const handleClickModify = (record: FileDetail) => {
    setModifyFile(record)
    setShow(true)
    setIsUpload(false)
  }

  const handleOverallExtensionInfo = async () => {
    await Axios({
      method: 'GET',
      url: '/api/overallExtensionInfo'
    }).then(res => {
      if (res.status === 200) {
        setExtensionOptions(Object.keys(res.data).map(k => {
          return {
            key: k,
            label: k,
            value: k
          }
        }))
      }
    }).catch(err => {
      console.warn(err)
    })
  }

  const handleOverallLabelInfo = async () => {
    await Axios({
      method: 'GET',
      url: '/api/overallLabelInfo'
    }).then(res => {
      if (res.status === 200) {
        setLabelOptions(Object.keys(res.data).map(k => {
          return {
            key: k,
            label: k,
            value: k
          }
        }))
      }
    }).catch(err => {
      console.warn(err)
    })
  }

  const handleListFilesByCondition = async (payload: {
    filename: string
    extension: string
    labels: string
    tags: string
    startTime: string
    endTime: string
  } | void) => {
    await Axios({
      method: 'GET',
      url: '/api/listFilesByCondition',
      params: {
        filename: payload?.filename ?? '',
        extension: payload?.extension ?? '',
        labels: payload?.labels ?? '',
        tags: payload?.tags ?? '',
        startTime: payload?.startTime ?? '',
        endTime: payload?.endTime ?? ''
      }
    }).then(res => {
      if (res.status === 200) {
        setData(res.data?.nasuFiles.map((f: { Id: any, Filename: any, Extension: any, Labels: string, Tags: string, UploadTime: any, Size: any, Location: any }) => {
          return {
            id: f.Id,
            filename: f.Filename,
            extension: f.Extension,
            labels: f.Labels.split(','),
            tags: f.Tags.split(','),
            uploadTime: f.UploadTime.split('T')[0],
            size: f.Size,
            location: f.Location
          }
        }))
      }
    }).catch(err => {
      console.warn(err)
    })
  }

  useEffect(() => {
    handleOverallLabelInfo()
    handleOverallExtensionInfo()
    handleListFilesByCondition()
  }, [])
  return (
        <div style={{ margin: '20px' }}>
            {props.show &&
                <>
                    <Row>
                        <Col span={8}>
                            <Space direction='horizontal'>
                                <label>文件名：</label>
                                <Input placeholder="文件名" onChange={handleChangeFilename} style={{ width: '250px' }}/>
                            </Space>
                        </Col>
                        <Col span={8}>
                            <Space direction='horizontal'>
                                <label>扩展名：</label>
                                <Select style={{ width: '250px' }} allowClear placeholder="扩展名" options={extensionOptions} onChange={handleChangeExtension}/>
                            </Space>
                        </Col>

                        <Col span={8}>
                            <Space direction='horizontal'>
                                <label>上传日期：</label>
                                <DatePicker.RangePicker onChange={handleChangeDate}/>
                            </Space>

                        </Col>
                    </Row>
                    <Row style={{ marginTop: '10px' }}>
                        <Col span={8}>
                            <Space direction='horizontal'>
                                <label>标签：&nbsp;&nbsp;&nbsp;</label>
                                <Select placeholder="标签"
                                        mode="multiple"
                                        allowClear
                                        style={{ width: '250px' }}
                                        options={labelOptions}
                                        onChange={handleChangeLabels}
                                />
                            </Space>
                        </Col>
                        <Col span={8}>
                            <Space direction='horizontal'>
                                <label>标记：&nbsp;&nbsp;&nbsp;</label>
                                <Select placeholder="标记"
                                        mode="multiple"
                                        allowClear
                                        tagRender={tagRender}
                                        style={{ width: '250px' }}
                                        options={tagOptions}
                                        onChange={handleChangeTags}
                                />
                            </Space>
                        </Col>
                        <Col span={5}>
                            <Button type='default' style={{ float: 'right' }} onClick={async () => {
                              await handleListFilesByCondition({
                                filename, endTime, extension, labels, startTime, tags
                              })
                            }}>查询</Button>
                        </Col>
                        <Col span={2} style={{ marginLeft: '25px' }}>
                            <Button type='primary' onClick={handleClickUpload}>上传</Button>
                        </Col>
                    </Row>
                    <Row style={{ marginTop: '10px' }} >
                        <Col span={24}>
                            <Table columns={tableColumns} dataSource={data} style={{ width: '1190px' }}/>
                        </Col>
                    </Row>
                    <FileDrawer show={show} setShow={setShow} labelOptions={labelOptions} refresh={handleListFilesByCondition} isUpload={isUpload} record={modifyFile}/>
                </>
            }
        </div>
  )
}

export default Files
