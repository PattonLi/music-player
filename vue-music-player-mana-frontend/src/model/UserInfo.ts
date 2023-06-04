export interface UserInfo {
  nickName: string
  loginUserName: string
}

export interface CustomerInfo {
  userId: number;
  username: string;
  gender: string;
  age: number;
  email: string;
  password: string;
  nickname: string;
  phone: string;
  picUrl: string
}

export interface AdminInfo {
  adminId: number;
  adminname: string;
  password: string
}