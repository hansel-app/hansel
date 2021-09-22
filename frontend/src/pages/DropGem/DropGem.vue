<template>
  <div>
    <NavBar
      v-bind:left-arrow="currentStage > 0"
      @click-left="prevStage"
      class="nav-bar"
    />
    <router-view
      :name="currentStageName"
      @nextStage="nextStage"
      @SetReceiverEvent="setReceiverId"
      @SetMessageEvent="setMessage"
      @SetGemColorEvent="setGemColor"
    ></router-view>

    <div>
      <van-button
        round
        type="info"
        v-if="shouldShowNextButton"
        @click="$emit('nextStage')"
      >
        Next
      </van-button>
      <van-button v-if="currentStage === numStages - 1" v-on:click="dropMyGem">
        Drop Gem
      </van-button>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { NavBar } from "vant";
import { GemMessage, GemColor } from "@/protobuf/gem_pb";
import { useStore } from "vuex";

enum DropGemStage {
  Friend,
  Media,
  Message,
}

export default defineComponent({
  data() {
    return {
      currentStage: DropGemStage.Friend as DropGemStage,
      draftGem: new GemMessage(),
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
    // Tried to make this inline, couldn't figure out how to pass in value.
    setReceiverId(value: number): void {
      this.draftGem.setReceiverId(value);
    },
    setMessage(value: string): void {
      this.draftGem.setMessage(value);
    },
    setGemColor(value: GemColor): void {
      this.draftGem.setColor(value);
    },
    dropMyGem() {
      const dropGem = () => this.store.dispatch("dropGem", this.draftGem);
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
</style>
