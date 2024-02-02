<template>
  <view class="audio-content">
    <!-- 调速拉条区 -->
    <view class="audio-progress">
      <image src="../../static/icons/get-back.svg" @click="onSeek(-15)"></image>
      <view class="audio-bar">
        <slider
          class="audio-slider"
          :value="currentTime"
          :min="0"
          :max="duration"
          @change="sliderChange"
          @changing="sliderChanging"
          activeColor="#00CA86"
          backgroundColor="#DAE0E6"
          block-color="#00CA86"
          block-size="18"
        />
      </view>
      <image
        src="../../static/icons/fast-forward.svg"
        @click="onSeek(15)"
      ></image>
    </view>
    <!-- 显示时间进度区 -->
    <view class="bar-tips">
      <text>{{ minTimeFmt(currentTime) }}</text>
      <text>{{ minTimeFmt(duration) }}</text>
    </view>
    <!-- 按钮区 -->
    <view class="audio-btn">
      <!-- <image v-if="!showAudioListIcon"></image> -->
      <image src="../../static/icons/play-list.svg" @tap="onOpenList"></image>
      <view class="play-box">
        <image
          src="../../static/icons/last-episode.svg"
          :style="{ opacity: showLeftBtnOpacity }"
          @tap="onSwitchAudio(audioActiveIndex - 1)"
        ></image>
        <view
          class="play-btn"
          :style="{ backgroundColor: !isPaused ?  '#475266':'#00CA86' }"
          @click="onPlayPause"
        >
          <image v-show="isPaused" src="../../static/icons/suspend.svg"></image>
          <image v-show="!isPaused" src="../../static/icons/start.svg"></image>
        </view>
        <image
          src="../../static/icons/next-set.svg"
          :style="{ opacity: showRightBtnOpacity }"
          @tap="onSwitchAudio(audioActiveIndex + 1)"
        ></image>
      </view>
      <image
        v-if="showAudioSpeedIcon"
        src="../../static/icons/multiple.svg"
        @tap="onSpeed"
      ></image>
      <text v-if="showAudioListIcon" class="play-list-text" @tap="onOpenList"
        >播放列表</text
      >
      <text v-if="showAudioSpeedIcon" class="speed-text" @tap="onSpeed"
        >x{{ speed }}</text
      >
    </view>
  </view>
</template>

<script lang="ts">
import Vue from 'vue'
import {innerAudioContext} from '@/store/gloabal'

export default Vue.extend({
  name: "musicPlayer",
  data() {
    return {
      isPaused: true, //是否暂停中
      duration: 0, //音频时长
      currentTime: 0, //当前时长
      audioActiveIndex: 0, //当前播放索引
      speed: 1, //倍速
      isSlidering: false, //是否移动中
      isEndAcudio: false, //最后一个音频结束
      audioList: [
        {
          title:"测试",
          fileUrl:'https://web-ext-storage.dcloud.net.cn/uni-app/ForElise.mp3',
          image:"https://img-cdn-qiniu.dcloud.net.cn/uniapp/audio/music.jpg",
          enpame:"专辑名",
          singer:"歌手名",
        },
        {
          title:"测试",
          fileUrl:'http://music.163.com/song/media/outer/url?id=34341360.mp3',
          image:"https://img-cdn-qiniu.dcloud.net.cn/uniapp/audio/music.jpg",
          enpame:"专辑名",
          singer:"歌手名",
        }
      ],
      isAutoplay: false,
      showAudioSpeedIcon:true,
      showAudioListIcon:true
    }
  },
  watch: {

  },
  computed: {
    showLeftBtnOpacity():number {
      return this.audioActiveIndex === 0 ? 0.5 : 1;
    },
    showRightBtnOpacity() :number{
      return this.audioActiveIndex + 1 === this.audioList.length ? 0.5 : 1;
    }
  },
  mounted() {
    this.$nextTick(() => {
      this.startPlay();
      this.setAudioInfo();
    });
  },
  methods:{
    minTimeFmt(val:any) {
      let minute = Math.floor(val / 60);
      let seconds = Math.ceil(val % 60);
      return `${minute >= 10 ? minute : "0" + minute}:${
        seconds >= 10 ? seconds : "0" + seconds
      }`;
    },
    onOpenList() {
      console.log("打开播放列表")
      this.$emit("onOpenAudioList");
    },
    startPlay() {
      innerAudioContext.autoplay = false;
      innerAudioContext.onPlay(()=>{
        this.onPlay()
      })
      innerAudioContext.onPause(()=>{
        this.onPause()
      });
      innerAudioContext.onCanplay((event)=>{
        this.onCanplay(event)
      });
      innerAudioContext.onError((res)=>{
        this.onError(res)
      });
      innerAudioContext.onEnded(()=>{
        this.onEnded()
      });
      innerAudioContext.onSeeked((num)=>{
        this.onSeek(num)
      });
      innerAudioContext.onTimeUpdate(()=>{
        this.setPlayData()
      });
    },
    // 暂停播放切换
    onPlayPause() {
      // const innerAudioContext = uni.createInnerAudioContext();
      // innerAudioContext.autoplay = true;
      // innerAudioContext.src = 'https://web-ext-storage.dcloud.net.cn/uni-app/ForElise.mp3';
      console.log('切换播放状态 ->',this.isPaused,innerAudioContext,">>>",innerAudioContext.src)

      if (this.isPaused) {
        innerAudioContext.play();
        this.isPaused = false;
      } else {
        innerAudioContext.pause();
        this.isPaused = true;
      }
    },

    // 切换音频
    onSwitchAudio(index:number) {
      console.log('切换音频',index,this.audioActiveIndex)
      if (index < 0 || index >= this.audioList.length ||this.audioActiveIndex === index) {
        return;
      }
      this.audioActiveIndex = index;

      // innerAudioContext.pause();
      innerAudioContext.stop();
      this.isPaused = true;
      innerAudioContext.seek(0);
      this.currentTime=0
      this.duration=0
      this.$emit("onAudioChange", this.audioList[this.audioActiveIndex], index);
      this.setAudioInfo();

      if (this.isPaused) {
        innerAudioContext.play();
        this.isPaused = false;
        this.isEndAcudio = false;
        this.isSlidering = false;
      }
    },
    // 调整播放倍速
    onSpeed() {
      console.log('调整播放倍速') 
      if (this.speed==1){
        this.speed = 1.5;
        innerAudioContext.playbackRate = 1.5;
      }else if(this.speed==1.5){
        this.speed = 2;
        innerAudioContext.playbackRate = 2;
      }else{
        this.speed = 1;
        innerAudioContext.playbackRate = 1;
      }
    },
    // 调整播放位置
    onSeek(num:number) {
      console.log('调整播放位置',num)
      if (!innerAudioContext.currentTime) return;
      let seekNum = num + innerAudioContext.currentTime;
      this.isSlidering = true;
      if (seekNum < 0) {
        // 调整后时间小于0
        this.currentTime = 0;
        innerAudioContext.seek(0);
      } else if (seekNum > innerAudioContext.duration) {
        // 调整后时间大于总时长
        this.currentTime = innerAudioContext.duration;
        innerAudioContext.seek(innerAudioContext.duration);
      } else {
        this.currentTime = seekNum;
        innerAudioContext.seek(seekNum);
      }
    },
    // 滑块滚动到的位置
    sliderChange(e:any) {
      console.log("滑块滚动到的位置",e)
      this.isSlidering = false;
      this.currentTime = e.detail.value;
      innerAudioContext.seek(e.detail.value);
    },
    // 滑块滚动到的位置 实时
    sliderChanging(e:any) {
      this.isSlidering = true;
      this.currentTime = e.detail.value;
    },
    // 设置以及转换信息
    setPlayData() {
      console.log('音频播放进度更新事件')
      if (this.isSlidering) return;
      if (!innerAudioContext.duration && !innerAudioContext.currentTime) return;
      this.duration = innerAudioContext.duration || 0;
      this.currentTime = innerAudioContext.currentTime || 0;
    },
    // 设置播放
    setAudioInfo() {
      if (innerAudioContext && this.audioActiveIndex > -1 &&this.audioList[this.audioActiveIndex]?.fileUrl) {
        // innerAudioContext.title = this.audioList[this.audioActiveIndex]?.title; //音频标题，用于做原生音频播放器音频标题。原生音频播放器中的分享功能，分享出去的卡片标题，也将使用该值。
        // innerAudioContext.coverImgUrl =this.audioList[this.audioActiveIndex]?.image; //封面图url，用于做原生音频播放器背景图。原生音频播放器中的分享功能，分享出去的卡片配图及背景也将使用该图。
        innerAudioContext.src = this.audioList[this.audioActiveIndex].fileUrl; //音频的数据链接，用于直接播放。
        // innerAudioContext.epname =this.audioList[this.audioActiveIndex]?.epname; //专辑名，原生音频播放器中的分享功能，分享出去的卡片简介，也将使用该值。
        // innerAudioContext.singer =this.audioList[this.audioActiveIndex]?.singer; //歌手名，原生音频播放器中的分享功能，分享出去的卡片简介，也将使用该值。
        // innerAudioContext.webUrl =this.audioList[this.audioActiveIndex]?.webUrl; //页面链接，原生音频播放器中的分享功能，分享出去的卡片简介，也将使用该值。
        // if (this.isAutoplay || this.autoplay) {
        //   innerAudioContext.play();
        // } else {
        //   setTimeout(() => {
        //     innerAudioContext.pause();
        //   }, 100);
        //   this.isAutoplay = true;
        // }
      }
      // if (!this.audioList[this.audioActiveIndex]?.fileUrl) {
      //   this.duration = 0;
      //   this.currentTime = 0;
      //   this.isPaused = true;
      //   innerAudioContext.pause();
      // }
    },
    // 播放事件
    onPlay() {
      console.log('音频播放事件');
      this.isEndAcudio = false;
      this.isSlidering = false;
      this.isPaused = false;
    },
    // 暂停事件
    onPause() {
      console.log("音频暂停事件");
      this.isPaused = true;
    },
    // 播放结束事件
    onEnded() {
      // 是否最后一个音频
      console.log("音频播放结束事件")
      if (this.audioActiveIndex + 1 >= this.audioList.length) {
        console.log("最后一个音频")
        this.currentTime = 0;
        this.isPaused = true;
        this.isEndAcudio = true;
        return;
      }
      // 切换到下一个音频
      this.onSwitchAudio(this.audioActiveIndex + 1);
    },
    // 音频进入可以播放状态
    onCanplay(event:any) {
      console.log('音频进入可以播放状态，但不保证后面可以流畅播放',event);
      console.log("音频时长：", innerAudioContext.duration)
      this.isSlidering = false;
      this.setPlayData();
    },
    // 播放失败
    onError(res:any) {
      console.log("音频播放错误事件", res);
      this.duration = 0;
      this.currentTime = 0;
      this.isPaused = true;
      innerAudioContext.pause();
    },
  },
})

</script>

<style lang="scss" scoped>
.audio-content {
  // width: 100vw;
  padding: 0 30rpx;
  box-sizing: border-box;

  .audio-progress {
    display: flex;
    align-items: center;
    justify-content: space-between;

    image {
      width: 48rpx;
      height: 48rpx;
    }

    .audio-bar {
      flex: 1;
      padding: 0 26rpx;
      box-sizing: border-box;

      .audio-slider {
        margin: 0 !important;
      }
    }
  }

  .bar-tips {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 0 74rpx;
    margin-bottom: 32rpx;
    color: #b8becc;
    font-size: 24rpx;
    line-height: 10rpx;
  }

  .audio-btn {
    position: relative;
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-bottom: 64rpx;

    image {
      width: 48rpx;
      height: 48rpx;
    }

    .play-box {
      display: flex;
      align-items: center;
    }

    .play-btn {
      width: 96rpx;
      height: 96rpx;
      display: flex;
      align-items: center;
      justify-content: center;
      margin: 0 64rpx;
      background: #475266;
      border-radius: 50%;

      image {
        width: 96rpx;
        height: 96rpx;
      }
    }

    .play-list-text {
      position: absolute;
      bottom: 0;
      left: -14rpx;
      color: #475266;
      font-size: 20rpx;
    }

    .speed-text {
      position: absolute;
      top: 20rpx;
      right: 0;
      color: #475266;
      font-size: 16rpx;
      font-weight: 600;
    }
  }
}
</style>