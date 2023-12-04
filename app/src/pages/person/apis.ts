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

