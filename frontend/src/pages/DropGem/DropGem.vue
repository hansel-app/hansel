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
import { defineComponent, ref, Ref } from "vue";
import { NavBar } from "vant";
import { useStore } from "vuex";
import { GemColor } from "@/interfaces";

enum DropGemStage {
  Friend,
  Media,
  Message,
}

export default defineComponent({
  components: {
    NavBar,
  },
  setup() {
    const store = useStore();
    store.commit("clearDropGemFormState");

    const currentStage: Ref<DropGemStage> = ref(DropGemStage.Friend);

    return {
      currentStage,
      store,
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
      this.store.commit("updateDropGemFormState", { receiverId: value });
    },
    setMessage(value: string): void {
      this.store.commit("updateDropGemFormState", { message: value });
    },
    setGemColor(value: GemColor): void {
      this.store.commit("updateDropGemFormState", { color: value });
    },
    dropMyGem() {
      const dropGem = () => this.store.dispatch("dropGem");
      dropGem().catch((err) => console.log(err, "Failed to drop gem"));
    },
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
