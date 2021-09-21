<template>
  <NavBar
    v-bind:left-arrow="currentStage > 0"
    @click-left="prevStage"
    class="nav-bar"
  />
  <router-view :name="currentStageName" @nextStage="nextStage"></router-view>

  <div>
    <van-button
      round
      type="info"
      v-if="shouldShowNextButton"
      @click="$emit('nextStage')"
    >
      Next
    </van-button>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { NavBar } from "vant";

enum DropGemStage {
  Friend,
  Media,
  Message,
}

export default defineComponent({
  data() {
    return {
      currentStage: DropGemStage.Friend as DropGemStage,
    };
  },
  computed: {
    numStages(): number {
      return Object.keys(DropGemStage).length / 2;
    },
    currentStageName(): string {
      return DropGemStage[this.currentStage];
    },
    shouldShowNextButton(): boolean {
      return [DropGemStage.Media].includes(this.currentStage);
    },
  },
  methods: {
    nextStage() {
      if (this.currentStage < this.numStages - 1) {
        this.currentStage += 1;
      }
    },
    prevStage() {
      if (this.currentStage > 0) {
        this.currentStage -= 1;
      }
    },
  },
  components: {
    NavBar,
  },
});
</script>

<style scoped>
.nav-bar {
  position: absolute;
}
</style>
