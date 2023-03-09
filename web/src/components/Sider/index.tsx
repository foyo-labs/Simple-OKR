import React from "react";
import { Breadcrumb, Menu, MenuProps } from "antd";
import { useRouter } from "next/router";
import { Container, SubMenuStyle } from "./index.style";
import menus from "./index.json";
import {
  LaptopOutlined,
  NotificationOutlined,
  UserOutlined,
  SisternodeOutlined,
} from "@ant-design/icons";
export default function SiderContainer() {
  const router = useRouter();

  console.log(menus, "menus");
  const items2: MenuProps["items"] = [
    UserOutlined,
    LaptopOutlined,
    NotificationOutlined,
  ].map((icon, index) => {
    const key = String(index + 1);

    return {
      key: `sub${key}`,
      icon: React.createElement(icon),
      label: `subnav ${key}`,

      children: new Array(4).fill(null).map((_, j) => {
        const subKey = index * 4 + j + 1;
        return {
          key: subKey,
          label: `option${subKey}`,
        };
      }),
    };
  });
  return (
    <Container>
      <Menu
        mode="inline"
        defaultSelectedKeys={["1"]}
        defaultOpenKeys={["sub1"]}
        style={{
          height: "100%",
          borderRight: 0,
          backgroundColor: "transparent",
        }}
        // items={items2}
      >
        {menus
          ? menus.map((m, index) => {
              console.log(m, "mm");
              return m.children && m.children.length > 0 ? (
                <Menu.SubMenu
                  key={m.key}
                  style={SubMenuStyle}
                  title={
                    <span>
                      <SisternodeOutlined />
                      <span>{m.name}</span>
                    </span>
                  }
                >
                  {m.children.map((m) => {
                    return (
                      <Menu.Item key={m.key} style={SubMenuStyle}>
                        <a>{m.name}</a>
                      </Menu.Item>
                    );
                  })}
                </Menu.SubMenu>
              ) : (
                <Menu.Item key={m.key} style={SubMenuStyle}>
                  <a>{m.name}</a>
                </Menu.Item>
              );
            })
          : ""}
      </Menu>
    </Container>
  );
}
