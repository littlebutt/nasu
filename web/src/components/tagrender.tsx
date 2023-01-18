import {Tag} from 'antd';
import type { CustomTagProps } from 'rc-select/lib/BaseSelect';

const tagRender = (props: CustomTagProps) => {
    const { label, value, closable, onClose } = props;
    const onPreventMouseDown = (event: React.MouseEvent<HTMLSpanElement>) => {
        event.preventDefault();
        event.stopPropagation();
    };
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
    );
}

export default tagRender;