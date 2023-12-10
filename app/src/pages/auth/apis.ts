import {BurningApis} from "@/services/api";
import BaseApi from "@/services/base";
export interface LoginRequestForm {
    phone: "",
    password: "",
    email: "",
    count: "",
}


export interface LoginResponsePayload {
  username: string,
  email: string,
  phone: string,
  avatar: string,
  role: null,
  is_super_admin: boolean,
  age: number,
  gender: number,
  access_token: string,
  refresh_token: string
} 

export interface RegisterForm {
  username: string,
  password: string,
  email: string,
  checkcode: string,
}

const AuthApis = new BaseApi()

export function Login(loginForm:LoginRequestForm, successCallback: (res: any) => void, failCallback: (err: any) => void) {
  // uni.request({
  //   // url: "http://43.128.110.230:30080/usercenter/api/v1/login",
  //   url:BurningApis.auth.login.url,
  //   data: loginForm,
  //   header: {
  //     "Content-Type": "application/x-www-form-urlencoded",
  //   },
  //   method: BurningApis.auth.login.method,
  //   // sslVerify: true,
  //   success:(res)=>{
  //     console.log("登录成功", res);
  //     successCallback(res.data)
  //   },
  //   fail: (error) => {
  //     console.log("登录失败", error);
  //     failCallback(error)
  //   },
  // });
  AuthApis.request<LoginResponsePayload>({
    url: BurningApis.auth.login.url,
    method: BurningApis.auth.login.method,
    contentType: "application/x-www-form-urlencoded",
    data: loginForm,
    success: successCallback,
    fail: failCallback
  })
}
