<template>
  <view class="container">
    <uni-section
      type="line"
      title="输入标题信息"
      class="container__page-header"
    >
      <view class="container__page-header__row">
        <input
          class="container__page-header__row__title-input"
          placeholder="输入该次训练标题"
          maxlength="10"
          v-model="trainHistory.title"
        />
        <uni-title
          type="h1"
          :title="format_consume_time"
          color="#027fff"
          class="container__page-header__row__consume-time"
        ></uni-title>
        <button
          type="default"
          plain="true"
          size="mini"
          class="container__page-header__row__start-button"
          @click="start"
        >
          {{ started ? "暂停" : "Go!" }}
        </button>
      </view>
    </uni-section>
    <uni-section type="line" title="动作" class="container__train-actions">
      <scroll-view scroll-y>
        <trainContent
          :key="index"
          v-for="(item, index) in trainHistory.trainActionList"
          :index="index"
          :actionName="item.action_name"
          :initTrainContent="item.action_content"
          :trainHistoryId="trainHistory.trainHistoryId"
          :actionInstrument="item.action_instrument"
          class="container__train-actions__train-item"
          ref="trainContent"
          @trainContentListUpdate="onTrainActionDetailUpdate"
          @deleteActionContent="onDeleteActionContent"
        ></trainContent>
      </scroll-view>
    </uni-section>
    <view class="container__footer">
      <view class="tool">
        <button type="default" plain="true" size="mini" @click="minimize">
          最小化
        </button>
        <button type="default" plain="true" size="mini" @click="addAction">
          添加动作
        </button>
        <button type="default" plain="true" size="mini" @click="addComment">
          添加心得
        </button>
        <button type="default" plain="true" size="mini" @click="onfinishTrain">
          完成
        </button>
      </view>
    </view>

    <!-- 添加心得 -->
    <uni-popup ref="popup" type="center">
      <uni-section title="输入你本次训练的心得吧" type="line" padding>
        <uni-easyinput
          type="textarea"
          autoHeight
          v-model="trainHistory.comment"
          placeholder="请输入内容"
        ></uni-easyinput>
      </uni-section>
    </uni-popup>
  </view>
</template>
<script lang="ts">
import Vue from "vue";
import {
  TrainContent,
  AddTrainHistoryResponse,
  ActionDetail,
  FinishTrain,
  TrainHistory,
  TrainContentToActionDetail,
} from "./apis";
import trainContent from "@/conpoments/history/trainContent.vue";
import { Action } from "@/pages/action/apis";
import {
  setDoingTrain,
  getDoingTrain,
  clearDoingTrain,
  getAllDoingTrain,
} from "@/store/local";

export default Vue.extend({
  components: { trainContent },
  data() {
    var timer: any;
    var trainActionList: ActionDetail[] = [];
    return {
      format_consume_time: "00:00:00",
      started: false,
      timer,
      trainHistory: {
        trainHistoryId: 0,
        consume_time: 0,
        title: "",
        comment: "",
        trainActionList,
        created_at: "",
        // unWatchTrainHistory:null
      },
      status: "created", // 新建记录情况下为created 修改记录情况为edit edit时不会清空/更新本地缓存
      date: "",
    };
  },
  mounted() {
    console.log("mounted....");
    // this.unWatchTrainHistory
  },
  onLoad() {
    // this.unWatchTrainHistory
  },
  watch: {
    trainHistory: {
      handler(new_value, _) {
        console.log("train history change ...", new_value.trainHistoryId);
        this.format_consume_time = this.formatConsumeTime(
          new_value.consume_time
        );
        if (this.status == "created") {
          setDoingTrain(new_value, this.date);
        }
      },
      deep: true,
    },
  },
  methods: {
    formatConsumeTime(seconds: number) {
      let hour = Math.floor(seconds / 3600);
      let minute = Math.floor((seconds - hour * 3600) / 60);
      let second = seconds - hour * 3600 - minute * 60;
      let paddedHours = hour.toString().padStart(2, "0");
      let paddedMinutes = minute.toString().padStart(2, "0");
      let paddedSeconds = second.toString().padStart(2, "0");

      return `${paddedHours}:${paddedMinutes}:${paddedSeconds}`;
    },
    start() {
      /* 训练时长计数器 */
      if (this.trainHistory.title == "") {
        uni.showToast({
          title: "请输入标题",
          icon: "error",
        });
        return;
      }
      if (this.started) {
        if (this.timer) {
          clearInterval(this.timer);
        }
        this.started = false;
      } else {
        this.started = true;
        this.timer = setInterval(() => {
          this.trainHistory.consume_time += 1;
        }, 1000);
      }
    },
    minimize() {
      // this.$refs["newRecord"].close();
      uni.$emit("minimizeDrawer"); //close drawer
    },
    addAction() {
      console.log("addAction");
      uni.navigateTo({
        url: "/pages/history/AddAction",
      });
    },
    addComment() {
      let ele = this.$refs["popup"] as any;
      if (ele != null) {
        ele.open();
      }
    },
    complete() {},
    getWidth() {
      let width = uni.getSystemInfoSync().windowWidth;
      console.log(width);
      return width;
    },
    ActionSelect(item: Action) {
      // 判断是否存在相同动作
      console.log("选择了动作", item, this.trainHistory);

      for (
        let index = 0;
        index < this.trainHistory.trainActionList.length;
        index++
      ) {
        const element = this.trainHistory.trainActionList[index];
        if (element.action_name == item.action_name) {
          console.log("存在相同动作");
          uni.showToast({
            title: "当前计划存在相同动作,已经合并",
            icon: "error",
            duration: 2000,
          });
          return;
        }
      }
      let actionDetail: ActionDetail = {
        action_name: item.action_name,
        action_content: [
          {
            left_weight: "0",
            right_weight: "0",
            total_weight: "0",
            number: "0",
            finish: false,
            action_name: item.action_name,
            training_history_id: this.trainHistory.trainHistoryId,
            action_instrument: item.action_instrument,
            consume_time: 0,
            action_id: item.id,
            // user_id:0,
          },
        ],
        consume_time: 0,
        action_instrument: item.action_instrument,
        action_id: item.id,
      };
      this.trainHistory.trainActionList.push(actionDetail);
    },
    onfinishTrain() {
      console.log("onfinishTrain", this.trainHistory.trainActionList);
      uni.showModal({
        title: "",
        content:
          this.status == "created" ? "是否结束当前训练?" : "是否保存当前训练?",
        success: (res) => {
          if (res.confirm) {
            if (this.trainHistory.title == "") {
              uni.showToast({
                title: "请输入标题",
                icon: "error",
              });
              return;
            }
            if (this.timer && this.started) {
              clearInterval(this.timer);
            }

            let h = {
              comment: this.trainHistory.comment,
              total_time: this.trainHistory.consume_time,
              title: this.trainHistory.title,
              finish: true,
              id: this.trainHistory.trainHistoryId,
            };
            uni.showLoading({ title: "提交中" });
            FinishTrain(
              h as TrainHistory,
              this.trainHistory.trainActionList,
              (res) => {
                this.started = false;
                this.trainHistory.consume_time = 0;
                this.trainHistory.title = "";
                this.trainHistory.trainActionList = [];
                this.trainHistory.trainHistoryId = 0;
                uni.$emit("finishTrain"); //close drawer
                uni.hideLoading();
                if (this.status == "created") {
                  clearDoingTrain();
                  this.status = ""; // 防止watch里面在进行修改
                }
                console.log("local store clear");
              },
              (err) => {
                console.log("finish train error", err);
                uni.hideLoading();
              }
            );
          } else {
            return;
          }
        },
      });
    },
    initData(data: AddTrainHistoryResponse, date: string) {
      console.log("initData", data, date);
      this.status = "created";
      //加载本地保留的未完成的trainContent
      if (data == null || data.existed) {
        let unfinishedTrainRecord = getDoingTrain(date) as any;
        console.log("有未完成的记录,加载本地记录 -> ", unfinishedTrainRecord);
        if (unfinishedTrainRecord != null) {
          // this.trainHistory.trainActionList = this.trainHistory.trainActionList.concat(unfinishedTrainContent);
          this.trainHistory = unfinishedTrainRecord;
          return;
        }
      }
      console.log("新建一条新的训练记录");
      this.date = data.train_history.created_at;
      this.trainHistory.title = data.train_history.title;
      this.trainHistory.trainHistoryId = data.train_history.id;
      this.trainHistory.comment = data.train_history.comment;
      this.trainHistory.consume_time = data.train_history.total_time;
      this.trainHistory.created_at = date;
      this.trainHistory.trainActionList = [];
    },
    editData(
      trainHistory: TrainHistory,
      trainContent: TrainContent[],
      date: string
    ) {
      this.date = trainHistory.created_at;
      if (trainHistory.finish) {
        this.status = "edit";
        let ActionDetailList = TrainContentToActionDetail(trainContent);
        this.trainHistory = {
          trainHistoryId: trainHistory.id,
          consume_time: trainHistory.total_time,
          title: trainHistory.title,
          comment: trainHistory.comment,
          trainActionList: ActionDetailList,
          created_at: trainHistory.created_at,
        };
      } else {
        this.status = "created";
        //加载本地保留的未完成的trainContent
        console.log(
          "加载本地保留的未完成的trainContent",
          trainHistory,
          trainContent
        );
        let unfinishedTrainRecord = getDoingTrain(
          trainHistory.created_at
        ) as any;
        // console.log("有未完成的记录,加载本地记录 -> ", unfinishedTrainRecord);
        if (unfinishedTrainRecord != null) {
          // this.trainHistory.trainActionList = this.trainHistory.trainActionList.concat(unfinishedTrainContent);
          this.trainHistory = unfinishedTrainRecord;
        } else {
          this.trainHistory = {
            trainHistoryId: trainHistory.id,
            consume_time: trainHistory.total_time,
            title: trainHistory.title,
            comment: trainHistory.comment,
            trainActionList: TrainContentToActionDetail(trainContent),
            created_at: trainHistory.created_at,
          };
        }
      }
    },
    onTrainActionDetailUpdate(item: TrainContent[], index: number) {
      this.trainHistory.trainActionList[index].action_content = item;
      if (this.status == "created") {
        setDoingTrain(this.trainHistory, this.date);
      }
      console.log("onTrainContentUpdate", item, index, this.trainHistory);
    },
    onDeleteActionContent(item: TrainContent) {
      console.log("onDeleteActionContent", item, this.trainHistory);
      // this.trainHistory.trainActionList.splice(item, 1);
    },
  },
});
</script>

<style lang="scss" scoped>
.container {
  display: flex;
  flex-direction: column;
  height: 100%;
  .container__page-header {
    flex: 2;
    width: 100%;
    .container__page-header__row {
      display: flex;
      justify-content: center;
      align-items: center;
      .container__page-header__row__title-input {
        flex: 8;
        margin-left: 10px;
      }

      .container__page-header__row__consume-time {
        flex: 2;
      }
      .container__page-header__row__start-button {
        flex: 1;
        margin: 15px 15px;
        font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
      }
    }
  }
  .container__train-actions {
    flex: 25;
    margin-bottom: 10%;
    z-index: 1;
    width: 100%;
    .container__train-actions__train-item {
      margin-top: 10px;
    }
  }
  .container__footer {
    flex: 1;
    bottom: 0;
    position: sticky;
    z-index: 10;
    width: 100%;
    background-color: rgba(255, 255, 255, 1);
    .tool {
      display: flex;
      margin-top: 10px;
      margin-bottom: 2%;
      margin-left: 20px;
      margin-right: 20px;
    }
  }
}
</style>