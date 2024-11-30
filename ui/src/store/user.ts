import {createAsyncThunk, createSlice} from "@reduxjs/toolkit";
import {EmptyHTTPAction, HTTPAction} from "./types.ts";
import {useAPI} from "../api/useAPI.tsx";
import {UserWithToken} from "../types/serializer.ts";

export const loginUser = createAsyncThunk(
    'users/login',
    async () => {
        const api = useAPI();
        return await api.login();
    }
)

type userState = {
    auth: HTTPAction<UserWithToken>;
}

const initialState: userState = {
    auth: EmptyHTTPAction<UserWithToken>('LOADING')
}

export const userSlice = createSlice({
    name: 'user',
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder.addCase(loginUser.pending, (state) => {
            state.auth.state = 'LOADING'
        });
        builder.addCase(loginUser.fulfilled, (state, action) => {
            state.auth.state = 'SUCCESS'
            state.auth.value = action.payload;
        });
        builder.addCase(loginUser.rejected, (state, action) => {
            state.auth.state = 'ERROR'
            console.log("error is ", action.error);
        });
    }
})


export default userSlice.reducer