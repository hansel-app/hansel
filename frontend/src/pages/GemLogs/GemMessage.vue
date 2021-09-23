<template>
  <van-row class="base">
    <GemMessageOpen
      :gem="gem"
      :gemMessageStyle="gemMessageStyle"
      :justify="justify"
      v-if="hasPickedUpGem"
    />
    <van-row v-else>
      <GemMessageClosed
        :gemMessageStyle="gemMessageStyle"
        :gem="gem"
        :justify="justify"
      />
    </van-row>
  </van-row>
</template>

<script lang="ts">
import { PropType, defineComponent } from "vue";
import { Gem, HexCode } from "@/interfaces";
import GemMessageOpen from "./GemMessageOpened.vue";
import GemMessageClosed from "./GemMessageClosed.vue";

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
  components: {
    GemMessageOpen,
    GemMessageClosed,
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
</style>
