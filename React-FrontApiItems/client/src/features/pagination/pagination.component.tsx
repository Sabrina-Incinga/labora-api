import { FC } from 'react'
import { useAppDispatch, useAppSelector } from '../../hooks/hooks'
import { pageAction } from '../../redux/itemSlice'

const PaginationComponent : FC = () => {
    const dispatch = useAppDispatch()
    const {page, totalItems, itemsPerPage} = useAppSelector(state => state.items)

    const onNextClicked = () => {
        dispatch(pageAction(1))
    }

    const onPrevClicked = () => {
        dispatch(pageAction(-1))
    }

    return(
        <div className='row col-12 justify-content-around gy-4'>
            <button className='col-3 btn btn-outline-primary' type='button' onClick={onPrevClicked} disabled={page === 1}>Prev</button>
            <button className='col-3 btn btn-outline-primary' type='button' onClick={onNextClicked} disabled={totalItems < page*itemsPerPage}>Next</button>
        </div>
    )
}

export default PaginationComponent