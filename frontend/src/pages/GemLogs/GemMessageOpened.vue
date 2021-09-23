<template>
  <van-row justify="end">
    <van-image class="gem-image" :src="gem.attachment" />
  </van-row>
  <van-row justify="end">
    <p class="gem-message" v-bind:style="gemMessageStyle">
      {{ gem.message }}
    </p>
  </van-row>
</template>

<script lang="ts">
import { defineComponent, PropType } from "vue";
import { Gem } from "@/interfaces";

export default defineComponent({
  props: {
    gem: {
      type: Object as PropType<Gem>,
      required: true,
    },
    // pseudo-CSS
    gemMessageStyle: {
      type: String,
      required: true,
    },
    isSentBySelf: {
      type: Boolean,
      required: true,
    },
  },
});
</script>

<style scoped lang="less">
.gem-image {
  padding: 1em 1em;
}

.gem-message {
  width: 70%;
  padding: 1em 0.5em;
  word-wrap: break-word;
  border-radius: @border-radius-lg;
}
</style>
