import React, {useEffect, useState} from 'react';
import {Button, Col, DatePicker, Input, Row, Select, Space, Table, Tag} from "antd";
import { SearchOutlined, FileAddTwoTone, EyeOutlined, EditOutlined, DeleteOutlined } from '@ant-design/icons';
import tagRender from "./tagrender";
import {ColumnsType} from "antd/es/table";


interface FileDetail {
    filename: string,
    extension: string,
    labels: string[],
    tags: string[],
    uploadTime: string,
    size: string
}

interface IFiles {
    show: boolean
}

const Files: React.FC<IFiles> = (props) => {
    const [extensionOptions, setExtensionOptions] = useState([]);
    const [labelOptions, setLabelOptions] = useState([]);
    const [data, setData] = useState([]);
    const tagOptions = [
        {
            param: 0,
            value: '#ffccc7',
            label: " "
        },{
            param: 1,
            value: '#fff1b8',
            label: " "
        },{
            param: 2,
            value: '#f4ffb8',
            label: " "
        },{
            param: 3,
            value: '#b5f5ec',
            label: " "
        },{
            param: 4,
            value: '#bae0ff',
            label: " "
        }
    ];

    const tableColumns: ColumnsType<FileDetail> = [
        {
            title: "文件名",
            dataIndex: 'filename'
        },{
            title: "大小",
            dataIndex: 'size'
        }, {
            title: "上传日期",
            dataIndex: 'uploadTime'
        },{
            title: "标签",
            dataIndex: 'labels',
            render: (_, {labels}) => (
                <>
                    {
                        labels.map(label => {
                            return (
                                <Tag key={label}>
                                    {label}
                                </Tag>
                            )
                        })
                    }
                </>

            )
        },{
            title: "标记",
            dataIndex: 'tags',
            render: (_, {tags}) => (
                <>
                    {
                        tags.map((tag) => {
                            let color = "null";
                            switch (tag) {
                                case '0': color = '#ffccc7';break;
                                case '1': color = '#fff1b8';break;
                                case '2': color = '#f4ffb8';break;
                                case '3': color = '#b5f5ec';break;
                                case '4': color = '#bae0ff';break;
                            }
                            return (
                                <>
                                    {color !== 'null' &&
                                        <Tag color={color} key={tag}>{" "}</Tag>
                                    }
                                </>
                            )
                        })
                    }
                </>
            )
        },{
            title: "操作",
            dataIndex: 'operation',
            render: (_, record) => (
                <Space size='middle'>
                    <Button shape='circle' icon={<EyeOutlined />}/>
                    <Button shape='circle' icon={<EditOutlined />}/>
                    <Button shape='circle' icon={<DeleteOutlined />}/>
                </Space>
            )
        }
    ]
    return (
        <div style={{margin: '20px'}}>
            {props.show &&
                <>
                    <Row>
                        <Col span={6}>
                            <Space direction='horizontal'>
                                <label>文件名：</label>
                                <Input placeholder="文件名"/>
                            </Space>
                        </Col>
                        <Col span={6}>
                            <Space direction='horizontal'>
                                <label>扩展名：</label>
                                <Select style={{width: '180px'}} placeholder="扩展名" options={extensionOptions}/>
                            </Space>
                        </Col>

                        <Col span={8}>
                            <Space direction='horizontal'>
                                <label>上传日期：</label>
                                <DatePicker.RangePicker />
                            </Space>

                        </Col>
                    </Row>
                    <Row style={{marginTop: '10px'}}>
                        <Col span={6}>
                            <Space direction='horizontal'>
                                <label>标签：&nbsp;&nbsp;&nbsp;</label>
                                <Select placeholder="标签"
                                        mode="multiple"
                                        allowClear
                                        style={{width: '180px'}}
                                        options={labelOptions}/>
                            </Space>
                        </Col>
                        <Col span={6}>
                            <Space direction='horizontal'>
                                <label>标记：&nbsp;&nbsp;&nbsp;</label>
                                <Select placeholder="标记"
                                        mode="multiple"
                                        allowClear
                                        tagRender={tagRender}
                                        style={{width: '180px'}}
                                        options={tagOptions}/>
                            </Space>
                        </Col>
                        <Col span={5}>
                            <Button type='default' icon={<SearchOutlined />} style={{float: 'right'}}>查询</Button>
                        </Col>
                        <Col span={4} style={{marginLeft: '25px'}}>
                            <Button type='primary' icon={<FileAddTwoTone />}>上传</Button>
                        </Col>
                    </Row>
                    <Row style={{marginTop: '10px'}} >
                        <Col span={24}>
                            <Table columns={tableColumns} dataSource={data}/>
                        </Col>
                    </Row>
                </>
            }
        </div>
    )
}

export default Files;