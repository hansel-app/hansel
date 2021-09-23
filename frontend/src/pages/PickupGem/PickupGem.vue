<template>
  <router-view
    :name="currentStageName"
    :gem="gem"
    @next-stage="nextStage"
  ></router-view>
</template>

<script lang="ts">
import { defineComponent, PropType } from "vue";
import { Gem } from "@/interfaces";

enum PickupGemStage {
  FoundGem,
  MessageDisplay,
}

export default defineComponent({
  props: {
    gem: {
      type: Object as PropType<Gem>,
      required: true,
    },
  },
  data() {
    return {
      currentStage: PickupGemStage.FoundGem as PickupGemStage,
    };
  },
  computed: {
    currentStageName(): string {
      return PickupGemStage[this.currentStage];
    },
  },
  methods: {
    nextStage() {
      this.currentStage = PickupGemStage.MessageDisplay;
    },
  },
});
</script>
