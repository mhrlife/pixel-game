import {createAsyncThunk, createSlice} from "@reduxjs/toolkit";
import {EmptyHTTPAction, HTTPAction} from "./types.ts";
import {useApi} from "../api/useApi.tsx";
import {HypeSerializer, UserWithToken} from "../types/serializer.ts";

export const loginUser = createAsyncThunk(
    'users/login',
    async () => {
        const api = useApi();
        return await api.login();
    }
)

export const fetchUserHype = createAsyncThunk(
    'users/fetchHype',
    async () => {
        const api = useApi();
        return await api.getHype();
    }
)

type userState = {
    auth: HTTPAction<UserWithToken>;
    hype: HTTPAction<HypeSerializer>;
}

const initialState: userState = {
    auth: EmptyHTTPAction<UserWithToken>('LOADING'),
    hype: EmptyHTTPAction<HypeSerializer>('IDLE')
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
            localStorage.setItem('pixel_jwt', action.payload.token);
        });
        builder.addCase(loginUser.rejected, (state, action) => {
            state.auth.state = 'ERROR'
            console.log("error is ", action.error);
        });
        builder.addCase(fetchUserHype.pending, (state) => {
            state.hype.state = 'LOADING'
        });
        builder.addCase(fetchUserHype.fulfilled, (state, action) => {
            state.hype.state = 'SUCCESS'
            state.hype.value = action.payload;
        });
        builder.addCase(fetchUserHype.rejected, (state, action) => {
            state.hype.state = 'ERROR'
            console.log("error is ", action.error);
        });
    }
})

export default userSlice.reducer
