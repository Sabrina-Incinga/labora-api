import { FC } from 'react';

interface Props {
    text: string;
}

export const CardText: FC<Props> = ({ text }) => {
  return <div className='card-text'>{text}</div>
};