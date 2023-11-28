// api.ts
import axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse } from "axios";
import { getToken } from "@/store/local"; 

interface AxiosRequestConfigPlus extends AxiosRequestConfig {
  // interceptorHooks?: InterceptorHooks;
  // showingLoading?:boolean;
  requiredLogin?: boolean
}

const baseURL = "127.0.0.1"

class BaseApi {
  private axiosInstance: AxiosInstance;

  constructor() {
    this.axiosInstance = axios.create({
      baseURL,
      timeout: 10000, // 设置请求超时时间
    });

    // 在此可以添加其他配置，如请求拦截器、响应拦截器等
  }

  request<T = any>(config: AxiosRequestConfigPlus): Promise<T> {
    if (config.requiredLogin) {
      config.headers = { Authorization: `Bearer ${getToken()?.access_token}` }
    }
    return new Promise((resolve, reject) => {
      this.axiosInstance
        .request<any, AxiosResponse<T>>(config)
        .then((res) => {
          // console.log(">>> 获取请求响应",res)
          resolve(res.data)
        })
        .catch((err) => {
          // console.log(">>> 获取请求响应异常",err)
          reject(err.response)
        })
    })
  }

  get<T = any>(config: AxiosRequestConfigPlus): Promise<T> {
    return this.request({ method: 'GET', ...config })
  }

  post<T = any>(config: AxiosRequestConfigPlus): Promise<T> {
    return this.request({ method: 'POST', ...config })
  }

  delete<T = any>(config: AxiosRequestConfigPlus): Promise<T> {
    return this.request({ method: 'DELETE', ...config })
  }

  patch<T = any>(config: AxiosRequestConfigPlus): Promise<T> {
    return this.request({ method: 'PATCH', ...config })
  }

  put<T = any>(config: AxiosRequestConfigPlus): Promise<T> {
    return this.request({ method: 'PUT', ...config })
  }

  // 可以添加其他 HTTP 方法，如 put、delete 等

  // 添加拦截器的方法
  public addInterceptor(
    onFulfilled?: (
      value: AxiosResponse<any>
    ) => AxiosResponse<any> | Promise<AxiosResponse<any>>,
    onRejected?: (error: any) => any
  ): number {
    return this.axiosInstance.interceptors.response.use(
      onFulfilled,
      onRejected
    );
  }

  // 移除拦截器的方法
  public removeInterceptor(interceptorId: number): void {
    this.axiosInstance.interceptors.response.eject(interceptorId);
  }
}

export default BaseApi;
