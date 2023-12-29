<template>
  <view class="diet-form-container" style="width: 100%;">
    <uni-section class="mb-10" :title='status=="update"?"更新饮食记录":"创建的饮食记录"'>
    <!-- 基础表单校验 -->
    <uni-forms ref="diet-form"  v-model="dietContentItem">
      <uni-forms-item label="时间" required name="created_at">
        <uni-datetime-picker type="datetime" :clear-icon="false" v-model="dietContentItem.created_at"/>
      </uni-forms-item>
      <uni-forms-item label="脂肪(g)" required name="fat">
        <uni-easyinput v-model="dietContentItem.fat" placeholder="请输入脂肪摄入量" />
      </uni-forms-item>
      <uni-forms-item label="碳(g)" required name="carbon">
        <uni-easyinput v-model="dietContentItem.carbon" placeholder="请输入碳摄入量" />
      </uni-forms-item>
      <uni-forms-item label="蛋白质(g)" required name="protein">
        <uni-easyinput v-model="dietContentItem.protein" placeholder="请输入蛋白质摄入量" />
      </uni-forms-item>
      <uni-forms-item label="水分(g)" required name="water">
        <uni-easyinput v-model="dietContentItem.water" placeholder="请输入水摄入量" />
      </uni-forms-item>
      <uni-forms-item label="热量(卡路里)" required name="calorie">
        <uni-easyinput v-model="dietContentItem.calorie" placeholder="请输入热量摄入量" />
      </uni-forms-item>
      <uni-forms-item >
        <button type="primary" @click="submit">提交</button>
      </uni-forms-item>
    </uni-forms>
    </uni-section>
  </view>
</template>


<script lang="ts">
import Vue from "vue";
import { DietContentItem,CreateNewDietHistory,UpdateDietHistory } from "@/pages/history/apis";

export default Vue.extend({
  data() {
    var dietContentItem: DietContentItem = {} as DietContentItem;
    return { 
      dietContentItem,
      status: "create", // create or update
    };
  },
  props: {
    // status: {
    //   type: String,
    //   default: "create"
    // },
    dietId: {
      type: Number,
      default: 0
    },
    dietContentItemProp: {
      type: Object,
      default: {}
    }
  },
  watch: {
    dietContentItemProp: {
      handler: function (val: DietContentItem, oldVal: DietContentItem) {
        // console.log("diet content card dietContentItemProp changed -> ", val, oldVal);
        this.dietContentItem = val;
        // this.$forceUpdate()
        // this.$emit('update:dietContentItemProp',this.dietContentItem)
      },
      immediate: true,
      deep:true
    },
  },
  onShow() {
    // if (this.status=="update"){
    //   this.dietContentItem = this.dietContentItem as DietContentItem;
    // }
  },
  onReady() {
    // if (this.status=="update"){
    //   console.log("update diet history dietContentItemProp -> ",this.dietContentItemProp);
    //   this.dietContentItem = this.dietContentItemProp as DietContentItem;
    // }
  },
  methods: {
    submit() {
      if (this.status=="create"){
        CreateNewDietHistory(this.dietContentItem,(res) => {
          console.log(res);
          uni.showToast({
            title: "创建成功",
            icon: "success",
            duration: 2000,
          });
          this.$emit('dietContentCreated',this.dietContentItem)
        },(err) => {
          console.log("create diet history failed -> ",err);
          uni.showToast({
            title: "创建失败",
            icon: "error",
            duration: 2000,
          });
        });
      }else {
        UpdateDietHistory(this.dietContentItem.id as number,this.dietContentItem,(res) => {
          console.log(res);
          uni.showToast({
            title: "更新成功",
            icon: "success",
            duration: 2000,
          });
          this.$emit('dietContentUpdated',this.dietContentItem)
        },(err) => {
          console.log("update diet history failed -> ",err);
          uni.showToast({
            title: "更新失败",
            icon: "error",
            duration: 2000,
          });
        });
      }

    },
  },
  mounted() {},
});
</script>