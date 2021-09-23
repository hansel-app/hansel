<template>
  <div class="stack-container">
    <div class="semi-circle" v-bind:style="{ backgroundColor: color }" />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  props: {
    color: String,
  },
});
</script>

<style scoped lang="less">
.stack-container {
  position: relative;
  height: 55vh;
  z-index: -1;
  overflow-x: hidden;

  .semi-circle {
    transform: translate(-29%, -50%);
    height: 100vh;
    width: 240vw;
    border-radius: 0 0 50% 50%;
    background-color: white;
    transition: background-color 1s ease;
  }
}
</style>
