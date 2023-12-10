export interface UserProfile {
  name: string;
  avatar: string;
  email: string;
  phone: string;
  gender: number;
  bodyInfo?: UserBodyInfo;
}

export interface UserBodyInfo {
  uuid: string;
  height: number;
  body_fat_rate: number; //体脂率
  chest_circumference: number; //胸围
  upper_arm_circumference: number; //臂围
  abdominal_circumference: number; //腹围
  thigh_circumference: number; //大腿围
  calf_circumference: number; //小腿围
  hip_line: number; //臀围
  waistline: number; //腰围
  weight: number; //体重
  shoulder_breadth: number; //肩宽
  days: number; //签到天数
}

interface AuthToken {
  access_token: string;
  refresh_token: string;
}

// export function UpdateUserProfile(params:{[key:string]:any}) {
//     const profile = getUserProfile()
//     if (profile==null) {
//         setUserProfile({
//             name: params.name,
//             avatar: params.avatar,
//             email: params.email,
//             phone: params.phone,
//             introduction: params.introduction,
//             roles: params.roles,
//             permissions: params.permissions,
//             id: params.id,
//             age: params.age
//         })
//     }else{
//         for (const key in params) {
//             profile[key] = params[key]
//         }
//     }
// }

export function setUserProfile(profile: UserProfile) {
  uni.setStorageSync("profile", JSON.stringify(profile));
}

export function getUserProfile(): UserProfile | null {
  const profile = uni.getStorageSync("profile");
  if (profile) {
    return JSON.parse(profile);
  }
  return null;
}

export function clearUserProfile() {
  uni.removeStorageSync("profile");
}

export function setToken(token: AuthToken) {
  uni.setStorageSync("token", JSON.stringify(token));
}

export function getToken(): AuthToken | null {
  const token = uni.getStorageSync("token");
  if (token) {
    return JSON.parse(token);
  }
  return null;
}

export function clearToken() {
  uni.removeStorageSync("token");
}

export function isLogin(): boolean {
  return !!getToken();
}

export function setDoingTrain(plan: object) {
  uni.setStorageSync("doingTrain", JSON.stringify(plan));
}

export function getDoingTrain(): object | null {
  const plan = uni.getStorageSync("doingTrain");
  if (plan) {
    return JSON.parse(plan);
  }
  return null;
}

export function clearDoingTrain() {
  uni.removeStorageSync("doingTrain");
}

export function setObject(key:string,val:object){
  uni.setStorageSync(key, JSON.stringify(val));
}

export function getObject(key:string){
  const val = uni.getStorageSync(key);
  if (val) {
    return JSON.parse(val);
  }
  return null;
}

export function setString(key:string,val:string){
  uni.setStorageSync(key, val);
} 

export function getString(key:string){
  return uni.getStorageSync(key);
}
