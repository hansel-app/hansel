<template>
  <NavBar
    v-bind:left-arrow="currentStage > 0"
    @click-left="prevStage"
    class="nav-bar"
  />
  <router-view
    :name="currentStageName"
    @next-stage="nextStage"
    @set-receiver-event="setReceiverId"
    @set-message-event="setMessage"
    @set-gem-color-event="setGemColor"
  ></router-view>

  <van-button
    class="form-navigation-button"
    round
    type="primary"
    v-if="shouldShowNextButton"
    @click="nextStage"
  >
    Next
  </van-button>
  <van-button
    class="form-navigation-button"
    round
    type="primary"
    v-if="currentStage === numStages - 1"
    v-on:click="dropMyGem"
  >
    Drop Gem
  </van-button>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { NavBar } from "vant";
import { useStore } from "vuex";
import { DropGemOject, GemColor } from "@/interfaces";

enum DropGemStage {
  Friend,
  Media,
  Message,
}

export default defineComponent({
  data() {
    return {
      currentStage: DropGemStage.Friend as DropGemStage,
      formState: {} as DropGemOject,
      store: useStore(),
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
    setReceiverId(value: number): void {
      this.formState.receiverId = value;
    },
    setMessage(value: string): void {
      this.formState.message = value;
    },
    setGemColor(value: GemColor): void {
      this.formState.color = value;
    },
    dropMyGem() {
      const dropGem = () => this.store.dispatch("dropGem", this.formState);
      dropGem().catch((err) => console.log(err, "Failed to drop gem"));
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
.form-navigation-button {
  width: 80vw;
}
</style>
