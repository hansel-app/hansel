<template>
  <van-row :justify="start">
    <p class="gem-message" v-bind:style="gemMessageStyle">
      They dropped a gem!
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
    gemMessageStyle: {
      type: String,
      required: true,
    },
    justify: {
      type: String,
      required: true,
    },
  },
});
</script>

<style scoped>
.gem-message {
  width: 70%;
  padding: 1em 0.5em;
  word-wrap: break-word;
  border-radius: @border-radius-lg;
}
</style>
