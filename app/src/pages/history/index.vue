<template>
  <view>
  <view class="content">
    <!-- <trainingNoticeBar
      style="width: 100%"
      ref="trainingNoticeBar"
    ></trainingNoticeBar> -->
    <!-- 日历模块 -->
    <view >
      <uni-calendar
        :insert="true"
        :lunar="true"
        :start-date="'2019-3-2'"
        :end-date="'3099-5-20'"
        @change="change"
        @monthSwitch ="switchMonth"
        :selected="transHistory"
      />
    </view>

    <!-- 当天训练日志 --> 
    <uni-section title="当天训练详情" type="line" style="width: 100%">
      <view v-if="!trainDetailList || trainDetailList.length == 0">
        <uni-card :is-shadow="true">
          <text class="uni-body">当天无训练日志</text>
        </uni-card>
      </view>
      <view style="width: 100%" v-if="trainDetailList && trainDetailList.length > 0">
        <TrainHistoryBriefCard
          v-for="(trainDetail, indextd) in trainDetailList"
          :key="indextd"
          :trainDetail="trainDetail"
          @showMenu = "onTrainCardClick(indextd)"
          class="train-detail-card"
        >
        </TrainHistoryBriefCard>
      </view>
    </uni-section>

    <!-- 当天饮食记录 -->
    <uni-section title="当天饮食详情" type="line" style="width: 100%">
      <view v-if="!dietHistoryList || dietHistoryList.length == 0">
        <uni-card :is-shadow="true">
          <text class="uni-body">当天无饮食日志</text>
        </uni-card>
      </view>
      <view class="diet-card"
        ref = "diet-card-list" 
        v-for="(item,index) in dietHistoryList"
        :key="index"
        @click.native="onDietCardClick(index)"
        >
        <dietCard 
          :dietContentItemProp="dietHistoryList[index]"
          >
        </dietCard>
      </view>
    </uni-section>

    <!-- 右下角悬浮按钮 -->
    <view class="button" >
      <uni-fab
        ref="fab"
        :pattern="pattern"
        :content="content"
        :horizontal="horizontal"
        :vertical="vertical"
        :direction="direction"
        @trigger="trigger"
        v-show="status != 'created'"
      />
    </view>

    <!-- 新建记录页面 -->
    <uni-drawer
      ref="newRecord"
      mode="right"
      :width="getWidth()"
      :mask-click="true"
      v-show="status == 'created' || status == 'edit'"
      @change="drawChange"
    >
      <scroll-view style="height: 100%" scroll-y="true">
        <NewRecord ref="newRecordEditPage"> </NewRecord>
      </scroll-view>
    </uni-drawer>

    <!-- 卡片点击 菜单弹框 -->
    <view >
      <uni-popup
        ref="card-popup"
        background-color="#fff"
      >
        <view class="card-popup__btn_group">
          <view class="uni-popup-share">
            <view class="uni-share-title"
              ><text class="uni-share-title-text">更多操作</text></view
            >
            <view class="uni-share-content">
              <view class="uni-share-content-box">
                <view
                  class="uni-share-content-item"
                  v-for="(item, index) in bottomData"
                  :key="index"
                  @click.stop="onCardBtnClick(item, index, selectedCardId)"
                >
                  <image
                    class="uni-share-image"
                    :src="item.icon"
                    mode="aspectFill"
                  ></image>
                  <text class="uni-share-text">{{ item.text }}</text>
                </view>
              </view>
            </view>
          </view>
        </view>
      </uni-popup>
    </view>

    <!-- 饮食记录编辑弹框 -->
    <uni-popup
      ref="diet-edit-form"
      title="修改饮食记录"
      background-color="#fff"
      style="z-index: 1;"
      >
      <dietContent 
        ref="diet-edit-form-dialog"
        @dietContentUpdated="onDietUpdate"
        @dietContentCreated="onDieCreated"
        :dietContentItemProp="dietHistoryList[selectedItemIndex]"
        > 
      </dietContent>
    </uni-popup>
   
  </view>
    <shareCard ref="shareCard" @cancel="onShareCardCancel" :posterData.sync="posterData" @previewImage='previewHandle' />
  </view>
</template>

<script lang="ts">
import Vue from "vue";
import NewRecord from "./NewRecord.vue";
import TrainHistoryBriefCard from "../../conpoments/history/trainHistoryBriefCard.vue";
import {
  getDate,
  GetTrainHistory,
  TrainHistoryBrief,
  AddTrainHistory,
  AddTrainHistoryResponse,
  DeleteTrainHistory,
  TrainHistoryDetail,
  DeleteDietHistory,
  getStartOfMonth,
  getMonthStartAndEndDate
} from "./apis";
import UniSection from "../../uni_modules/uni-section/components/uni-section/uni-section.vue";
import { getDoingTrain, clearDoingTrain } from "@/store/local";
import { isLogin } from "../../store/local";
import trainingNoticeBar from "@/conpoments/history/trainingNoticeBar.vue";
import dietContent from "@/conpoments/history/dietContent.vue";
import dietCard from "@/conpoments/history/dietCard.vue";
import { DietContentItem,GetDietHistory } from "@/pages/history/apis";
import shareCard from "@/conpoments/history/shareCard.vue";


export default Vue.extend({
  data() {
    var transHistory: TrainHistoryBrief[] = []; // 日志组件展示简要信息
    var trainHistoryMap: { [key: string]: TrainHistoryDetail[] } = {}; // {date: 训练列表}
    var trainDetailList: TrainHistoryDetail[] = []; // 选中的当天的所有训练记录列表
    var dietHistoryList: DietContentItem[] = [] // 选中的当天所有饮食记录
    return {
      title: "训练记录",
      transHistory,
      trainHistoryMap, 
      trainDetailList,
      horizontal: "left",
      vertical: "bottom",
      direction: "horizontal",
      pattern: {
        color: "#7A7E83",
        backgroundColor: "#fff",
        selectedColor: "#007AFF",
        buttonColor: "#007AFF",
        iconColor: "#fff",
      },
      is_color_type: false,
      
      // 悬浮按钮菜单项
      content: [
        {
          iconPath: "/static/icons/add.png",
          selectedIconPath: "/static/icons/add.png",
          text: "训练",
          active: false,
        },
        {
          iconPath: "/static/icons/add.png",
          selectedIconPath: "/static/icons/add.png",
          text: "饮食",
          active: false,
        },
      ],
      status: "",
      trainHistoryId: 0, //当前选中的历史记录ID
        
      //弹出框菜单
      bottomData: [
        {
          text: "分享",
          icon: "../../static/icons/share.png",
          name: "share",
        },
        {
          text: "删除",
          icon: "../../static/icons/del.png",
          name: "del",
        },
        {
          text: "储存为模板",
          icon: "../../static/icons/saveTem.png",
          name: "saveAsTem",
        },
        {
          text: "编辑",
          icon: "../../static/icons/edit.png",
          name: "edit",
        },
      ],

      selectedCardId: 0,
      selectedItemIndex:0,
      dietHistoryList,
      selectDate:"",
      selectCardType:"train", // train or diet
      
      //分享记录或者饮食
      shareImage:"",      
      posterData: {
        type:"",
        content:[],
        poster: {
          //根据屏幕大小自动生成海报背景大小
          // url: "../../static/icons/shareImageBG.png",//图片地址
          r: 10, //圆角半径
          w: 300, //海报宽度
          h: 600, //海报高度
          p: 20 //海报内边距padding
        },
        mainImg: {
          //海报主商品图
          url: "../../static/icons/shareImageBG.png",//图片地址
          r: 10, //圆角半径
          w: 300, //宽度
          h: 400 //高度
        },
        title: {
          //商品标题
          text: "今日上新水果，牛奶草莓，颗粒饱满，每盒 200g",//文本
          fontSize: 16, //字体大小
          color: "#000", //颜色
          lineHeight: 25, //行高
          mt: 20 //margin-top
        }
      }
    };
  },
  onLoad() {
    uni.$on("finishTrain", () => {
      console.log("finishTrain");
      let ele = this.$refs["newRecord"] as any;
      if (ele) {
        ele.close();
      }
      this.status = "";
      clearDoingTrain()
      this.refreshHistory();
    });
    uni.$on("minimizeDrawer", () => {
      console.log("最小化训练记录");
      let ele = this.$refs["newRecord"] as any;
      if (ele) {
        ele.close();
      }
      this.status = "";
      this.refreshHistory();
    });
  },
  onUnload() {
    uni.$off("finishTrain");
    uni.$off("minimizeDrawer");
  },
  onShow() {
    console.log("on train index page show");
    if (isLogin()) {
      //登录的情况直接拉取最新的训练记录
      this.refreshHistory();
    } else {
      this.trainDetailList = [];
      this.trainHistoryMap = {};
      this.transHistory = [];
      this.dietHistoryList = []
    }
    //从主页直接跳转到训练详情
    if (uni.getStorageSync("showDetail")){
      console.log("open training history detail")
      this.status = "created";
      let ele1 = this.$refs["newRecordEditPage"] as any;
      ele1.initData(null);
      let ele = this.$refs["newRecord"] as any;
      if (ele) {
        ele.open();
      }
      uni.removeStorageSync('showDetail')
    } 
  },
  methods: {
    updateNoticeBar() {
      let el = this.$refs["trainingNoticeBar"] as any;
      if (el) {
        el.refreshStatus();
      }
    },
    updateDietHistory(){
      if (this.selectDate == ""){
        this.selectDate = getDate(new Date(), 0).fullDate
      }
      // 获取当天的饮食记录
      let startTime = this.selectDate
      let endTime = getDate(new Date(startTime),1).fullDate
      GetDietHistory(
        startTime,
        endTime,
        (res) => {
          console.log("get diet history", res);
          this.dietHistoryList = res.data
        },
        (err) => {
          console.log("get diet history err", err);
        }
      )
    },
    refreshHistory() {
      uni.showLoading({
        title: "获取训练记录中...",
      });
      this.trainDetailList = [];
      this.trainHistoryMap = {};
      this.transHistory = [];
      this.dietHistoryList = []

      if (this.selectDate == ""){
        var start_time = getDate(getStartOfMonth(),-1).fullDate
        var end_time = getDate(new Date(),1).fullDate      
      }else{
        var [start_time,end_time] = getMonthStartAndEndDate(this.selectDate)
        // console.log("select date is ",this.selectDate,"start_time",start_time,"end_time",end_time)
      }

      //获取当天的饮食记录
      this.updateDietHistory()
      this.getHistory(start_time,end_time);
      //当前正在训练通知栏
      this.updateNoticeBar()  
    },
    change(e: any) {
      let _trainHistory = this.trainHistoryMap[e.fulldate];
      this.selectDate = e.fulldate
      console.log("select date is " + e.fulldate,"trainHistory",_trainHistory);
      if (_trainHistory) {
        this.trainDetailList = _trainHistory;
      } else {
        this.trainDetailList = [];
      }
      // 获取当天的饮食记录
      this.updateDietHistory()
    },
    switchMonth(e:any){
      this.trainDetailList = [];
      this.trainHistoryMap = {};
      this.transHistory = [];
      this.dietHistoryList = []
      console.log("switch month",e)
      let start_time = getDate(new Date(e.year,e.month-1),-1).fullDate
      let end_time = getDate(new Date(e.year,e.month),0).fullDate
      this.getHistory(start_time,end_time);
    },
    getHistory(start_time:string,end_time:string) {
      console.log("get train history from ",start_time," to " ,end_time)
      // 获取历史记录
      GetTrainHistory(
        start_time,
        end_time,
        (res) => {
          console.log("get train history", res);
          uni.hideLoading();
          if (!res.data) {
            return;
          }
          for (let i = 0; i < res.data.length; i++) {
            let item = res.data[i];
            if (!item.train_history.finish){
                //更新下当前正在训练的内容到当天的训练详情里面,因为当天未完成的训练
                // 只有在完成才会更新到后台
                // TODO 非当天的训练记录？
                let training = getDoingTrain(item.train_history.created_at)  as {
                  title:"",
                  consume_time:"",
                  comment:""
                }
                if (training){
                  // content会在进入详情页更新
                  item.train_history.title = training.title
                  item.train_history.total_time = training.consume_time
                  item.train_history.comment = training.comment
                }
            }
            let n_date = getDate(new Date(item.train_history.created_at),0).fullDate;
            if (n_date in this.trainHistoryMap) {
              this.trainHistoryMap[n_date].push(item);
            } else {
              this.trainHistoryMap[n_date] = [item];
            }
            // 日历组件简要数据展示
            this.transHistory.push({
              date: getDate(n_date, 0).fullDate,
              info:
                (item.train_history.title == ""||item.train_history.finish==false)
                  ? "进行中.."
                  : item.train_history.title,
            });
          }
          //默认更新下当天的训练详情
          this.trainDetailList  = this.trainHistoryMap[this.selectDate];
          console.log(`当天(${getDate(new Date(), 0).fullDate})训练详情 ->`,this.trainDetailList,this.trainHistoryMap);

        },
        (err) => {
          console.log("get train history err", err);
          uni.hideLoading();
        }
      );
    },

    /** 点击新建按钮中的菜单 */
    trigger(e: any) {
      console.log("添加训练记录,当前登录状态 ->", isLogin());
      if (this.content[e.index].text == "训练") {
        if (!isLogin()) {
          uni.showToast({
            title: "请先登录",
            icon: "error",
            duration: 2000,
          });
          return;
        }
        //先判断是否有未完成的训练记录 #
        uni.showLoading({
          title: "初始化训练中...",
        });
        AddTrainHistory(
          { comment: "",
            total_time: 0,
            title: "",
            finish: false,
            force: false,
            created_at: this.selectDate
          },
          (res) => {
            res = res.data as AddTrainHistoryResponse;
            if (res.existed) {
              uni.showModal({
                title: "",
                content: "当前有未完成的训练,请先完成再添加!",
                showCancel: false,
                success: function (_) {
                  console.log("打开未完成的训练记录 -> ", res);
                },
              });
            }
            console.log("add train history", res);
            this.trainHistoryId = res.train_history.id;
            this.status = "created";
            let ele1 = this.$refs["newRecordEditPage"] as any;
            ele1.initData(res,this.selectDate);
            let ele = this.$refs["newRecord"] as any;
            if (ele) {
              ele.open();
            }
            uni.hideLoading();
          },
          (err) => {
            uni.showToast({
              title: "添加历史记录失败",
              icon: "error",
            });
            console.log("add train history err", err);
            uni.hideLoading();
          }
        );
        
      }
      else if (this.content[e.index].text == "饮食") {
        let ele = this.$refs["diet-edit-form"] as any
            if (ele){
              // 设置 编辑界面的值
              let dietForm = this.$refs["diet-edit-form-dialog"] as any
              if (dietForm){
                dietForm.status="create"
                dietForm.dietContentItem={}
                // dietForm.dietContentItem=this.dietHistoryList[index]
              }
              ele.open("bottom")
        }
      }
    },

    drawChange(e: any) {
      if (!e) {
        this.status = "";
        uni.showTabBar();
      } else {
        uni.hideTabBar();
      }
    },
    getWidth() {
      let width = uni.getSystemInfoSync().windowWidth;
      console.log(width);
      return width;
    },

    OnActionSelect(item: any) {
      console.log("OnActionSelect", item);
      if (this.$refs.newRecordEditPage) {
        let ele = this.$refs.newRecordEditPage as any;
        ele.ActionSelect(item);
      }
    },

    onTrainCardClick(itemIndex: number) {
      console.log("onTrainCardClick", this.trainDetailList[itemIndex]);
      this.selectCardType = "train"
      this.selectedCardId =  this.trainDetailList[itemIndex].train_history.id;
      this.selectedItemIndex = itemIndex;
      let ele = this.$refs["card-popup"] as any;
      if (ele) {
        ele.open("bottom");
      }
    },
    onDietCardClick(itemIndex: number) {
      console.log("onDietCardClick", this.dietHistoryList[itemIndex]);
      this.selectCardType = "diet"
      this.selectedCardId =  this.dietHistoryList[itemIndex].id;
      this.selectedItemIndex = itemIndex;
      let ele = this.$refs["card-popup"] as any;
      if (ele) {
        ele.open("bottom");
      }
    },

    onCardBtnClick(item: any, index: any, recordId: number) {
      switch (item.name) {
        case "share":
          // uni.showToast({
          //   title: "暂不支持",
          //   icon: "error",
          // });
          if (this.selectCardType=="train"){
            this.posterData.type = "train"
            this.posterData.content = this.trainDetailList[this.selectedItemIndex] as any
          }else if (this.selectCardType=="diet"){
            this.posterData.type = "diet"
            this.posterData.content = this.dietHistoryList[this.selectedItemIndex] as any
          }
          console.log("share card",this.posterData)
          this.createsShareImage()
          let ele = this.$refs["card-popup"] as any;
          if (ele) {
            ele.close();
          }
          break;
        case "del":
          uni.showModal({
            title: "",
            content: "确定删除该记录吗?",
            success: (res) => {
              if (res.confirm) {
                uni.showLoading({
                  title: "删除中...",
                });
                if (this.selectCardType=="train"){ 
                    DeleteTrainHistory(
                      recordId,
                      (res) => {
                        uni.hideLoading();
                        uni.showToast({
                          title: "删除成功",
                          icon: "success",
                          duration: 2000,
                        });
                        let ele = this.$refs["card-popup"] as any;
                        if (ele) {
                          ele.close();
                        }
                        //删除的是进行中的训练,删除本地缓存
                        if (!this.trainDetailList[this.selectedItemIndex].train_history.finish) {
                          clearDoingTrain();
                        }

                        this.refreshHistory();
                      },
                      (error) => {
                        console.log("delete train history err", error)
                        uni.hideLoading();
                        let ele = this.$refs["card-popup"] as any;
                        if (ele) {
                          ele.close();
                        }
                      }
                    );
                }else{
                  DeleteDietHistory(
                    recordId,
                    (res) => {
                      uni.hideLoading();
                      uni.showToast({
                        title: "删除成功",
                        icon: "success",
                        duration: 2000,
                      });
                      let ele = this.$refs["card-popup"] as any;
                      if (ele) {
                        ele.close();
                      }
                      // this.refreshHistory();
                      this.updateDietHistory()
                    },
                    (error) => {
                      console.log("delete diet history err", error)
                      uni.hideLoading();
                      let ele = this.$refs["card-popup"] as any;
                      if (ele) {
                        ele.close();
                      }
                    }
                  );
                }
              } else if (res.cancel) {
                console.log("用户点击取消");
              }
            },
          });
          break;
        case "saveAsTem":
          if (!this.trainDetailList[this.selectedItemIndex].train_history.finish) {
            uni.showToast({
              title: "未完成的训练记录不能保存为模板",
              icon: "error",
            });
            return;
          } else {
            // todo 保存为模板
            uni.showToast({
              title: "暂不支持",
              icon: "error",
            });
            return;
          }
        case "edit": {
          if (this.selectCardType=="diet"){
            let ele = this.$refs["diet-edit-form"] as any
            if (ele){
              // 设置 编辑界面的值
              let dietForm = this.$refs["diet-edit-form-dialog"] as any
              if (dietForm){
                dietForm.status="update"
                // dietForm.dietContentItem=this.dietHistoryList[index]
              }
              ele.open("center")
            }
            return;
          }else{
            //判断当前编辑的是否为已经完成训练记录
            if (this.trainDetailList[this.selectedItemIndex].train_history.finish) {
              this.status = "edit";
            }else{
              this.status = "created";
            }
            let ele = this.$refs["newRecordEditPage"] as any;
            if (ele) {
              ele.editData(
                this.trainDetailList[this.selectedItemIndex].train_history,
                this.trainDetailList[this.selectedItemIndex].train_content
              );
            }
            let ele2 = this.$refs["newRecord"] as any;
            if (ele2) {
              ele2.open();
            }
          }
        }
        default:
          break;
      }
      console.log(item, index);
      let ele = this.$refs["card-popup"] as any;
      if (ele) {
        ele.close();
      }
    },
    onDietUpdate(item:DietContentItem){
      // this.dietHistoryList[this.selectedItemIndex] = item
      this.dietHistoryList.splice(this.selectedItemIndex,1,item)
      console.log("update diet history list",this.dietHistoryList,item)
      // this.updateDietHistory()
      this.closeDietDetail()
    },
    onDieCreated(item:DietContentItem){
      this.dietHistoryList.push(item)
      console.log("create diet history list",this.dietHistoryList,item)
      this.closeDietDetail()
    },
    closeDietDetail(){
      let ele = this.$refs["diet-edit-form"] as any
      if (ele){
        ele.close()
      }
    },
    /* 分享图片生成 */
		createsShareImage(){
			// console.log(this.$refs.canvas)
      // this.$refs.hchPoster.posterShow()
      uni.hideTabBar()
      let ele = this.$refs["shareCard"] as any
      if (ele){
        ele.posterData=this.posterData
        ele.posterShow()
      }
		},
		// 预览图片
		previewHandle(){
			uni.previewImage({
				urls: [this.shareImage],
			});
		},
		// 回调图片地址
		shareSuccess(e:any){
      console.log("share image success",e)
			this.shareImage = e
      uni.previewImage({
				urls: [this.shareImage],
        indicator: 'none',
			});
		},
		// 分享
		onShareAppMessage(res:any) {
			// if (res.from === 'button') {
			// 	console.log(res.target)
			// }
			return {
				title: 'canvas图片分享',
				path: this.shareImage
			}
		},
    onShareCardCancel(){
      
    },

  },
  components: {
    NewRecord,
    UniSection,
    TrainHistoryBriefCard,
    trainingNoticeBar,
    dietContent,
    dietCard,
    shareCard
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

.uni-fab__circle {
  position: absolute;
  bottom: 10px;
  left: 5px;
}

.train-detail-card {
  width: 100%;
  /* box-sizing: border-box;
  border-style:solid;
  border-color:rgba(103, 105, 98, 0.459); */
}

.card-popup__btn_group {
  margin: 10px 10px;
}

/* pop menu  */
.uni-popup-share {
  background-color: #fff;
  border-top-left-radius: 11px;
  border-top-right-radius: 11px;
}
.uni-share-title {
  /* #ifndef APP-NVUE */
  display: flex;
  /* #endif */
  flex-direction: row;
  align-items: center;
  justify-content: center;
  height: 40px;
}
.uni-share-title-text {
  font-size: 14px;
  color: #666;
}
.uni-share-content {
  /* #ifndef APP-NVUE */
  display: flex;
  /* #endif */
  flex-direction: row;
  justify-content: center;
  padding-top: 10px;
}

.uni-share-content-box {
  /* #ifndef APP-NVUE */
  display: flex;
  /* #endif */
  flex-direction: row;
  flex-wrap: wrap;
  width: 360px;
}

.uni-share-content-item {
  width: 90px;
  /* #ifndef APP-NVUE */
  display: flex;
  /* #endif */
  flex-direction: column;
  justify-content: center;
  padding: 10px 0;
  align-items: center;
}

.uni-share-content-item:active {
  background-color: #f5f5f5;
}

.uni-share-image {
  width: 30px;
  height: 30px;
}

.uni-share-text {
  margin-top: 10px;
  font-size: 14px;
  color: #3b4144;
}

.uni-share-button-box {
  /* #ifndef APP-NVUE */
  display: flex;
  /* #endif */
  flex-direction: row;
  padding: 10px 15px;
}

.uni-share-button {
  flex: 1;
  border-radius: 50px;
  color: #666;
  font-size: 16px;
}

.uni-share-button::after {
  border-radius: 50px;
}
</style>
