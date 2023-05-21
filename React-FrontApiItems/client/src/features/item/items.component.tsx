import { FC, useEffect } from 'react'
import { ItemCard } from './itemCard'
import PaginationComponent from '../pagination/pagination.component'
import { useAppDispatch, useAppSelector } from '../../hooks/hooks'
import { item } from '../types/item.types'
import { deleteItem, fetchItems, itemsPerPageAction } from '../../redux/itemSlice'
import {toastr} from 'react-redux-toastr'
import { Audio } from  'react-loader-spinner'
import { useNavigate } from 'react-router-dom'

const ItemsComponent : FC = () => {
    const dispatch = useAppDispatch()
    const {items, page, itemsPerPage, error, loading , response} = useAppSelector(state => state.items)
    const navigate = useNavigate()

    const itemMapper = (item: item) => 
    {
        const onDeleteClicked = () => dispatch(deleteItem(item.id))
        const onUpdateClicked = () => navigate('/')

        return (<ItemCard key={item.id}>
                <ItemCard.Header name={`Nombre de cliente: ${item.name}`}></ItemCard.Header>
                <ItemCard.Content>
                    <ItemCard.Text text={`Nombre del Producto: ${item.product}`}></ItemCard.Text>
                    <ItemCard.Text text={`Cantidad adquirida: ${item.quantity}`}></ItemCard.Text>
                </ItemCard.Content>
                <div className={'row justify-content-around'}>
                    <ItemCard.Button text='Eliminar' className='btn-outline-danger col-4' onClickAction={onDeleteClicked}></ItemCard.Button>
                    <ItemCard.Button text='Editar' className='btn-outline-primary col-4' onClickAction={onUpdateClicked}></ItemCard.Button>
                </div>
        </ItemCard>)
    }
    
    useEffect(() => {
        dispatch(itemsPerPageAction(100))
    }, [dispatch])

    useEffect(() => {
        dispatch(fetchItems({page: page-1, itemsPerPage}))
    }, [page, itemsPerPage, dispatch])


    useEffect(() => {
        if(error){
            toastr.error('Ha ocurrido un error',error)
        }else if(response){
            toastr.success('Soliciud exitosa', response)
        }
    }, [error, response])

    if (loading){
       return <Audio
        height = "80"
        width = "80"
        color = 'green'
        ariaLabel = 'three-dots-loading'     
    />
    }

    if(error){
        toastr.error('Ha ocurrido un error',error)
        return <div>
                    <p>No se pudo cargar la información</p>
               </div>
    }

    return (
        <div className='container row gy-4 justify-content-center'>
        <PaginationComponent></PaginationComponent>
          {items?.map(itemMapper)}
          <PaginationComponent></PaginationComponent>
        </div>
      )
}

export default ItemsComponent