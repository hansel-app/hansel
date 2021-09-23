<template>
  <van-row :justify="justify">
    <van-row :justify="justify">
      <van-image class="gem-image" :src="gem.attachment" radius="1em" />
    </van-row>
    <p class="gem-message" v-bind:style="gemMessageStyle">
      {{ gem.message }}
    </p>
  </van-row>
  <p class="open-message">{{ openMessage }}</p>
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
    // "start" | "end" sorry I made this weaker typing again haha
    justify: {
      type: String,
      required: true,
    },
    isSentBySelf: {
      type: Boolean,
      required: true,
    },
  },
  computed: {
    openMessage(): string {
      const creator = this.gem.createdBy;
      if (this.isSentBySelf) {
        return `${creator.displayName} opened your gem!`;
      } else return `You opened ${creator.displayName}'s gem!`;
    },
  },
});
</script>

<style scoped lang="less">
.gem-image {
  max-width: 200px;
  max-height: 200px;
  margin: auto;
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

.van-row {
  padding-top: 1vh;
}
</style>
