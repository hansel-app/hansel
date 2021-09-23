<template>
  <van-row class="gem-message" :justify="justify">
    <van-image :src="gem.attachment" />
    <p>{{ gem.message }}</p>
  </van-row>
</template>

<script lang="ts">
import { PropType, defineComponent } from "vue";
import { Gem, HexCode } from "@/interfaces";

export default defineComponent({
  props: {
    color: String as PropType<HexCode>,
    gem: {
      type: Object as PropType<Gem>,
      required: true,
    },
    justify: {
      type: String,
      default: () => "start",
      validator(val: string) {
        return ["start", "end"].includes(val);
      },
    },
  },
});
</script>

<style scoped lang="less">
.gem-message {
  width: 70%;
}
</style>
