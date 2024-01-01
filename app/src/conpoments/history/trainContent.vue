<template>
  <view class="detail-item">
    <uni-title
      type="h4"
      :title="actionName"
      class="detail-item-title"
    ></uni-title>

    <view class="detail-item-group detail-item-headline">
      <view class="detail-item-group-index item-headline">
        <text class="uni-h6 detail-item-group-left-weight__label"></text>
      </view>
      <view class="detail-item-group-left-weight item-headline">
        <text class="uni-h6 detail-item-group-left-weight__label"
          >左侧(kg)</text
        >

      </view>
      <view class="detail-item-group-right-weight item-headline">
        <text class="uni-h6 detail-item-group-right-weight__label"
          >右侧(kg)</text
        >

      </view>
      <view class="detail-item-group-count item-headline">
        <text class="uni-h6 detail-item-group-count__label">次数</text>
      </view>
      <view class="detail-item-group-finish item-headline">
        <label class="detail-item-group-finish__check">
        </label>
      </view>
      <view class="detail-item-group-menu">
        <view class="demo-view light"></view>
      </view>
    </view>

    <view
      class="detail-item-group detail-item-content"
      v-for="(item  , index) in trainContentList"
      :key="index"
    >
      <view class="detail-item-group-index">
        <uni-tag
          class="item-content-index"
          :text="index + 1"
          type="primary"
          style="margin-top: 5px"
        />
      </view>
      <view class="detail-item-group-left-weight">
        <uni-easyinput
          class="detail-item-group-left-weight__input"
          placeholder="0"
          v-model="item.left_weight"
          :clearable="false"
          trim="all"
        ></uni-easyinput>
      </view>
      <view class="detail-item-group-right-weight">
        <uni-easyinput
          class="detail-item-group-right-weight__input"
          placeholder="0"
          v-model="item.right_weight"
          :clearable="false"
          trim="all"
        ></uni-easyinput>
      </view>
      <view class="detail-item-group-count">
        <uni-easyinput
          class="detail-item-group-count__input"
          placeholder="0"
          v-model="item.number"
          :clearable="false"
          trim="all"
          type="number"
        ></uni-easyinput>
      </view>
      <view class="detail-item-group-finish">
        <label class="detail-item-group-finish__check">
          <checkbox-group @change="change">
            <checkbox
              value="cb"
              :checked="item.finish"
              class="detail-item-group-finish__checkbox"
            />
          </checkbox-group>
        </label>
      </view>
      <view class="detail-item-group-menu">
        <uni-icons
          class="detail-item-group-menu__more"
          type="minus"
          size="20"
          @click="deleteActionContent(item)"
        ></uni-icons>
      </view>
    </view>

    <view class="detail-item-add-btn">
      <uni-icons type="plus" size="20" @click="addActionContent"></uni-icons>
    </view>
  </view>
  <!-- </view> -->
</template>

<script lang="ts">
import Vue from "vue";
import {
  TrainContent,
} from "@/pages/history/apis";
export default Vue.extend({
  props: [
    "initTrainContent",
    "actionName",
    "index",
    "trainHistoryId",
    "actionInstrument",
    "actionId",
  ],
  data() {
    return {
      // action_name: "动作名称",
      trainContentList: this.initTrainContent
    };
  },
  watch: {
    trainContentList: {
      handler: function (val: TrainContent[]) {
        let value = JSON.parse(JSON.stringify(val));
        this.$emit("trainContentListUpdate",value,this.index);
      },
      deep: true,
    },
  },
  methods: {
    change(e: object, item: any) {
      item.finish = !item.finish;
    },
    getWidth() {
      uni.getSystemInfoSync().windowWidth;
    },
    addActionContent() {
      let trainContent: TrainContent = {
        left_weight: "0",
        right_weight: "0",
        total_weight: "0",
        number: "0",
        finish: false,
        action_name: this.actionName,
        training_history_id: this.trainHistoryId,
        action_instrument: this.actionInstrument,
        consume_time: 0,
        action_id: this.actionId,
      };
      this.trainContentList.push(trainContent);
      console.log("addActionContent", this.trainContentList);
    },
    // updateActionContent(item:TrainContent){
    //   UpdateTrainHistoryContent(this.trainHistoryId,item.id as number,item,(res:any)=>{
    //     console.log("update action success", res);
    //   },(err)=>{
    //     console.log("update action content error - >", err);
    //     uni.showToast({
    //       title: "更新失败.",
    //       icon: "error",
    //       duration: 2000,
    //     });
    //   })
    // },
    deleteActionContent(item:TrainContent){
      this.trainContentList.splice(this.trainContentList.indexOf(item),1);
      this.$emit("deleteActionContent",item);
    }
  },
});
</script>
<style lang="scss" scoped>
.detail-item-title {
  margin-left: 5px;
}

.detail-item {
  // .detail-item-title {
  //   margin-left: 5px;
  // }

  .detail-item-headline {
    .item-headline {
      margin-left: 10px;
    }
  }
  .detail-item-group {
    margin-left: 2%;
    display: flex;
    flex-direction: row;

    .detail-item-group-index {
      flex: 1;
      text-align: center;
      justify-content: center;
      margin-top: inherit;
    }
    .detail-item-group-left-weight {
      flex: 4;
      margin-left: 10px;
    }
    .detail-item-group-right-weight {
      flex: 4;
      margin-left: 10px;
    }
    .detail-item-group-count {
      flex: 4;
      margin-left: 10px;
    }
    .detail-item-group-finish {
      flex: 3;
      margin: inherit;
      text-align: center;
      margin-top: 6px;

      .detail-item-group-finish__check {
        // font-size: 12px;
        color: #8f8f94;
      }
      .detail-item-group-finish__checkbox {
        transform: scale(1.5);
      }
    }
    .detail-item-group-menu {
      flex: 12;
      margin-left: 3px;
    }
    // .{
    //   margin-top: 5px;
    // }
  }

  .detail-item-content {
    margin-top: 10px;
    .item-content-index {
      margin-top: 15px;
    }
    .detail-item-group-menu__more {
      position: absolute;
      right: 20px;
    }
  }

  .detail-item-add-btn {
    margin-top: 15px;
    text-align: center;
    justify-content: center;
  }
}



</style>