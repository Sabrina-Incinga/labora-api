import { FC, useEffect } from 'react'
import { ItemCard } from './itemCard'
import PaginationComponent from '../pagination/pagination.component'
import { useAppDispatch, useAppSelector } from '../../hooks/hooks'
import { item } from '../types/item.types'
import { fetchItems, itemsPerPageAction } from '../../redux/itemSlice'

const ItemsComponent : FC = () => {
    const dispatch = useAppDispatch()
    const {items, page, itemsPerPage, error} = useAppSelector(state => state.items)

    const itemMapper = (item: item) => (
        <ItemCard key={item.id}>
                <ItemCard.Header name={`Nombre de cliente: ${item.name}`}></ItemCard.Header>
                <ItemCard.Content>
                    <ItemCard.Text text={`Nombre del Producto: ${item.product}`}></ItemCard.Text>
                    <ItemCard.Text text={`Cantidad adquirida: ${item.quantity}`}></ItemCard.Text>
                </ItemCard.Content>
        </ItemCard>
    )

    useEffect(() => {
        dispatch(fetchItems({page: page-1, itemsPerPage}))
    }, [page, itemsPerPage, dispatch])

    useEffect(() => {
        dispatch(itemsPerPageAction(100))
    }, [dispatch])

    if(error){
        return <div>
                    <p>No se pudo cargar la información</p>
               </div>
    }

    return (
        <div className='container row gy-4 justify-content-center'>
          {items?.map(itemMapper)}
          <PaginationComponent></PaginationComponent>
        </div>
      )
}

export default ItemsComponent