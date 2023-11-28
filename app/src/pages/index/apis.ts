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

export function GetVideoInfo(request: VideoInfoRequest) {
  return HomeApis.get<VideoInfoResponse>({
    url: BurningApis.home.getVideoInfo.url,
    params: request,
    requiredLogin: BurningApis.home.getVideoInfo.authenticated,
  });
}

export function GetRandomMusic() {
  return HomeApis.get<Music>({
    url: BurningApis.home.getRandomMusic.url,
    requiredLogin: BurningApis.home.getRandomMusic.authenticated,
  });
}

export function GetFoodHistory(date: any) {
  return HomeApis.get<FoodRecord>({
    url: BurningApis.history.getHistory.url,
    params:{
      date:date
    },
    requiredLogin: BurningApis.history.getHistory.authenticated,
  });
}