<template>
  <view class="content">
    <uni-section title="今日已完成训练 " type="line">
      <TrainHistoryBriefCard
        v-for="(trainDetail,index) in trainDetailList"
        :trainDetail="trainDetail"
        :key="index"
      >
      {{ index }}
      </TrainHistoryBriefCard>  
    </uni-section>
    <uni-section title="今日饮食记录 TODO" type="line">
      <uni-card :title="date" :extra="day">
        <view class="food-detail">
          <view class="food-total">
            <text>已经进食{{ food_info.total }}卡路里</text>
          </view>
          <view class="type-detail">
            <text class="type protein">蛋白质:{{ food_info.protein }}g</text>
            <text class="type fat">脂肪:{{ food_info.fat }}g</text>
            <text class="type carbohydrate new-row"
              >碳:{{ food_info.carbohydrate }}g</text
            >
            <text class="type water new-row">水:{{ food_info.water }}g</text>
          </view>
        </view>
      </uni-card>
    </uni-section>
    <uni-section title="教程推荐 TODO" type="line" subTitle="数据更新于2022-11-27" padding>
      <view class="move-tag-list">
        <view class="tag-view" 
          v-for="(type) in typeList"
          :key="type">
          <uni-tag :text="type" type="primary" :inverted="isSelect(type)" @click="setType(type)"/>
        </view>
      </view>
      <view class="video-list">
        <uni-card 
          :title="video_detail.title" 
          :sub-title="video_detail.up" 
          extra="查看" 
          :thumbnail="video_detail.thumbnail" 
          class="video-card" 
          style="margin:0px 5px"
          v-for="(video_detail,index) in video_info.list"
          :key="index"
          >
          <text class="uni-body">{{video_detail.summary}}</text>
        </uni-card>
      </view>
			<uni-pagination 
        :total="video_info.total" 
        :pageSize="getVideoInfoRequest.page_size" 
        v-model="getVideoInfoRequest.page" 
        prev-text="前一页" 
        next-text="后一页" 
        @change="changeVideoList"
      />
    </uni-section>
    <uni-section title="来首音乐 TODO" type="line" padding></uni-section>
  </view>
</template>

<script  lang="ts">
import uniSection from "@/uni_modules/uni-section/components/uni-section/uni-section.vue";
import Vue from "vue";
import { format, getDay } from "date-fns";
import {VideoInfoRequest,VideoInfo,VideoInfoResponse,FoodRecord,GetRandomMusic,GetFoodHistory,Music} from "./apis"
import {GetVideoInfo} from "./apis"
import {GetTodayTrainHistory} from "@/pages/history/apis"
import  TrainHistoryBriefCard  from "@/conpoments/history/trainHistoryBriefCard.vue";
import {
  TrainHistoryDetail
} from "@/pages/history/apis";


export default Vue.extend({
  components: { uniSection,TrainHistoryBriefCard },
  data() {
    var trainDetailList:TrainHistoryDetail[] = []
    var video_info:VideoInfoResponse={
      list:[{
        title:"视频名称",
        url:"视频url",
        up:"up主",
        type:[],
        thumbnail:"../../static/image/travel/personal/tx.png" ,
        summary:"视频简介"
      }],
      pre:"string",
      next:"string",
      page:1,
      page_size:10,
      total:50,
      
    }

    var getVideoInfoRequest:VideoInfoRequest={
      page:1,
      page_size:10,
      type:[]
    }

    var food_info: FoodRecord={
      total:0,
      protein:0,
      fat:0,
      carbohydrate:0,
      water:0
    }
    var music:Music = {
      name:"歌曲名称",
      author:"歌手",
      url:"歌曲url",
      cover:"歌曲封面"
    }

    return {
      date: "",
      day: "",
      food_info,
      video_info,
      getVideoInfoRequest,
      typeList: ["胸", "背", "腿", "肩", "手臂", "腹肌", "有氧", "臀"],
      selectType:Array<string>(),
      trainDetailList
    };
  },
  onLoad() {
    // this.getDate();
    // this.getFoodRecord();
    this.getTodayTrainRecord();
  },
  methods: {
    getDate() {
      // let currentDateTime = format(new Date(), 'yyyy-MM-dd');
      var dict = {
        0: "日",
        1: "一",
        2: "二",
        3: "三",
        4: "四",
        5: "五",
        6: "六",
      };
      this.date = format(new Date(), "yyyy-MM-dd");
      this.day = "星期" + dict[getDay(new Date())];
    },
    setType(type:string) {
      if( this.selectType.indexOf(type) > -1) {
        this.selectType.splice(this.selectType.indexOf(type), 1)
      } else {
      this.selectType.push(type)
    }
    console.log(this.selectType)
    },
    isSelect(type:string){
      return this.selectType.indexOf(type) == -1
    },
    changeVideoList(e:any){
      console.log("change video info page",e)
      GetVideoInfo(this.getVideoInfoRequest,(res)=>{
        console.log("get video info",res) 
        // this.video_info=res.data #TODO
      },(err)=>{
        console.log("get video info err",err)
      })
    },
    getFoodRecord(){
      // TODO
      console.log("get food record")
      GetFoodHistory(this.date,(res)=>{
        console.log("get food history",res)
      },(err)=>{
        console.log("get food history err",err)
      })
    },
    getTodayTrainRecord(){
      GetTodayTrainHistory((res)=>{
        console.log("get today train history",res)
        for (let index = 0; index < res.data.length; index++) {
          const element = res.data[index];
          if (element.train_history.finish){
          this.trainDetailList.push({
            train_history:element.train_history,
            train_content:element.train_content
          })}
        }
        console.log("this train detail list",this.trainDetailList)
      },(err)=>{
        console.log("get today train history err",err)
      })
    }
  },
});
</script>

<style lang="scss">
.content {
  display: flex;
  flex-direction: column;
  margin: 10px 10px;
  /* align-items: center;
  justify-content: center; */
}

.food-detail {
  display: flex;
  flex-direction: row;

  .food-total {
    flex: 0 1 50%;
    margin-top: 20px;
  }
  .type-detail {
    flex: 0 1 50%;
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
    /* flex-start:  默认从头开始排列*/
    justify-content: flex-start;
    .type {
      flex: 0 1 50%;
      width: 50%;
      text-align: center;
    }
    .new-row {
      margin-top: 10px;
      text-align: center;
    }
  }
}

.move-tag-list {
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  justify-content: flex-start;
  .tag-view {
    margin-left: 5px;
  }

}
.video-card {
  margin-left: 0px;
  margin: 0px 0px;
}
</style>
