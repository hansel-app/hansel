<template>
  <van-image :src="getGemImageUrl" />
</template>

<script lang="ts">
import { defineComponent, PropType } from "vue";
import { getEnumKeyByEnumValue } from "@/utils/enum";
import { GemColor } from "@/interfaces";

export default defineComponent({
  props: {
    gemColor: {
      type: Object as PropType<GemColor>,
      required: true,
    },
  },
  computed: {
    getGemImageUrl(): string {
      const gemColorName = getEnumKeyByEnumValue(GemColor, this.gemColor);
      return require(`@/assets/images/${gemColorName.toLowerCase()}_64.png`);
    },
  },
});
</script>
