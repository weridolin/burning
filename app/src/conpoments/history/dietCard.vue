<template>
  <uni-card :title="dietContentItem.created_at" >
  <view class="food-detail">
    <view class="food-total">
      <text>已经进食{{ dietContentItem.calorie }}卡路里</text>
    </view>
    <view class="type-detail">
      <text class="type protein">蛋白质:{{ dietContentItem.protein }}g</text>
      <text class="type fat">脂肪:{{ dietContentItem.fat }}g</text>
      <text class="type carbohydrate new-row"
        >碳:{{ dietContentItem.carbon }}g</text
      >
      <text class="type water new-row">水:{{ dietContentItem.water }}g</text>
    </view>
  </view>
</uni-card>  
</template>

<script lang="ts">
import Vue from 'vue'
import { DietContentItem } from '@/pages/history/apis'
export default Vue.extend({
  data() {
    var dietContentItem: DietContentItem = {} as DietContentItem
    return {
      dietContentItem
    }
  },
  onReady() {
    console.log('dietContentItemProp -> ', this.dietContentItemProp)
    this.dietContentItem = this.dietContentItemProp as DietContentItem
  },
  props: {
    dietContentItemProp: {
      type: Object,
      default: {}
    }
  },
  watch:{
    dietContentItemProp: {
      handler: function (val: DietContentItem, oldVal: DietContentItem) {
        console.log("diet card dietContentItemProp changed -> ", val, oldVal);
        this.dietContentItem = val;
        // this.$forceUpdate()
        // this.$emit('update:dietContentItemProp',this.dietContentItem)
      },
      immediate: true,
      deep:true
    },
  },
  methods: {
    
  }
})
</script>

<style lang="scss" scoped>
.food-detail {
  display: flex;
  flex-direction: row;

  .food-total {
    flex: 0 1 50%;
    margin-top: 20px;
  }
  .type-detail {
    flex: 0 1 50%;
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
    /* flex-start:  默认从头开始排列*/
    justify-content: flex-start;
    .type {
      flex: 0 1 50%;
      width: 50%;
      text-align: center;
    }
    .new-row {
      margin-top: 10px;
      text-align: center;
    }
  }
}

</style>