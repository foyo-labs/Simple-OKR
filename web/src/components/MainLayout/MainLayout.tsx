import { useState } from "react";
import React from "react";
import { Breadcrumb, Layout, Menu, theme, MenuProps, Row, Col } from "antd";
const { Header, Content } = Layout;
import { Logo, MainNavStyle, MainNavLinkStyle } from "./MainLayout.styles";
import Link from "next/link";

type MainLayoutProps = {
  children: React.ReactNode;
};

export default function MainLayout({ children }: MainLayoutProps) {
  const {
    token: { colorBgContainer },
  } = theme.useToken();
  const [collapsed, setCollapsed] = useState(false);
  return (
    <Layout>
      <Header className="header" style={MainNavStyle}>
        <a href="/home" style={MainNavLinkStyle}><img src="../../static/images/sokr-logo.png" /></a>
        <a style={MainNavLinkStyle} href="/objectives" title="OKR" >OKR</a>
        <a style={MainNavLinkStyle} href="/objectives" title="报告" >报告</a>
        <a style={MainNavLinkStyle} href="/objectives" title="设置" >设置</a>
      </Header>
      <Layout>
        {children}
      </Layout>
    </Layout>
  );
}
