<!--
 * @Description: 生成海报组件
 * @Version: 1.0.0
 * @Autor: hch
 * @Date: 2020-08-07 14:48:41
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2021-07-31 18:11:35
 * 保存海报按钮和关闭按钮 在html代码中写出来 绑定点击方法然后透明 再用canvas 覆盖
-->

<template>
  <view
    class="canvas-content"
    v-if ="canvasShow"
  >
    <!-- 遮罩层 -->
    <view class="canvas-mask"   
      >
    </view>
    <!-- 海报 -->
    <!-- :width="system.w" :height="system.h" 支付宝必须要这样设置宽高才有效果 -->
    <canvas
      class="canvas"
      canvas-id="shareImage"
      id="shareImage"
      :style="style"

    ></canvas>
    <view class="button-wrapper">
      <!-- 保存海报按钮 -->
      <!-- #ifndef MP-QQ -->
      <!-- cover-view 标签qq小程序有问题 -->
      <!-- <cover-view class="save-btn" @tap="handleSaveCanvasImage">保存海报</cover-view> -->
      <!-- <cover-view class="save-btn" @tap="shareWeChat"  open-type="share"
        >分享</cover-view
      >
      <cover-view class="save-btn cancel-btn" @tap="handleCanvasCancel"
        >取消</cover-view
      > -->
      <!-- #endif -->
      <!-- #ifdef MP-QQ -->
      <!-- <view class="save-btn" @tap="handleSaveCanvasImage">保存海报</view> -->
      <!-- <view class="save-btn cancel-btn" @tap="handleCanvasCancel">取消</view> -->
      <!-- #endif -->
    </view>
  </view>
</template>

<script>
import { set } from 'date-fns';
// import { drawSquarePic, drawTextReturnH, getSystem } from "./utils";
// import { DietContentItem,GetTodayDietHistory,GetDietHistory } from "@/pages/history/apis";

export default {
  data() {
    return {
      system: {},
      canvasShow: false,
      sentence:[
        "即使在最黑的夜，也要有最热的血✊",
        "你的脚步，决定了你的高度✊",
        "你的努力，决定了你的成就✊",
        "你的坚持，决定了你的未来✊",
        "你的选择，决定了你的人生✊",
        "你的行为，决定了你的性格✊",
        "不开心的时候，流泪不如流汗✊",
        "不努力的时候，哭泣不如奋斗✊",
        "所有的可能性，都在你跑下去的路上✊",
        "哪有什么天生如此，只是我们天天坚持✊",
        "管它终点有多远,JUST DO IT✊",
        "拥有梦想的人，他们不做选择题，他们只做证明题✊",
        "为不凡，突破极限✊",
        "想放弃了的时候，想想当初为什么开始✊",
      ],
      canvas_h:0,
      style:"",
      // canvasCtx:any
    };
  },
  props: {
    posterData: {
      type: Object,
      default: () => {
        return {};
      },
    },
  },
  watch: {
    posterData: {
      handler: function (newVal, oldVal) {
        if (newVal) {
          this.posterData = newVal;
        }
      },
      deep: true,
    },
  },
  computed: {
    /**
     * @description: 计算海报背景数据
     * @param {*}
     * @return {*}
     * @author: hch
     */
    poster() {
      let data = this.posterData;
      let system = this.system;
      let posterBg = {
        r: data.poster.r * system.scale, //海报圆角半径
        w: data.poster.w * system.scale, // 海报宽度
        h: data.poster.h * system.scale, // 海报高度
        x: (system.w - data.poster.w * system.scale) / 2, // 海报x轴位置
        y: (system.h - data.poster.h * system.scale) / 2, // 海报y轴位置,(x,y)左上角坐标
        p: data.poster.p * system.scale, // padding
      };
      return posterBg;
    },
    /**
     * @description: 计算海报头部主图
     * @param {*}
     * @return {*}
     * @author: hch
     */
    mainImg() {
      let data = this.posterData;
      let system = this.system;
      let posterMain = {
        url: data.mainImg.url,
        r: data.mainImg.r * system.scale,
        w: data.mainImg.w * system.scale,
        h: data.mainImg.h * system.scale,

        // 主背景左上角坐标
        x: (system.w - data.mainImg.w * system.scale) / 2,
        y: this.poster.y + data.poster.p * system.scale,
      };
      return posterMain;
    },
    content() {
      let data = this.posterData;
      let system = this.system;
      return data.content;
    },
    type(){
      return this.posterData.type
    },
    contentFormat() {
      let data = this.posterData;
      let system = this.system;
      if (data.type == "diet") {
        return data.content;
      }
      let contentFormat = this.reformatTrainContent(data.content.train_content);
      return contentFormat;
    },
  },
  created() {
    // 获取设备信息
    this.system = this.getSystem();
    // this.ctx = uni.createCanvasContext("shareImage", this);
  },
  methods: {
    getSystem() {
      let system = wx.getSystemInfoSync()
      console.log("getSystem -> system", system)
      let scale = system.windowWidth / 375 //按照苹果留 375*667比例 其他型号手机等比例缩放 显示
      return { w: system.windowWidth, h: system.windowHeight, scale: scale }
    },
    /**
     * @description: 展示海报
     * @param {type}
     * @return {type}
     * @author: hch
     */
    posterShow() {
      this.canvasShow = true;
      this.$nextTick(function () {
        // 此时已经渲染完成
        console.log("nextTick -> this", this)
        this.createPoster();
        // this.shareWeChat()
      });
      // this.createPoster();
      // this.shareWeChat()
    },
    /**
     * @description: 生成海报
     * @author: hch
     */
    createPoster() {
      console.log("createPoster",this.posterData,this.system)
      uni.showLoading({
        title: "生成海报中...",
      });
      uni.pageScrollTo({ scrollTop: 0,duration:0 })
      const ctx = uni.createCanvasContext("shareImage", this);
      
      // this.canvasCtx = ctx;
      ctx.clearRect(0, 0, this.system.w, this.system.h); //清空之前的海报
      ctx.draw(); //清空之前的海报

      // 根据设备屏幕大小和距离屏幕上下左右距离，及圆角绘制背景
      let poster = this.poster;
      let mainImg = this.mainImg;
      let content= this.content
      let contentType =this.type
      console.log
      this.drawPicture(
        ctx,
        poster.x,
        poster.y,
        mainImg.w,
        mainImg.h,
        mainImg.r,
        mainImg.url
      );
      // 绘制内容      
      this.drawContent(
        ctx,
        contentType,
        content,
        poster.x,
        poster.y + mainImg.h,
        mainImg.w,
        poster.h - mainImg.h,
        mainImg.r
      );

      // 绘制 icon
      ctx.beginPath();
      ctx.fillStyle = "#fff";
      ctx.font = "20px Arial";
      ctx.fillText("Air-body", poster.x + 20, poster.y + 30);

      ctx.draw(true);

      // 绘制小程序码
      this.drawPicture(
        ctx,
        poster.x + mainImg.w - 85,
        poster.y +5  ,
        80,
        80,
        mainImg.r,
        "../../static/icons/qrcode.jpg"
      )

      // 绘制语录
      this.drawSentence(ctx,poster.w/3 - 20,(poster.y + mainImg.h)/3+10) 
      uni.hideLoading();
      // ctx.destroy()
    },
    // 绘制背景图片
    drawPicture(ctx, x, y, w, h, r, url) {
      // 画图
      ctx.save();
      ctx.beginPath();
      ctx.moveTo(x, y);
      // 绘制左上角圆弧
      ctx.arc(x + r, y + r, r, Math.PI, Math.PI * 1.5);
      // 绘制border-top
      // 画一条线 x终点、y终点
      ctx.lineTo(x + w - r, y);
      // 绘制右上角圆弧
      ctx.arc(x + w - r, y + r, r, Math.PI * 1.5, Math.PI * 2);
      // 绘制border-right
      ctx.lineTo(x + w, y + h);
      ctx.lineTo(x,y+h)
      // 绘制右下角圆弧
      // ctx.arc(x + w - r, y + h - r, r, 0, Math.PI * 0.5);
      // 绘制左下角圆弧
      // ctx.arc(x + r, y + h - r, r, Math.PI * 0.5, Math.PI);
      // 绘制border-left
      ctx.lineTo(x, y + r);
      // 填充颜色(需要可以自行修改)
      ctx.setFillStyle("#ffffff");
      ctx.fill();
      // 剪切，剪切之后的绘画绘制剪切区域内进行，需要save与restore 这个很重要 不然没办法保存
      ctx.clip();
      ctx.drawImage(url, x, y, w, h);
      ctx.restore(); //恢复之前被切割的canvas，否则切割之外的就没办法用
      ctx.draw(true);
    },
    // 绘制具体的内容n b
    drawContent(ctx, type,content, x, y, w, h, r) {
      h=h+this.getContentHeight()
      this.style = `width:${this.system.w}px;height:${this.system.h+this.getContentHeight()}px;`
      // console.log("drawContent -> content", this.contentFormat);
      // 画图
      ctx.save();
      ctx.beginPath();
      // // 绘制左上角圆弧
      // 画一条线 x终点、y终点
      ctx.moveTo(x, y);
      ctx.lineTo(x + w , y);
      // 绘制border-right
      ctx.lineTo(x + w, y + h);
      // 绘制右下角圆弧
      ctx.arc(x + w - r, y + h - r, r, 0, Math.PI * 0.5);
      // 绘制左下角圆弧
      ctx.arc(x + r, y + h - r, r, Math.PI * 0.5, Math.PI);
      // 绘制border-left
      ctx.lineTo(x, y + r);
      ctx.setFillStyle("#ffffff");
      ctx.fill();
      ctx.draw(true);
      // 绘制内容
      if (type=="train"){
        // 训练时间
        // 内容标题
        ctx.moveTo(x + 10, y + 30);
        ctx.font = "bold 14px Arial";
        ctx.setFillStyle("#000000");
        ctx.fillText(`${content.train_history.title}`, x + 10, y + 30);

        ctx.moveTo(x + 10, y + 50);
        ctx.font = "12px Arial";
        ctx.fillText(`训练时间:${content.train_history.created_at}`, x + 10, y + 50);
        // 
        var  axis_y = y + 80
        console.log("contentFormat", this.contentFormat)
        for (const key in this.contentFormat) {
          if (Object.hasOwnProperty.call(this.contentFormat, key)) {
            axis_y = axis_y + 20
            const element = this.contentFormat[key];
            console.log("drawContent -> element", element,axis_y)
            ctx.moveTo(x + 10, axis_y);
            ctx.font = "bold 12px Arial";
            ctx.fillText(`${key} * ${element.length}`, x + 10, axis_y);
            ctx.font = "12px Arial";
            axis_y = axis_y + 20
            let loop_count = Math.ceil(element.length/3)
            for (let index2 = 0; index2 < loop_count; index2++) {
              let content = ""
              for (let index3 = 0; index3 < 3; index3++) {
                const item = element[index2*3+index3];
                if (item){
                  content = content + `${item.left_weight}kg*${item.right_weight}kg*${item.number}次   `
                }
              }
              ctx.fillText(content, x + 10, axis_y);
              axis_y = axis_y + 20
            }
          }
        }
        // this.style = `width:${this.system.w}px;height:${this.system.h+axis_y}px;`

      }else if (type=="diet"){
        // 饮食记录时间
        // 内容标题
        // ctx.moveTo(x + 10, y + 10);
        ctx.font = "bold 14px Arial";
        ctx.setFillStyle("#000000");
        ctx.fillText(`饮食记录时间:${content.created_at}`, x + 10, y + 30);
        // 饮食内容
        ctx.fillText(`蛋白质:${content.protein}(g)`, x + 10, y + 50 + 20);
        // 饮食内容
        ctx.fillText(`脂肪:${content.fat}(g)`, x + 200, y + 50 + 20);
        ctx.fillText(`碳:${content.carbon}(g)`, x + 10, y + 70 + 30);
        ctx.fillText(`水:${content.fat}(g)`, x + 200, y + 70 + 30);
        ctx.fillText(`已经进食:${content.calorie}(卡路里)`, x + 10, y + 90 + 40);
      }
      ctx.draw(true)
    },
    // 绘制语录
    drawSentence(ctx,x,y){
      let sentence = this.sentence
      let index = Math.floor(Math.random() * sentence.length)
      // ctx.moveTo(10, 10);
      ctx.font = "bold 18px Arial";
      ctx.setFillStyle("#fff");
      let arr = sentence[index].split("，")
      // console.log("drawSentence -> arr", arr)
      for (let index = 0; index < arr.length; index++) {
        const element = arr[index];
        ctx.fillText(`${element}`, x + index*30, y + index * 30);
      }
      ctx.draw(true,(res)=>{
        console.log("draw share image finish")
        this.shareWeChat()
      })
    },

    reformatTrainContent(contentList){
      //生成训练内容的简要信息
      let res = {};
      for (let index = 0; index < contentList.length; index++) {
        const content = contentList[index];
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

    shareWeChat(){
      console.log("shareWeChatImage")
      // 分享canvas生成的图片
      let that = this
      // that.handleCanvasCancel()
      // var shareImage = uni.createSelectorQuery().in(this).select('#shareImage')
      // console.log("shareWeChat -> canvas", shareImage)
      wx.canvasToTempFilePath({
        canvasId: 'shareImage',
        // canvas: shareImage,
        success: (res) => { // 成功回调
         // 在H5平台下，tempFilePath 为 base64
          console.log('share file path', res.tempFilePath)
          wx.showShareImageMenu({
            path: res.tempFilePath,
            success: function (res) {
              console.log('分享成功', res)
              that.handleCanvasCancel()
            },
            fail: function (res) {
              console.log('分享失败', res)
              that.handleCanvasCancel()
            }
          })
        },
        fail: (err) => { // 失败回调
          console.log('share file path err', err)
          that.handleCanvasCancel()
        }
      },that)
      // uni.canvasToTempFilePath({
      //   x: 0,
      //   y: 0,
      //   width: this.system.w,
      //   height: this.canvas_h,
      //   destWidth: this.system.w,
      //   destHeight: this.canvas_h,
      //   canvasId: 'shareImage',
      //   success: function(res) {
      //     // 在H5平台下，tempFilePath 为 base64
      //     console.log("save canvas",res.tempFilePath)
      //     that.handleCanvasCancel()
      //     },
      //   fail: (err) => { // 失败回调
      //     console.log('share file path err', err,this.system)
      //     that.handleCanvasCancel()
      //   }
      // },that)    
    },
    getContentHeight(){
      let axis_y = 0
      for (const key in this.contentFormat) {
          if (Object.hasOwnProperty.call(this.contentFormat, key)) {
            axis_y = axis_y + 20
            const element = this.contentFormat[key];
            let loop_count = Math.ceil(element.length/3)
            for (let index2 = 0; index2 < loop_count; index2++) {
              axis_y = axis_y + 20
            }
          }
        }      
      return axis_y
    },
    handleCanvasCancel(){
      this.canvasShow = false
      // console.log("handleCanvasCancel",this.canvasCtx)
      // this.style = "width:0px;height:0px;"
      // this.$emit('cancel', true)
    },
    onLoadReady(){
      console.log("on canvas content load ready ->")
      this.createPoster();
      this.shareWeChat()
    }

  },
};
</script>

<style lang="scss">
.content {
  height: 100%;
  text-align: center;
}

.canvas-content {
  position: absolute;
  top: 0;

  .canvas-mask {
    position: fixed;
    top: 0;
    right: 0;
    bottom: 0;
    left: 0;
    z-index: 9;
    width: 100%;
    height: 100%;
    background: rgba(0, 0, 0, 0.5);
  }

  .canvas {
    z-index: 10;
  }

  .button-wrapper {
    position: fixed;
    bottom: 40rpx;
    z-index: 16;
    display: flex;
    width: 100%;
    height: 72rpx;
    justify-content: space-around;
  }

  .save-btn {
    z-index: 16;
    width: 40%;
    height: 100%;
    font-size: 30rpx;
    line-height: 72rpx;
    color: #fff;
    text-align: center;
    background: #007aff;
    border-radius: 45rpx;
    border-radius: 36rpx;
  }

  .cancel-btn {
    color: #007aff;
    background: #fff;
  }

  .canvas-close-btn {
    position: fixed;
    top: 30rpx;
    right: 0;
    z-index: 12;
    width: 60rpx;
    height: 60rpx;
    padding: 20rpx;
  }
}
</style>
