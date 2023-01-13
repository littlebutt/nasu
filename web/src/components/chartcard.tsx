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
    const _data = new Array(props.data.length);

    const options = {
        series: [
            {
                type: 'pie',

            }
        ]
    }
    return (
        <ReactEcharts option={options}/>
    )
}
