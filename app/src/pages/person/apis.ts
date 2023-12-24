import BaseApi from "@/services/base";
import { BurningApis } from "@/services/api";
// import {UserProfole}

const AuthApi = new BaseApi()

export interface GetUserProfilePayload {
  height:number
  body_fat_rate:number   //体脂率
  chest_circumference:number //胸围
  upper_arm_circumference:number //臂围
  abdominal_circumference:number //腹围
  thigh_circumference:number //大腿围
  calf_circumference:number //小腿围
  hip_line:number //臀围
  waistline:number //腰围
  weight:number //体重
  shoulder_breadth:number //肩宽
  days:number //签到天数
}


//

export function GetUserProfile(successCallback: (res: any) => void, failCallback: (err: any) => void) {
  return AuthApi.request<GetUserProfilePayload>({
    url: BurningApis.auth.getProfile.url,
    requiredLogin: BurningApis.auth.getProfile.authenticated,
    method: BurningApis.auth.getProfile.method,
    success: successCallback,
    fail: failCallback
  })
}

export function Sign(successCallback: (res: any) => void, failCallback: (err: any) => void) {
  return AuthApi.request<GetUserProfilePayload>({
    url: BurningApis.auth.sign.url,
    requiredLogin: BurningApis.auth.sign.authenticated,
    method: BurningApis.auth.sign.method,
    success: successCallback,
    fail: failCallback
  })
}

export function GetLastSign(successCallback: (res: any) => void, failCallback: (err: any) => void) {
  return AuthApi.request({
    url: BurningApis.auth.getLastSign.url,
    requiredLogin: BurningApis.auth.getLastSign.authenticated,
    method: BurningApis.auth.getLastSign.method,
    success: successCallback,
    fail: failCallback
  })
}



export function GetUserBodyInfo(date:string,successCallback: (res: any) => void, failCallback: (err: any) => void) {
  return AuthApi.request({
    url: BurningApis.auth.getBodyInfo.url,
    requiredLogin: BurningApis.auth.getBodyInfo.authenticated,
    method: BurningApis.auth.getBodyInfo.method,
    data:{
      "date":date
    },
    success: successCallback,
    fail: failCallback
  })
}

export function GetAllBodyInfo(successCallback: (res: any) => void, failCallback: (err: any) => void) {
  return AuthApi.request({
    url: BurningApis.auth.getBodyInfo.url,
    requiredLogin: BurningApis.auth.getBodyInfo.authenticated,
    method: BurningApis.auth.getBodyInfo.method,
    success: successCallback,
    fail: failCallback
  })
}


export function CreateUserBodyInfo(data:any,successCallback: (res: any) => void, failCallback: (err: any) => void) {
  return AuthApi.request({
    url: BurningApis.auth.createBodyInfo.url,
    requiredLogin: BurningApis.auth.createBodyInfo.authenticated,
    method: BurningApis.auth.createBodyInfo.method,
    data:data,
    success: successCallback,
    fail: failCallback
  })
}

export function UpdateUserBodyInfo(data:any,successCallback: (res: any) => void, failCallback: (err: any) => void) {
  return AuthApi.request({
    url: BurningApis.auth.updateBodyInfo.url,
    requiredLogin: BurningApis.auth.updateBodyInfo.authenticated,
    method: BurningApis.auth.updateBodyInfo.method,
    data:data,
    success: successCallback,
    fail: failCallback
  })
}