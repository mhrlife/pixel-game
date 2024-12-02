import {createAsyncThunk, createSlice} from '@reduxjs/toolkit';
import {useApi} from '../api/useApi.tsx';
import {EmptyHTTPAction, HTTPAction, HTTPError} from "./types.ts";

interface StatsState {
    onlineUsers: HTTPAction<number>;
}

const initialState: StatsState = {
    onlineUsers: EmptyHTTPAction<number>('IDLE'),
};

export const fetchOnlineUsersCount = createAsyncThunk<number, void>(
    'stats/fetchOnlineUsersCount',
    async () => {
        const api = useApi()
        return await api.getOnlineUsersCount()
    }
);

const statsSlice = createSlice({
    name: 'stats',
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder.addCase(fetchOnlineUsersCount.pending, (state) => {
            state.onlineUsers.state = 'LOADING';
        });
        builder.addCase(fetchOnlineUsersCount.fulfilled, (state, action) => {
            state.onlineUsers.state = 'SUCCESS';
            state.onlineUsers.value = action.payload;
        });
        builder.addCase(fetchOnlineUsersCount.rejected, (state, action) => {
            state.onlineUsers.state = 'ERROR';
            state.onlineUsers.error = action.error as HTTPError;
        });
    },
});

export default statsSlice.reducer;

