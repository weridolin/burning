<template>
  <uni-notice-bar
    show-icon
    scrollable
    text="当前有正在进行的训练计划，点击查看详情"
    @click="switchToTrainDetail"
    v-show="isTrainingPlanExist"
  >
  </uni-notice-bar>
</template>

<script lang="ts">
import Vue from "vue";
import {getDoingTrain} from "@/store/local"
import { get } from "lodash";
export default Vue.extend({
  data() {
    return {
      isTrainingPlanExist: false,
      // trainingPlan: getDoingTrain,
    };
  },
  onShow() {
    this.refreshStatus()
  },
  created() {
  },
  methods:{
    switchToTrainDetail(){
      console.log("跳转到训练详情页")
    },
    refreshStatus(){
      if (getDoingTrain() == null){
        this.isTrainingPlanExist = false
      }
      this.isTrainingPlanExist = true
    }
  },
});
</script>