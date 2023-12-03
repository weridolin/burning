<template>
  <view style="background-color: #fff; display: flex">
    <view style="width: 20%; background-color: #fff" class="action-type-list">
      <swiper
        :duration="500"
        :vertical="true"
        :display-multiple-items="navCount"
        :style="'height:' + hh + 'px'"
      >
        <swiper-item
          v-for="(item, type) in actionsList"
          :key="type"
          class="action-type-item"
        >
          <view
            class="action-type-item-content"
            :style="
              currentSelectType == type
                ? 'background-color:#f1eeee;'
                : 'background-color: #fff;'
            "
            @click="onNav(e, type)"
          >
            {{ type }}
          </view>
        </swiper-item>
      </swiper>
    </view>

    <scroll-view
      :scroll-y="true"
      :scroll-with-animation="true"
      :enable-back-to-top="true"
      :style="'height:' + hh + 'px'"
      :scroll-into-view="scrollintoview"
      class="actions-list"
    >
      <uni-section
        type="circle"
        titleFontSize="16px"
        padding
        v-for="(list, name) in selectActionsList"
        :title="name"
        :key="name"
      >
        <uni-grid
          :column="2"
          :highlight="true"
          @change="change"
          class="action-machine"
          :showBorder="false"
        >
          <uni-grid-item
            v-for="(item, index) in list"
            :index="index"
            :key="index"
            class="grid-item"
          >
            <view class="grid-item-box" @click="onActionClick(item)">
              <image
                :src="item.action_img"
                mode="aspectFit"
                class="cover"
                style="flex-basis: 0; flex-grow: 4"
              />
              <text style="flex-basis: 0; flex-grow: 1">{{
                item.action_name
              }}</text>
            </view>
          </uni-grid-item>
        </uni-grid>
      </uni-section>
    </scroll-view>
  </view>
</template>

<script lang="ts">
import Vue from "vue";
import { GetActionList, Action } from "@/pages/action/apis";

export default Vue.extend({
  data() {
    var actionsList: { [key: string]: { [key: string]: Action[] } } = {};
    var selectActionsList: { [key: string]: Action[] } = {};
    return {
      list: [],
      tabCur: 0,
      mainCur: 0,
      verticalNavTop: 0,
      load: true,

      heightleft: 100,

      hh: 0,
      scrolltop: 0,
      scrollintoview: "",

      mainList: [],
      navList: [],

      navCount: 0,
      navNow: 0,
      vue_all_list: [],

      my_hua_dong_num: 0,

      navScroll: "",
      selectActionsList,
      actionsList,
      currentSelectType: "",
      // showMode:"normal"
    };
  },
  props: ["showMode"],
  onLoad() {
    uni.showLoading({
      title: "加载中...",
      mask: true,
    });
  },
  onReady() {
    this._onReadyApi();
    this.getActionList();
  },
  methods: {
    change() {},
    onNav(e: any, kk: string | number) {
      console.log("选择了" + kk, e, this.actionsList);
      this.currentSelectType = kk as string;
      this.selectActionsList = this.actionsList[kk];
    },
    _onReadyApi() {
      console.log("--------------------------------");
      var windowHeight = uni.getSystemInfoSync().windowHeight;
      console.log(windowHeight);
      var that = this;
      that.hh = windowHeight;
      var arr = Object.keys(that.actionsList);
      that.navCount = arr.length;
    },
    getActionList() {
      let that = this;
      GetActionList(
        (res) => {
          console.log("get action list", res);
          // this.actionsList = res;
          let data = res.data;
          for (let index = 0; index < data.length; index++) {
            const element = data[index];
            if (element.action_type in that.actionsList) {
              if (
                element.action_instrument in
                that.actionsList[element.action_type]
              ) {
                that.actionsList[element.action_type][
                  element.action_instrument
                ].push(element);
              } else {
                that.actionsList[element.action_type][
                  element.action_instrument
                ] = [element];
              }
            } else {
              that.actionsList[element.action_type] = {};
              that.actionsList[element.action_type][element.action_instrument] =
                [element];
            }
          }
          uni.hideLoading();
          console.log(that.actionsList, ">>>");
          this._onReadyApi();
        },
        (err) => {
          console.log(err);
          uni.hideLoading();
        }
      );
    },
    onActionClick(item: Action) {
      if (this.showMode == "select") {
        console.log("onActionClick", item);
        uni.$emit("actionSelect", item);
        // uni.navigateBack();
      }
    },
  },

  components: {
    // cover,
  },
});
</script>

<style>
/* .aaaa{ */

/* } */
.fixed {
  position: fixed;
  z-index: 99;
}

.VerticalNav.nav {
  width: 200upx;
  white-space: initial;
}

.VerticalNav.nav .cu-item {
  width: 100%;
  text-align: center;
  background-color: #fff;
  margin: 0;
  border: none;
  height: 50px;
  position: relative;
}

.VerticalNav.nav .cu-item.cur {
  background-color: #f1f1f1;
}

.VerticalNav.nav .cu-item.cur::after {
  content: "";
  width: 8upx;
  height: 30upx;
  border-radius: 10upx 0 0 10upx;
  position: absolute;
  background-color: currentColor;
  top: 0;
  right: 0upx;
  bottom: 0;
  margin: auto;
}

.VerticalBox {
  display: flex;
}

.VerticalMain {
  background-color: #f1f1f1;
  flex: 1;
}

.action-type-item-content {
  width: 100%;
  height: 100rpx;
  line-height: 80rpx;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #8d8a8a;
  font-weight: blo;
}
.grid-item-box {
  margin: 5px 5px;
  background-color: #fff;
  flex-basis: 0;
  flex-grow: 1;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}
.grid-item {
  margin-bottom: 5px;
  margin-top: 5px;
  background-color: #dad5d517;
  border-radius: 15px 15px;
}
.action-type-list {
  /* border-right-style: solid; */
  margin-top: 10px;
  margin-left: 5px;
}

.cover {
  margin-top: 30px;
  margin-left: 50px;
  margin-right: 50px;
  margin-bottom: 20px;
}
</style>
