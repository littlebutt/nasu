import React from 'react'

import ReactEcharts from 'echarts-for-react'
import { Empty } from 'antd'

interface IChartCard {
  width: number
  height: number
  title: string
  labels: string[]
  data: number[]
}

const ChartCard: React.FC<IChartCard> = (props) => {
  const _data = new Array(props.labels.length)
  for (let i = 0; i < props.labels.length; i++) {
    _data[i] = {
      value: props.data[i],
      name: props.labels[i]
    }
  }

  const options = {
    tooltip: {
      trigger: 'item',
      formatter: '{a} <br/>{b}: {c} ({d}%)'
    },
    color: ['#ffccc7', '#fff1b8', '#f4ffb8', '#b5f5ec', '#bae0ff'],
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
        <div style={{ height: '300px' }}>
            <h2>{props.title}</h2>
            {props.labels.length === 0
              ? <Empty image={Empty.PRESENTED_IMAGE_SIMPLE} />
              : <ReactEcharts option={options} style={{ width: props.width + 'px', height: props.height + 'px' }}/>}
        </div>

  )
}

export default ChartCard
