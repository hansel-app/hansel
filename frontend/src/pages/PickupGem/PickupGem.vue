<template>
  <router-view :name="currentStageName" :gemPackage="gemPackage" @nextStage="nextStage"></router-view>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import dayjs from "dayjs"; 

// eslint-disable-next-line no-unused-vars
enum PickupGemStage {
  // eslint-disable-next-line no-unused-vars
  FoundGem,
  // eslint-disable-next-line no-unused-vars
  MessageDisplay,
}

export default defineComponent({
  setup() {
    const gemPackage = {
      message: "Me gusta los camarones",
      position: { lat: 1.2966, lng: 103.7764 },
      dropper: "Goose",
      receiver: "Wimp",
      dropTime: dayjs(),
      color: "purple",
    }

    return { gemPackage }
  },
  data() {
    return {
      currentStage: PickupGemStage.FoundGem as PickupGemStage,
    };
  },
  computed: {
    currentStageName(): String {
      return PickupGemStage[this.currentStage];
    },
  },
  methods: {
    nextStage() {
        this.currentStage = PickupGemStage.MessageDisplay;
    }
  }
});
</script>
