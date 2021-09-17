<template>
  <router-view :name="currentStageName" @nextStage="nextStage" :gem="gemInfo"></router-view>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import dayjs from 'dayjs';

// eslint-disable-next-line no-unused-vars
enum PickupGemStage {
  // eslint-disable-next-line no-unused-vars
  FoundGem,
  // eslint-disable-next-line no-unused-vars
  MessageDisplay,
}

export default defineComponent({
  setup() {
    // TODO: same old same old
    const gemInfo = {
        dropperName: "Country boy",
        dropTime: dayjs(),
    }
    return gemInfo;
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
