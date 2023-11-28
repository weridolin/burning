<template>
  <view>
    <!-- <image
      src="https://qiniu-web-assets.dcloud.net.cn/unidoc/zh/shuijiao.jpg"
      mode="aspectFill"
      class="cover"
    /> -->
    <text class="text">动作名称</text>
  </view>
</template>
  
  <script lang="ts">
import Vue from "vue";
export default Vue.extend({
  name: "cover",
  props: {
    // msg: String
  },
  data() {
    return {
      //
    };
  },
  methods: {
    //
  },
  computed: {
    //
  },
  watch: {
    //
  },
  created() {
    //
  },
  mounted() {
    //
  },
  beforeDestroy() {
    //
  },
});
</script>
  
  
<style lang="scss">
.cover {
  flex: 0 0 80%;
  min-width: 0;
  min-height: 0;

  overflow: hidden;
  word-break: break-all;
}
.text {
  flex: 0 0 20%;
}
</style>