<template>
  <router-view
    :name="currentStageName"
    :gem="gem"
    @next-stage="nextStage"
  ></router-view>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { Gem } from "@/interfaces";
import { mapState } from "vuex";

enum PickupGemStage {
  FoundGem,
  MessageDisplay,
}

export default defineComponent({
  data() {
    return {
      currentStage: PickupGemStage.FoundGem as PickupGemStage,
    };
  },
  computed: {
    ...mapState({
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      gem: (state: any) => state.gems.lastPickedUpGem as Gem,
    }),
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
