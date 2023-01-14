import React, {useEffect, useState} from 'react';
import InfiniteScroll from 'react-infinite-scroll-component';
import {Divider, List, Skeleton} from "antd";

type FileInfo = {
    id: number,
    filename: string,
    labels: string[],
    tags: string[],
    size: string,
    uploadTime: string,
    extension: string
}

interface IListCard {
    title: string,
    filenames: string[],
    fetchLoadMore: (filename: string) => FileInfo
}

const ListCard: React.FC<IListCard> = (props) => {
    const [loading, setLoading] = useState(false);
    const [data, setData] = useState<FileInfo[]>([]);
    const [idx, setIdx] = useState(0);
    const loadMore = () => {
        let _data: FileInfo[] = [];
        for (let i = idx; i < Math.min(idx + 5, props.filenames.length); i ++) {
            _data.push(props.fetchLoadMore(props.filenames[i]));
        }
        setIdx(Math.min(idx + 5, props.filenames.length));
        return _data;
    }
    const more = () => {
        if (loading) {
            return;
        }
        setLoading(true);
        setData([...data, ...loadMore()]);
        setLoading(false);
    }

    useEffect(() => {
        more();
    }, []);

    return (
        <>
            <h1>{props.title}</h1>
        <div
            id="scrollableDiv"
            style={{
                height: 400,
                overflow: 'auto',
                padding: '0 16px',
                border: '1px solid rgba(140, 140, 140, 0.35)',
            }}
        >
            <InfiniteScroll
                dataLength={data.length}
                next={more}
                hasMore={data.length < 50}
                loader={<Skeleton avatar paragraph={{ rows: 1 }} active />}
                endMessage={<Divider plain>到此为止了 🤐</Divider>}
                scrollableTarget="scrollableDiv"
            >
                <List
                    dataSource={data}
                    renderItem={(item) => (
                        <List.Item key={item.id}>
                            <List.Item.Meta
                                title={<a >{item.filename}</a>}
                                description={item.labels.join(',')}
                            />
                            <div>上传时间：{item.uploadTime}</div>
                        </List.Item>
                    )}
                />
            </InfiniteScroll>
        </div>
        </>
    )
}

export default ListCard