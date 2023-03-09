import Head from "next/head";
import Image from "next/image";
import React, { ReactElement, useEffect } from "react";
import { Inter } from "@next/font/google";
import styles from "@/styles/Home.module.css";
import { Breadcrumb, Button, DatePicker, Layout } from "antd";
import { useRouter } from "next/router";
import MainLayout from "@/components/MainLayout/MainLayout";
import { useDispatch } from "react-redux";
import { objectiveAction } from "../../redux/objectivesSlice";

const inter = Inter({ subsets: ["latin"] });

export default function Home() {
  const router = useRouter();
  const dispatch = useDispatch();
  useEffect(() => {
    dispatch(objectiveAction.loadObjectivesRequest({}));
  });
  return (
    <>
      <header>header...</header>
    </>
  );
}

// 指定Layout.
Home.getLayout = (page: ReactElement) => <MainLayout>{page}</MainLayout>;
