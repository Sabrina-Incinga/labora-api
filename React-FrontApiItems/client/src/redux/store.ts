import {combineReducers, configureStore} from '@reduxjs/toolkit'
import itemReducer from './itemSlice'
import {reducer as toastrReducer} from 'react-redux-toastr'


const reducers = {
    items: itemReducer,
    toastr: toastrReducer 
  }

const reducer = combineReducers(reducers)

export const store = configureStore({reducer})

export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch
