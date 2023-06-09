import {createAsyncThunk, createSlice} from '@reduxjs/toolkit'
import { item, itemDTO } from '../features/types/item.types'
import { deleteItemRequest, getItems, updateItemRequest } from '../features/item/item.endpoints'

export const fetchItems = createAsyncThunk("/fetchItems", 
    async ({page, itemsPerPage}:{page: number, itemsPerPage: number}) => {
        const res = await getItems(page, itemsPerPage)

        return res
    }
)
export const deleteItem:any = createAsyncThunk("/deleteItem/{id}", 
    async (id:number) => {
        const res = await deleteItemRequest(id)
        
        return {res, id}
    }
)
export const updateItem:any = createAsyncThunk("/updateItem/{id}", 
    async ({id, payload}:{id: number, payload: itemDTO}) => {
        const res = await updateItemRequest(id, payload)

        return {res, id}
    }
)

export interface initialType{
    items: item[]
    page: number
    itemsPerPage: number
    totalItems: number
    error: string | undefined
    loading: boolean
    response: string | undefined
}

const initialState : initialType = {
    items: [],
    page: 1,
    itemsPerPage: 0,
    totalItems: 0,
    error: undefined,
    loading: false,
    response: undefined
}


export const itemSlice = createSlice({
    name: 'items',
    initialState,
    reducers: {
        pageAction: (state, action) => {
            state.page += action.payload
        },
        itemsPerPageAction: (state, action) => {
            state.itemsPerPage = action.payload
        },
    },
    extraReducers: builder => {
        builder.addCase(fetchItems.fulfilled, (state, action) => {
            state.loading = false
            state.items = action.payload.Items
            state.totalItems = action.payload.ItemCount
        })
        builder.addCase(fetchItems.pending, (state) => {
            state.loading = true
        })
        builder.addCase(fetchItems.rejected, (state, action) => {
            state.loading = false
            state.error = action.error.message
        })
        builder.addCase(deleteItem.fulfilled, (state, action) => {
            state.items = state.items.filter(item => item.id != action.payload.id)
            state.loading = false
            state.response = action.payload.res
        })
        builder.addCase(deleteItem.rejected, (state, action) => {
            state.loading = false
            state.error = action.error.message
        })
        builder.addCase(updateItem.fulfilled, (state, action) => {
            state.loading = false
            state.response = action.payload.res
        })
        builder.addCase(updateItem.pending, (state) => {
            state.loading = true
        })
        builder.addCase(updateItem.rejected, (state, action) => {
            state.loading = false
            state.error = action.error.message
        })
    }
})

export const {pageAction, itemsPerPageAction} = itemSlice.actions
export default itemSlice.reducer