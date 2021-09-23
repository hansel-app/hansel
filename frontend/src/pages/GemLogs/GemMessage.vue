<template>
  <van-row :justify="justify">
    <van-image :src="gem.attachment" />
  </van-row>
  <van-row class="gem-message" :justify="justify">
    <p>{{ gem.message }}</p>
  </van-row>
  {{ isSentBySelf ? "ME" : "YOU" }}
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
    isSentBySelf: {
      type: Boolean,
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

<style scoped lang="less"></style>
