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
			<uni-card :is-shadow="true" v-show="trainDetail!=''">
				<text class="uni-body">{{ trainDetail }}</text>
			</uni-card> 
		<!-- </uni-section> -->
  </view>
</template>

<script lang="ts">
import Vue from "vue";
import {getDate,GetTrainHistory,GetTrainHistoryDetail,TrainHistoryBrief,TrainHistory} from "./apis";
function getTrainDetail(date:string){
  return "训练详情"
}


export default Vue.extend({
  data() {
    var transHistory:TrainHistoryBrief[]=[]
    var trainHistoryMap:{[key: string]:TrainHistory}={}
    return {
      title: "训练记录",
      transHistory,
      trainHistoryMap,
      trainDetail:"当天暂无训练日志"
    };
  },
  onLoad() {},
  methods: {
    change(e: any) {
      console.log("select date is " + e.fulldate);
      this.trainDetail = getTrainDetail(e.fulldate)
      let trainHistory = this.trainHistoryMap[e.fulldate];
      if (trainHistory) {
        GetTrainHistoryDetail(trainHistory.id).then((res) => {
          console.log("get change history detail -> ",res);
          // this.trainDetail = res;
        })
      } else {
        this.trainDetail = "当天暂无训练日志";
      }
    },
    getHistory() {
      // 获取历史记录
      GetTrainHistory().then((res) => {
        console.log(res);
        for (let i = 0; i < res.length; i++) {
          let item = res[i];
          this.trainHistoryMap[item.created_at] = item;
          this.transHistory.push({
            date: item.created_at,
            info: item.title,
          });
        }
      });
    },
    
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
</style>
