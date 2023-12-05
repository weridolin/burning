<template>
  <view class="edit-page">
    <uni-section type="line" title="输入标题信息" class="edit-page-header">
      <view class="header">
        <input
          class="title-input"
          placeholder="输入该次训练标题"
          maxlength="10"
          v-model="title"
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
          v-for="(item, index) in trainActionList"
          :index="index"
          :actionName="item.action_name"
          :initTrainContent="item.action_content"
          :trainHistoryId="trainHistoryId"
          :actionInstrument="item.action_instrument"
          class="train-item"
          ref="trainContent"
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
		<uni-section title="输入你本次训练的心得吧"  type="line" padding>
			<uni-easyinput type="textarea" autoHeight v-model="comment" placeholder="请输入内容"></uni-easyinput>
		</uni-section>
  </uni-popup>
  </view>
</template>
<script lang="ts">
import Vue from "vue";
import { TrainContent,AddTrainHistoryResponse } from "./apis";
import trainContent from "@/conpoments/history/trainContent.vue";
import { Action } from "@/pages/action/apis";
interface ActionDetail {
  action_name: string;
  action_content: TrainContent[];
  consume_time: number;
  action_instrument: string;
  action_id: number;
}

export default Vue.extend({
  components: { trainContent },
  data() {
    var timer: any;
    var trainActionList: ActionDetail[] = [];
    return {
      title: "",
      comment: "该次训练备注",
      consume_time: 0,
      format_consume_time: "00:00:00",
      started: false,
      timer,
      trainActionList,
      trainHistoryId:0,
    };
  },
  // MOU() {
  //   console.log("onShow");
  // },
  mounted() {
    console.log("mounted....");
  },
  watch: {
    consume_time: function (val) {
      this.format_consume_time = this.formatConsumeTime(val);
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
      if (this.title == "") {
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
          this.consume_time += 1;
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
      let ele = this.$refs["popup"] as any
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
      console.log("选择了动作", item);
      let actionDetail: ActionDetail = {
        action_name: item.action_name,
        action_content: [
          {
            left_weight: "0",
            right_weight: "0",
            total_weight: "0",
            number: 0,
            finish: false,
            action_name: item.action_name,
            training_history_id: this.trainHistoryId,
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
      this.trainActionList.push(actionDetail);
    },
    onTrainActionDetailAdd(item: TrainContent) {
      console.log("onTrainContentAdd", item);
    },
    onTrainActionDetailDelete(item: TrainContent) {
      console.log("onTrainContentDelete", item);
    },
    onfinishTrain(){
      console.log("onfinishTrain");
      if (this.title == ""){
        uni.showToast({
          title: "请输入标题",
          icon: "error",
        });
        return;
      }
      if (this.timer && this.started) {
        clearInterval(this.timer);
      }
      this.started = false;
      this.consume_time = 0;
      this.title = "";
      this.trainActionList = [];
      uni.$emit("finishTrain"); //close drawer
      },
    initData(data: AddTrainHistoryResponse) {
      console.log("initData", data);
      this.title = data.train_history.title;
      this.trainHistoryId = data.train_history.id;
      this.comment=data.train_history.comment;
      this.consume_time = data.train_history.total_time
      let map:{[key:string]:TrainContent[]} = {};
      if (data.train_content){
        for (let index = 0; index < data.train_content.length; index++) {
          let item = data.train_content[index];
          if (item.action_name in map){
            map[item.action_name].push(item);
          }else{
            map[item.action_name] = [item];
          }
        }
      }
      for (const key in map) {
        if (Object.prototype.hasOwnProperty.call(map, key)) {
          const element = map[key];
          let actionDetail: ActionDetail = {
            action_name: key,
            action_content: element,
            consume_time: 0,
            action_instrument: element[0].action_instrument,
            action_id: element[0].action_id,
          };
          this.trainActionList.push(actionDetail);
        }
      }
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