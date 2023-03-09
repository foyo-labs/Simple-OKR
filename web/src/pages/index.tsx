import Head from 'next/head'
import Image from 'next/image'
import React, { ReactElement } from "react";
import { Inter } from '@next/font/google'
import styles from '@/styles/main.module.css';
import { useRouter } from "next/router";
import MainLayout from '@/components/MainLayout/MainLayout';
import Login from './login/index'
import Link from 'next/link';
const inter = Inter({ subsets: ['latin'] })

export default function Home() {
  const router = useRouter();
  return (
    <>
      <main className={styles.main}>
        <Link href={'/login'}>Go Login.</Link>
      </main>
    </>
  )
}

// 指定Layout.
Home.getLayout = (page: ReactElement) => <div>{page}</div>;
