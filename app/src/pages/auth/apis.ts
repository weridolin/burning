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


export function Login(loginForm:LoginRequestForm, successCallback: (res: any) => void, failCallback: (err: any) => void) {
  uni.request({
    url: "http://43.128.110.230:30080/usercenter/api/v1/login",
    data: loginForm,
    header: {
      "Content-Type": "application/x-www-form-urlencoded",
    },
    method: "POST",
    // sslVerify: true,
    success:(res)=>{
      console.log("登录成功", res);
      successCallback(res.data)
    },
    fail: (error) => {
      console.log("登录失败", error);
      failCallback(error)
    },
  });
}
