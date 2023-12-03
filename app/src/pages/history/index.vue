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
    <!-- <uni-section title="详细信息" type="line"> -->
    <uni-card :is-shadow="true" v-show="trainDetail != ''">
      <text class="uni-body">{{ trainDetail }}</text>
    </uni-card>
    <!-- </uni-section> -->
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
        <NewRecord :trainHistoryId="trainHistoryId" ref="newRecordEditPage">
        </NewRecord>
      </scroll-view>
    </uni-drawer>
  </view>
</template>

<script lang="ts">
import Vue from "vue";
import NewRecord from "./NewRecord.vue";
// import ActionGrid from "@/conpoments/actions/actionGrid.vue";
import {
  getDate,
  GetTrainHistory,
  GetTrainHistoryDetail,
  TrainHistoryBrief,
  TrainHistory,
  AddTrainHistory,
} from "./apis";
function getTrainDetail(date: string) {
  return "训练详情";
}

export default Vue.extend({
  data() {
    var transHistory: TrainHistoryBrief[] = [];
    var trainHistoryMap: { [key: string]: TrainHistory } = {};
    return {
      title: "训练记录",
      transHistory,
      trainHistoryMap,
      trainDetail: "当天暂无训练日志",
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
      test: [{
						date: "2023-12-03",
						info: '打卡'
					},
					{
						date: getDate(new Date(),-1).fullDate,
						info: '已打卡'
					}
				]
    };
  },
  onLoad() {
    uni.$on("finishTrain", ()=>{
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
  onShow(){
    uni.showLoading({
      title: "获取训练记录中...",
    });
    console.log(">>>>>",getDate(new Date(),-2).fullDate)
    this.getHistory()
  },
  methods: {
    change(e: any) {
      console.log("select date is " + e.fulldate);
      this.trainDetail = getTrainDetail(e.fulldate);
      let trainHistory = this.trainHistoryMap[e.fulldate];
      if (trainHistory) {
        GetTrainHistoryDetail(
          trainHistory.id,
          (res) => {
            console.log("get change history detail -> ", res);
            // this.trainDetail = res;
          },
          (err) => {
            console.log("get change history detail err -> ", err);
          }
        );
      } else {
        this.trainDetail = "当天暂无训练日志";
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
            this.trainHistoryMap[item.created_at] = item;
            let n_date = new Date(item.train_history.CreatedAt);
            this.transHistory.push({
              date: getDate(n_date, 0).fullDate,
              info: item.train_history.title,
            });
          }
          console.log("transHistory", this.transHistory)
        },
        (err) => {
          console.log("get train history err", err);
          uni.hideLoading();
        }
      );
    },
    trigger(e: any) {
      if (this.content[e.index].text == "添加") {
        AddTrainHistory(
          {
            comment: "",
            total_time: 0,
            title: "",
            finish: false,
          },
          (res: TrainHistory) => {
            console.log("add train history", res);
            this.trainHistoryId = res.id;
            this.status = "created";
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
      if (this.$refs.newRecordEditPage){
        let ele = this.$refs.newRecordEditPage as any
        ele.ActionSelect(item);
      }
    },

  },
  components: {
    NewRecord,
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
</style>
