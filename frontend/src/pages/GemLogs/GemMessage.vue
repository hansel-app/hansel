<template>
  <van-row :justify="justify">
    <van-image
      v-if="isSentBySelf || hasPickedUpGem"
      class="gem-image"
      :src="`data:image/png;base64,${gem.attachment}`"
    />
  </van-row>
  <van-row :justify="justify">
    <p
      v-if="isSentBySelf || hasPickedUpGem"
      class="gem-message"
      v-bind:style="gemMessageStyle"
    >
      {{ gem.message }}
    </p>
    <p v-else class="gem-message" v-bind:style="gemMessageStyle">
      They dropped a gem!
    </p>
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
    isSentBySelf: {
      type: Boolean,
      required: true,
    },
  },
  computed: {
    hasPickedUpGem() {
      return this.gem.receivedAt !== null;
    },
    justify(): "start" | "end" {
      if (this.isSentBySelf) {
        return "end";
      } else {
        return "start";
      }
    },
    gemMessageStyle() {
      if (this.isSentBySelf) {
        return {
          backgroundColor: "#585969",
          color: "white",
          textAlign: "end",
        };
      } else {
        return {
          backgroundColor: "#f7f8fa",
          textAlign: "start",
        };
      }
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

.gem-image {
  padding: 1em 1em;
}
</style>
