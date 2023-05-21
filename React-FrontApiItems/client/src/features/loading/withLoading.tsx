import { FC } from 'react'
import {toastr} from 'react-redux-toastr'
import { Audio } from  'react-loader-spinner'

interface WrappedComponentProps {
  isLoading: boolean
  error: string
  errorText: string
}

export const withLoading = <T extends WrappedComponentProps>(WrappedComponent: FC<T>) => {
  return function (props: T) {
    if (props.error) {
        toastr.error('Ha ocurrido un error', props.error)
        return <div>
                    <p>{props.errorText}</p>
               </div>
    }

    if (props.isLoading) return <Audio
                                    height = "80"
                                    width = "80"
                                    color = 'green'
                                    ariaLabel = 'three-dots-loading'     
                                />

    return <WrappedComponent {...props} />
  }
}