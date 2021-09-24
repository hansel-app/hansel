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
import SemiCircleBg from "@/components/SemiCircleBg.vue";

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
.gem-image {
  width: 28vh;
  height: 28vh;
  margin-top: 16vh;
}
</style>
