<template>
  <van-row justify="center">
    <p class="big-header overlay">{{ title }}</p>
  </van-row>
  <van-row justify="center">
    <Image class="gem-image overlay" :src="gemImgSrc" :show-loading="false" />
  </van-row>
  <SemiCircleBg :color="color" />
</template>
<script lang="ts">
import { defineComponent, PropType } from "vue";
import { Image } from "vant";
import { GemColor, HexCode } from "@/interfaces";
import { getEnumKeyByEnumValue } from "@/utils/enum";
import SemiCircleBg from "./SemiCircleBg.vue";

export default defineComponent({
  components: { SemiCircleBg, Image },
  props: {
    title: {
      type: String,
      required: true,
    },
    color: {
      type: String as PropType<HexCode>,
      required: true,
    },
  },

  computed: {
    gemImgSrc() {
      const gemColorName = getEnumKeyByEnumValue(GemColor, this.color);
      return require(`@/assets/images/${gemColorName.toLowerCase()}_2048.png`);
    },
  },
});
</script>

<style scoped lang="less">
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

.gem-image {
  width: 20vh;
  height: 20vh;
  margin-top: 15vh;
}
</style>
