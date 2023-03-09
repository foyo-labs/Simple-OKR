import * as yup from 'yup';
import { yupResolver } from '@hookform/resolvers/yup';


export const LoginFormSchema = yupResolver(
    yup.object().shape({
        email: yup.string().required("登录邮箱是必须的"),
        password: yup.string().required("密码是必须的"),
    }),
);