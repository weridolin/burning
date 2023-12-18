import BaseApi from "@/services/base";
import {BurningApis} from "@/services/api";


const ActionApis= new BaseApi()

export interface Action {
  uuid: string;
  action_name:string; //动作名称 比如 杠铃卧推
  action_instrument:string; //器械
  action_type:string; //类型 比如胸肩背
  action_desc:string;
  action_img:string;
  action_video:string;
  id:number;
  created_at:string;
}


export function GetActionList(successCallback: (res: any) => void, failCallback: (err: any) => void):UniApp.RequestTask {
  return ActionApis.request<Action[]>({
    url: BurningApis.actions.getActions.url,
    method: BurningApis.actions.getActions.method,
    requiredLogin: BurningApis.actions.getActions.authenticated,
    success: successCallback,
    fail: failCallback
  })
}

export function GetCustomActions(successCallback: (res: any) => void, failCallback: (err: any) => void):UniApp.RequestTask {
  return ActionApis.request<Action[]>({
    url: BurningApis.actions.getCustomActions.url,
    method: BurningApis.actions.getCustomActions.method,
    requiredLogin: BurningApis.actions.getCustomActions.authenticated,
    success: successCallback,
    fail: failCallback
  })
}

export function AddCustomAction(data: Action, successCallback: (res: any) => void, failCallback: (err: any) => void):UniApp.RequestTask {
  return ActionApis.request<Action>({
    url: BurningApis.actions.addCustomAction.url,
    method: BurningApis.actions.addCustomAction.method,
    data,
    requiredLogin: BurningApis.actions.addCustomAction.authenticated,
    success: successCallback,
    fail: failCallback
  })
}

export function DeleteCustomAction(id:number, successCallback: (res: any) => void, failCallback: (err: any) => void):UniApp.RequestTask {
  return ActionApis.request<Action>({
    url: BurningApis.actions.deleteCustomAction.url(id),
    method: BurningApis.actions.deleteCustomAction.method,
    requiredLogin: BurningApis.actions.deleteCustomAction.authenticated,
    success: successCallback,
    fail: failCallback
  })
}

export function UpdateCustomAction(id:number, data: Action, successCallback: (res: any) => void, failCallback: (err: any) => void):UniApp.RequestTask {
  return ActionApis.request<Action>({
    url: BurningApis.actions.updateCustomAction.url(id),
    method: BurningApis.actions.updateCustomAction.method,
    data,
    requiredLogin: BurningApis.actions.updateCustomAction.authenticated,
    success: successCallback,
    fail: failCallback
  })
}


// export function  getActionDetail(id:number) {
//   return ActionApis.get<Action>({
//     url: BurningApis.actions.getActionDetail.url,
//     params: {id},
//     requiredLogin: BurningApis.actions.getActionDetail.authenticated,
//   })
// }