import { UserInfo } from "@/types/shared";
import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import moment from "moment";
import { persistReducer } from 'redux-persist';
import storage from 'redux-persist/lib/storage';

interface UserState {
    errors?: string[];
    username: string | null;
    auth: boolean;
    loading: boolean;
    userInfo: UserInfo | null;
    isLogin: boolean;
}

const initialState: UserState = {
    username: null,
    auth: false,
    loading: false,
    userInfo: null,
    isLogin: false
};

const userslice = createSlice({
    name: 'user',
    initialState,
    reducers: {
        loginRequest(state, action) {
            state.auth = action.payload;
        },
        logInSuccess: (state, { payload: { data } }) => {
            let session = global.localStorage;
            session.setItem('token', data.access_token)
            session.setItem('expire_at', moment().add(data.expires_at, 'seconds').format('X'));
            if (data.user_info && data.user_info.profile && data.user_info.profile.name) {
                session.setItem('userName', data.user_info.profile.name)
            }
            // localStorage.setItem()
            state.userInfo = data;
            state.isLogin = true

        },
        logInFail: (state) => {
            state.isLogin = false
        },
        loginOut: (state) => {
            if (global.localStorage !== undefined) {
                let session = global.localStorage;
                session.clear();
                state.isLogin = false;
                state.userInfo = null;
                state.username = null
            }
        }
    }
});

export const userAction = userslice.actions;
export default userslice.reducer;