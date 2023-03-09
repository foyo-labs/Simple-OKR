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

    const url = "http://192.168.0.134:10088/api/v1/user/login"
    try {
      const userInfo = yield select(state => state.users.auth);
      const res = yield call(axios.post, url, {
        ...userInfo
      });
      if (res && res.status && res.status == 200) {
        // 登录成功
        yield put(userAction.logInSuccess(res))
      } else {
        // 登录失败
        yield put(userAction.logInFail())
      }
    } catch (errors) {
    }
  });
}

export function* loginOut() {
  yield put(userAction.loginOut())
}