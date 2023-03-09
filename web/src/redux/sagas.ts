import { all, takeEvery, put } from "@redux-saga/core/effects";
import { PayloadAction } from "@reduxjs/toolkit";
import { watchLogin, loginOut } from "./usersSaga"

export const BaseURL = process.env.PUBLIC_API;
export function* rootSaga() {
    yield all([
        watchLogin(), loginOut()
    ]);
}