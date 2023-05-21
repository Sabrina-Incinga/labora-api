import { FC } from "react"
import { useAppDispatch } from "../../../../hooks/hooks"
import {  itemDTO } from '../features/types/item.types'
import { updateItem } from "../../item.endpoints"

interface WrappedComponentProps  {
    text: string
    onClickAction: ({id, payload}:{id:number, payload?:itemDTO}) => void
    className: string
}

const buttonHoc = <T extends WrappedComponentProps >(WrappedComponent : FC<T>) => {
    const Wrapper = (props: T) => {
        const clickHandler = () => {
            props.dispatch(props.onClickAction(props.id, props.payload?))
        }

        return <WrappedComponent {...props} />;
    }
}