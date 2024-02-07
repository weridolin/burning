// api.ts
// import axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse } from "axios";
import { getToken, clearToken } from "@/store/local";

// interface RequestConfig extends AxiosRequestConfig {
//   // interceptorHooks?: InterceptorHooks;
//   // showingLoading?:boolean;
//   requiredLogin?: boolean
// }

interface RequestConfig {
  // interceptorHooks?: InterceptorHooks;
  // showingLoading?:boolean;
  requiredLogin?: boolean;
  method: string;
  url: string;
  data?: any;
  headers?: any;
  success?: (res: any) => void;
  fail?: (err: any) => void;

  /*
    对于 GET 方法，会将数据转换为 query string。例如 { name: 'name', age: 18 } 转换后的结果是 name=name&age=18。
    对于 POST 方法且 header['content-type'] 为 application/json 的数据，会进行 JSON 序列化。
    对于 POST 方法且 header['content-type'] 为 application/x-www-form-urlencoded 的数据，会将数据转换为 query string。
  */
  contentType?: string;
}

export interface ResponseBase {
  code: number;
  msg: string;
  data: any;
}

const baseURL = "https://www.weridolin.cn";

class BaseApi {
  // private axiosInstance: AxiosInstance;

  constructor() {
    // this.axiosInstance = axios.create({
    //   baseURL,
    //   timeout: 10000, // 设置请求超时时间
    // });
    // 在此可以添加其他配置，如请求拦截器、响应拦截器等
  }

  request<T = any>(config: RequestConfig): UniApp.RequestTask {
    config.headers = config.headers || {};
    config.headers["Content-Type"] = config.contentType || "application/json";
    if (config.requiredLogin) {
      config.headers["Authorization"] = `Bearer ${getToken()?.access_token}`;
    }
    // return new Promise((resolve, reject) => {
    //   this.axiosInstance
    //     .request<any, AxiosResponse<T>>(config)
    //     .then((res) => {
    //       // console.log(">>> 获取请求响应",res)
    //       resolve(res.data)
    //     })
    //     .catch((err) => {
    //       // console.log(">>> 获取请求响应异常",err)
    //       reject(err.response)
    //     })
    // })
    // console.log(">>> 发起请求", config);
    return uni.request({
      url: `${baseURL}${config.url}`,
      method: config.method as
        | "OPTIONS"
        | "GET"
        | "HEAD"
        | "POST"
        | "PUT"
        | "DELETE"
        | "TRACE"
        | "CONNECT"
        | undefined,
      data: config.data,
      header: config.headers,
      success: (res) => {
        // console.log(">>> 获取请求响应", res);
        if (res.statusCode == 200) {
          config.success && config.success(res.data as ResponseBase);
        } else {
          if (res.statusCode == 401) {
              clearToken();
              uni.showToast({
                title: "请先登录",
                icon: "error",
                duration: 3000,
              });
              // uni.showModal({
              //   title: "未登录",
              //   content: "请重新登录",
              //   // showCancel: ,
              //   success: (res) => {
              //     if (res.confirm) {
              //       clearToken();
              //       uni.navigateTo({
              //         url: "/pages/auth/login",
              //       });
              //     }else{
              //       uni.reLaunch({
              //         url: "/pages/index/index",
              //       });
              //     }
              //   },
              // })
          } else {
            uni.showToast({
              title: "API请求失败",
              duration: 2000,
              icon: "error",
            });
            config.fail && config.fail(res.data as ResponseBase);
          }
        }
      },
      fail: (err) => {
        console.log(">>> 获取请求响应异常", err);
        config.fail && config.fail(err);
      },
    });
  }

  // 可以添加其他 HTTP 方法，如 put、delete 等

  // // 添加拦截器的方法
  // public addInterceptor(
  //   onFulfilled?: (
  //     value: AxiosResponse<any>
  //   ) => AxiosResponse<any> | Promise<AxiosResponse<any>>,
  //   onRejected?: (error: any) => any
  // ): number {
  //   return this.axiosInstance.interceptors.response.use(
  //     onFulfilled,
  //     onRejected
  //   );
  // }

  // // 移除拦截器的方法
  // public removeInterceptor(interceptorId: number): void {
  //   this.axiosInstance.interceptors.response.eject(interceptorId);
  // }
}

export default BaseApi;
