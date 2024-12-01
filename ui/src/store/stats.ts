import {createAsyncThunk, createSlice} from '@reduxjs/toolkit';
import {useApi} from '../api/useApi.tsx';

interface StatsState {
    onlineUsers: {
        state: 'IDLE' | 'LOADING' | 'SUCCESS' | 'ERROR';
        value: number | null;
        error: string | null;
    };
}

const initialState: StatsState = {
    onlineUsers: {
        state: 'IDLE',
        value: null,
        error: null,
    },
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
            state.onlineUsers.error = null;
        });
        builder.addCase(fetchOnlineUsersCount.fulfilled, (state, action) => {
            state.onlineUsers.state = 'SUCCESS';
            state.onlineUsers.value = action.payload;
        });
        builder.addCase(fetchOnlineUsersCount.rejected, (state, action) => {
            state.onlineUsers.state = 'ERROR';
            state.onlineUsers.error = action.payload as string;
        });
    },
});

export default statsSlice.reducer;

