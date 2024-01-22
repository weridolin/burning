<template>
  <view @click="showMenu">
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
        v-for="(briefInfoList, actionName) in getShowList"
        :key="actionName"
      >
        <view
          class="brief-info-content-list"
          style="
            display: flex;
            flex-flow: row wrap;
            gap: 10px;
            margin-left: 10px;
          "
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
      <view class="more" @click.stop="showMore">
        {{ showDetail ? "收起" : "查看更多" }}
      </view>
    </uni-card>
  </view>
</template>
<script lang="ts">
import { TrainHistory, TrainContent } from "../../pages/history/apis";
import UniSection from "../../uni_modules/uni-section/components/uni-section/uni-section.vue";
import Vue from "vue";

interface BriefItem {
  action_name: string;
  number: string;
  left_weight: string;
  right_weight: string;
  total_weight: string;
}
export default Vue.extend({
  data() {
    var TrainContentBriefInfoTotal: { [key: string]: BriefItem[] } = {} as {
      [key: string]: BriefItem[];
    };
    var TrainContentBriefInfoTop2: { [key: string]: BriefItem[] } = {} as {
      [key: string]: BriefItem[];
    };
    return {
      showDetail: false,
      TrainContentBriefInfoTotal,
      TrainContentBriefInfoTop2,
    };
  },
  props: ["trainDetail"],
  onReady() {
    this.getTrainContentBriefInfo(this.trainDetail.train_content);
  },
  watch: {
  },
  computed: {
    getShowList(): {[key: string]: BriefItem[]}{
      if (this.showDetail) {
        return this.TrainContentBriefInfoTotal;
      } else {
        return this.TrainContentBriefInfoTop2;
      }
    },
  },
  methods: {
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
      this.TrainContentBriefInfoTotal = res;
      let index = 0;
      for (const [key, value] of Object.entries(res)) {
        this.TrainContentBriefInfoTop2[key] = value;
        index++;
        if (index >= 1) {
          break;
        }
      }
      console.log("getTrainContentBriefInfo", res);
      return res;
    },
    showMore() {
      this.showDetail = !this.showDetail;
    },
    showMenu() {
      console.log("showMenu");
      this.$emit("showMenu");
    },
  },
});
</script>

<style lang="scss" scoped>
.more {
  font-size: smaller;
  justify-content: center;
  text-align: center;
}
</style>