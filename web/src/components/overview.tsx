import React, {useEffect, useState} from 'react';
import {Col, Empty, Row, Space} from 'antd';
import ChartCard from "./chartcard";
import ListCard, {FileInfo} from "./listcard";
import Axios from "../axios";
import {useNavigate} from "react-router-dom";


interface IOverview {
    show: boolean
}

const Overview: React.FC<IOverview> = (props) => {
    const [labelLabels, setLabelLabels] = useState<Array<string>>([]);
    const [labelData, setLabelData] = useState<Array<number>>([]);
    const [tagLabels, setTagLabels] = useState<Array<string>>([]);
    const [tagData, setTagData] = useState<Array<number>>([]);
    const [extensionLabels, setExtensionLabels] = useState<Array<string>>([]);
    const [extensionData, setExtensionData] = useState<Array<number>>([]);
    const [filenames, setFilenames] = useState<string[]>([]);
    const [fileIdx, setFileIdx] = useState<number>(0);
    const [files, setFiles] = useState<Array<FileInfo>>([]);
    const [fileLoading, setFileLoading] = useState<boolean>(false);
    const [dataFlag, setDataFlag] = useState<boolean>(false);
    const navigate = useNavigate();

    const handleOverallLabelInfo = () => Axios({
        method: 'GET',
        url: '/api/overallLabelInfo',
        }).then(res => {
            if (res.status === 200) {
                setLabelLabels(Object.keys(res.data));
                setLabelData(Object.values(res.data));
            }
        }).catch(err => {
            console.warn(err);
    });
    const handleOverallTagInfo = () => Axios({
        method: 'GET',
        url: '/api/overallTagInfo'
    }).then(res => {
        if (res.status === 200) {
            setTagLabels(Object.keys(res.data));
            setTagData(Object.values(res.data));
        }
    }).catch(err => {
        console.warn(err);
    });
    const handleOverallExtensionInfo = () => Axios({
        method: 'GET',
        url: '/api/overallExtensionInfo'
    }).then(res => {
        if (res.status === 200) {
            setExtensionLabels(Object.keys(res.data));
            setExtensionData(Object.values(res.data));
        }
    }).catch(err => {
        console.warn(err);
    });
    const buildFilenames = () => Axios({
        method: 'GET',
        url: '/api/overallFileInfo'
    }).then(res => {
        if (res.status === 200) {
            setFilenames(res.data?.filename);
            setDataFlag(true);
        }
    }).catch(err => {
        console.warn(err);
    });
    const handleLoadMore = () => {
        setFileLoading(true);
        for (let i = fileIdx;i < Math.min(filenames.length, fileIdx + 5); i ++) {
            Axios({
                method: 'GET',
                url: '/api/listFilesByCondition',
                params: {
                    filename: filenames[i]
                }
            }).then(res => {
                if (res.status === 200) {
                    let _f = res.data?.nasuFiles[0];
                    let file: FileInfo = {
                        filename: _f?.Filename,
                        id: _f?.Id,
                        labels: _f?.Labels,
                        uploadTime: _f?.UploadTime
                    };
                    setFiles(files => [...files, file]);
                }
            }).catch(err => {
                console.warn(err);
            });

        }
        setFileLoading(false);
        setFileIdx(fileIdx => Math.min(filenames.length, fileIdx + 5));
    }


    useEffect(() => {
        handleOverallLabelInfo();
        handleOverallTagInfo();
        handleOverallExtensionInfo();
        buildFilenames();
        handleLoadMore();
    }, [dataFlag]);
    return (
        <div style={{margin: '20px'}}>
        {props.show &&
        <>
            <Row>
                <Col span={8}>
                        <ChartCard width={300} height={250} title={"标签数据"} labels={labelLabels} data={labelData}/>
                </Col>
                <Col span={8}>

                        <ChartCard width={300} height={250} title={"标记数据"} labels={tagLabels} data={tagData}/>
                </Col>
                <Col span={8}>
                        <ChartCard width={300} height={250} title={"类型数据"} labels={extensionLabels} data={extensionData}/>
                </Col>
            </Row>
            <Row>
                <Col span={12}>
                    <ListCard title="最近上传" files={files} loadMore={handleLoadMore} loading={fileLoading}/>
                </Col>
                <Col span={12}>
                    <>
                        <h2>服务器状态</h2>
                        <Empty description='构建中' image={Empty.PRESENTED_IMAGE_SIMPLE}/>
                    </>
                </Col>
            </Row>
        </>}
        </div>
    )
}

export default Overview;