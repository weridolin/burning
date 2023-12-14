<template>
  <view class="content">
    <trainingNoticeBar
      style="width: 100%"
      ref="trainingNoticeBar"
    ></trainingNoticeBar>
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
    <view v-if="trainDetailList.length == 0">
      <uni-card :is-shadow="true">
        <text class="uni-body">当天无训练日志</text>
      </uni-card>
    </view>
    <view style="width: 100%" v-if="trainDetailList.length > 0">
      <TrainHistoryBriefCard
        v-for="(trainDetail, indextd) in trainDetailList"
        :key="indextd"
        :trainDetail="trainDetail"
        @click.native="onCardClick(trainDetail)"
        class="train-detail-card"
      >
      </TrainHistoryBriefCard>
    </view>

    <!-- 右下角悬浮按钮 -->
    <view class="button">
      <uni-fab
        ref="fab"
        :pattern="pattern"
        :content="content"
        :horizontal="horizontal"
        :vertical="vertical"
        :direction="direction"
        @trigger="trigger"
        v-show="status != 'created'"
      />
    </view>

    <!-- 新建记录页面 -->
    <uni-drawer
      ref="newRecord"
      mode="right"
      :width="getWidth()"
      :mask-click="true"
      v-show="status == 'created' || status == 'edit'"
      @change="drawChange"
    >
      <scroll-view style="height: 100%" scroll-y="true">
        <NewRecord ref="newRecordEditPage"> </NewRecord>
      </scroll-view>
    </uni-drawer>

    <!-- 卡片点击弹出菜单 -->
    <view>
      <uni-popup
        ref="card-popup"
        background-color="#fff"
        @change="cardPopMenuChange"
      >
        <view class="card-popup__btn_group">
          <view class="uni-popup-share">
            <view class="uni-share-title"
              ><text class="uni-share-title-text">更多操作</text></view
            >
            <view class="uni-share-content">
              <view class="uni-share-content-box">
                <view
                  class="uni-share-content-item"
                  v-for="(item, index) in bottomData"
                  :key="index"
                  @click.stop="onCardBtnClick(item, index, selectedCardId)"
                >
                  <image
                    class="uni-share-image"
                    :src="item.icon"
                    mode="aspectFill"
                  ></image>
                  <text class="uni-share-text">{{ item.text }}</text>
                </view>
              </view>
            </view>
          </view>
        </view>
      </uni-popup>
    </view>
  </view>
</template>

<script lang="ts">
import Vue from "vue";
import NewRecord from "./NewRecord.vue";
import TrainHistoryBriefCard from "../../conpoments/history/trainHistoryBriefCard.vue";
import {
  getDate,
  GetRecentMonthTrainHistory,
  TrainHistoryBrief,
  AddTrainHistory,
  AddTrainHistoryResponse,
  DeleteTrainHistory,
  TrainHistoryDetail,
} from "./apis";
import UniSection from "../../uni_modules/uni-section/components/uni-section/uni-section.vue";
import { getDoingTrain, clearDoingTrain } from "@/store/local";
import { isLogin } from "../../store/local";
import trainingNoticeBar from "@/conpoments/history/trainingNoticeBar.vue";

export default Vue.extend({
  data() {
    var transHistory: TrainHistoryBrief[] = []; //
    var trainHistoryMap: { [key: string]: TrainHistoryDetail[] } = {}; // {date: 训练列表}
    var trainDetailList: TrainHistoryDetail[] = []; // 当天的所有训练记录列表
    var selectedItem: TrainHistoryDetail = {} as TrainHistoryDetail;
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
      bottomData: [
        {
          text: "分享",
          icon: "../../static/icons/share.png",
          name: "share",
        },
        {
          text: "删除",
          icon: "../../static/icons/del.png",
          name: "del",
        },
        {
          text: "储存为模板",
          icon: "../../static/icons/saveTem.png",
          name: "saveAsTem",
        },
        {
          text: "编辑",
          icon: "../../static/icons/edit.png",
          name: "edit",
        },
      ],
      selectedCardId: 0,
      selectedItem,
    };
  },
  onLoad() {
    uni.$on("finishTrain", () => {
      console.log("finishTrain");
      let ele = this.$refs["newRecord"] as any;
      if (ele) {
        ele.close();
      }
      this.status = "";
      this.refreshHistory();
    });
    uni.$on("minimizeDrawer", () => {
      console.log("最小化训练记录");
      let ele = this.$refs["newRecord"] as any;
      if (ele) {
        ele.close();
      }
      this.status = "";
      this.refreshHistory();
    });
    // if (isLogin()){
    //   this.refreshHistory();
    // }
  },
  onUnload() {
    uni.$off("finishTrain");
    uni.$off("minimizeDrawer");
  },
  onShow() {
    console.log("on train index page show");
    if (isLogin()) {
      this.refreshHistory();
    } else {
      this.trainDetailList = [];
      this.trainHistoryMap = {};
      this.transHistory = [];
    }
    if (uni.getStorageSync("showDetail")){
      console.log("open training history detail")
      this.status = "created";
      let ele1 = this.$refs["newRecordEditPage"] as any;
      ele1.initData(null);
      let ele = this.$refs["newRecord"] as any;
      if (ele) {
        ele.open();
      }
      uni.removeStorageSync('showDetail')
    }
  },
  methods: {
    updateNoticeBar() {
      let el = this.$refs["trainingNoticeBar"] as any;
      if (el) {
        el.refreshStatus();
      }
    },
    refreshHistory() {
      uni.showLoading({
        title: "获取训练记录中...",
      });
      this.trainDetailList = [];
      this.trainHistoryMap = {};
      this.transHistory = [];
      this.getHistory();
      this.updateNoticeBar()
      //
    },
    change(e: any) {
      let _trainHistory = this.trainHistoryMap[e.fulldate];
      console.log(
        "select date is " + e.fulldate,
        "trainHistory",
        _trainHistory
      );
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
        this.trainDetailList = [];
      }
    },
    getHistory() {
      // 获取历史记录
      GetRecentMonthTrainHistory(
        (res) => {
          console.log("get train history", res);
          uni.hideLoading();
          if (!res.data) {
            return;
          }
          for (let i = 0; i < res.data.length; i++) {
            let item = res.data[i];
            if (!item.train_history.finish){
                //更新下当前正在训练的内容到当天的训练详情里面 
                let training = getDoingTrain()  as {
                  title:"",
                  consume_time:"",
                  comment:""
                }
                if (training){
                  // console.log(training,">>") 
                  // content会在进入详情页更新
                  item.train_history.title = training.title
                  item.train_history.total_time = training.consume_time
                  item.train_history.comment = training.comment
                }
            }
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
            console.log(
              "now day train detail ->",
              this.trainDetailList,
              this.trainHistoryMap,
              getDate(new Date(), 0).fullDate
            );
          }


        },
        (err) => {
          console.log("get train history err", err);
          uni.hideLoading();
        }
      );
    },
    trigger(e: any) {
      console.log("添加训练记录,当前登录状态 ->", isLogin());
      if (this.content[e.index].text == "添加") {
        if (!isLogin()) {
          uni.showToast({
            title: "请先登录",
            icon: "error",
            duration: 2000,
          });
          return;
        }
        //先判断是否有未完成的训练记录 #
        uni.showLoading({
          title: "初始化训练中...",
        });
        AddTrainHistory(
          { comment: "",
            total_time: 0,
            title: "",
            finish: false,
            force: false,
          },
          (res) => {
            res = res.data as AddTrainHistoryResponse;
            if (res.existed) {
              uni.showModal({
                title: "",
                content: "当前有未完成的训练,请先完成再添加!",
                showCancel: false,
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
            uni.hideLoading();
          },
          (err) => {
            uni.showToast({
              title: "添加历史记录失败",
              icon: "error",
            });
            console.log("add train history err", err);
            uni.hideLoading();
          }
        );
        
      }
    },
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
    onCardClick(item: TrainHistoryDetail) {
      console.log("onCardClick", item);
      this.selectedCardId = item.train_history.id;
      this.selectedItem = item;
      let ele = this.$refs["card-popup"] as any;
      if (ele) {
        ele.open("bottom");
      }
    },
    onCardBtnClick(item: any, index: any, trainID: number) {
      switch (item.name) {
        case "share":
          uni.showToast({
            title: "暂不支持",
            icon: "error",
          });
          let ele = this.$refs["card-popup"] as any;
          if (ele) {
            ele.close();
          }
          break;
        case "del":
          uni.showModal({
            title: "",
            content: "确定删除该训练记录吗?",
            success: (res) => {
              if (res.confirm) {
                uni.showLoading({
                  title: "删除中...",
                });
                DeleteTrainHistory(
                  trainID,
                  (res) => {
                    uni.hideLoading();
                    uni.showToast({
                      title: "删除成功",
                      icon: "success",
                      duration: 2000,
                    });
                    let ele = this.$refs["card-popup"] as any;
                    if (ele) {
                      ele.close();
                    }
                    //删除的是进行中的训练,删除本地缓存
                    if (!this.selectedItem.train_history.finish) {
                      clearDoingTrain();
                    }

                    this.refreshHistory();
                  },
                  (error) => {
                    uni.hideLoading();
                    let ele = this.$refs["card-popup"] as any;
                    if (ele) {
                      ele.close();
                    }
                  }
                );
              } else if (res.cancel) {
                console.log("用户点击取消");
              }
            },
          });
          break;
        case "saveAsTem":
          if (!this.selectedItem.train_history.finish) {
            uni.showToast({
              title: "未完成的训练记录不能保存为模板",
              icon: "error",
            });
            return;
          } else {
            // todo 保存为模板
            uni.showToast({
              title: "暂不支持",
              icon: "error",
            });
            return;
          }
        case "edit": {
          //判断当前编辑的是否为已经完成训练记录
          if (this.selectedItem.train_history.finish) {
            this.status = "edit";
          } else {
            this.status = "created";
          }
          let ele = this.$refs["newRecordEditPage"] as any;
          if (ele) {
            ele.editData(
              this.selectedItem.train_history,
              this.selectedItem.train_content
            );
          }
          let ele2 = this.$refs["newRecord"] as any;
          if (ele2) {
            ele2.open();
          }
        }
        default:
          break;
      }
      console.log(item, index);
      let ele = this.$refs["card-popup"] as any;
      if (ele) {
        ele.close();
      }
    },
    cardPopMenuChange() {
      // console.log("cardPopMenuChange");
    },
  },
  components: {
    NewRecord,
    UniSection,
    TrainHistoryBriefCard,
    trainingNoticeBar,
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
  /* box-sizing: border-box;
  border-style:solid;
  border-color:rgba(103, 105, 98, 0.459); */
}
/* .train-detail-card::before, .train-detail-card::after {
  box-sizing: inherit;
  position: absolute;
  content: '';
  border: 2px solid transparent;
  width: 0;
  height: 0;
}

.train-detail-card:hover::before {
  border-top-color: #4361ee;
  border-right-color: #4361ee;
  transition: width 0.3s ease-out, height 0.3s ease-out 0.3s;
}

.train-detail-card:hover::after {
  border-bottom-color: #4361ee;
  border-left-color: #4361ee;
  transition: border-color 0s ease-out 0.6s, width 0.3s ease-out 0.6s, height 0.3s ease-out 1s;
} */

.card-popup__btn_group {
  margin: 10px 10px;
}

/* pop menu  */
.uni-popup-share {
  background-color: #fff;
  border-top-left-radius: 11px;
  border-top-right-radius: 11px;
}
.uni-share-title {
  /* #ifndef APP-NVUE */
  display: flex;
  /* #endif */
  flex-direction: row;
  align-items: center;
  justify-content: center;
  height: 40px;
}
.uni-share-title-text {
  font-size: 14px;
  color: #666;
}
.uni-share-content {
  /* #ifndef APP-NVUE */
  display: flex;
  /* #endif */
  flex-direction: row;
  justify-content: center;
  padding-top: 10px;
}

.uni-share-content-box {
  /* #ifndef APP-NVUE */
  display: flex;
  /* #endif */
  flex-direction: row;
  flex-wrap: wrap;
  width: 360px;
}

.uni-share-content-item {
  width: 90px;
  /* #ifndef APP-NVUE */
  display: flex;
  /* #endif */
  flex-direction: column;
  justify-content: center;
  padding: 10px 0;
  align-items: center;
}

.uni-share-content-item:active {
  background-color: #f5f5f5;
}

.uni-share-image {
  width: 30px;
  height: 30px;
}

.uni-share-text {
  margin-top: 10px;
  font-size: 14px;
  color: #3b4144;
}

.uni-share-button-box {
  /* #ifndef APP-NVUE */
  display: flex;
  /* #endif */
  flex-direction: row;
  padding: 10px 15px;
}

.uni-share-button {
  flex: 1;
  border-radius: 50px;
  color: #666;
  font-size: 16px;
}

.uni-share-button::after {
  border-radius: 50px;
}
</style>
