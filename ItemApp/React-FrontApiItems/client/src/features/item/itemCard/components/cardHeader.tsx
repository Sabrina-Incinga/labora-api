import { FC } from 'react';

interface Props {
  name: string;
}

export const CardHeader: FC<Props> = ({ name }) => {
  return <h2 className='card-title'>{name}</h2>
};