import React, { useEffect } from "react";
import { useRouter } from "next/router";
import {
  Controller,
  useForm,
  FormProvider,
  SubmitHandler,
  FieldErrors,
} from "react-hook-form";
import { IFormLogin } from "@/types/shared";
import { AnyAction, bindActionCreators, Dispatch } from "redux";
import { connect, useDispatch, useSelector } from "react-redux";
import { userAction } from "../../redux/usersSlice";
import { Form as AntdForm } from "antd";
import { Form, Input, Button, Alert } from 'antd';
import { LoginFormSchema } from "@/utils/validation";
import moment from "moment";

const getErrorsMessages = (errors: FieldErrors<FormData>) => {
  const errorValues = Object.values(errors);
  return errorValues
    .map((error) => error?.message)
    .filter((message) => message);
};

function LoginForm(props: any, context: any) {
  const dispatch = useDispatch();
  const router = useRouter();
  const methods = useForm<IFormLogin>({
    resolver: LoginFormSchema,
    defaultValues: {
      email: "",
      password: "",
    },
  });

  const onSubmit: SubmitHandler<IFormLogin> = async (data: IFormLogin) => {
    console.log(data, "====")
    await dispatch(userAction.loginRequest(data));
  };

  const onFinishFailed = (errorInfo: any) => {
    alert(errorInfo);
  };

  useEffect(() => {
    if (props.isLogin) {
      router.push("/Home");
    } else {
      // router.push("/");
    }
  });

  const errorsMessages = getErrorsMessages(methods.formState.errors);
  return (
    <AntdForm
      name="basic"
      layout="vertical"
      onFinish={methods.handleSubmit(onSubmit)}
      onFinishFailed={onFinishFailed}
      autoComplete="off"
    >
      {errorsMessages.map((error) => (
        <Alert
          className="ale"
          key={error}
          message={error}
          type="error"
          showIcon
        />
      ))}
      <Controller
        control={methods.control}
        name="email"
        render={({ field }) => (
          <Form.Item label="登录邮箱">
            <Input {...field} />
          </Form.Item>
        )}
      />

      <Controller
        control={methods.control}
        name="password"
        render={({ field }) => (
          <Form.Item label="登录密码">
            <Input type="password" {...field} />
          </Form.Item>
        )}
      />

      <Form.Item>
        <Button type="primary" htmlType="submit">
          登 录
        </Button>
      </Form.Item>
    </AntdForm>
  );
}
const mapStateToProps = (state: { users: any; isLogin: boolean }) => {
  return {
    isLogin: state.users.isLogin,
  };
};
function mapDispatchToProps(dispatch: Dispatch<AnyAction>) {
  return bindActionCreators({}, dispatch);
}

export default connect(mapStateToProps, mapDispatchToProps)(LoginForm);
