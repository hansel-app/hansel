<template>
  <van-row class="gem-message" :justify="justify">
    <van-image />
    <p>{{ message }}</p>
  </van-row>
</template>

<script lang="ts">
import { PropType, defineComponent } from "vue";
import { HexCode } from "@/interfaces";

export default defineComponent({
  props: {
    color: String as PropType<HexCode>,
    message: {
      type: String,
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
