import {configureStore} from "@reduxjs/toolkit";
import userReducers from "./user.ts"
import statsReducers from "./stats.ts"
import {useDispatch, useSelector} from "react-redux";


export const store = configureStore({
    reducer: {
        user: userReducers,
        stats: statsReducers,
    }
})


export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch

export const useAppDispatch = useDispatch.withTypes<AppDispatch>()
export const useAppSelector = useSelector.withTypes<RootState>()