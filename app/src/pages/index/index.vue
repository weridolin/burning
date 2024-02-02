<template>
  <view>
    <!-- 当前训练提示框 -->
    <trainingNoticeBar
      style="width: 100%"
      ref="trainingNoticeBar"
    ></trainingNoticeBar>

    <view class="container">
      <!-- 训练日志 -->
      <uni-section
        title="今日已完成训练 "
        type="line"
        class="container__finish-train"
      >
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
      <uni-section title="今日饮食记录 " type="line" class="container__diet">
        <view
          class="container__diet_card"
          ref="diet-card-list"
          v-for="(item, index) in dietHistoryList"
          :key="index"
        >
          <dietCard :dietContentItemProp="item"> </dietCard>
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
        class="container__video"
      >
        <view class="container__video__tag_list">
          <view class="tag-view" :key="type" v-for="type in typeList">
            <uni-tag
              :text="type"
              type="primary"
              :inverted="isSelect(type)"
              @click="setType(type)"
            />
          </view>
        </view>
        <view>
          <uni-card
            :title="video_detail.title"
            :sub-title="video_detail.up"
            extra="查看"
            :thumbnail="video_detail.thumbnail"
            class="container__video__card"
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
      <uni-section
        title="来首音乐"
        type="line"
        padding
        class="container__music"
      >
        <!-- <view class="container__music__player" >
          <audio style="text-align: left;width: 100%;" :src="current.src" :poster="current.poster" :name="current.name" :author="current.author" :action="audioAction" controls></audio>
        </view>     -->
        <MusicPlayer></MusicPlayer>
      </uni-section>
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
import {
  GetTodayTrainHistory,
  GetTodayDietHistory,
} from "@/pages/history/apis";
import TrainHistoryBriefCard from "@/conpoments/history/trainHistoryBriefCard.vue";
import { TrainHistoryDetail, DietContentItem } from "@/pages/history/apis";
import { isLogin } from "@/store/local";
import trainingNoticeBar from "@/conpoments/history/trainingNoticeBar.vue";
import dietCard from "@/conpoments/history/dietCard.vue";
import dietContent from "@/conpoments/history/dietContent.vue";
import MusicPlayer from "@/conpoments/index/musicPlayer.vue";

export default Vue.extend({
  components: {
    uniSection,
    TrainHistoryBriefCard,
    trainingNoticeBar,
    dietCard,
    dietContent,
    MusicPlayer
    // LeAudio,
  },
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

    var dietHistoryList: DietContentItem[] = [];

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
      dietSelectedIndex: 0,
      current: {
				poster: 'https://qiniu-web-assets.dcloud.net.cn/unidoc/zh/music-a.png',
				name: '致爱丽丝',
				author: '暂无',
				src: 'https://bjetxgzv.cdn.bspapp.com/VKCEYUGU-hello-uniapp/2cc220e0-c27a-11ea-9dfb-6da8e309e0d8.mp3',
			},
			audioAction: {
				method: 'pause'
			}
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
    let el = this.$refs["trainingNoticeBar"] as any;
    if (el) {
      el.refreshStatus();
    }
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
    getTodayDietRecord() {
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
.container {
  display: flex;
  flex-direction: column;
  margin: 10px 10px;
  .container__video {
    .container__video__tag_list {
      display: flex;
      flex-direction: row;
      flex-wrap: wrap;
      justify-content: space-around;
      .tag-view {
        margin-left: 5px;
      }
    }
    .container__video__card {
      margin-left: 0px;
      margin: 0px 0px;
    }
  }
}
</style>
