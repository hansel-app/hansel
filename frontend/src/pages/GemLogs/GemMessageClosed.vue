<template>
  <van-row :justify="justify" class="fill-width">
    <p v-if="isSentBySelf" class="gem-message" v-bind:style="gemMessageStyle">
      New gem!
      <van-icon
        class="circle-button-icon-sm"
        :name="require('@/assets/images/purple_64.png')"
        @click="goToDropGem"
      />
    </p>
    <p v-else class="gem-message" v-bind:style="gemMessageStyle">
      <van-icon
        class="circle-button-icon-sm"
        :name="require('@/assets/images/purple_64.png')"
        @click="goToDropGem"
      />
      New gem!
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
    isSentBySelf: {
      type: Boolean,
      required: true,
    },
  },
});
</script>

<style scoped lang="less">
.gem-message {
  width: 70%;
  padding: 1em 0.5em;
  word-wrap: break-word;
  border-radius: @border-radius-lg;
}

.fill-width {
  width: 100%;
}
</style>
