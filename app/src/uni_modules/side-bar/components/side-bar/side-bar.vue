<template>
  <!-- 主体部分 -->
  <view class="main">
    <!-- 左侧选项区 -->
    <swiper
      class="left"
      :duration="500"
      vertical
      :display-multiple-items="navCount"
      :style="'height:' + height + 'px'"
      :current="current"
    >
      <swiper-item
        class="nav-item"
        :class="{ ac: active == index }"
        v-for="(item, index) in tabData"
        :key="item.name"
        :item-id="item.id"
      >
        <view @click="onNav(item, index)">
          {{ item.name }}
        </view>
      </swiper-item>
    </swiper>
    <!-- 右侧内容区 -->
    <scroll-view
      class="right"
      scroll-y
      scroll-with-animation
      enable-back-to-top
      :scroll-top="scrolltop"
      :style="'height:' + height + 'px'"
      @scroll="scroll"
    >
      <view class="goods-item" v-for="(item, index) in tabData" :key="index">
        <text style="padding-left: 10rpx">{{ item.title[0] }}</text>
        <text
          style="font-family: SimHei; font-weight: bold; padding-left: 10rpx"
          >{{ item.title[1] }}</text
        >
        <u-divider
          :text="item.name"
          textPosition="left"
          textColor="#000"
        ></u-divider>
        <navigator v-for="i in item.goods" :key="i.url" :url="i.url">
          <image class="image" :src="i.img" />
        </navigator>
      </view>
    </scroll-view>
  </view>
</template>
<script>
export default {
  name: "side-bar",
  props: {
    tabData: {
      type: Array,
      default: () => [
        {
          id: "1",
          name: "女装",
          title: ["新品热卖", "New here！"],
          goods: [
            {
              img: "https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/e25dd7488f695942cfbfd9468ac2c3e7.jpg?thumb=1&w=720&h=360",
              url: "/pages/test1/test2",
            },
            {
              img: "https://m.360buyimg.com/babel/jfs/t1/156951/13/28971/148632/62eb3642E6b6fa4f1/e2f497bfb4990d74.png.webp",
              url: "/pages/index/index",
            },
          ],
        },
        {
          id: "2",
          name: "男装",
          title: ["家有萌宠", "ID Photo"],
          goods: [
            {
              img: "https://m.360buyimg.com/babel/jfs/t1/205859/5/25851/179223/62ea3417E8a4ebd7a/e9306624cd6c835c.png.webp",
              url: "/pages/test1/test2",
            },
            {
              img: "https://m.360buyimg.com/babel/jfs/t1/156951/13/28971/148632/62eb3642E6b6fa4f1/e2f497bfb4990d74.png.webp",
              url: "/pages/index/index",
            },
          ],
        },
        {
          id: "3",
          name: "奢品",
          title: ["家有萌宠", "ID Photo"],
          goods: [
            {
              img: "https://m.360buyimg.com/babel/jfs/t1/205859/5/25851/179223/62ea3417E8a4ebd7a/e9306624cd6c835c.png.webp",
              url: "/pages/test1/test2",
            },
            {
              img: "https://m.360buyimg.com/babel/jfs/t1/156951/13/28971/148632/62eb3642E6b6fa4f1/e2f497bfb4990d74.png.webp",
              url: "/pages/index/index",
            },
          ],
        },
        {
          id: "4",
          name: "女鞋",
          title: ["家有萌宠", "ID Photo"],
          goods: [
            {
              img: "https://m.360buyimg.com/babel/jfs/t1/205859/5/25851/179223/62ea3417E8a4ebd7a/e9306624cd6c835c.png.webp",
              url: "/pages/test1/test2",
            },
            {
              img: "https://m.360buyimg.com/babel/jfs/t1/156951/13/28971/148632/62eb3642E6b6fa4f1/e2f497bfb4990d74.png.webp",
              url: "/pages/index/index",
            },
          ],
        },
        {
          id: "5",
          name: "男鞋",
          title: ["家有萌宠", "ID Photo"],
          goods: [
            {
              img: "https://m.360buyimg.com/babel/jfs/t1/205859/5/25851/179223/62ea3417E8a4ebd7a/e9306624cd6c835c.png.webp",
              url: "/pages/test1/test2",
            },
            {
              img: "https://m.360buyimg.com/babel/jfs/t1/156951/13/28971/148632/62eb3642E6b6fa4f1/e2f497bfb4990d74.png.webp",
              url: "/pages/index/index",
            },
          ],
        },
        {
          id: "6",
          name: "箱包",
          title: ["家有萌宠", "ID Photo"],
          goods: [
            {
              img: "https://m.360buyimg.com/babel/jfs/t1/205859/5/25851/179223/62ea3417E8a4ebd7a/e9306624cd6c835c.png.webp",
              url: "/pages/test1/test2",
            },
            {
              img: "https://m.360buyimg.com/babel/jfs/t1/156951/13/28971/148632/62eb3642E6b6fa4f1/e2f497bfb4990d74.png.webp",
              url: "/pages/index/index",
            },
          ],
        },
        {
          id: "7",
          name: "美妆",
          title: ["家有萌宠", "ID Photo"],
          goods: [
            {
              img: "https://m.360buyimg.com/babel/jfs/t1/205859/5/25851/179223/62ea3417E8a4ebd7a/e9306624cd6c835c.png.webp",
              url: "/pages/test1/test2",
            },
            {
              img: "https://m.360buyimg.com/babel/jfs/t1/156951/13/28971/148632/62eb3642E6b6fa4f1/e2f497bfb4990d74.png.webp",
              url: "/pages/index/index",
            },
          ],
        },
        {
          id: "8",
          name: "饰品",
          title: ["家有萌宠", "ID Photo"],
          goods: [
            {
              img: "https://m.360buyimg.com/babel/jfs/t1/205859/5/25851/179223/62ea3417E8a4ebd7a/e9306624cd6c835c.png.webp",
              url: "/pages/test1/test2",
            },
            {
              img: "https://m.360buyimg.com/babel/jfs/t1/156951/13/28971/148632/62eb3642E6b6fa4f1/e2f497bfb4990d74.png.webp",
              url: "/pages/index/index",
            },
          ],
        },
        {
          id: "9",
          name: "洗护",
          title: ["家有萌宠", "ID Photo"],
          goods: [
            {
              img: "https://m.360buyimg.com/babel/jfs/t1/205859/5/25851/179223/62ea3417E8a4ebd7a/e9306624cd6c835c.png.webp",
              url: "/pages/test1/test2",
            },
            {
              img: "https://m.360buyimg.com/babel/jfs/t1/156951/13/28971/148632/62eb3642E6b6fa4f1/e2f497bfb4990d74.png.webp",
              url: "/pages/index/index",
            },
          ],
        },
        {
          id: "10",
          name: "运动",
          title: ["家有萌宠", "ID Photo"],
          goods: [
            {
              img: "https://m.360buyimg.com/babel/jfs/t1/205859/5/25851/179223/62ea3417E8a4ebd7a/e9306624cd6c835c.png.webp",
              url: "/pages/test1/test2",
            },
            {
              img: "https://m.360buyimg.com/babel/jfs/t1/156951/13/28971/148632/62eb3642E6b6fa4f1/e2f497bfb4990d74.png.webp",
              url: "/pages/index/index",
            },
          ],
        },
        {
          id: "11",
          name: "百货",
          title: ["家有萌宠", "ID Photo"],
          goods: [
            {
              img: "https://m.360buyimg.com/babel/jfs/t1/205859/5/25851/179223/62ea3417E8a4ebd7a/e9306624cd6c835c.png.webp",
              url: "/pages/test1/test2",
            },
            {
              img: "https://m.360buyimg.com/babel/jfs/t1/156951/13/28971/148632/62eb3642E6b6fa4f1/e2f497bfb4990d74.png.webp",
              url: "/pages/index/index",
            },
          ],
        },
        {
          id: "12",
          name: "手机",
          title: ["家有萌宠", "ID Photo"],
          goods: [
            {
              img: "https://m.360buyimg.com/babel/jfs/t1/205859/5/25851/179223/62ea3417E8a4ebd7a/e9306624cd6c835c.png.webp",
              url: "/pages/test1/test2",
            },
            {
              img: "https://m.360buyimg.com/babel/jfs/t1/156951/13/28971/148632/62eb3642E6b6fa4f1/e2f497bfb4990d74.png.webp",
              url: "/pages/index/index",
            },
          ],
        },
        {
          id: "13",
          name: "数码",
          title: ["家有萌宠", "ID Photo"],
          goods: [
            {
              img: "https://m.360buyimg.com/babel/jfs/t1/205859/5/25851/179223/62ea3417E8a4ebd7a/e9306624cd6c835c.png.webp",
              url: "/pages/test1/test2",
            },
            {
              img: "https://m.360buyimg.com/babel/jfs/t1/156951/13/28971/148632/62eb3642E6b6fa4f1/e2f497bfb4990d74.png.webp",
              url: "/pages/index/index",
            },
          ],
        },
        {
          id: "14",
          name: "企业礼品",
          title: ["家有萌宠", "ID Photo"],
          goods: [
            {
              img: "https://m.360buyimg.com/babel/jfs/t1/205859/5/25851/179223/62ea3417E8a4ebd7a/e9306624cd6c835c.png.webp",
              url: "/pages/test1/test2",
            },
            {
              img: "https://m.360buyimg.com/babel/jfs/t1/156951/13/28971/148632/62eb3642E6b6fa4f1/e2f497bfb4990d74.png.webp",
              url: "/pages/index/index",
            },
          ],
        },
        {
          id: "15",
          name: "家装",
          title: ["家有萌宠", "ID Photo"],
          goods: [
            {
              img: "https://m.360buyimg.com/babel/jfs/t1/205859/5/25851/179223/62ea3417E8a4ebd7a/e9306624cd6c835c.png.webp",
              url: "/pages/test1/test2",
            },
            {
              img: "https://m.360buyimg.com/babel/jfs/t1/156951/13/28971/148632/62eb3642E6b6fa4f1/e2f497bfb4990d74.png.webp",
              url: "/pages/index/index",
            },
          ],
        },
      ],
    },
    topheight: {
      type: Number,
      default: 0,
    },
  },
  data() {
    return {
      arr: [],
      scrolltop: 0,
      active: 0,
      height: "",
      navCount: 0,
      navScroll: "",
      current: 0,
    };
  },
  created() {
    this._ini();
  },
  methods: {
    onNav(item, index) {
      // 当前点击的选项
      this.active = index;
      // 当前点击的选项的对应的产品的高度
      this.scrolltop = this.arr[index];
      // console.log('当前高度:',this.arr[index]);
    },
    scroll(e) {
      //两个分类直接的相差一定要超过100px,不然要修改这里的一百,目前我们这个相差大概是在400px左右,所以我给了100px,多给点也没事
      //这一步主要是右侧滑动的时候,左侧的选项也会跟着滑动
      const index = this.arr.findIndex(
        (ele) => e.detail.scrollTop > ele && ele + 100 > e.detail.scrollTop
      );
      if (index > 0) {
        this.active = index;
        this.current = index;
      }
    },
    _ini() {
      setTimeout(() => {
        this.$nextTick(() => {
          // 若是上面需要增加搜索框等其他组件需要把他的高度获取出来 这样才能计算出我们分类的高度
          const wid = uni.getSystemInfoSync();
          // console.log('可使用窗口高度:',wid.windowHeight);
          // console.log('顶部组件this.topheight:',this.topheight);
          this.height = wid.windowHeight - this.topheight;
          // console.log('选项导航高度:',this.height);
          this.navCount = Math.round(this.height / 50);
          // 获取每个分类的高度
          uni
            .createSelectorQuery()
            .in(this)
            .selectAll(".goods-item")
            .boundingClientRect((rects) => {
              // console.log('rects:',JSON.stringify(rects));
              rects.forEach((rect) => {
                // console.log('rect.top:',rect.top);
                this.arr.push(rect.top - rects[0].top);
              });
            })
            .exec();
        });
      }, 100);
    },
  },
};
</script>
<style lang="scss" scoped>
// 左侧导航背景颜色
$leftbackground: #f7f8f9;
//左侧导航文字颜色
$leftcolor: #505660;
//左侧选中背景
$leftacbg: #48ce55;
//左侧选中文字颜色
$leftacclo: #fff;
////////////////////////////
.main {
  background-color: #f7f8f9;
  display: flex;
  .left {
    width: 30%;
    background-color: $leftbackground;
    .nav-item {
      // height: 200%;
      line-height: 104rpx;
      font-size: 28rpx;
      color: $leftcolor;
      text-align: center;
    }
    .ac {
      color: $leftacclo;
      font-size: 30rpx;
      background: $leftacbg;
    }
  }
  .right {
    .goods-item {
      margin-bottom: 20rpx;
      padding-top: 20rpx;
      background-color: #fff;
      .image {
        margin: 0 5%;
        width: 90%;
        height: 300rpx;
        border-radius: 20rpx;
        margin-bottom: 5rpx;
      }
    }
  }
}
</style>

