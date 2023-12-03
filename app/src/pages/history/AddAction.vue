<template>
  <view class="action-select">
    <actionGrid :showMode="'select'"></actionGrid>
  </view>
</template>

<script lang="ts">
import Vue from "vue";
import actionGrid from "@/conpoments/actions/actionGrid.vue";
export default Vue.extend({
  data() {
    return {};
  },
  methods: {},
  computed: {},
  watch: {},
  created() {},
  mounted() {},
  beforeDestroy() {},
  components: {
    actionGrid,
  },
  onLoad() {
    // 监听事件
    uni.$on("actionSelect", function (item) {
      let pages = getCurrentPages();
      let prevPage = pages[pages.length - 2];
      // prevPage.$emit("actionSelect", item);
      // console.log("actionSelect", prevPage.$vm,pages);
      prevPage.$vm.OnActionSelect(item);
      uni.navigateBack(
        {
          delta: 1,
        }
      );
    });
  },
  onUnload() {
    // 移除监听事件
    uni.$off("actionSelect");
  },
});
</script>