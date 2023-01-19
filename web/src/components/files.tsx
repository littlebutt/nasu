import React, {useEffect, useState} from 'react';
import {Button, Col, DatePicker, Input, Row, Select, Space, Table, Tag} from "antd";
import { SearchOutlined, FileAddTwoTone, EyeOutlined, EditOutlined, DeleteOutlined } from '@ant-design/icons';
import tagRender from "./tagrender";
import {ColumnsType} from "antd/es/table";
import Axios from "../axios";


interface FileDetail {
    filename: string,
    extension: string,
    labels: string[],
    tags: string[],
    uploadTime: string,
    size: string
}

interface IOption {
    key: string,
    label: string,
    value: string
}

interface IFiles {
    show: boolean
}

const Files: React.FC<IFiles> = (props) => {
    const [extensionOptions, setExtensionOptions] = useState<Array<IOption>>([]);
    const [labelOptions, setLabelOptions] = useState<Array<IOption>>([]);
    const [data, setData] = useState<FileDetail[]>([]);
    const [filename, setFilename] = useState("");
    const [extension, setExtension] = useState("");
    const [labels, setLabels] = useState("");
    const [tags, setTags] = useState("")
    const tagOptions = [
        {
            param: 0,
            value: '#ffccc7',
            label: "红"
        },{
            param: 1,
            value: '#fff1b8',
            label: "黄"
        },{
            param: 2,
            value: '#f4ffb8',
            label: "绿"
        },{
            param: 3,
            value: '#b5f5ec',
            label: "蓝"
        },{
            param: 4,
            value: '#bae0ff',
            label: "紫"
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
    ];

    const handleChangeFilename = (e: { target: { value: React.SetStateAction<string>; }; }) => {
        setFilename(e.target.value);
    }

    const handleChangeExtension = (v: any, event: any) => {
        setExtension(v);
    }

    const handleChangeLabels = (v: any, event: any) => {
        setLabels(v.join(':'));
    }

    const handleChangeTags = (v: any, event: any) => {
        setTags(v.join(':'));
    }

    const handleOverallExtensionInfo = () => Axios({
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
            }));
        }
    }).catch(err => {
            console.warn(err);
        })

    const handleOverallLabelInfo = () => Axios({
        method: 'GET',
        url: '/api/overallLabelInfo',
    }).then(res => {
        if (res.status === 200) {
            setLabelOptions(Object.keys(res.data).map(k => {
                return {
                    key: k,
                    label: k,
                    value: k
                }
            }));
        }
    }).catch(err => {
        console.warn(err);
    })

    const handleListFilesByCondition = (payload: {
        filename: string,
        extension: string,
        labels: string,
        tags: string,
        startTime: string,
        endTime: string
    } | void) => Axios({
        method: 'GET',
        url: '/api/listFilesByCondition',
        params: {
            filename: payload?.filename??'',
            extension: payload?.extension??'',
            labels: payload?.labels??'',
            tags: payload?.tags??'',
            startTime: payload?.startTime??'',
            endTime: payload?.endTime??''
        }
    }).then(res => {
        if (res.status === 200) {
            console.log(res.data)
            setData(res.data?.nasuFiles.map((f: { Filename: any; Extension: any; Labels: string; Tags: string; UploadTime: any; Size: any; }) => {
                return {
                    filename: f.Filename,
                    extension: f.Extension,
                    labels: f.Labels.split(','),
                    tags: f.Tags.split(','),
                    uploadTime: f.UploadTime.split('T')[0],
                    size: f.Size
                }
            }));
        }
    }).catch(err => {
        console.warn(err);
    })

    useEffect(() => {
        handleOverallLabelInfo();
        handleOverallExtensionInfo();
        handleListFilesByCondition();
    }, []);
    return (
        <div style={{margin: '20px'}}>
            {props.show &&
                <>
                    <Row>
                        <Col span={8}>
                            <Space direction='horizontal'>
                                <label>文件名：</label>
                                <Input placeholder="文件名" onChange={handleChangeFilename} style={{width: '250px'}}/>
                            </Space>
                        </Col>
                        <Col span={8}>
                            <Space direction='horizontal'>
                                <label>扩展名：</label>
                                <Select style={{width: '250px'}} placeholder="扩展名" options={extensionOptions} onChange={handleChangeExtension}/>
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
                        <Col span={8}>
                            <Space direction='horizontal'>
                                <label>标签：&nbsp;&nbsp;&nbsp;</label>
                                <Select placeholder="标签"
                                        mode="multiple"
                                        allowClear
                                        style={{width: '250px'}}
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
                                        style={{width: '250px'}}
                                        options={tagOptions}
                                        onChange={handleChangeTags}
                                />
                            </Space>
                        </Col>
                        <Col span={5}>
                            <Button type='default' icon={<SearchOutlined />} style={{float: 'right'}}>查询</Button>
                        </Col>
                        <Col span={2} style={{marginLeft: '25px'}}>
                            <Button type='primary' icon={<FileAddTwoTone />}>上传</Button>
                        </Col>
                    </Row>
                    <Row style={{marginTop: '10px'}} >
                        <Col span={24}>
                            <Table columns={tableColumns} dataSource={data} style={{width: '1190px'}}/>
                        </Col>
                    </Row>
                </>
            }
        </div>
    )
}

export default Files;