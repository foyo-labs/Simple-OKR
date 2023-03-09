import { useState } from "react";
import React from "react";
import { Breadcrumb, Layout, Menu, theme, MenuProps, Row, Col } from "antd";
import SiderContainer from "../Sider/index";
const { Header, Content } = Layout;
import { Logo, headerStyle, MenuStyle, LayoutStyle } from "./MainLayout.styles";
type MainLayoutProps = {
  children: React.ReactNode;
};
const items1: MenuProps["items"] = [
  "销售订单",
  "采购订单",
  "库存管理",
  "财务管理",
].map((key) => ({
  key,
  label: `${key}`,
}));
export default function MainLayout({ children }: MainLayoutProps) {
  const {
    token: { colorBgContainer },
  } = theme.useToken();
  const [collapsed, setCollapsed] = useState(false);
  return (
    <Layout>
      <Header className="header" style={headerStyle}>
        <Logo />
        <Menu
          theme="dark"
          mode="horizontal"
          defaultSelectedKeys={["2"]}
          items={items1}
          style={MenuStyle}
        />
      </Header>
      <Layout style={LayoutStyle}>
        <Row style={{ display: "flex", flexDirection: "row" }}>
          <Col
            style={{
              flex: "0 0 200px",
              width: 200,

              overflow: "hidden",
            }}
          >
            <SiderContainer />
          </Col>
          <Col
            style={{ flex: "auto", overflow: "auto" }}
            className="main-content"
          >
            {children}
          </Col>
        </Row>

        {/* 左侧菜单 */}

        {/* 上方菜单 */}
      </Layout>
    </Layout>
  );
}
