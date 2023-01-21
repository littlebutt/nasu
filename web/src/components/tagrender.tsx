import { Tag } from 'antd'
import type { CustomTagProps } from 'rc-select/lib/BaseSelect'
import React from "react";

const tagRender = (props: CustomTagProps) => {
  const { label, value, closable, onClose } = props
  const onPreventMouseDown = (event: React.MouseEvent<HTMLSpanElement>) => {
    event.preventDefault()
    event.stopPropagation()
  }
  return (
        <Tag
            color={value}
            onMouseDown={onPreventMouseDown}
            closable={closable}
            onClose={onClose}
            style={{ marginRight: 3, borderRadius: '100%' }}
        >
            {label}
        </Tag>
  )
}

export default tagRender

export const tagOptions = [
  {
    param: 0,
    value: '#ffccc7',
    label: '红'
  }, {
    param: 1,
    value: '#fff1b8',
    label: '黄'
  }, {
    param: 2,
    value: '#f4ffb8',
    label: '绿'
  }, {
    param: 3,
    value: '#b5f5ec',
    label: '蓝'
  }, {
    param: 4,
    value: '#bae0ff',
    label: '紫'
  }
]

export const toParam = (value: string) => {
  for (const tagOption of tagOptions) {
    if (tagOption.value === value) {
      return tagOption.param
    }
  }
  return ''
}
