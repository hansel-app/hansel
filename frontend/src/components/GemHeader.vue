<template>
  <van-row justify="center">
    <p class="big-header overlay">{{ title }}</p>
  </van-row>
  <van-row justify="center">
    <Image class="gem-image overlay" :src="gemImgSrc" :show-loading="false" />
  </van-row>
  <div class="stack-container">
    <div class="semi-circle" v-bind:style="backgroundColor" />
  </div>
</template>
<script lang="ts">
import { defineComponent, PropType } from "vue";
import { Image } from "vant";
import { GemColor } from "@/interfaces";
import { getEnumKeyByEnumValue } from "@/utils/enum";

export default defineComponent({
  components: { Image },
  props: {
    title: {
      type: String,
      required: true,
    },
    color: {
      type: String as PropType<GemColor>,
      required: true,
    },
  },

  computed: {
    gemImgSrc() {
      const gemColorName = getEnumKeyByEnumValue(GemColor, this.color);
      return require(`@/assets/images/${gemColorName.toLowerCase()}_2048.png`);
    },
    backgroundColor() {
      return {
        backgroundColor: this.color,
      };
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

.overlay {
  position: absolute;
  left: 0;
  right: 0;
  margin-left: auto;
  margin-right: auto;
}

.gem-image {
  width: 20vh;
  height: 20vh;
  margin-top: 15vh;
}
</style>
