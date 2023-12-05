<template>
  <view class="content">
    <view>
      <uni-calendar
        :insert="true"
        :lunar="true"
        :start-date="'2019-3-2'"
        :end-date="'3099-5-20'"
        @change="change"
        :selected="transHistory"
      />
    </view>
    <!-- <view   >
    <uni-card :is-shadow="true" >
      <text class="uni-body">当天无训练日志</text>
    </uni-card>
    </view> -->
    <view>
      <uni-card
        :is-shadow="true"
        v-for="(trainDetail, indextd) in trainDetailList"
        :key="indextd"
        :title="trainDetail.train_history.title?trainDetail.train_history.title:'未定义标题'"
        :sub-title="trainDetail.train_history.comment?trainDetail.train_history.comment:'本次训练似乎没什么心得...'"
        :extra="trainDetail.train_history.finish ? '已完成' : '进行中...'"
        thumbnail="/static/image/travel/personal/attendance.png"
        class="train-detail-card"
      >
        <uni-section
          type="line"
          :title="actionName"
          titleFontSize="14px"
          v-for="(briefInfoList, actionName) in getTrainContentBriefInfo(
            trainDetail.train_content
          )"
          :key="actionName"
          class
        >
          <view
            class="brief-info-content-list"
            style="
              display: flex;
              flex-flow: row wrap;
              gap: 10px;
              margin-left: 10px;
            "
          >
            <uni-title
              type="h5"
              :title="
                item.left_weight +
                'kg*' +
                item.right_weight +
                'kg*' +
                item.number
              "
              v-for="(item, index) in briefInfoList"
              :key="index"
            >
            </uni-title>
          </view>
        </uni-section>
      </uni-card>
    </view>

    <view class="button">
      <uni-fab
        ref="fab"
        :pattern="pattern"
        :content="content"
        :horizontal="horizontal"
        :vertical="vertical"
        :direction="direction"
        @trigger="trigger"
        @fabClick="fabClick"
        v-show="status != 'created'"
      />
    </view>
    <uni-drawer
      ref="newRecord"
      mode="right"
      :width="getWidth()"
      :mask-click="true"
      v-show="status == 'created'"
      @change="drawChange"
    >
      <scroll-view style="height: 100%" scroll-y="true">
        <NewRecord ref="newRecordEditPage"> </NewRecord>
      </scroll-view>
    </uni-drawer>
  </view>
</template>

<script lang="ts">
import Vue from "vue";
import NewRecord from "./NewRecord.vue";
import {
  getDate,
  GetTrainHistory,
  GetTrainHistoryDetail,
  TrainHistoryBrief,
  TrainHistory,
  TrainContent,
  AddTrainHistory,
  AddTrainHistoryResponse,
} from "./apis";
import UniSection from "../../uni_modules/uni-section/components/uni-section/uni-section.vue";

interface TrainHistoryDetail {
  train_history: TrainHistory;
  train_content: TrainContent[];
}
interface BriefItem {
  action_name: string;
  number: number;
  left_weight: string;
  right_weight: string;
  total_weight: string;
}

export default Vue.extend({
  data() {
    var transHistory: TrainHistoryBrief[] = [];
    var trainHistoryMap: { [key: string]: TrainHistoryDetail[] } = {};
    var trainDetailList: TrainHistoryDetail[] = [];
    return {
      title: "训练记录",
      transHistory,
      trainHistoryMap,
      trainDetailList,
      directionStr: "垂直",
      horizontal: "left",
      vertical: "bottom",
      direction: "horizontal",
      pattern: {
        color: "#7A7E83",
        backgroundColor: "#fff",
        selectedColor: "#007AFF",
        buttonColor: "#007AFF",
        iconColor: "#fff",
      },
      is_color_type: false,
      content: [
        {
          iconPath: "/static/icons/add.png",
          selectedIconPath: "/static/icons/add.png",
          text: "添加",
          active: false,
        },
      ],
      status: "",
      trainHistoryId: 0,
      test: [
        {
          date: "2023-12-03",
          info: "打卡",
        },
        {
          date: getDate(new Date(), -1).fullDate,
          info: "已打卡",
        },
      ],
    };
  },
  onLoad() {
    uni.$on("finishTrain", () => {
      console.log("finishTrain");
      let ele = this.$refs["newRecord"] as any;
      if (ele) {
        ele.close();
      }
    });
  },
  onUnload() {
    uni.$off("finishTrain");
  },
  onShow() {
    uni.showLoading({
      title: "获取训练记录中...",
    });
    this.getHistory();
  },
  methods: {
    change(e: any) {
      let _trainHistory = this.trainHistoryMap[e.fulldate];
      console.log("select date is " + e.fulldate,"trainHistory", _trainHistory);
      if (_trainHistory) {
        this.trainDetailList = _trainHistory;
        // GetTrainHistoryDetail(
        //   trainHistory.id,
        //   (res) => {
        //     console.log("get change history detail -> ", res);
        //     // this.trainDetail = res;
        //   },
        //   (err) => {
        //     console.log("get change history detail err -> ", err);
        //     this.trainDetail = "获取训练日志出错";
        //   }
        // );
      } else {
        this.trainDetailList = [] ;
      }
    },
    getHistory() {
      // 获取历史记录
      GetTrainHistory(
        (res) => {
          console.log("get train history", res);
          uni.hideLoading();
          for (let i = 0; i < res.data.length; i++) {
            let item = res.data[i];
            let n_date = getDate(
              new Date(item.train_history.created_at),
              0
            ).fullDate;
            if (n_date in this.trainHistoryMap) {
              this.trainHistoryMap[n_date].push(item);
            } else {
              this.trainHistoryMap[n_date] = [item];
            }
            this.transHistory.push({
              date: getDate(n_date, 0).fullDate,
              info:
                item.train_history.title == ""
                  ? "进行中.."
                  : item.train_history.title,
            });
          }
          //默认更新下当天的训练详情
          let _trainHistory =
            this.trainHistoryMap[getDate(new Date(), 0).fullDate];
          if (_trainHistory) {
            this.trainDetailList = _trainHistory;
            console.log("now day train detail ->",this.trainDetailList,this.trainHistoryMap,getDate(new Date(), 0).fullDate)
          }
        },
        (err) => {
          console.log("get train history err", err);
          uni.hideLoading();
        }
      );
    },
    trigger(e: any) {
      if (this.content[e.index].text == "添加") {
        //先判断是否有未完成的训练记录 #TODO
        AddTrainHistory(
          {
            comment: "",
            total_time: 0,
            title: "",
            finish: false,
            force: false,
          },
          (res) => {
            res = res.data as AddTrainHistoryResponse;
            console.log("created a new train history", res);
            if (res.existed) {
              uni.showModal({
                title: "",
                content: "当前有未完成的训练,请先完成再添加!",
                // showCancel:false,
                success: function (_) {
                  console.log("打开未完成的训练记录 -> ", res);
                },
              });
            }
            console.log("add train history", res);
            this.trainHistoryId = res.train_history.id;
            this.status = "created";
            let ele1 = this.$refs["newRecordEditPage"] as any;
            ele1.initData(res);
            let ele = this.$refs["newRecord"] as any;
            if (ele) {
              ele.open();
            }
          },
          (err) => {
            uni.showToast({
              title: "添加历史记录失败",
              icon: "error",
            });
            console.log("add train history err", err);
          }
        );
      }
    },
    fabClick() {},
    drawChange(e: any) {
      if (!e) {
        this.status = "";
        uni.showTabBar();
      } else {
        uni.hideTabBar();
      }
    },
    getWidth() {
      let width = uni.getSystemInfoSync().windowWidth;
      console.log(width);
      return width;
    },

    OnActionSelect(item: any) {
      console.log("OnActionSelect", item);
      if (this.$refs.newRecordEditPage) {
        let ele = this.$refs.newRecordEditPage as any;
        ele.ActionSelect(item);
      }
    },
    getTrainContentBriefInfo(items: TrainContent[]): {
      [key: string]: BriefItem[];
    } {
      //生成训练内容的简要信息
      let res: { [key: string]: BriefItem[] } = {};
      for (let index = 0; index < items.length; index++) {
        const content = items[index];
        if (content.action_name in res) {
          res[content.action_name].push({
            action_name: content.action_name,
            number: content.number,
            left_weight: content.left_weight,
            right_weight: content.right_weight,
            total_weight: content.total_weight,
          });
        } else {
          res[content.action_name] = [
            {
              action_name: content.action_name,
              number: content.number,
              left_weight: content.left_weight,
              right_weight: content.right_weight,
              total_weight: content.total_weight,
            },
          ];
        }
      }
      return res;
    },
  },
  components: {
    NewRecord,
    UniSection,
    // ActionGrid
  },
});
</script>

<style>
.content {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

.logo {
  height: 200rpx;
  width: 200rpx;
  margin: 200rpx auto 50rpx auto;
}

.text-area {
  display: flex;
  justify-content: center;
}

.title {
  font-size: 36rpx;
  color: #8f8f94;
}

.uni-fab__circle {
  position: absolute;
  bottom: 10px;
  left: 5px;
}

.train-detail-card {
  width: 100%;
}
</style>
