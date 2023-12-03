export const actionsList =  {
  胸: {
    杠铃: [
      {
        name: "杠铃卧推",
        desc: "杠铃卧推",
        id: "barbellBenchPress",
        cover: "../../static/actions/ganglingwotui.png",
      },
      {
        name: "上斜杠铃卧推",
        desc: "上斜杠铃卧推",
        id: "UpbarbellSquat",
        cover: "../../static/actions/shangxieganglingwotui.png",
      },
    ],
  },
  背: {
    杠铃: [
      {
        name: "杠铃划船",
        desc: "杠铃划船",
        id: "BarbellRow",
        cover: "../../static/actions/ganglinghuachuan.png",
      },
    ],
  },
  腿: {
    杠铃: [
      {
        name: "杠铃深蹲",
        desc: "杠铃深蹲",
        id: "barbellSquat",
        cover: "../../static/actions/ganglingshendun.png",
      },
      {
        name: "硬拉",
        desc: "硬拉",
        id: "deadlift",
        cover: "../../static/actions/yingla.png",
      },
    ],
  },
  肩: {
    杠铃: [
      {
        name: "杠铃推举",
        desc: "杠铃推举",
        id: "BarbellPress",
        cover: "../../static/actions/ganglingtuijv.png",
      },
    ],
  },
  手臂: {
    杠铃: [
      {
        name: "杠铃弯举",
        desc: "杠铃弯举",
        id: "BarbellCurl",
        cover: "../../static/actions/ganglingwanju.png",
      },
    ],
  },
  臀: {
    杠铃: [
      {
        name: "杠铃臀桥",
        desc: "杠铃臀桥",
        id: "BarbellHipThrust",
        cover: "../../static/actions/ganglingtunqiao.png",
      },
    ],
  },
  腹: {
    杠铃: [
      {
        name: "杠铃卷腹",
        desc: "杠铃卷腹",
        id: "BarbellRollout",
        cover: "../../static/actions/ganglingjuanfu.png",
      },
    ],
  },
  有氧: {
    杠铃: [
      {
        name: "杠铃卷腹",
        desc: "杠铃卷腹",
        id: "BarbellRollout",
        cover: "../../static/actions/ganglingjuanfu.png",
      },
    ],
  },
  体能: {
    杠铃: [
      {
        name: "杠铃卷腹",
        desc: "杠铃卷腹",
        id: "BarbellRollout",
        cover: "../../static/actions/ganglingjuanfu.png",
      },
    ],
  }
}

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

// export function  getActionDetail(id:number) {
//   return ActionApis.get<Action>({
//     url: BurningApis.actions.getActionDetail.url,
//     params: {id},
//     requiredLogin: BurningApis.actions.getActionDetail.authenticated,
//   })
// }