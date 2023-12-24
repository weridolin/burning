<template>
  <view>
    <view class="action-search-bar">
      <uni-search-bar
        placeholder="搜索动作"
        @confirm="onSearch"
        clearButton="none"
        cancelButton="none"
        class="action-search-bar__search_bar"
      >
        <uni-icons slot="searchIcon" color="#999999" size="18" type="home" />
      </uni-search-bar>
      <view class="action-search-bar__add_button">
        <button type="default" plain="true" size="mini" @click="addNewAction">
          add
        </button>
      </view>
    </view>
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

    <!-- 添加新动作对话框 -->
    <!-- <view> -->
      <uni-popup ref="newActionInputDialog" type="dialog" style="width: 100%;">
        <uni-section class="mb-10" title="新增自定义动作" type="line">
          <uni-forms ref="newActionForm" :modelValue="newAction">
					<uni-forms-item label="名称" required class="new-action-form-item" style="{padding: 10px 10px;}">
						<uni-easyinput v-model="newAction.action_name" placeholder="请输入动作" />
					</uni-forms-item>
					<uni-forms-item label="类型" required>
						<!-- <uni-easyinput v-model="newAction.action_type" placeholder="请输入年龄" /> -->
            <uni-data-select
              v-model="newAction.action_type"
              :localdata="actiontypeList"
              :clear="false"
            ></uni-data-select>
					</uni-forms-item>
					<uni-forms-item label="描述" >
            <uni-easyinput type="textarea" v-model="newAction.action_desc" placeholder="动作描述"></uni-easyinput>
					</uni-forms-item>
          <uni-forms-item >
            <button type="primary" @click="addAction">提交</button>
					</uni-forms-item>
				</uni-forms>
        </uni-section>
      </uni-popup>
    <!-- </view> -->
  </view>
</template>
<script lang="ts">
import Vue from "vue";
import { GetActionList, Action,GetCustomActions,DeleteCustomAction,UpdateCustomAction,AddCustomAction } from "@/pages/action/apis";
import { indexOf } from "lodash";

export default Vue.extend({
  data() {
    var actionsList: { [key: string]: { [key: string]: Action[] } } = {};
    var selectActionsList: { [key: string]: Action[] } = {};
    var actiontypeList:{value:string,text:string}[] = []
    var newAction: Action = {
      uuid: "",
      action_name: "", //动作名称 比如 杠铃卧推
      action_instrument: "自定义", //器械
      action_type: "", //类型 比如胸肩背
      action_desc: "",
      action_img: "",
      action_video: "",
      id: 0,
      created_at: "",
    };
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
      newAction, //自定义动作
      actiontypeList
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
    this.getCustomActionList();
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
      this.actiontypeList = []
      for (var i = 0; i < arr.length; i++) {
        that.actiontypeList.push({value:arr[i],text:arr[i]})
      }
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
    getCustomActionList(){
      let that = this;
      GetCustomActions(
        (res) => {
          console.log("get custom action list", res);
          // this.actionsList = res;
          let data = res.data;
          for (let index = 0; index < data.length; index++) {
            const element = data[index];
            if (element.action_type in that.actionsList) {
              if (
                element.action_instrument in
                that.actionsList[element.action_type]
              ) {
                // if(that.actionsList[element.action_type][element.action_instrument].indexOf(element)!=-1)
                // { console.log("add element to custom actionList -> ",element)
                let isExist=false
                for (let index = 0; index < that.actionsList[element.action_type][element.action_instrument].length; index++) {
                  const element2 = that.actionsList[element.action_type][element.action_instrument][index];
                  if (element2.action_name==element.action_name){
                    isExist=true
                    break
                  }
                }
                if (!isExist){

                  that.actionsList[element.action_type][element.action_instrument].push(element);
                
                }

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
    onSearch(){},
    addNewAction() {
      let ele= this.$refs["newActionInputDialog"] as any
      if (ele){
        ele.open("top")
      }
    },
    addAction(){
      uni.showLoading({
        title: "提交中...",
        mask: true,
      });
      AddCustomAction(this.newAction,(res)=>{
        uni.hideLoading();
        uni.showToast({
          title: "提交成功",
          icon: "none",
        });
        this.getCustomActionList()
        let ele= this.$refs["newActionInputDialog"] as any
        if (ele){
          ele.close()
        }
      },(err)=>{
        console.log("add custom action error",err)
        uni.hideLoading();
        uni.showToast({
          title: "提交失败",
          icon: "none",
        });
      })
    }
  },

  components: {
    // cover,
  },
});
</script>

<style lang="scss" scoped>
.action-search-bar {
  display: flex;
  flex-direction: row;
  text-align: center;
  justify-content: center;

  .action-search-bar__search_bar {
    flex: 12;
    margin-left: 10px;
    // margin-right: 10px;
  }
  .action-search-bar__add_button {
    position: relative;
    padding: 12px;
    flex: 1;
    font-size: 5px;
  }
}

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
