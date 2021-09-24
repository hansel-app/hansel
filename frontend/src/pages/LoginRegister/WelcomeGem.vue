<template>
  <van-image :src="getGemImageUrl" width="32vw" height="32vw" />
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { getEnumKeyByEnumValue } from "@/utils/enum";
import { GemColor } from "@/interfaces";

export default defineComponent({
  props: {
    gemColor: {
      type: String,
      required: true,
    },
  },
  computed: {
    getGemImageUrl(): string {
      const gemColorName = getEnumKeyByEnumValue(GemColor, this.gemColor);
      return require(`@/assets/images/${gemColorName.toLowerCase()}_2048.png`);
    },
  },
});
</script>

<style scoped lang="less">
.welcome-gem {
  font-size: @circle-button-icon-size-lg;
}
</style>
