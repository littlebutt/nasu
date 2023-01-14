import React from 'react';

import ReactEcharts from "echarts-for-react"



interface IChartCard {
    width: number,
    height: number,
    title: string,
    labels: string[],
    data: number[]
}

const ChartCard: React.FC<IChartCard> = (props) => {
    let _data = new Array(props.labels.length)
    for (let i = 0; i < props.labels.length; i++) {
        _data[i] = {
            value: props.data[i],
            name: props.labels[i]
        }
    }

    const options = {
        title : {
          text: props.title,
          left: 'left'
        },
        tooltip: {
            trigger: "item",
            formatter: "{a} <br/>{b}: {c} ({d}%)"
        },
        color: ['#f5222d', '#fadb14', '#a0d911', '#13c2c2', '#722ed1'],
        series: [
            {
                name: props.title,
                type: 'pie',
                stillShowZeroSum: false,
                data: _data
            }
        ]
    }
    return (
        <ReactEcharts option={options} style={{width: props.width + "px", height: props.height + "px"}}/>
    )
}

export default ChartCard;
