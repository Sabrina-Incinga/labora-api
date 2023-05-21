import { FC, ReactNode } from 'react';

interface Props {
    children: ReactNode;
}

export const CardContent: FC<Props> = ({ children }:Props) => {
  return <div className="card-body">
            {children}
        </div>
};