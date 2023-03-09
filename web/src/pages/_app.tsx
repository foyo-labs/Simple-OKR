import "@/styles/globals.css";
import type { AppProps } from "next/app";
import { NextPage } from "next";
import { Provider } from "react-redux";
import { persistStore, persistReducer } from "redux-persist";
import { persistor, store } from "../redux/store";
import { PersistGate } from "redux-persist/lib/integration/react";
import { NextPageWithLayout } from "@/types/shared";
import { useEffect } from "react";

type AppPropsWithLayout = AppProps & {
  Component: NextPageWithLayout;
};

function App({ Component, pageProps }: AppPropsWithLayout) {
  const getLayout = Component.getLayout ?? ((page: NextPage) => page);
  return (
    <Provider store={store}>
      <PersistGate loading={null} persistor={persistor}>
        {getLayout(<Component {...pageProps} />)}
      </PersistGate>
    </Provider>
  );
}

export default App;
