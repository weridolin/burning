<template>
  <view class="edit-page">
    <uni-section type="line" title="输入标题信息" class="edit-page-header">
      <view class="header">
        <input
          class="title-input"
          placeholder="输入该次训练标题"
          maxlength="10"
          v-model="trainHistory.title"
        />
        <uni-title
          type="h1"
          :title="format_consume_time"
          color="#027fff"
          class="consume-time"
        ></uni-title>
        <button
          type="default"
          plain="true"
          size="mini"
          class="start-button"
          @click="start"
        >
          {{ started ? "暂停" : "Go!" }}
        </button>
      </view>
    </uni-section>
    <uni-section type="line" title="动作" class="edit-page-actions">
      <scroll-view scroll-y>
        <trainContent
          :key="index"
          v-for="(item, index) in trainHistory.trainActionList"
          :index="index"
          :actionName="item.action_name"
          :initTrainContent="item.action_content"
          :trainHistoryId="trainHistory.trainHistoryId"
          :actionInstrument="item.action_instrument"
          class="train-item"
          ref="trainContent"
          @trainContentListUpdate="onTrainActionDetailUpdate"
        ></trainContent>
      </scroll-view>
    </uni-section>
    <view class="edit-page-footer">
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
import { setDoingTrain, getDoingTrain, clearDoingTrain } from "@/store/local";

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
      },
      status: "created", // 新建记录情况下为created 修改记录情况为edit edit时不会清空/更新本地缓存
    };
  },
  mounted() {
    console.log("mounted....");
  },
  watch: {
    trainHistory: {
      handler(new_value, _) {
        console.log("train history change ...", new_value.trainHistoryId);
        this.format_consume_time = this.formatConsumeTime(
          new_value.consume_time
        );
        if (this.status == "created") {
          setDoingTrain(new_value);
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
            // duration:2000
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
        content: this.status == "created"?"是否结束当前训练?":"是否保存当前训练?",
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
                }
                console.log("local store clear", getDoingTrain);
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
    initData(data: AddTrainHistoryResponse) {
      console.log("initData", data);
      this.status = "created";
      //加载本地保留的未完成的trainContent
      if (data.existed) {
        let unfinishedTrainRecord = getDoingTrain() as any;
        console.log("有未完成的记录,加载本地记录 -> ", unfinishedTrainRecord);
        if (unfinishedTrainRecord != null) {
          // this.trainHistory.trainActionList = this.trainHistory.trainActionList.concat(unfinishedTrainContent);
          this.trainHistory = unfinishedTrainRecord;
        }
      } else {
        console.log("新建一条新的训练记录");
        this.trainHistory.title = data.train_history.title;
        this.trainHistory.trainHistoryId = data.train_history.id;
        this.trainHistory.comment = data.train_history.comment;
        this.trainHistory.consume_time = data.train_history.total_time;
      }
    },
    editData(trainHistory: TrainHistory, trainContent: TrainContent[]) {
      if (trainHistory.finish){
        this.status = "edit";
        let ActionDetailList = TrainContentToActionDetail(trainContent);
        this.trainHistory = {
          trainHistoryId: trainHistory.id,
          consume_time: trainHistory.total_time,
          title: trainHistory.title,
          comment: trainHistory.comment,
          trainActionList: ActionDetailList,
        };
      }else{
        this.status = "created";
        //加载本地保留的未完成的trainContent
        let unfinishedTrainRecord = getDoingTrain() as any;
        console.log("有未完成的记录,加载本地记录 -> ", unfinishedTrainRecord);
        if (unfinishedTrainRecord != null) {
          // this.trainHistory.trainActionList = this.trainHistory.trainActionList.concat(unfinishedTrainContent);
          this.trainHistory = unfinishedTrainRecord;
        }
      }

    },
    onTrainActionDetailUpdate(item: TrainContent[], index: number) {
      this.trainHistory.trainActionList[index].action_content = item;
      if (this.status == "created") {
        setDoingTrain(this.trainHistory);
      }
      console.log("onTrainContentUpdate", item, index, this.trainHistory);
    },
  },
});
</script>

<style lang="scss" scoped>
.header {
  display: flex;
  justify-content: center;
  align-items: center;
}

.title-input {
  flex: 8;
  margin-left: 10px;
}

.train-item {
  margin-top: 10px;
  // margin-bottom: 10px;
}
.consume-time {
  flex: 2;
}
.start-button {
  flex: 1;
  margin: 15px 15px;
  font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
}

.edit-page {
  display: flex;
  flex-direction: column;
  height: 100%;
  .edit-page-header {
    flex: 2;
  }
  .edit-page-actions {
    flex: 25;
    margin-bottom: 10%;
    z-index: 1;
  }
  .edit-page-footer {
    // display: flex;
    // margin-top: 10px;
    // margin-bottom: 8%;
    // margin-left: 20px;
    // margin-right: 20px;
    flex: 1;
    // bottom: 2%;
    bottom: 0;
    position: sticky;
    z-index: 10;
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