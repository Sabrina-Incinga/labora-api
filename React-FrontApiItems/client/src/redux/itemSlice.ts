import {createAsyncThunk, createSlice} from '@reduxjs/toolkit'
import { item } from '../features/types/item.types'
import { getItems } from '../features/item/item.endpoints'

export const fetchItems = createAsyncThunk("/fetchItems", 
    async ({page, itemsPerPage}:{page: number, itemsPerPage: number}) => {
        const response = await getItems(page, itemsPerPage)

        return response
    }
)

interface initialType{
    items: item[]
    page: number
    itemsPerPage: number
    totalItems: number
    error: string | undefined
}

const initialState : initialType = {
    items: [],
    page: 1,
    itemsPerPage: 0,
    totalItems: 0,
    error: undefined
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
        }
    },
    extraReducers: builder => {
        builder.addCase(fetchItems.fulfilled, (state, action) => {
            state.items = action.payload.Items
            state.totalItems = action.payload.ItemCount
        })
        builder.addCase(fetchItems.rejected, (state, action) => {
            state.error = action.error.message
        })
    }
})

export const {pageAction, itemsPerPageAction} = itemSlice.actions
export default itemSlice.reducer