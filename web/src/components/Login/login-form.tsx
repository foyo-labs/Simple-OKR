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
// redux\usersSaga.ts
import { Alert, Button, Form as AntdForm, Input } from "antd";
import { LoginFormSchema } from "@/utils/validation";
import { Container, Test, Title, Welcome } from "./login-form.style";
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
      userName: "",
      password: "",
    },
  });

  const onSubmit: SubmitHandler<IFormLogin> = async (data: IFormLogin) => {
    await dispatch(userAction.loginRequest(data));
  };

  //判断是否在登录状态
  const checkInLogin = () => {
    let session = global.localStorage;
    let expire_at = session.getItem("expire_at");
    if (
      session !== undefined &&
      session.getItem("token") !== undefined &&
      session.getItem("userName") !== undefined
    ) {
      if (expire_at) {
        if (moment().format("X") <= expire_at) {
          return true;
        } else {
          return false;
        }
      }
    } else {
      return false;
    }
  };

  useEffect(() => {
    if (props.isLogin) {
      router.push("/Home");
    } else {
      // router.push("/");
    }
  });

  const onFinishFailed = (errorInfo: any) => {
    alert(errorInfo);
  };

  const errorsMessages = getErrorsMessages(methods.formState.errors);
  return (
    <Container>
      <Title>登录</Title>
      <Welcome>您好！欢迎来到胜利达对账管理系统</Welcome>
      {/* <FormProvider {...methods}> */}
      <AntdForm
        name="basic"
        onFinish={methods.handleSubmit(onSubmit)}
        onFinishFailed={onFinishFailed}
        autoComplete="off"
      >
        <Test>
          <Controller
            control={methods.control}
            name="userName"
            render={({ field }) => (
              <Input className="ipt style-margin" {...field} />
            )}
          />
        </Test>
        <Test>
          <Controller
            control={methods.control}
            name="password"
            render={({ field }) => (
              <Input.Password
                className="ipt"
                {...field}
                onChange={(e) => {
                  field.onChange(e.target.value);
                }}
              />
            )}
          />{" "}
          <div className="pw">记住密码</div>
        </Test>

        {/* <Btn  type="submit">登录</Btn> */}
        <Test>
          {errorsMessages.map((error) => (
            <Alert
              className="ale"
              key={error}
              message={error}
              type="error"
              showIcon
            />
          ))}
        </Test>
        <Test>
          <Button htmlType="submit" type="primary" block className="btn">
            登录
          </Button>
        </Test>
      </AntdForm>

      {/* </FormProvider> */}
    </Container>
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
