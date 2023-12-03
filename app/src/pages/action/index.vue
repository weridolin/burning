<template>
  <view id="content">
    <ActionGrid ref="actionGrid"></ActionGrid>  
  </view>
</template>


<script lang="ts">
import Vue from "vue";
import ActionGrid from "@/conpoments/actions/actionGrid.vue";
export default Vue.extend( {
  data() {
    return {

    }
  },
  components: {
    ActionGrid
  },
});
</script>

