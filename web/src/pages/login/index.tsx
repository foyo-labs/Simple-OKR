import React from "react";
import { useState, ReactElement } from "react";
import styles from '@/styles/main.module.css';
import { useRouter } from "next/router";
import { useDispatch, useSelector } from "react-redux";
import { Row, Col } from "antd";
import LoginForm from "../../components/Login/login-form";
import LeftPage from "../../components/Login/LeftPage";
import styled from "styled-components";
import {
  LoginWrapper,
} from "./index.style";
export default function LoginPage() {
  return (
    <LoginWrapper>
      Login Form....
    </LoginWrapper>
  );
}

LoginPage.getLayout = (page: ReactElement) => <div className={styles.main}>{page}</div>;