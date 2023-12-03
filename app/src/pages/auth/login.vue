<template>
  <view class="content">
    <view class="header">
      <image src="/static/icons/login.png"></image>
    </view>
    <view class="list">
      <view class="list-call">
        <input
          class="sl-input"
          v-model="loginForm.email"
          type="number"
          maxlength="11"
          placeholder="输入邮箱"
        />
      </view>
      <view class="list-call">
        <input
          class="sl-input"
          v-model="loginForm.count"
          type="number"
          maxlength="11"
          placeholder="输入账号"
        />
      </view>
      <view class="list-call">
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
import { setToken, setUserProfile } from "../../store/local";
export default Vue.extend({
  data() {
    return {
      loginForm: {
        phone: "",
        password: "",
        email: "",
        count: "",
      },
    };
  },
  methods: {
    bindLogin() {
      uni.request({
        url: "http://43.128.110.230:30080/usercenter/api/v1/login",
        data: this.loginForm,
        header: {
          // Accept: 'application/json',
          "Content-Type": "application/x-www-form-urlencoded",
          // 'X-Requested-With': 'XMLHttpRequest'
        },
        method: "POST",
        // sslVerify: true,
        success: ({ data, statusCode, header }) => {
          // access_token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDIxMjQ2NTUsImlkIjoxLCJzdXBlckFkbWluIjpmYWxzZSwidXNlcm5hbWUiOiJ3ZXJpZG8ifQ.vz2wCIED7ApRfs6A7AeqhG7YCDrPLsAoOhaIxEDVqn8"
          // age: 0
          // avatar: ""
          // email: "35920664312@qq.com"
          // gender: 0
          // is_super_admin: false
          // phone: ""
          // refresh_token: ""
          // role: null
          // username: "werido"
          console.log("登录成功", data);
          setToken({
            access_token: data.data.access_token,
            refresh_token: data.data.refresh_token,
          });
          uni.switchTab({
            url: "/pages/index/index",
          });
        },
        fail: (error) => {
          console.log("登录失败", error);
        },
      });
      console.log(this.state.login);
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
  /* width: 161rpx;
		height: 161rpx; */
  border-radius: 50%;
  margin-top: 30rpx;
  margin-left: auto;
  margin-right: auto;
}

.header image {
  width: 100rpx;
  height: 100rpx;
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
