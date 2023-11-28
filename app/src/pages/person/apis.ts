import BaseApi from "@/services/base";
import { BurningApis } from "@/services/api";


const AuthApi = new BaseApi()

export interface UserProfile {
  user_name: string
  user_id: number
  user_email: string
  user_avatar: string
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

export function GetUserProfile():Promise<UserProfile> {
  return AuthApi.get<UserProfile>({
    url: BurningApis.auth.getProfile.url,
    requiredLogin: BurningApis.auth.getProfile.authenticated,
  })
}