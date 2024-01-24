<template>
  <uni-notice-bar
    show-icon
    :scrollable="false"
    :text="text"
    @click="switchToTrainDetail"
    v-if="isTrainingPlanExist"
  >
  </uni-notice-bar>
</template>

<script lang="ts">
import Vue from "vue";
import { isLogin } from "@/store/local";
import { getAllDoingTrain } from "../../store/local";
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
      // uni.setStorageSync("showDetail",true)
      //判断当前页
      // let pages = getCurrentPages();
      // if (pages[pages.length - 1].$page.fullPath == "/pages/history/index") {
      //   this.$forceUpdate();
      // }else{
      uni.switchTab({
          url: "/pages/history/index",
        });
      // }

    
    },
    refreshStatus() {
      //
      // let date = getDate(new Date(),0).fullDate
      // console.log("refresh notice bar..",getDoingTrain(date) )
      console.log("refresh notice bar..",getAllDoingTrain())
      this.text="存在未完成的训练计划，点击查看详情"
      if (!isLogin()){
        this.isTrainingPlanExist=false
        return
      }
      if (getAllDoingTrain() == null) {
        this.isTrainingPlanExist = false;
      }else{
      this.isTrainingPlanExist = true;}
    },
  },
});
</script>