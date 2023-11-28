<template>
  <view id="content">
    <!-- navScroll -->
    <!-- navNow -->
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
              <view class="grid-item-box">
                <!-- <uni-icons type="image" :size="30" color="#777" class="cover" style="flex-basis: 0; flex-grow: 4"/> -->
                <image
                  :src="item.actions_img"
                  mode="aspectFit"
                  class="cover"
                  style="flex-basis: 0; flex-grow: 4"
                />
                <text style="flex-basis: 0; flex-grow: 1">{{ item.action_name }}</text>
              </view>
            </uni-grid-item>
          </uni-grid>
        </uni-section>
      </scroll-view>
    </view>
  </view>
</template>

<script>
import cover from "@/conpoments/actions/cover";
// import ActionCover from "@/components/actions/cover";
import {GetActionList} from "./apis";
import type {Action} from "@/pages/action/apis.ts";
export default {
  data() {

    var action_list:{[key:string]:{[key:string]:Action[]}} 

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
      selectActionsList: {},
      actionsList,
      currentSelectType: "",
    };
  },

  onLoad() {

    uni.showLoading({
      title: "加载中...",
      mask: true,
    });

  },
  onReady() {
    this._onReadyApi();
  },
  methods: {
    onNav(e, kk) {
      console.log("选择了" + kk);
      this.currentSelectType = kk;
      this.selectActionsList = this.actionsList[kk];

    },
    _onReadyApi() {
      console.log("--------------------------------");
      var windowHeight = uni.getSystemInfoSync().windowHeight;
      console.log(windowHeight);
      var that = this;
      that.hh = windowHeight;
      var arr= Object.keys(that.actionsList); 
      that.navCount = arr.length
    },
    getActionList(){
      GetActionList().then((res) => {
        console.log("get action list",res);
        // this.actionsList = res;
        for (let index = 0; index < res.length; index++) {
          const element = res[index];
          if (element.action_type in this.actionsList.keys()) {
            if element.action_instrument in this.actionsList[element.action_type].keys() {
              this.actionsList[element.action_type][element.action_instrument].push(element);
            } else {
              this.actionsList[element.action_type][element.action_instrument] = [element];
            }
          } else {
            this.actionsList[element.action_type][element.action_instrument] = [element];
          }
        }
      });
    }
  },
  components: {
    cover,
  },
};
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
