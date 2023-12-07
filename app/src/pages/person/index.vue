<template>
  <view class="homepage">
    <view class="bar">
      <u-navbar :placeholder="true" bgColor="#efeefd">
        <view slot="left"></view>
        <view slot="right" class="right">
          <view>
            <image src="../../static/image/travel/personal/edit.png"></image>
          </view>
          <!-- <view>
						<image src="../../static/image/travel/personal/set.png"></image>
					</view> -->
        </view>
      </u-navbar>
    </view>
    <view class="people">
      <view class="headImg">
        <image src="../../static/image/travel/personal/tx.png" />
      </view>
      <view class="info">
        <view class="nick">
          <text>{{ profile.name }}</text>
          <view class="sex">
            <u-icon name="man" color="#ffffff" size="24"></u-icon>
          </view>
        </view>
        <view class="userId">
          <image src="../../static/image/travel/personal/id.png" />
          <view class="number">
            <text>{{ user_uuid }}</text>
            <text>复制</text>
          </view>
        </view>
      </view>
      <view class="space"> </view>
    </view>
    <view class="infos">
      <view class="open-vip attendance">
        <image src="../../static/image/travel/personal/attendance.png" />
        <text class="text">目前已经连续签到{{ sign_days}}天</text>
        <text class="button" @click="sign">{{alreadySign?"已签到":"签到"}}</text>
      </view>
      <view class="tool">
        <view>
          <image src="../../static/image/travel/personal/profile.png" />
          <text>身体数据</text>
        </view>
        <view>
          <image src="../../static/image/travel/personal/photos.png" />
          <text>照片记录</text>
        </view>
        <view>
          <image src="../../static/image/travel/personal/house.png" />
          <text>我的收藏</text>
        </view>
        <view>
          <image src="../../static/image/travel/personal/setting.png" />
          <text>账号设置</text>
        </view>
      </view>
      <view class="service">
        <view>
          <image
            class="icon"
            src="../../static/image/travel/personal/pic07.png"
          />
          <text>在线客服</text>
          <image
            class="right"
            src="../../static/image/travel/personal/Clipped.png"
          />
        </view>
        <view>
          <image
            class="icon"
            src="../../static/image/travel/personal/share.png"
          />
          <text>分享小程序</text>
          <image
            class="right"
            src="../../static/image/travel/personal/Clipped.png"
          />
        </view>
      </view>
    </view>
  </view>
</template>

<script lang="ts">
import Vue from "vue";
import {
  UserProfile,
  isLogin,
  getUserProfile,
  UserBodyInfo,
} from "@/store/local";
import { GetUserProfile, Sign } from "@/pages/person/apis";
import { setUserProfile,getString,setString } from "../../store/local";
import {getDate} from "@/pages/history/apis"
 
export default Vue.extend({
  data() {
    // var profile: UserProfile = getUserProfile();
    var profile: UserProfile = {} as UserProfile;
    return {
      value: 1,
      value1: 0,
      alreadySign: false,
      profile,
    };
  },
  computed: {
    user_uuid() {
      return this.profile.bodyInfo?.uuid.slice(1, 7);
    },
    sign_days() {
      return this.profile.bodyInfo?.days;
    },
  },
  onLoad() {
    let lastSignDate = getString("lastSignDate");
    let today = getDate(new Date(),0).fullDate;
    if (today==lastSignDate){
      this.alreadySign = true;
    }else{
      this.alreadySign = false;
    }
  },
  onShow() {
    if (!isLogin()) {
      uni.navigateTo({
        url: "/pages/auth/login",
      });
    } else {
      // 获取用户档案信息
      GetUserProfile(
        (res) => {
          console.log("获取用户档案信息 -> ", res);
          let userBodyProfile = res.data as UserBodyInfo;
          // this.profile = userBodyProfile;
          let profile = getUserProfile();
          if (profile != null) {
            profile.bodyInfo = userBodyProfile;
            setUserProfile(profile);
            this.profile = profile;
          }
          console.log("user profile -> ", profile);
        },
        (err) => {
          console.log("获取用户档案信息失败 -> ", err);
          uni.showToast({
            title: "更新用户信息失败",
            icon: "error",
          });
          uni.navigateTo({
            url: "/pages/auth/login",
          });
        }
      );
    }
  },
  methods: {
    sign() {
      Sign(
        (res) => {
          console.log("签到成功 -> ", res);
          uni.showToast({
            title: "签到成功",
            icon: "success",
          });
          let profile = getUserProfile();
          if (profile != null && profile.bodyInfo != null) {
            profile.bodyInfo.days = profile.bodyInfo.days + 1;
            setUserProfile(profile);
            this.profile = profile;
          }
          // let lastSignDate = getString("lastSignDate");
          this.alreadySign=true;
          let today = getDate(new Date(),0).fullDate;
          setString("lastSignDate",today);
        },
        (err) => {
          console.log("签到失败 -> ", err);
          uni.showToast({
            title: "签到失败",
            icon: "error",
          });
        }
      );
    },
  },
});
</script>
<style lang="scss" scoped>
.homepage {
  width: 100%;
  height: 1624rpx;
  background: linear-gradient(
    1deg,
    rgba(247, 240, 255, 0.19) 0%,
    rgba(229, 232, 255, 0.2) 68%,
    rgba(231, 228, 255, 0.66) 100%
  );

  & text {
    color: #333333;
    font-family: PingFangSC-Semibold, PingFang SC;
  }

  .bar {
    .right {
      display: flex;

      & view {
        width: 52rpx;
        height: 52rpx;
        margin-left: 26rpx;
        border-radius: 26rpx;
        background: #f0eeff;
        display: flex;
        justify-content: center;
        align-items: center;
        box-shadow: 0rpx -2rpx 2rpx 4rpx rgba(255, 255, 255, 0.5),
          0rpx 4rpx 4rpx 0rpx rgba(197, 183, 211, 0.5),
          inset 0rpx 2rpx 6rpx 0rpx rgba(255, 255, 255, 0.5);

        > image {
          width: 32rpx;
          height: 32rpx;
        }
      }
    }
  }

  .people {
    padding: 0 42rpx 28rpx 32rpx;
    display: flex;
    align-items: center;

    .headImg {
      > image {
        width: 166rpx;
        height: 166rpx;
        border-radius: 83rpx;
      }
    }

    .info {
      flex: 1;

      .nick {
        display: flex;

        > text {
          font-size: 36rpx;
          font-weight: 600;
          line-height: 50rpx;
        }

        .sex {
          width: 24rpx;
          height: 24rpx;
          border-radius: 12rpx;
          background: #61c9fd;
        }
      }

      .grade {
        display: flex;
        align-items: center;

        > view {
          display: flex;
          align-items: center;
          margin-right: 12rpx;

          & text {
            font-size: 20rpx;
            font-weight: 600;
            color: #ffffff;
            line-height: 28rpx;
            text-shadow: 0rpx 2rpx 4rpx rgba(0, 0, 0, 0.5);
          }

          & image {
            width: 28rpx;
            height: 30rpx;
          }

          &:last-child {
            > image {
              width: 40rpx;
              height: 40rpx;
            }

            > text {
              margin-left: -6rpx;
            }
          }
        }
      }

      .userId {
        width: 220rpx;
        display: flex;
        background: #f5f5ff;
        border-radius: 8rpx;
        box-shadow: 0rpx 2rpx 6rpx 0rpx rgba(0, 0, 0, 0.14),
          0rpx -4rpx 6rpx 0rpx #ffffff;
        margin-top: 10px;
        > image {
          width: 36rpx;
          height: 40rpx;
        }

        .number {
          flex: 1;
          display: flex;
          justify-content: center;

          > text {
            font-size: 24rpx;
            font-weight: 600;
            line-height: 40rpx;

            &:last-child {
              font-weight: 500;
              font-size: 22rpx;
              margin-left: 8rpx;
            }
          }
        }
      }
    }

    .space {
      display: flex;
      align-items: center;

      > text {
        font-size: 28rpx;
        line-height: 40rpx;
      }
    }
  }

  .list {
    width: 100%;
    display: flex;
    padding: 0 44rpx;
    box-sizing: border-box;

    .item {
      width: 25%;
      display: flex;
      justify-content: space-evenly;
      align-items: center;

      .text {
        display: flex;
        flex-direction: column;
        align-items: center;

        > text:first-child {
          font-size: 36rpx;
          font-family: CloudHeiChaoGBK;
          line-height: 48rpx;
          font-weight: 600;
        }

        > text:last-child {
          font-size: 24rpx;
          color: #999999;
          line-height: 34rpx;
        }
      }
    }
  }

  .infos {
    padding: 0 40rpx;

    .open-vip {
      width: 100%;
      height: 72rpx;
      background: linear-gradient(
        180deg,
        #a7c1b3 0%,
        #91d6b0 2%,
        #18b7c7cc 100%
      );
      border-radius: 49rpx;
      display: flex;
      align-items: center;
      margin-top: 36rpx;
      padding: 0 24rpx 0 34rpx;
      box-sizing: border-box;

      > image {
        width: 48rpx;
        height: 48rpx;
      }

      .text {
        flex: 1;
        font-size: 24rpx;
        line-height: 34rpx;
        margin-left: 14rpx;
      }

      .button {
        width: 128rpx;
        height: 42rpx;
        background: linear-gradient(90deg, #4d4d4d 0%, #151515 100%);
        border-radius: 22rpx;
        font-size: 22rpx;
        color: #ffdfa9;
        line-height: 42rpx;
        text-align: center;
      }
    }

    .tool {
      display: flex;
      width: 100%;
      height: 172rpx;
      background: #ffffff;
      box-shadow: 0rpx 2rpx 22rpx 0rpx rgba(114, 118, 206, 0.27);
      border-radius: 28rpx;
      justify-content: space-evenly;
      margin: 22rpx 0;

      > view {
        display: flex;
        flex-direction: column;

        & text {
          font-size: 22rpx;
          font-weight: 600;
          color: #666666;
          line-height: 32rpx;
          text-shadow: 0rpx 4rpx 14rpx rgba(222, 148, 247, 0.6);
          margin-left: 3px;
        }

        & image {
          width: 98rpx;
          height: 96rpx;
          margin-top: 10rpx;
        }
      }
    }

    .set {
      width: 100%;
      padding: 34rpx 24rpx 44rpx 34rpx;
      background: #ffffff;
      box-shadow: 0rpx 2rpx 28rpx 0rpx rgba(142, 146, 230, 0.27);
      border-radius: 28rpx;
      display: flex;
      flex-direction: column;
      justify-content: space-between;
      box-sizing: border-box;

      > view {
        display: flex;
        align-items: center;
        margin-bottom: 40rpx;

        & text {
          flex: 1;
          font-size: 28rpx;
          line-height: 40rpx;
          margin-left: 30rpx;
        }

        .icon {
          width: 36rpx;
          height: 36rpx;
        }

        .right {
          width: 40rpx;
          height: 40rpx;
        }
      }
    }

    .service {
      background: #ffffff;
      box-shadow: 0rpx 2rpx 28rpx 0rpx rgba(142, 146, 230, 0.27);
      border-radius: 28rpx;
      margin-top: 26rpx;
      padding: 34rpx 24rpx 44rpx 34rpx;
      display: flex;
      flex-direction: column;
      justify-content: space-between;

      > view {
        display: flex;
        align-items: center;
        margin-bottom: 40rpx;

        & text {
          flex: 1;
          font-size: 28rpx;
          line-height: 40rpx;
          margin-left: 30rpx;
        }

        .icon {
          width: 36rpx;
          height: 36rpx;
        }

        .right {
          width: 40rpx;
          height: 40rpx;
        }
      }
    }
  }
}
</style>