import BaseApi from "@/services/base";
import { BurningApis } from "@/services/api";


export interface Music{
  title : string;//标题
  singer : string;//歌手
  image : string;//封面
  fileUrl : string;//文件
  enable: boolean;//是否可用
  id: number;//id
}

export interface GetMusicListRequest {
  page: number;
  size: number;
}

export interface GetMusicListResponse {
  code:number;
  data:{
    list: Music[];
    total: number;
    page: number;
    pageSize: number;
  }

}

const MusicApis  = new BaseApi()

export function getMusicList(data: GetMusicListRequest,successCallback: (res: any) => void, failCallback: (err: any) => void) {
  return MusicApis.request<GetMusicListResponse>({
    method: BurningApis.media.getMusicList.method,
    url: BurningApis.media.getMusicList.url,
    requiredLogin: BurningApis.media.getMusicList.authenticated,
    data,
    success: successCallback,
    fail: failCallback,
  });
}