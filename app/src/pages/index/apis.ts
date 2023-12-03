import BaseApi from "@/services/base";
import { BurningApis } from "@/services/api";
export interface VideoInfo {
  title: string;
  url: string;
  up: string;
  type: string[];
  thumbnail: string;
  summary: string;
}

export interface VideoInfoResponse {
  list: VideoInfo[];
  pre: string;
  next: string;
  page: number;
  page_size: number;
  total: number;
}

export interface VideoInfoRequest {
  page: number;
  page_size: number;
  type: string[];
  up?: string;
  start_time?: string;
  end_time?: string;
}

export interface FoodRecord {
  total: number;
  protein: number;
  fat: number;
  carbohydrate: number;
  water: number;
  // tagType: "default",
}

export interface Music {
  url: string;
  author: string;
  name: string;
  cover: string;
}

const HomeApis = new BaseApi();

export function GetVideoInfo(request: VideoInfoRequest,successCallback: (res: any) => void, failCallback: (err: any) => void) {
  return HomeApis.request<VideoInfoResponse>({
    url: BurningApis.home.getVideoInfo.url,
    data: request,
    requiredLogin: BurningApis.home.getVideoInfo.authenticated,
    method: BurningApis.home.getVideoInfo.method,
    success: successCallback,
    fail: failCallback
  });
}

export function GetRandomMusic(successCallback: (res: any) => void, failCallback: (err: any) => void) {
  return HomeApis.request<Music>({
    url: BurningApis.home.getRandomMusic.url,
    requiredLogin: BurningApis.home.getRandomMusic.authenticated,
    method: BurningApis.home.getRandomMusic.method,
    success: successCallback,
    fail: failCallback
  });
}

export function GetFoodHistory(date: any,successCallback: (res: any) => void, failCallback: (err: any) => void) {
  return HomeApis.request<FoodRecord>({
    url: BurningApis.history.getHistory.url,
    data:{
      date:date
    },
    method: BurningApis.history.getHistory.method,
    requiredLogin: BurningApis.history.getHistory.authenticated,
    success: successCallback,
    fail: failCallback
  });
}