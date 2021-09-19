<template>
  <router-view
    :name="currentStageName"
    :gemPackage="gemPackage"
    @nextStage="nextStage"
  ></router-view>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import dayjs from "dayjs";

enum PickupGemStage {
  FoundGem,
  MessageDisplay,
}

export default defineComponent({
  setup() {
    // TODO: replace later~
    const gemPackage = {
      message: "Me gusta los camarones",
      position: { lat: 1.2966, lng: 103.7764 },
      dropper: "Goose",
      receiver: "Wimp",
      dropTime: dayjs(),
      color: "purple",
    };

    return { gemPackage };
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
