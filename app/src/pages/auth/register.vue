<template>
  <view class="content">
    <view class="header">
      <image src="/static/images/login/logo.jpg"></image>
    </view>

    <view class="list">
      <view class="list-call">
        <input
          class="sl-input"
          v-model="registerForm.email"
          type="number"
          maxlength="11"
          placeholder="邮箱"
        />
      </view>
      <view class="list-call">
        <input
          class="sl-input"
          v-model="registerForm.username"
          type="text"
          maxlength="11"
          placeholder="用户名"
        />
      </view>
      <view class="list-call">
        <input
          class="sl-input"
          v-model="registerForm.password"
          type="text"
          maxlength="32"
          placeholder="登录密码"
          :password="!showPassword"
        />
      </view>
      <view class="list-call">
        <input
          class="sl-input"
          v-model="registerForm.checkcode"
          type="number"
          maxlength="4"
          placeholder="验证码"
        />
        <view class="yzm" :class="{ yzms: state.second > 0 }" @tap="getcode">{{
          yanzhengma
        }}</view>
      </view>
    </view>

    <view class="button-login" hover-class="button-hover" @tap="register">
      <text>注册</text>
    </view>
  </view>
</template>

<script>
import Vue from "vue";
import { RegisterForm } from "./apis";
export default Vue.extend({
  data() {
    var registerForm = {
      email: "",
      password: "",
      username: "",
      checkcode: "",
    };
    return {
      registerForm,
      showPassword: false,
      state: { second: 0 },
    };
  },
  computed: {
    yanzhengma() {
      if (this.state.second > 0) {
        if (this.state.second < 10) {
          return "重新获取0" + this.state.second;
        } else {
          return "重新获取" + this.state.second;
        }
      } else {
        return "获取验证码";
      }
    },
  },
  methods: {
    register() {
      uni.showToast({ title: "测试阶段不开放注册", icon: "error" ,duration:10000 });
      return;
    },
    getcode() {
      uni.showToast({ title: "测试阶段不开放注册", icon: "error" ,duration:10000 });
      return;
      if (this.state.second > 0) {
        return;
      }
      this.state.second = 60;
      //获取验证码 #TODO
      setInterval(() => {
        let that = this;
        that.state.second--;
        if (that.state.second == 0) {
          // this.state.clear()
        }
      }, 1000);
      // uni.reque
      console.log("获取邮箱验证码");
    },
  },
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
  color: #ff7d13;
  font-size: 24rpx;
  line-height: 64rpx;
  padding-left: 10rpx;
  padding-right: 10rpx;
  width: auto;
  height: 64rpx;
  border: 1rpx solid #ffa800;
  border-radius: 50rpx;
}

.yzms {
  color: #999999 !important;
  border: 1rpx solid #999999;
}

.button-login {
  color: #ffffff;
  font-size: 34rpx;
  width: 470rpx;
  height: 100rpx;
  line-height: 100rpx;
  background: linear-gradient(
    -90deg,
    rgba(193, 25, 32, 1),
    rgba(238, 38, 38, 1)
  );
  border-radius: 50rpx;
  text-align: center;
  margin-left: auto;
  margin-right: auto;
  margin-top: 100rpx;
}

.button-hover {
  background: linear-gradient(
    -90deg,
    rgba(193, 25, 32, 0.8),
    rgba(238, 38, 38, 0.8)
  );
}
</style>
