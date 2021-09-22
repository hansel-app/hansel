<template>
  <div class="stack-container">
    <div class="semi-circle" v-bind:style="backgroundColor" />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
    props: {
        color: String
    }
})
</script>


<style scoped>
.stack-container {
  position: relative;
  height: 45vh;
  z-index: -1;
  overflow-x: hidden;

  .semi-circle {
    transform: translate(-25%, -50%);
    height: 80vh;
    width: 200vw;
    border-radius: 0 0 50% 50%;
    background-color: #e4bfbf;
    transition: background-color 1s ease;
  }
}
</style>
