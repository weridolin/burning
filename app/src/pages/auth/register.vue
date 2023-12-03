<template>
	<view class="content">
		<view class="header">
			<image src="/static/images/login/logo.jpg"></image>
		</view>

		<view class="list">
			<view class="list-call">
				<input class="sl-input" v-model="phone" type="number" maxlength="11" placeholder="手机号" />
			</view>
			<view class="list-call">
				<input class="sl-input" v-model="password" type="text" maxlength="32" placeholder="登录密码"
					:password="!showPassword" />
			</view>
			<view class="list-call">
				<input class="sl-input" v-model="code" type="number" maxlength="4" placeholder="验证码" />
				<view class="yzm" :class="{ yzms: second>0 }" @tap="getcode">{{yanzhengma}}</view>
			</view>

		</view>

		<view class="button-login" hover-class="button-hover" @tap="bindLogin">
			<text>注册</text>
		</view>
	</view>
</template>

<script setup>
	import {
		reactive,
		onUnmounted,
		computed
	} from 'vue';
	const state = reactive({
		phone: '',
		password: '',
		code: '',
		showPassword: false,
		second: 0
	});

	function clear() {
		clearInterval(js)
		js = null
		state.second = 0
	}

	function display() {
		state.showPassword = !state.showPassword
	}

	onUnmounted(() => {
		clear()
	});
	// 注册
	function bindLogin() {

	}

	function getcode() {
		if (state.phone.length != 11) {
			uni.showToast({
				icon: 'none',
				title: '手机号不正确'
			});
			return;
		}
		if (state.second > 0) {
			return;
		}
		state.second = 60;
		//请求业务
		js = setInterval(function() {
			state.second--;
			if (state.second == 0) {
				state.clear()
			}
		}, 1000)
		// uni.request({
		//   url: 'http://***/getcode.html', //仅为示例，并非真实接口地址。
		//   data: {
		//     phone: this.phone,
		//     type: 'reg'
		//   },
		//   method: 'POST',
		//   dataType: 'json',
		//   success: (res) => {
		//     if (res.data.code != 200) {
		//       uni.showToast({
		//         title: res.data.msg,
		//         icon: 'none'
		//       });
		//     } else {
		//       uni.showToast({
		//         title: res.data.msg
		//       });
		//       js = setInterval(function() {
		//         _this.second--;
		//         if (_this.second == 0) {
		//           _this.clear()
		//         }
		//       }, 1000)
		//     }
		//   },
		//   fail() {
		//     this.second == 0
		//   }
		// });
	}
	const yanzhengma = computed(() => {
		if (state.second == 0) {
			return '获取验证码';
		} else {
			if (state.second < 10) {
				return '重新获取0' + state.second;
			} else {
				return '重新获取' + state.second;
			}
		}
	});
</script>

<style>
	page {
		background-color: #fff;
	}

	.content {
		display: flex;
		flex-direction: column;
		justify-content: center;
	}

	.header {
		width: 161rpx;
		height: 161rpx;
		border-radius: 50%;
		margin-top: 30rpx;
		margin-left: auto;
		margin-right: auto;
	}

	.header image {
		width: 161rpx;
		height: 161rpx;
		border-radius: 50%;
	}

	.list {
		display: flex;
		flex-direction: column;
		padding-top: 50rpx;
		padding-left: 70rpx;
		padding-right: 70rpx;
	}

	.list-call {
		display: flex;
		flex-direction: row;
		justify-content: space-between;
		align-items: center;
		height: 100rpx;
		color: #333333;
		border-bottom: 0.5px solid #e2e2e2;
	}



	.list-call .sl-input {
		flex: 1;
		text-align: left;
		font-size: 32rpx;
		margin-left: 16rpx;
	}

	.yzm {
		color: #FF7D13;
		font-size: 24rpx;
		line-height: 64rpx;
		padding-left: 10rpx;
		padding-right: 10rpx;
		width: auto;
		height: 64rpx;
		border: 1rpx solid #FFA800;
		border-radius: 50rpx;
	}

	.yzms {
		color: #999999 !important;
		border: 1rpx solid #999999;
	}

	.button-login {
		color: #FFFFFF;
		font-size: 34rpx;
		width: 470rpx;
		height: 100rpx;
		line-height: 100rpx;
		background: linear-gradient(-90deg, rgba(193, 25, 32, 1), rgba(238, 38, 38, 1));
		border-radius: 50rpx;
		text-align: center;
		margin-left: auto;
		margin-right: auto;
		margin-top: 100rpx;
	}

	.button-hover {
		background: linear-gradient(-90deg, rgba(193, 25, 32, 0.8), rgba(238, 38, 38, 0.8));
	}
</style>
