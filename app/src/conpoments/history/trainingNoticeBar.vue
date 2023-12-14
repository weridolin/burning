<template>
  <uni-notice-bar
    show-icon
    :scrollable="false"
    :text="text"
    @click="switchToTrainDetail"
    v-show="isTrainingPlanExist"
  >
  </uni-notice-bar>
</template>

<script lang="ts">
import Vue from "vue";
import { getDoingTrain } from "@/store/local";
import { isLogin } from "@/store/local";
export default Vue.extend({
  data() {
    return {
      isTrainingPlanExist: true,
      text:""
    };
  },
  onShow() {
    console.log("check training status...")
    this.refreshStatus();
  },
  onLoad() {
  },
  created() {},
  methods: {
    switchToTrainDetail() {
      console.log("跳转到训练详情页");
      uni.setStorageSync("showDetail",true)
      uni.switchTab({
            url: "/pages/history/index",
          });
    },
    refreshStatus() {
      console.log("refresh notice bar..",getDoingTrain() )
      this.text="当前有正在进行的训练计划，点击查看详情"
      if (!isLogin()){
        this.isTrainingPlanExist=false
        return
      }
      if (getDoingTrain() == null) {
        this.isTrainingPlanExist = false;
      }else{
      this.isTrainingPlanExist = true;}
    },
  },
});
</script>