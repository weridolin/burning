<template>
  <view class="container">
    <view class="container__header">
      <image src="/static/icons/login1.png"></image>
    </view>
    <view class="container__list">
      <!-- <view class="list-call">
        <input
          class="sl-input"
          v-model="loginForm.email"
          type="number"
          maxlength="11"
          placeholder="输入邮箱"
        />
      </view> -->
      <view class="container__list-call">
        <input
          class="sl-input"
          v-model="loginForm.count"
          type="text"
          maxlength="11"
          placeholder="输入账号/邮箱/手机号"
        />
      </view>
      <view class="container__list-call">
        <input
          class="sl-input"
          v-model="loginForm.password"
          type="text"
          maxlength="32"
          placeholder="输入密码"
          password="true"
        />
      </view>
    </view>

    <view class="button-login" hover-class="button-hover" @tap="bindLogin()">
      <text>登录</text>
    </view>
    <view class="agreenment">
      <navigator url="/pages/auth/register" open-type="navigate"
        >注册账户</navigator
      >
    </view>
  </view>
</template>

<script lang="ts">
import Vue from "vue";
import { setToken, setUserProfile, UserBodyInfo } from "../../store/local";
import {
  LoginRequestForm,
  Login,
  LoginResponsePayload,
  GetUserProfileResponsePayload,
} from "./apis";
import { GetUserProfile } from "@/pages/person/apis";
// import {}

export default Vue.extend({
  data() {
    var loginForm: LoginRequestForm = {
      phone: "",
      password: "",
      email: "",
      count: "",
    };
    return {
      loginForm,
    };
  },
  methods: {
    bindLogin() {
      uni.showLoading({
        title: '登陆中',
        mask: true
      })
      Login(
        this.loginForm,
        (res) => {
          console.log("登录请求结果 -> ", res);
          uni.hideLoading()
          if (res.code != 0) {
            uni.showToast({
              title: "账号或密码错误",
              icon: "error",
              duration: 2000,
            });
            return;
          }
          let loginData = res.data as LoginResponsePayload;
          setToken({
            access_token: loginData.access_token,
            refresh_token: loginData.refresh_token,
          });
          GetUserProfile(
            (res) => {
              console.log("获取用户信息 -> ", res);
              let userBodyProfile = res.data as GetUserProfileResponsePayload;
              setUserProfile({
                name: loginData.username,
                avatar: loginData.avatar,
                email: loginData.email,
                phone: loginData.phone,
                gender: loginData.gender,
                uuid: userBodyProfile.uuid,
                days: userBodyProfile.days,
                // bodyInfo: userBodyProfile
              });
              uni.showToast({
                title: "登录成功",
                icon: "success",
                duration: 2000,
              });
              uni.navigateBack({ delta: 1 });
            },
            (error) => {
              console.log("获取用户信息异常 -> ", error);
              uni.showToast({
                title: "获取用户信息异常",
                icon: "error",
                duration: 2000,
              });
            }
          );
        },
        (error) => {
          console.log("登录异常 -> ", error);
          uni.hideLoading()
          uni.showToast({
            title: "登录异常",
            icon: "error",
            duration: 2000,
          });
        }
      );
    },
  },
});
</script>

<style lang="scss">
page {
  background-color: #fff;
}

.container {
  display: flex;
  flex-direction: column;
  justify-content: center;
  .container__header {
    /* width: 161rpx;
		height: 161rpx; */
    border-radius: 50%;
    margin-top: 30rpx;
    margin-left: auto;
    margin-right: auto;
  }

  .container__header image {
    width: 100rpx;
    height: 100rpx;
    border-radius: 50%;
  }

  .container__list {
    display: flex;
    flex-direction: column;
    padding-top: 50rpx;
    padding-left: 70rpx;
    padding-right: 70rpx;
    .container__list-call {
      display: flex;
      flex-direction: row;
      justify-content: space-between;
      align-items: center;
      height: 100rpx;
      color: #333333;
      border-bottom: 0.5px solid #e2e2e2;
    }

    .container__list-call .sl-input {
      flex: 1;
      text-align: left;
      font-size: 32rpx;
      margin-left: 16rpx;
    }
  }
}

.button-login {
  color: #ffffff;
  font-size: 34rpx;
  width: 470rpx;
  height: 100rpx;
  background: linear-gradient(
    -90deg,
    rgba(193, 25, 32, 1),
    rgba(238, 38, 38, 1)
  );
  border-radius: 50rpx;
  line-height: 100rpx;
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
.agreenment {
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
  font-size: 30rpx;
  margin-top: 80rpx;
  color: #ffa800;
  text-align: center;
  height: 40rpx;
  line-height: 40rpx;
}

.agreenment text {
  font-size: 24rpx;
  margin-left: 15rpx;
  margin-right: 15rpx;
}
</style>
