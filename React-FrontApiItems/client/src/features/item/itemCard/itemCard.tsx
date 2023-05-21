import { FC, ReactNode } from 'react';

interface Props {
  children: ReactNode;
}

export const ItemCard: FC<Props> = ({ children }) => {
  return <div className={'card col-3 ms-2 border bg-light'}>{children}</div>;
};