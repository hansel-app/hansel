<template>
  <div class="notification center-text">
    <div v-if="requestCount < 10">
      {{ requestCount }}
    </div>
    <div v-else>
      9+
    </div>
  </div>
</template>
<script>
import { defineComponent } from "vue";

export default defineComponent({
  props: {
    requestCount: {
      type: Number,
      required: true,
    },
  },
});
</script>
<style scoped lang="less">
.notification {
  background-color: white;
  color: black;
  height: 1.5em;
  width: 1.5em;
  margin: 0.8em 0 0 1em;
  border-radius: 0.3em;
  vertical-align: middle;
  display: inline-block;
  position: fixed;
}

.center-text {
  text-align: middle;
  line-height: 1.5em;
}
</style>
