import { configureStore, Reducer, applyMiddleware, combineReducers } from '@reduxjs/toolkit';
import {
  persistStore,
  persistReducer,
} from "redux-persist";
import { createStore } from 'redux'
import createSagaMiddleware from "redux-saga";
import storage from 'redux-persist/lib/storage';
import thunk from "redux-thunk";
import usersReducer from './usersSlice'
import logger from 'redux-logger';
import { rootSaga } from './sagas';

const sagaMiddleware = createSagaMiddleware();

const reducers = combineReducers({
  users: usersReducer,
});

const persistConfig = {
  key: "root",
  storage,
};

const persistedReducer = persistReducer(persistConfig, reducers);
// 构建store 
// createStore (reducer,{},中间件)
export default createStore(persistedReducer, {}, applyMiddleware(sagaMiddleware))
export const store = configureStore({
  reducer: persistedReducer,
  middleware: [thunk, sagaMiddleware, logger],
});

export const persistor = persistStore(store);
sagaMiddleware.run(rootSaga);