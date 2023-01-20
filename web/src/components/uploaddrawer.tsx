import React, {useState} from 'react';
import {Button, Drawer, Form, Input, message, Select, Upload, UploadFile} from "antd";
import { UploadOutlined } from '@ant-design/icons';
import tagRender, {tagOptions, toParam} from "./tagrender";
import Axios from "../axios";
import {RcFile} from "antd/es/upload";
import moment from 'moment';

interface IUploadDrawer {
    show: boolean,
    setShow: (b: boolean) => void,
    labelOptions: any[],
    refresh: () => void
}

const UploadDrawer: React.FC<IUploadDrawer> = (props) => {
    const [form] = Form.useForm();
    const innerLabelOptions = props.labelOptions;
    const [fileList, setFileList] = useState<UploadFile[]>([]);
    const [uploading, setUploading] = useState(false);
    const [filename, setFilename] = useState("");
    const [labels, setLabels] = useState("");
    const [tags, setTags] = useState("");

    const handleChangeFilename = (e: { target: { value: React.SetStateAction<string>; }; }) => {
        setFilename(e.target.value);
    }

    const handleChangeLabels = (v: any, event: any) => {
        setLabels(v.join(':'));
    }

    const handleChangeTags = (v: string[], event: any) => {
        setTags(v.map(t => toParam(t)).join(':'));
    }

    const upload = (data: FormData) => Axios({
        method: 'POST',
        url: '/api/uploadFile',
        headers: { "Content-Type": "multipart/form-data" },
        data: data,
        timeout: 600000
    }).then(res => {
        console.log(res);
            if (res.status === 200 && res.data?.success == true) {
                message.success("上传成功");
            } else {
                message.error(res.data?.reason??"内部错误");
            }
        }).catch(err => {
            console.warn(err);
        }).finally(() => {
            setUploading(false);
            props.setShow(false);
            props.refresh();
    })

    const handleUpload = () => {
        setUploading(true);
        let data = new FormData();
        data.append('file', fileList[0] as RcFile);
        data.append('filename', filename);
        data.append('labels', labels);
        data.append('tags', tags);
        const now = moment().format('yyyy-MM-DD HH:mm:ss');
        data.append('uploadTime', now);
        upload(data);
    }

    return (
        <Drawer
            title="上传文件"
            width={720}
            open={props.show}
            onClose={() => props.setShow(false)}
            bodyStyle={{ paddingBottom: 80 }}
            extra={
            <Button type='primary' loading={uploading} onClick={handleUpload}>上传</Button>
            }>
            <Form form={form} name="basic" autoComplete='off'>
                <Form.Item
                    label="文件名"
                    name='filename'
                    rules={[{ required: true, message: '请输入文件名' }]}>
                    <Input onChange={handleChangeFilename}/>
                </Form.Item>
                <Form.Item
                    label="标签"
                    name='labels'>
                    <Select mode="tags"
                            style={{width: '600px', marginLeft: '25px'}}
                            options={innerLabelOptions}
                            onChange={handleChangeLabels}
                    />
                </Form.Item>
                <Form.Item
                    label="标记"
                    name='tags'>
                    <Select mode="multiple"
                            allowClear
                            tagRender={tagRender}
                            style={{width: '600px', marginLeft: '25px'}}
                            options={tagOptions}
                            onChange={handleChangeTags}
                    />
                </Form.Item>
                <Form.Item
                    label="文件"
                    name='file'
                    rules={[{ required: true, message: '请选择文件' }]}>
                    <Upload onRemove={(file) => {
                        const index = fileList.indexOf(file);
                        const newFileList = fileList.slice();
                        newFileList.splice(index, 1);
                        setFileList(newFileList);
                    }}
                            beforeUpload={(file) => {
                                setFileList([...fileList, file]);
                                return false;
                            }
                    }
                            fileList={fileList}
                            maxCount={1}
                    >
                        <Button icon={<UploadOutlined />}>选择文件</Button>
                    </Upload>
                </Form.Item>

            </Form>
        </Drawer>
    )
}

export default UploadDrawer;