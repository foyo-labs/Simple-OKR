import React from "react";
import { useState, ReactElement } from "react";
import styles from '@/styles/main.module.css';
import { Card, Form, Input, Button } from 'antd';
import { useRouter } from "next/router";
import { useDispatch, useSelector } from "react-redux";
import { Row, Col } from "antd";
import LoginForm from "../../components/Login/login-form";
import styled from "styled-components";
import {
  LoginWrapper,
} from "./index.style";
import Image from "rc-image";
export default function LoginPage() {
  return (
    <LoginWrapper>
      <Card title={'登录到Simple-OKR'} bordered={false} style={{ width: 500 }}>
        <LoginForm />
      </Card>
    </LoginWrapper>
  );
}

LoginPage.getLayout = (page: ReactElement) => <div className={styles.main}>{page}</div>;