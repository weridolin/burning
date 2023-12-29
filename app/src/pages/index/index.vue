<template>
  <view>
    <!-- 当前训练提示框 -->
    <trainingNoticeBar
      style="width: 100%"
      ref="trainingNoticeBar"
    ></trainingNoticeBar>

    <view class="content">
    <!-- 训练日志 -->
      <uni-section title="今日已完成训练 " type="line">
        <view
          v-show="isLogin() && trainDetailList && trainDetailList.length > 0"
        >
          <TrainHistoryBriefCard
            v-for="(trainDetail, index) in trainDetailList"
            :trainDetail="trainDetail"
            :key="index"
          >
            {{ index }}
          </TrainHistoryBriefCard>
        </view>
        <uni-card :is-shadow="true" v-show="!isLogin()">
          <text class="uni-h6">登录后可查看</text>
        </uni-card>
        <uni-card
          :is-shadow="true"
          v-show="isLogin() && trainDetailList && trainDetailList.length == 0"
        >
          <text class="uni-h6">暂无完成的训练记录</text>
        </uni-card>
      </uni-section>

      <!-- 饮食记录 -->
      <uni-section title="今日饮食记录 " type="line">
        <view class="diet-card"
          ref = "diet-card-list" 
          v-for="(item,index) in dietHistoryList"
          :key="index"
          >
          <dietCard 
            :dietContentItemProp="item"
            >
          </dietCard>
        </view>
        <uni-card :is-shadow="true" v-show="!isLogin()">
          <text class="uni-h6">登录后可查看</text>
        </uni-card>
        <uni-card
          :is-shadow="true"
          v-show="isLogin() && dietHistoryList && dietHistoryList.length == 0"
        >
          <text class="uni-h6">暂无饮食记录</text>
        </uni-card>
      </uni-section>
      
      <!-- 教学视频 -->
      <uni-section
        title="教程推荐 TODO"
        type="line"
        subTitle="数据更新于2022-11-27"
        padding
      >
        <view class="move-tag-list">
          <view class="tag-view" v-for="type in typeList" :key="type">
            <uni-tag
              :text="type"
              type="primary"
              :inverted="isSelect(type)"
              @click="setType(type)"
            />
          </view>
        </view>
        <view class="video-list">
          <uni-card
            :title="video_detail.title"
            :sub-title="video_detail.up"
            extra="查看"
            :thumbnail="video_detail.thumbnail"
            class="video-card"
            style="margin: 0px 5px"
            v-for="(video_detail, index) in video_info.list"
            :key="index"
          >
            <text class="uni-body">{{ video_detail.summary }}</text>
          </uni-card>
        </view>
        <uni-pagination
          :total="video_info.total"
          :pageSize="getVideoInfoRequest.page_size"
          v-model="getVideoInfoRequest.page"
          prev-text="前一页"
          next-text="后一页"
          @change="changeVideoList"
        >
        </uni-pagination>
      </uni-section>

      <!-- 音乐推荐 -->
      <uni-section title="来首音乐 TODO" type="line" padding></uni-section>
    </view>
  </view>
</template>

<script  lang="ts">
import uniSection from "@/uni_modules/uni-section/components/uni-section/uni-section.vue";
import Vue from "vue";
import { format, getDay } from "date-fns";
import {
  VideoInfoRequest,
  VideoInfo,
  VideoInfoResponse,
  FoodRecord,
  GetRandomMusic,
  GetFoodHistory,
  Music,
} from "./apis";
import { GetVideoInfo } from "./apis";
import { GetTodayTrainHistory,GetTodayDietHistory } from "@/pages/history/apis";
import TrainHistoryBriefCard from "@/conpoments/history/trainHistoryBriefCard.vue";
import { TrainHistoryDetail,DietContentItem } from "@/pages/history/apis";
import { isLogin } from "@/store/local";
import trainingNoticeBar from "@/conpoments/history/trainingNoticeBar.vue";
import dietCard from "@/conpoments/history/dietCard.vue";
import dietContent from  "@/conpoments/history/dietContent.vue"

export default Vue.extend({
  components: { uniSection, TrainHistoryBriefCard, trainingNoticeBar, dietCard,dietContent },
  data() {
    var trainDetailList: TrainHistoryDetail[] = [];
    var video_info: VideoInfoResponse = {
      list: [
        {
          title: "视频名称",
          url: "视频url",
          up: "up主",
          type: [],
          thumbnail: "../../static/image/travel/personal/tx.png",
          summary: "视频简介",
        },
      ],
      pre: "string",
      next: "string",
      page: 1,
      page_size: 10,
      total: 50,
    };

    var getVideoInfoRequest: VideoInfoRequest = {
      page: 1,
      page_size: 10,
      type: [],
    };

    var dietHistoryList: DietContentItem[] = []

    var music: Music = {
      name: "歌曲名称",
      author: "歌手",
      url: "歌曲url",
      cover: "歌曲封面",
    };

    return {
      date: "",
      day: "",
      dietHistoryList,
      video_info,
      getVideoInfoRequest,
      typeList: ["胸", "背", "腿", "肩", "手臂", "腹肌", "有氧", "臀"],
      selectType: Array<string>(),
      trainDetailList,
      dietSelectedIndex:0
    };
  },
  onLoad() {
    //监听dietContent更新事件
    // this.$on('dietContentUpdated',(dietContentItem:DietContentItem)=>{
    //   console.log("dietContentUpdated -> ",dietContentItem)
    //   this.dietHistoryList[this.dietSelectedIndex]=dietContentItem
    //   this.closeDietDetail()
    // })

  },
  onShow() {
    this.getTodayTrainRecord();
    this.getTodayDietRecord();
    this.$forceUpdate();
    let el = this.$refs["trainingNoticeBar"] as any
    if (el){
      el.refreshStatus()
    }
    console.debug("on index page show");

  },
  computed: {
    is_login() {
      return isLogin();
    },
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
    setType(type: string) {
      if (this.selectType.indexOf(type) > -1) {
        this.selectType.splice(this.selectType.indexOf(type), 1);
      } else {
        this.selectType.push(type);
      }
      console.log(this.selectType);
    },
    isSelect(type: string) {
      return this.selectType.indexOf(type) == -1;
    },
    changeVideoList(e: any) {
      console.log("change video info page", e);
      GetVideoInfo(
        this.getVideoInfoRequest,
        (res) => {
          console.log("get video info", res);
          // this.video_info=res.data #TODO
        },
        (err) => {
          console.log("get video info err", err);
        }
      );
    },
    getFoodRecord() {
      // TODO
      console.log("get food record");
      GetFoodHistory(
        this.date,
        (res) => {
          console.log("get food history", res);
        },
        (err) => {
          console.log("get food history err", err);
        }
      );
    },
    getTodayTrainRecord() {
      GetTodayTrainHistory(
        (res) => {
          console.log("get today train history", res);
          this.trainDetailList = [];
          for (let index = 0; index < res.data.length; index++) {
            const element = res.data[index];
            if (element.train_history.finish) {
              this.trainDetailList.push({
                train_history: element.train_history,
                train_content: element.train_content,
              });
            }
          }
          console.log("this train detail list", this.trainDetailList);
        },
        (err) => {
          console.log("get today train history err", err);
        }
      );
    },
    getTodayDietRecord(){
      GetTodayDietHistory(
        (res) => {
          console.log("get today diet history", res);
          this.dietHistoryList = res.data as DietContentItem[];
        },
        (err) => {
          console.log("get today diet history err", err);
        }
      );
    },
    isLogin() {
      return isLogin();
    },
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

// .food-detail {
//   display: flex;
//   flex-direction: row;

//   .food-total {
//     flex: 0 1 50%;
//     margin-top: 20px;
//   }
//   .type-detail {
//     flex: 0 1 50%;
//     display: flex;
//     flex-direction: row;
//     flex-wrap: wrap;
//     /* flex-start:  默认从头开始排列*/
//     justify-content: flex-start;
//     .type {
//       flex: 0 1 50%;
//       width: 50%;
//       text-align: center;
//     }
//     .new-row {
//       margin-top: 10px;
//       text-align: center;
//     }
//   }
// }

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
