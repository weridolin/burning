<template>
  <uni-card
    :is-shadow="true"
    :title="
      trainDetail.train_history.title
        ? trainDetail.train_history.title
        : '未定义标题'
    "
    :sub-title="
      trainDetail.train_history.comment
        ? trainDetail.train_history.comment
        : '本次训练似乎没什么心得...'
    "
    :extra="trainDetail.train_history.finish ? '已完成' : '进行中...'"
    :thumbnail="
      trainDetail.train_history.finish
        ? '/static/icons/finish.png'
        : '/static/icons/doing.png'
    "
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
        style="display: flex; flex-flow: row wrap; gap: 10px; margin-left: 10px"
      >
        <uni-title
          type="h5"
          :title="
            item.left_weight + 'kg*' + item.right_weight + 'kg*' + item.number
          "
          v-for="(item, index) in briefInfoList"
          :key="index"
        >
        </uni-title>
      </view>
    </uni-section>
  </uni-card>
</template>
<script lang="ts">
import {
  TrainHistory,
  TrainContent,
} from "../../pages/history/apis";
import UniSection from "../../uni_modules/uni-section/components/uni-section/uni-section.vue";
import Vue from "vue";
interface TrainHistoryDetail {
  train_history: TrainHistory;
  train_content: TrainContent[];
}
interface BriefItem {
  action_name: string;
  number: string;
  left_weight: string;
  right_weight: string;
  total_weight: string;
}
export default Vue.extend({
  props: ["trainDetail","onCardClick"],
  methods:{
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
  }
});
</script>