<template>
  <SemiCircleBg class="semi-circle overlay" />
  <div class="blue-gradient-bg" />
</template>

<script lang="ts">
import { defineComponent } from "vue";
import SemiCircleBg from "../../components/SemiCircleBg.vue";

export default defineComponent({
  components: {
    SemiCircleBg,
  },
});
</script>

<style SemiCircleBg scoped lang="less">
.blue-gradient-bg {
  background: linear-gradient(to bottom, #c0c2e5, #e43d9700);
  height: 100vh;
  position: relative;
}

.semi-circle {
  background-color: #ffffff;
  position: absolute;
}
</style>
