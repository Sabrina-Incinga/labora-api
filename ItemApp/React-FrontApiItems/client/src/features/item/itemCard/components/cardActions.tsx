import { FC } from 'react'

interface Props {
    text: string
    onClickAction: () => void
    className: string
}

export const CardActions: FC<Props> = ({ text, onClickAction, className }:Props) => {
  return <button type='button' onClick={onClickAction} className={`btn ${className}`}>{text}</button>
};