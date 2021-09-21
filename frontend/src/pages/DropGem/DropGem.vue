<template>
  <NavBar
    v-bind:left-arrow="currentStage > 0"
    @click-left="prevStage"
    class="nav-bar"
  />
  <router-view class="select-friend" :name="currentStageName"></router-view>
  <div>
    <van-button
      round
      type="info"
      v-if="shouldShowNextButton"
      @click="nextStage"
    >
      Next
    </van-button>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { NavBar } from "vant";

enum DropGemStage {
  friend,
  media,
  message,
}

export default defineComponent({
  data() {
    return {
      currentStage: DropGemStage.friend as DropGemStage,
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
      return [DropGemStage.media].includes(this.currentStage);
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
