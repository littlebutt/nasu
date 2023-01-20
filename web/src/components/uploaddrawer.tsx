import React, {useState} from 'react';
import {Button, Drawer, Form, Input, Select} from "antd";

interface IUploadDrawer {
    show: boolean,
    setShow: (b: boolean) => void,
    labelOptions: any[]
}

const UploadDrawer: React.FC<IUploadDrawer> = (props) => {
    const [form] = Form.useForm();
    const [innerLabelOptions, setInnerLabelOptions] = useState(props.labelOptions);
    return (
        <Drawer
            title="上传文件"
            width={720}
            open={props.show}
            bodyStyle={{ paddingBottom: 80 }}
            extra={
            <Button type='primary'>上传</Button>
            }>
            <Form form={form} name="basic" autoComplete='off'>
                <Form.Item
                    label="文件名"
                    name='filename'
                    rules={[{ required: true, message: '请输入文件名' }]}>
                    <Input />
                </Form.Item>
                <Form.Item
                    label="标签"
                    name='labels'>
                    <Select mode="tags"
                            options={innerLabelOptions}/>
                </Form.Item>
                <Form.Item/>

            </Form>
        </Drawer>
    )
}

export default UploadDrawer;