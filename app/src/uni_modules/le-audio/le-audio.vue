<template>
	<view class="audio-content">
		<!-- #ifdef H5 -->
			<video
				v-if="audioList[audioActiveIndex]"
				id="le-audio" 
				style="display: none;"
				:autoplay="autoplay"
				:src="audioList[audioActiveIndex].fileUrl" 
				@error="onError" 
				@timeupdate="setPlayData"
				@loadedmetadata="onCanplay"
				@play="onPlay"
				@pause="onPause"
				@ended="onEnded"
			>
			</video>
		<!-- #endif -->
		<!-- 调速拉条区 -->
		<view class="audio-progress">
			<image src="./icon/get-back.svg" @click="onSeek(-15)"></image>
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
					block-size="18" />
			</view>
			<image src="./icon/fast-forward.svg" @click="onSeek(15)"></image>
		</view>
		<!-- 显示时间进度区 -->
		<view class="bar-tips">
			<text>{{minTimeFmt(currentTime)}}</text>
			<text>{{minTimeFmt(duration)}}</text>
		</view>
		<!-- 按钮区 -->
		<view class="audio-btn">
			<image v-if="!showAudioListIcon"></image>
			<image v-if="showAudioListIcon" src="./icon/play-list.svg" @tap="onOpenList"></image>
			<view class="play-box">
				<image src="./icon/last-episode.svg" :style="{'opacity':showLeftBtnOpacity}"
					@tap="onSwitchAudio(audioActiveIndex - 1)"></image>
				<view class="play-btn" :style="{backgroundColor:isPaused?'#00CA86':'#475266'}" @click="onPlayPause">
					<image v-show="isPaused" src="./icon/suspend.svg"></image>
					<image v-show="!isPaused" src="./icon/start.svg"></image>
				</view>
				<image src="./icon/next-set.svg" :style="{'opacity':showRightBtnOpacity}"
					@tap="onSwitchAudio(audioActiveIndex + 1)"></image>
			</view>
			<image v-if="!showAudioSpeedIcon"></image>
			<image v-if="showAudioSpeedIcon" src="./icon/multiple.svg" @tap="onSpeed"></image>
			<text v-if="showAudioListIcon" class="play-list-text" @tap="onOpenList">播放列表</text>
			<text v-if="showAudioSpeedIcon" class="speed-text" @tap="onSpeed">x{{speed}}</text>
		</view>
	</view>
</template>

<script>
	let innerAudioContext = null;
	export default {
		name: "le-audio",
		props: {
			/**
			 * audioData 播放列表
			 * @property {String} title 播放名称
			 * @property {String} fileUrl 播放链接
			 * */
			audioData: {
				type: Array,
				default: () => {
					return []
				}
			},
			// 当前播放的位置索引
			activeIndex: {
				type: Number,
				default: -1
			},
			// 是否显示播放列表
			showAudioListIcon: {
				type: Boolean,
				default: false
			},
			// 是否显示播放倍速
			showAudioSpeedIcon: {
				type: Boolean,
				default: false
			},
			// 是否自动播放
			autoplay:{
				type: Boolean,
				default: true
			}
		},
		data() {
			return {
				isPaused: true, //是否暂停中
				duration: 0, //音频时长
				currentTime: 0, //当前时长
				audioActiveIndex: -1, //当前播放索引
				speed: 1, //倍速
				isSlidering: false, //是否移动中
				isEndAcudio: false, //最后一个音频结束
				audioList: [
					// {title:'',image:'',fileUrl:''}
				],
				isAutoplay:false
			}
		},
		watch: {
			activeIndex() {
				this.audioActiveIndex = this.activeIndex;
				this.setAudioInfo();
			},
			audioData: {
				handler: function() {
					this.audioList = this.audioData;
					this.setAudioInfo();
				},
				deep: true
			}
		},
		computed: {
			showLeftBtnOpacity() {
				return this.audioActiveIndex === 0 ? 0.5 : 1;
			},
			showRightBtnOpacity() {
				return this.audioActiveIndex + 1 === this.audioList.length ? 0.5 : 1;
			},
		},
		mounted() {
			this.$nextTick(() => {
				this.startPlay();
				this.audioActiveIndex = this.activeIndex;
				this.audioList = this.audioData;
				this.setAudioInfo();
			})
		},
		beforeDestroy() {
			// #ifdef MP-ALIPAY || MP-LARK
				innerAudioContext.destroy();
			// #endif
			// #ifdef H5
				innerAudioContext.pause();
			// #endif
		},
		methods: {
			minTimeFmt(val) {
				let minute = parseInt(val / 60);
				let seconds = Math.ceil(val % 60)
				return `${minute>=10?minute:'0'+minute}:${seconds>=10?seconds:'0'+seconds}`
			},
			onOpenList() {
				this.$emit("onOpenAudioList")
			},
			startPlay() {
				// #ifdef H5
					innerAudioContext = uni.createVideoContext('le-audio');
				// #endif
				// #ifdef MP-ALIPAY || MP-LARK
					innerAudioContext = uni.createInnerAudioContext();
				// #endif
				// #ifndef H5 || MP-ALIPAY || MP-LARK
					innerAudioContext = uni.getBackgroundAudioManager();
				// #endif
				/*
				play		播放（H5端部分浏览器需在用户交互时进行）	
				pause		暂停	
				stop		停止	
				seek	position	跳转到指定位置，单位 s	
				destroy		销毁当前实例
				*/
				innerAudioContext.startTime = 0; //开始播放的位置（单位：s）
				this.setAudioInfo();
				// #ifdef H5
					innerAudioContext.playbackRate(1.0);
				// #endif
				// #ifndef H5
					innerAudioContext.playbackRate = this.speed; //放的倍率。可取值：0.5/0.8/1.0/1.25/1.5/2.0
					// 音频播放事件
					innerAudioContext.onPlay(() => {
						this.onPlay();
					});
					// 音频暂停事件
					innerAudioContext.onPause(() => {
						this.onPause();
					});
					// 音频进入可以播放状态，但不保证后面可以流畅播放
					innerAudioContext.onCanplay(() => {
						this.onCanplay();
					});
					// 音频自然播放结束事件
					innerAudioContext.onEnded(() => {
						this.onEnded();
					});
					// 音频播放错误事件
					innerAudioContext.onError((res) => {
						this.onError(res)
					});
					// 音频播放进度更新事件
					innerAudioContext.onTimeUpdate(() => {
						this.setPlayData();
					});
				// #endif
				// #ifndef H5 || MP-ALIPAY || MP-LARK
					// 用户在系统音乐播放面板点击上一曲事件（iOS only）
					innerAudioContext.onPrev(() => {
						// 切换到上一个音频
						this.onSwitchAudio(this.audioActiveIndex - 1);
						// console.log("用户在系统音乐播放面板点击下一曲事件（iOS only）");
					});
					// 用户在系统音乐播放面板点击下一曲事件（iOS only）
					innerAudioContext.onNext(() => {
						// 切换到下一个音频
						this.onSwitchAudio(this.audioActiveIndex + 1);
						// console.log("用户在系统音乐播放面板点击下一曲事件（iOS only）");
					});
				// #endif
			},
			// 暂停播放切换
			onPlayPause() {
				if (this.isPaused) {
					if (this.isEndAcudio) {
						innerAudioContext.title = this.audioList[this.audioActiveIndex]?.title;
						innerAudioContext.coverImgUrl = this.audioList[this.audioActiveIndex]?.image;
						innerAudioContext.src = encodeURI(this.audioList[this.audioActiveIndex]?.fileUrl).replace(/\+/g, "%2B"); //音频的数据链接，用于直接播放。
					}
					innerAudioContext.play()
					this.isPaused = false;
				} else {
					innerAudioContext.pause()
					this.isPaused = true;
				}
			},
			// 切换音频
			onSwitchAudio(index) {
				if (index < 0 || index >= this.audioList.length || this.audioActiveIndex === index) {
					return;
				}
				this.audioActiveIndex = index;
				this.$emit("onAudioChange", this.audioList[this.audioActiveIndex], index)
				this.setAudioInfo();
			},
			// 调整播放倍速
			onSpeed() {
				if (this.speed == 1) {
					// #ifdef H5
						innerAudioContext.playbackRate(1.5);
						this.speed = 1.5;
					// #endif
					// #ifndef H5
						innerAudioContext.playbackRate = 2.0;
						this.speed = 2;
					// #endif
				} else {
					// #ifdef H5
					innerAudioContext.playbackRate(1.0);
					// #endif
					// #ifndef H5
					innerAudioContext.playbackRate = 1.0;
					// #endif
					this.speed = 1;
				}
			},
			// 调整播放位置
			onSeek(num) {
				if (!innerAudioContext.currentTime) return;
				let seekNum = num + innerAudioContext.currentTime;
				this.isSlidering = true;
				if (seekNum < 0) {
					// 调整后时间小于0
					this.currentTime = 0;
					innerAudioContext.seek(0)
				} else if (seekNum > innerAudioContext.duration) {
					// 调整后时间大于总时长
					this.currentTime = innerAudioContext.duration;
					innerAudioContext.seek(innerAudioContext.duration)
				} else {
					this.currentTime = seekNum;
					innerAudioContext.seek(seekNum)
				}
			},
			// 滑块滚动到的位置
			sliderChange(e) {
				this.isSlidering = false;
				this.currentTime = e.detail.value;
				innerAudioContext.seek(e.detail.value);
			},
			// 滑块滚动到的位置 实时
			sliderChanging(e) {
				this.isSlidering = true;
				this.currentTime = e.detail.value;
			},
			// 设置以及转换信息
			setPlayData(event) {
				// #ifndef H5
					if (this.isSlidering) return;
					if (!innerAudioContext.duration && !innerAudioContext.currentTime) return;
					this.duration = innerAudioContext.duration || 0;
					this.currentTime = innerAudioContext.currentTime || 0;
				// #endif
				// #ifdef H5
					if (this.isSlidering) return;
					this.duration = event.detail.duration  || 0;
					this.currentTime = event.detail.currentTime || 0;
				// #endif
			},
			// 设置播放
			setAudioInfo() {
				if (innerAudioContext && this.audioActiveIndex > -1 && this.audioList[this.audioActiveIndex]?.fileUrl) {
					innerAudioContext.title = this.audioList[this.audioActiveIndex]?.title;//音频标题，用于做原生音频播放器音频标题。原生音频播放器中的分享功能，分享出去的卡片标题，也将使用该值。
					innerAudioContext.coverImgUrl = this.audioList[this.audioActiveIndex]?.image;//封面图url，用于做原生音频播放器背景图。原生音频播放器中的分享功能，分享出去的卡片配图及背景也将使用该图。
					innerAudioContext.src = encodeURI(this.audioList[this.audioActiveIndex]?.fileUrl).replace(/\+/g, "%2B"); //音频的数据链接，用于直接播放。
					innerAudioContext.epname = this.audioList[this.audioActiveIndex]?.epname;//专辑名，原生音频播放器中的分享功能，分享出去的卡片简介，也将使用该值。
					innerAudioContext.singer = this.audioList[this.audioActiveIndex]?.singer;//歌手名，原生音频播放器中的分享功能，分享出去的卡片简介，也将使用该值。
					innerAudioContext.webUrl = this.audioList[this.audioActiveIndex]?.webUrl;//页面链接，原生音频播放器中的分享功能，分享出去的卡片简介，也将使用该值。
					if(this.isAutoplay||this.autoplay){
						innerAudioContext.play();
					}else{
						setTimeout(()=>{
							innerAudioContext.pause();
						},100)
						this.isAutoplay = true;
					}
				}
				if (!this.audioList[this.audioActiveIndex]?.fileUrl) {
					this.duration = 0;
					this.currentTime = 0;
					this.isPaused = true;
					innerAudioContext.pause();
				}
			},
			// 播放事件
			onPlay(){
				// console.log('音频播放事件');
				this.isEndAcudio = false;
				this.isSlidering = false;
				this.isPaused = false;
			},
			// 暂停事件
			onPause(){
				// console.log("音频暂停事件");
				this.isPaused = true;
			},
			// 播放结束事件
			onEnded(){
				// 是否最后一个音频
				if (this.audioActiveIndex + 1 >= this.audioList.length) {
					// console.log("最后一个音频")
					this.currentTime = 0;
					this.isPaused = true;
					this.isEndAcudio = true;
					return;
				}
				// 切换到下一个音频
				this.onSwitchAudio(this.audioActiveIndex + 1);
				// console.log("音频自然播放结束事件");
			},
			// 音频进入可以播放状态
			onCanplay(event){
				// console.log('音频进入可以播放状态，但不保证后面可以流畅播放');
				// console.log("音频时长：", innerAudioContext.duration)
				this.isSlidering = false;
				this.setPlayData(event);
			},
			// 播放失败
			onError(res) {
				// console.log("音频播放错误事件", res);
				this.duration = 0;
				this.currentTime = 0;
				this.isPaused = true;
				innerAudioContext.pause();
			},
			
		}
	}
</script>

<style lang="scss" scoped>
	.audio-content {
		width: 100vw;
		padding: 0 64rpx;
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
			color: #B8BECC;
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