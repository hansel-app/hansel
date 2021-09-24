<template>
  <van-row class="base">
    <GemMessageOpen
      v-if="hasPickedUpGem"
      :gem="gem"
      :gemMessageStyle="gemMessageStyle"
      :justify="justify"
      :isSentBySelf="isSentBySelf"
    />
    <GemMessageClosed
      v-else
      :gemMessageStyle="gemMessageStyle"
      :gem="gem"
      :justify="justify"
      :isSentBySelf="isSentBySelf"
    />
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
