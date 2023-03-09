export type NextPageWithLayout = NextPage & {
  getLayout?: (page: ReactElement) => ReactNode;
};


interface UserInfo{
  username: string | null;
  id: string | null;
  token: string | null;
}

export interface IFormLogin {
  userName: string;
  password: string;
}