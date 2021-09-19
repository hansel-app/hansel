<template>
  <router-view
    :name="currentStageName"
    :gem="gemPackage"
    @nextStage="nextStage"
  ></router-view>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { mockGems } from "@/interfaces/mockData";
import { Gem } from "@/interfaces";

enum PickupGemStage {
  FoundGem,
  MessageDisplay,
}

export default defineComponent({
  setup() {
    // TODO: replace later~
    const gem: Gem = mockGems[0];
    return { gem };
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
