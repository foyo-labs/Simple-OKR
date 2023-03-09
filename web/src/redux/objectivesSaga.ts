import { select, takeEvery, put, call } from "@redux-saga/core/effects";
import { PayloadAction } from "@reduxjs/toolkit";
import { objectiveAction } from "./objectivesSlice";
import { BaseURL } from "./sagas";
import axios from "axios";
import getConfig from 'next/config'

const { apiUrl } = getConfig().publicRuntimeConfig

export function* loadObjectives() {
  yield takeEvery(objectiveAction.loadObjectivesRequest.type, function* (): any {
    const url = apiUrl + "/objectives"
    yield put(objectiveAction.loadObjectivesSuccess([]))
  });
}
