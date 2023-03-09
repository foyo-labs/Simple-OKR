export type NextPageWithLayout = NextPage & {
  getLayout?: (page: ReactElement) => ReactNode;
};


interface UserInfo{
  email: string | null;
  id: string | null;
  token: string | null;
}

export interface IFormLogin {
  email: string;
  password: string;
}


interface Objective {
  name: string | null
}