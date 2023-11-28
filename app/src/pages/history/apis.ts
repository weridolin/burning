import BaseApi from "@/services/base";
import {BurningApis} from "@/services/api";

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

export interface TrainHistoryBrief {
  date: string,
  info:string
}

export interface TrainContent {
  action_name: string;
  left_weight: number;
  right_weight: number;
  total_weight: number;
  number: number;
  user_id: number;
  training_history_id: number;
  id: number;
  finish: boolean;
  action_id: number;
  action_instrument:string
  consume_time:number // 单位秒
  created_at:string
}



const HistoryApis = new BaseApi();

export function GetTrainHistory() {
  return HistoryApis.get<TrainHistory[]>({
    url: BurningApis.history.getHistory.url,
    params: {
      start_time: getDate(getStartOfMonth()).fullDate,
      end_time: getDate(new Date(),0).fullDate,
    },
    requiredLogin: BurningApis.history.getHistory.authenticated,
  });
}

export function GetTrainHistoryDetail(trainID: number) {
  return HistoryApis.get<TrainContent[]>({
    url: BurningApis.history.getTrainHistoryDetail.url(trainID),
    requiredLogin: BurningApis.history.getTrainHistoryDetail.authenticated,
  });
}

export function AddTrainHistory(data: TrainHistory[]) {
  return HistoryApis.post({
    url: BurningApis.history.addHistory.url,
    data: data,
    requiredLogin: BurningApis.history.addHistory.authenticated,
  });
}

export function DeleteTrainHistory(trainID: number) {
  return HistoryApis.delete({
    url: BurningApis.history.deleteHistory.url(trainID),
    requiredLogin: BurningApis.history.deleteHistory.authenticated,
  });
}