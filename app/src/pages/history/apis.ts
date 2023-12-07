import BaseApi from "@/services/base";
import {BurningApis} from "@/services/api";
import { forEach } from "lodash";

export function getDate(date: any, AddDayCount = 0) {
  if (!date) {
    date = new Date();
  }
  if (typeof date !== "object") {
    date = date.replace(/-/g, "/");
  }
  const dd = new Date(date);

  dd.setDate(dd.getDate() + AddDayCount); // 获取AddDayCount天后的日期

  const y = dd.getFullYear();
  const m =
    dd.getMonth() + 1 < 10 ? "0" + (dd.getMonth() + 1) : dd.getMonth() + 1; // 获取当前月份的日期，不足10补0
  const d = dd.getDate() < 10 ? "0" + dd.getDate() : dd.getDate(); // 获取当前几号，不足10补0
  return {
    fullDate: y + "-" + m + "-" + d,
    year: y,
    month: m,
    date: d,
    day: dd.getDay(),
  };
}

export function getStartOfMonth(): Date {
  const now = new Date();
  const startOfMonth = new Date(now.getFullYear(), now.getMonth(), 1);
  return startOfMonth;
}

export interface TrainHistory {
  comment: string;
  total_time: number;
  title: string;
  id: number;
  finish: boolean;
  user_id: number;
  created_at:string
}

export interface AddTrainHistoryRequest {
  comment: string;
  total_time: number;
  title: string;
  finish: boolean;
  force: boolean; //是否强制新建一个新的记录，如果为false，则如果存在未完成的记录，则不会新建，而是返回该记录
}

export interface AddTrainHistoryResponse  {
  train_history:TrainHistory
  train_content:TrainContent[]
  existed:boolean
}


export interface TrainHistoryBrief {
  date: string,
  info:string
}

export interface TrainContent {
  action_name: string;
  left_weight: string;
  right_weight: string;
  total_weight: string;
  number: string;
  user_id?: number;
  training_history_id: number;
  id?: number;
  finish: boolean;
  action_id: number;
  action_instrument:string
  consume_time:number // 单位秒
  created_at?:string
}


export interface ActionDetail {
  action_name: string;
  action_content: TrainContent[];
  consume_time: number;
  action_instrument: string;
  action_id: number;
}

export interface  TrainHistoryDetail {
  train_history: TrainHistory;
  train_content: TrainContent[];
}


const HistoryApis = new BaseApi();


export function TrainContentToActionDetail(trainContent: TrainContent[]):ActionDetail[] {
  let actionDetail:ActionDetail[] = []
  trainContent.forEach((item)=>{
    let index = actionDetail.findIndex((action)=>action.action_name==item.action_name)
    if (index==-1) {
      actionDetail.push({
        action_name: item.action_name,
        action_content: [item],
        consume_time: item.consume_time,
        action_instrument: item.action_instrument,
        action_id: item.action_id,
      })
    }else{
      actionDetail[index].action_content.push(item)
      actionDetail[index].consume_time += item.consume_time
    }
  })
  return actionDetail
}

export function GetRecentMonthTrainHistory(successCallback: (res: any) => void, failCallback: (err: any) => void) {
  return HistoryApis.request<TrainHistory[]>({
    url: BurningApis.history.getHistory.url,
    data: {
      start_time: getDate(getStartOfMonth(),-1).fullDate,
      end_time: getDate(new Date(),1).fullDate,
    },
    requiredLogin: BurningApis.history.getHistory.authenticated,
    method: BurningApis.history.getHistory.method,
    success: successCallback,
    fail: failCallback,
  });
}

export function GetTodayTrainHistory(successCallback: (res: any) => void, failCallback: (err: any) => void) {
  return HistoryApis.request<TrainHistory[]>({
    url: BurningApis.history.getHistory.url,
    data: {
      start_time: getDate(new Date(),-1).fullDate,
      end_time: getDate(new Date(),1).fullDate,
    },
    requiredLogin: BurningApis.history.getHistory.authenticated,
    method: BurningApis.history.getHistory.method,
    success: successCallback,
    fail: failCallback,
  });
}


export function GetTrainHistoryDetail(trainID: number,successCallback: (res: any) => void, failCallback: (err: any) => void) {
  return HistoryApis.request<TrainContent[]>({
    url: BurningApis.history.getTrainHistoryDetail.url(trainID),
    requiredLogin: BurningApis.history.getTrainHistoryDetail.authenticated,
    method: BurningApis.history.getTrainHistoryDetail.method,
    success: successCallback,
    fail: failCallback
  });
}

export function AddTrainHistory(data: AddTrainHistoryRequest,successCallback: (res: any) => void, failCallback: (err: any) => void) {
  return HistoryApis.request({
    url: BurningApis.history.addHistory.url,
    data: data,
    requiredLogin: BurningApis.history.addHistory.authenticated,
    method: BurningApis.history.addHistory.method,
    success: successCallback,
    fail: failCallback,
    contentType: "application/json"
  });
}

export function DeleteTrainHistory(trainID: number,successCallback: (res: any) => void, failCallback: (err: any) => void) {
  console.log("delete train history -> ",trainID)
  return HistoryApis.request({
    url: BurningApis.history.deleteHistory.url(trainID),
    requiredLogin: BurningApis.history.deleteHistory.authenticated,
    method: BurningApis.history.deleteHistory.method,
    success: successCallback,
    fail: failCallback
  });
}

export function UpdateTrainHistory(trainID: number,data: AddTrainHistoryRequest,successCallback: (res: any) => void, failCallback: (err: any) => void) {
  return HistoryApis.request({
    url: BurningApis.history.updateHistory.url(trainID),

    data: data,
    requiredLogin: BurningApis.history.updateHistory.authenticated,
    method: BurningApis.history.updateHistory.method,
    success: successCallback,
    fail: failCallback,
    contentType: "application/json"
  });
}

export function AddTrainHistoryContent(trainID:number,data: TrainContent,successCallback: (res: any) => void, failCallback: (err: any) => void) {
  return HistoryApis.request({
    url: BurningApis.history.addNewTrainContent.url(trainID),
    data: data,
    requiredLogin: BurningApis.history.addNewTrainContent.authenticated,
    method: BurningApis.history.addNewTrainContent.method,
    success: successCallback,
    fail: failCallback,
    contentType: "application/json"
  });
}


export function UpdateTrainHistoryContent(trainID:number,contentID:number,data: TrainContent,successCallback: (res: any) => void, failCallback: (err: any) => void) {
  return HistoryApis.request({
    url: BurningApis.history.updateTrainContent.url(trainID,contentID),
    data: data,
    requiredLogin: BurningApis.history.updateTrainContent.authenticated,
    method: BurningApis.history.updateTrainContent.method,
    success: successCallback,
    fail: failCallback,
    contentType: "application/json"
  });
}

export function DeleteTrainHistoryContent(trainID:number,contentID:number,successCallback: (res: any) => void, failCallback: (err: any) => void) {
  return HistoryApis.request({
    url: BurningApis.history.deleteTrainContent.url(trainID,contentID),
    requiredLogin: BurningApis.history.deleteTrainContent.authenticated,
    method: BurningApis.history.deleteTrainContent.method,
    success: successCallback,
    fail: failCallback
  });
}

export function FinishTrain(TrainHistory: TrainHistory,actionDetail:ActionDetail[] ,successCallback: (res: any) => void, failCallback: (err: any) => void) {
  let contentList:TrainContent[] = []
  actionDetail.forEach((item)=>{
    contentList = contentList.concat(item.action_content)
  })
  console.log(contentList,TrainHistory,"finish train")
  return HistoryApis.request({
    url: BurningApis.history.finishTrain.url(TrainHistory.id),
    data: {
      train_history: TrainHistory,
      train_content: contentList,
    },
    requiredLogin: BurningApis.history.finishTrain.authenticated,
    method: BurningApis.history.finishTrain.method,
    success: successCallback,
    fail: failCallback,
    contentType: "application/json"
  });
}