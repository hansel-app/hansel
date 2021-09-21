<template>
  <NavBar
    v-bind:left-arrow="currentStage > 0"
    @click-left="prevStage"
    class="nav-bar"
  />
  <router-view
    class="select-friend"
    :name="currentStageName"
    @SetReceiverEvent="setReceiverId"
    @SetMessageEvent="setMessage"
    @SetGemColorEvent="setGemColor"
  ></router-view>
  <div>
    <button v-if="currentStage < numStages - 1" v-on:click="nextStage">
      Next
    </button>
    <button v-if="currentStage === numStages - 1" v-on:click="dropMyGem">
      Drop Gem
    </button>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";
import { NavBar } from "vant";
import { GemMessage, GemColor } from "@/protobuf/gem_pb";
import { useStore } from "vuex";

enum DropGemStage {
  friend,
  media,
  message,
}

export default defineComponent({
  data() {
    return {
      currentStage: DropGemStage.friend as DropGemStage,
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
