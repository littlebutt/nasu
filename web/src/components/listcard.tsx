import React, {useEffect} from 'react';
import {Button, Empty, List, Skeleton} from "antd";

type FileInfo = {
    id: number,
    filename: string,
    labels: string,
    uploadTime: string
}

interface IListCard {
    title: string,
    files: FileInfo[],
    loadMore: () => void,
    loading: boolean
}

const ListCard: React.FC<IListCard> = (props) => {

    const onLoadMore =
        !props.loading ? (
            <div style={{
                textAlign: 'center',
                marginTop: 12,
                height: 32,
                lineHeight: '32px',
            }}>
                <Button onClick={props.loadMore}>加载更多</Button>
            </div>
        ) : null

    return (
        <div style={{height: '250px'}}>
            <h2>{props.title}</h2>
            {props.files.length === 0 ? <Empty image={Empty.PRESENTED_IMAGE_SIMPLE} /> :
                <List
                    itemLayout='horizontal'
                    loadMore={onLoadMore}
                    dataSource={props.files}
                    renderItem={(item) => (
                        <List.Item>
                            <Skeleton avatar title={false} loading={props.loading} active>
                                <List.Item.Meta
                                    title={<a href="">{item.filename}</a>}
                                    description={item?.labels}
                                />
                                <div>{item.uploadTime.split('T')[0]}</div>
                            </Skeleton>
                        </List.Item>
                    )
                    }
                />
            }
        </div>
    )
}




export default ListCard;
export type {FileInfo};