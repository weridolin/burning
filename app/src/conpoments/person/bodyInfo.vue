<template>
  <view class="container__createBodyInfo">
    <uni-calendar ref="create-form-calendar" class="uni-calendar--hook" :clear-date="true" 
        :date="date" :insert="false"  @confirm="switchDate" :selected="selected"  />
    <view class="container__date" style="padding-left: 10px; padding-right: 10px">
      <button  plain="default" type="default" size="mini" @click="openCalendar" style="width: 80%;">选择日期</button>
      <button  plain="default" type="default" size="mini" @click="addNewBodyInfo">新建</button>
    </view>
    <uni-section class="mb-10" :title="'基本数据(' + (date==''?'请先选择日期':date )+ ')'" type="line">
      <view class="bodyImage">
        <canvas
          canvas-id="body"
          class="canvas-body"
          id="body"
          style="width: 300px; height: 550px; background-color: #eeeeee"
          @click="onCanvasClick"
        >
        </canvas>
      </view>
    </uni-section>
    <uni-section class="mb-10" :title="'基本数据(' + (date==''?'请先选择日期':date )+ ')'" type="line">
      <view class="synthetic-data">
        <view class="synthetic-data__item" >
          <uni-title type="h3" title="输入身高:"></uni-title>
          <uni-number-box :value="userBodyInfo.height" background="#2979FF" color="#fff" @change="updateUserInfo" />
        <!-- <uni-easyinput v-model="userBodyInfo.weight" :styles="styles" :placeholderStyle="placeholderStyle" placeholder="输入体重" @input="updateUserInfo"></uni-easyinput> -->
        </view>
        <view class="synthetic-data__item" >
          <uni-title type="h3" title="输入体重:"></uni-title>
          <uni-number-box :value="userBodyInfo.weight" background="#2979FF" color="#fff"  @change="updateUserInfo"/>
        <!-- <uni-easyinput v-model="userBodyInfo.weight" :styles="styles" :placeholderStyle="placeholderStyle" placeholder="输入体重" @input="updateUserInfo"></uni-easyinput> -->
        </view>
        <view class="synthetic-data__item" >
          <uni-title type="h3" title="体脂比"></uni-title>
          <uni-number-box :value="userBodyInfo.body_fat_rate" background="#2979FF" color="#fff"  @change="updateUserInfo" />
        </view>
      </view>
    </uni-section>

    <!-- 输入框示例 -->
    <uni-popup ref="inputDialog" type="dialog">
      <uni-popup-dialog ref="inputClose"  mode="input" :title="flag" :value="'输入'+flag+'数据'"
        placeholder="请输入内容" @confirm="dialogInputConfirm"></uni-popup-dialog>
    </uni-popup>

  <!-- 新建档案 -->
    <view class="container__createBodyInfo_form" >
      <uni-drawer ref="create-form" mode="left" :width="getWidth()" >
				<uni-forms ref="baseForm" :modelValue="createForm" >
          <uni-forms-item label="日期" required class="container__create_form_item">
						<uni-datetime-picker type="date" v-model="createForm.date"/>
					</uni-forms-item>
          <uni-forms-item label="身高" class="container__create_form_item">
            <uni-number-box :value="createForm.height" background="#2979FF" color="#fff" />
					</uni-forms-item>
          <uni-forms-item label="体重" >
            <uni-number-box :value="createForm.waistline" background="#2979FF" color="#fff"/>
					</uni-forms-item>
          <uni-forms-item label="体脂" >
            <uni-number-box :value="createForm.body_fat_rate" background="#2979FF" color="#fff"/>
					</uni-forms-item>
					<uni-forms-item label="臂围" >
            <uni-number-box :value="createForm.upper_arm_circumference" background="#2979FF" color="#fff"/>
					</uni-forms-item>
          <uni-forms-item label="胸围" >
            <uni-number-box :value="createForm.chest_circumference" background="#2979FF" color="#fff"/>
					</uni-forms-item>
          <uni-forms-item label="臀围" >
						<uni-number-box :value="createForm.hip_line" background="#2979FF" color="#fff"/>
					</uni-forms-item>
          <uni-forms-item label="肩宽" >
            <uni-number-box :value="createForm.shoulder_breadth" background="#2979FF" color="#fff"/>
					</uni-forms-item>
          <uni-forms-item label="腰围" >
            <uni-number-box :value="createForm.waistline" background="#2979FF" color="#fff"/>
					</uni-forms-item>
          <uni-forms-item label="大腿围" >
            <uni-number-box :value="createForm.thigh_circumference" background="#2979FF" color="#fff"/>
					</uni-forms-item>
          <uni-forms-item label="小腿围" >
            <uni-number-box :value="createForm.calf_circumference" background="#2979FF" color="#fff"/>
					</uni-forms-item>
          <uni-forms-item>
            <button type="primary" @click="createUserInfo">提交</button>
          </uni-forms-item>
				</uni-forms>
      </uni-drawer>
    </view>  

  </view>
</template>
<script lang="ts">
import Vue from "vue";
import { UserBodyInfo } from "../../store/local";
import {
  GetUserBodyInfo,
  UpdateUserBodyInfo,
  CreateUserBodyInfo,
  GetAllBodyInfo
} from "../../pages/person/apis";

export default Vue.extend({
  data() {
    let dateBodyInfoMap:{[key:string]:UserBodyInfo}= {} as {[key:string]:UserBodyInfo};
    let userBodyInfo = {} as UserBodyInfo;
    let createForm = {} as UserBodyInfo;
    let selected = [] as any;
    return {
        userBodyInfo, 
        date: "" ,				
        placeholderStyle: "color:#2979FF;font-size:14px",
				styles: {
					color: '#2979FF',
					borderColor: '#2979FF'
				},
        bodyInfoFlagPosition:{
          chest_circumference:{ //胸围
            x: 220,
            y: 103
          },
          upper_arm_circumference:{ // 臂围
            x: 2,
            y: 161
          },
          shoulder_width:{ // 肩宽
            x: 2,
            y: 95
          },
          thigh_circumference:{  // 大腿围
            x: 221,
            y: 352
          },
          calf_circumference:{ // 小腿围
            x: 7,
            y: 443
          },
          waistline:{ // 腰围
            x: 227,
            y: 236
          },
          hip_line:{ // 臀围
            x: 24,
            y: 276
          },
        },
        createForm,
        flag:"",//当前编辑的数据
        dateBodyInfoMap,
        selected
      };
  },
  onReady() {
    this.drawInfo();
  },
  onUnload() {
  },
  onShow() {
    this.getAllBodyInfo();
  },
  methods: {
    drawInfo() {
      let ctx = uni.createCanvasContext("body", this);
      ctx.drawImage("../../static/icons/body.png", 0, 0, 300, 550);
      ctx.beginPath();
      ctx.fillStyle = "#9c9796";
      ctx.font = "12px Arial";
      // ctx.fillText(`脖维:${this.userBodyInfo.height || "无数据"}`, 10, 50);
      // this.drawLine(ctx, 15, 65, 135, 65);

      // 臂围
      ctx.fillText(`臂围:${this.userBodyInfo.upper_arm_circumference || "无数据"}`,
        this.bodyInfoFlagPosition.upper_arm_circumference.x,this.bodyInfoFlagPosition.upper_arm_circumference.y);
        this.drawLine(ctx,this.bodyInfoFlagPosition.upper_arm_circumference.x, 
                      this.bodyInfoFlagPosition.upper_arm_circumference.y+10, 
                      79,
                      this.bodyInfoFlagPosition.upper_arm_circumference.y+10);

      // 肩宽
      ctx.fillText(`肩宽:${this.userBodyInfo.shoulder_breadth || "无数据"}`, 
      this.bodyInfoFlagPosition.shoulder_width.x, this.bodyInfoFlagPosition.shoulder_width.y);
      this.drawLine(ctx, this.bodyInfoFlagPosition.shoulder_width.x,
        this.bodyInfoFlagPosition.shoulder_width.y+10, 
        85, 
        this.bodyInfoFlagPosition.shoulder_width.y+10);

      //胸围
      ctx.fillText(`胸围:${this.userBodyInfo.chest_circumference || "无数据"}`,
        this.bodyInfoFlagPosition.chest_circumference.x, this.bodyInfoFlagPosition.chest_circumference.y);
      this.drawLine(ctx, 292, 112, 164, 112);

      // 腰围
      ctx.fillText(`腰围:${this.userBodyInfo.waistline || "无数据"}`,
        this.bodyInfoFlagPosition.waistline.x, this.bodyInfoFlagPosition.waistline.y);
      this.drawLine(ctx, 
        this.bodyInfoFlagPosition.waistline.x, 
        this.bodyInfoFlagPosition.waistline.y+10, 
        151, 
        this.bodyInfoFlagPosition.waistline.y+10);

      // 臀围
      ctx.fillText(`臀围:${this.userBodyInfo.hip_line || "无数据"}`, 
        this.bodyInfoFlagPosition.hip_line.x, this.bodyInfoFlagPosition.hip_line.y);
      this.drawLine(ctx, 
        this.bodyInfoFlagPosition.hip_line.x, 
        this.bodyInfoFlagPosition.hip_line.y+10, 
        102, 
        this.bodyInfoFlagPosition.hip_line.y+10);

      //  大腿围
      ctx.fillText(`大腿围:${this.userBodyInfo.thigh_circumference || "无数据"}`, 
        this.bodyInfoFlagPosition.thigh_circumference.x, this.bodyInfoFlagPosition.thigh_circumference.y);
      this.drawLine(ctx, 231, 362, 186, 362);

      // 小腿围
      ctx.fillText(`小腿围:${this.userBodyInfo.calf_circumference || "无数据"}`, 
        this.bodyInfoFlagPosition.calf_circumference.x,this.bodyInfoFlagPosition.calf_circumference.y);
      this.drawLine(ctx, 10, 453, 103, 453);

      ctx.lineWidth = 0.5; // 设置线的宽度
      ctx.strokeStyle = "#9c9796"; // 设置线的颜色
      ctx.lineCap = "round";
      ctx.stroke();
      ctx.draw();
    },
    drawLine(ctx: any, x1: number, y1: number, x2: number, y2: number) {
      ctx.moveTo(x1, y1);
      ctx.lineTo(x2, y2);
      ctx.closePath();
    },
    getUserInfo() {
      console.log("(body info detail page) get body info -> ", this.date);
      uni.showLoading({
        title: "获取身体数据中...",
      });
      GetUserBodyInfo(
          this.date,
          (res) => {
            console.log("(body info detail page) get body info success -> ", res);
            this.userBodyInfo = res.data;
            this.dateBodyInfoMap[this.date] = res.data;
            this.selected.push({
              date: this.date,
              info: '已记录'
            });
            this.drawInfo();
            uni.hideLoading();
          },
          (error) => {
            console.log("(body info detail page) get body info error -> ",error);
            uni.hideLoading();
          }
        );
    },
    updateUserInfo() {
      uni.showLoading({
        title: "更新身体数据中...",
      });
      UpdateUserBodyInfo(
        this.userBodyInfo,
        (res) => {
          console.log("(body info detail page) update body info success -> ", res);
          uni.hideLoading();
        },
        (error) => {
          console.log("(body info detail page) update body info error -> ",error);
          uni.hideLoading();
        }
      );
    },
    createUserInfo() {
      uni.showLoading({
        title: "创建身体数据中...",
      });
      CreateUserBodyInfo(
        this.createForm,
        (res) => {
          console.log("(body info detail page) create body info success -> ", res);
          this.userBodyInfo = this.createForm
          this.dateBodyInfoMap[this.date] = this.createForm;
          this.selected.push({
              date: this.createForm.date,
              info: '已记录'
            });
          uni.hideLoading();
        },
        (error) => {
          console.log("(body info detail page) create body info error -> ",error);
          uni.hideLoading();
        }
      );
    },
    switchDate(date:any) {
      this.date = date.fulldate;
      if (this.dateBodyInfoMap[this.date] != null){
        this.userBodyInfo = this.dateBodyInfoMap[this.date]
        this.drawInfo();
      }else{
          this.getUserInfo();
      }
    
    },
    change(value:string){
      this.updateUserInfo();
    },
    addNewBodyInfo(){
      // this.userBodyInfo = {} as UserBodyInfo;
      if (this.date == ""){
        uni.showToast({
          title: '请先选择日期',
          icon: 'error'
        });
        return;
      }
      // this.createUserInfo();
      var createForm = this.$refs["create-form"] as any;
      if (createForm == null) {
        console.log("(body info detail page) createForm is null");
        return;
      }else{
        createForm.open();
      }
    },
    onCanvasClick(e:any){
      // 303 193
    
      uni.createSelectorQuery().in(this).select(`.canvas-body`).boundingClientRect().exec(res => {
      console.log("(body info detail page) canvas click position -> ", 
        e.detail.x, e.detail.y,res);
      
      let x = e.detail.x - res[0].left;
      let y = e.detail.y - res[0].top;
      if (x <= this.bodyInfoFlagPosition.chest_circumference.x + 60 && x >= this.bodyInfoFlagPosition.chest_circumference.x - 10
        && y <= this.bodyInfoFlagPosition.chest_circumference.y + 30 && y >= this.bodyInfoFlagPosition.chest_circumference.y - 10){
        console.log("(body info detail page) 胸围编辑");
        this.flag = "胸围";
      }else if (
        x <= this.bodyInfoFlagPosition.upper_arm_circumference.x + 60 && x >= this.bodyInfoFlagPosition.upper_arm_circumference.x - 10
        && y <= this.bodyInfoFlagPosition.upper_arm_circumference.y + 30 && y >= this.bodyInfoFlagPosition.upper_arm_circumference.y - 10){
        console.log("(body info detail page) 臂围编辑");
        this.flag = "臂围";
      }else if (
        x <= this.bodyInfoFlagPosition.shoulder_width.x + 60 && x >= this.bodyInfoFlagPosition.shoulder_width.x - 10
        && y <= this.bodyInfoFlagPosition.shoulder_width.y + 30 && y >= this.bodyInfoFlagPosition.shoulder_width.y - 10){
        console.log("(body info detail page) 肩宽编辑");
        this.flag = "肩宽";
      }else if (
        x <= this.bodyInfoFlagPosition.thigh_circumference.x + 60 && x >= this.bodyInfoFlagPosition.thigh_circumference.x - 10
        && y <= this.bodyInfoFlagPosition.thigh_circumference.y + 30 && y >= this.bodyInfoFlagPosition.thigh_circumference.y - 10){
        console.log("(body info detail page) 大腿围编辑");
        this.flag = "大腿围";
      }else if (
        x <= this.bodyInfoFlagPosition.calf_circumference.x + 60 && x >= this.bodyInfoFlagPosition.calf_circumference.x - 10
        && y <= this.bodyInfoFlagPosition.calf_circumference.y + 30 && y >= this.bodyInfoFlagPosition.calf_circumference.y - 10){
        console.log("(body info detail page) 小腿围编辑");
        this.flag = "小腿围";
      }else if (
        x <= this.bodyInfoFlagPosition.waistline.x + 60 && x >= this.bodyInfoFlagPosition.waistline.x - 10
        && y <= this.bodyInfoFlagPosition.waistline.y + 30 && y >= this.bodyInfoFlagPosition.waistline.y - 10){
        console.log("(body info detail page) 腰围编辑");
        this.flag = "腰围";
      }else if (
        x <= this.bodyInfoFlagPosition.hip_line.x + 60 && x >= this.bodyInfoFlagPosition.hip_line.x - 10
        && y <= this.bodyInfoFlagPosition.hip_line.y + 30 && y >= this.bodyInfoFlagPosition.hip_line.y - 10){
        console.log("(body info detail page) 臀围编辑");
        this.flag = "臀围";
        }     
      })

      var inputDialog = this.$refs.inputDialog as any;
      if (inputDialog == null) {
        console.log("(body info detail page) inputDialog is null");
        return;
      }else{
        inputDialog.open();
      }
    },
    dialogInputConfirm(e:any){
      console.log("(body info detail page) dialogInputConfirm -> ", e);
      // 判断输入值是否为数字
      if (isNaN(parseFloat(e))){
        uni.showToast({
          title: '请输入数字',
          icon: 'error'
        });
        return;
      }
      switch (this.flag) {
        case "胸围":
          this.userBodyInfo.chest_circumference = e;
          break;
        case "臂围":
          this.userBodyInfo.upper_arm_circumference = e;
          break;
        case "肩宽":
          this.userBodyInfo.shoulder_breadth = e;
          break;
        case "大腿围":
          this.userBodyInfo.thigh_circumference = e;
          break;
        case "小腿围":
          this.userBodyInfo.calf_circumference = e;
          break;
        case "腰围":
          this.userBodyInfo.waistline = e;
          break;
        case "臀围":
          this.userBodyInfo.hip_line = e;
          break;
        default:
          break;
      }
      this.updateUserInfo();
      this.drawInfo();
    },
    getWidth() {
      let width = uni.getSystemInfoSync().windowWidth * 0.8;
      console.log(width);
      return width;
    },
    openCalendar(){
      // 打开日历选择框
      var calendar = this.$refs["create-form-calendar"] as any;
      if (calendar == null) {
        console.log("(body info detail page) calendar is null");
        return;
      }else{
          calendar.open();
      }
    },
    getAllBodyInfo(){
      // 获取所有用户的身体数据记录
      GetAllBodyInfo(
        (res) => {
          console.log("(body info detail page) get all body info success -> ", res);
          this.dateBodyInfoMap = {} ;
          this.selected = [];
          let data = res.data as UserBodyInfo[];
          for (let i = 0; i < data.length; i++) {
            let item = data[i];
            let date = item.date;
            this.dateBodyInfoMap[date] = item;
            this.selected.push({
              date: date,
              info: '已记录'
            });
          }
        },
        (error) => {
          console.log("(body info detail page) get all body info error -> ",error);
        }
      );
    } 
  },
});
</script>

<style lang="scss" scoped>
.synthetic-data {
  .synthetic-data__item {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-between;
    // padding: 10px;
    padding-right: 10px;
    padding-left: 10px;
    border-bottom: 1px solid #eeeeee;

  }
}
.container__createBodyInfo {
  padding: 10px;

  .container__date{
    padding-left: 10px;
    padding-right: 10px;
    display: flex;
    gap: 10px;
  }
}
.bodyImage {
  //居中显示
  display: flex;
  justify-content: center;
  align-items: center;
}

</style>