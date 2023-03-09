import { select, takeEvery, put, call } from "@redux-saga/core/effects";
import { PayloadAction } from "@reduxjs/toolkit";
import { userAction } from "./usersSlice";
import { BaseURL } from "./sagas";
import axios from "axios";
import getConfig from 'next/config'

const { apiUrl } = getConfig().publicRuntimeConfig

export function* logIn(action: any) {
  // yield put(userAction.logInSuccess(null));
}

export function* watchLogin() {
  yield takeEvery(userAction.loginRequest.type, function* (): any {
    const url = apiUrl + "/users/login"
    try {
      const userInfo = yield select(state => state.users.auth);
      const res = yield call(axios.post, url, {
        ...userInfo
      });
      if (res && res.status && res.status == 200) {
        yield put(userAction.logInSuccess(res))
      } else {
        yield put(userAction.logInFail())
      }
    } catch (errors) {
    }
  });
}

export function* loginOut() {
  yield put(userAction.loginOut())
}