<template>
  <van-row :justify="justify">
    <van-row :justify="justify">
      <van-image class="gem-image" :src="gem.attachment" />
    </van-row>
    <van-row :justify="justify">
      <p class="gem-message" v-bind:style="gemMessageStyle">
        {{ gem.message }}
      </p>
    </van-row>
  </van-row>
  <p class="open-message">Groundskeeper opened your gem!</p>
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
    // "start" | "end"
    justify: {
      type: String,
      required: true,
    },
    isSentBySelf: {
      type: Boolean,
      required: true,
    },
  },
  computed: {},
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

.open-message {
  text-align: center;
  margin: 0 auto;
  padding-top: 2vh;
  padding-bottom: 2vh;
}
</style>
