<template>
  <NavBar v-bind:left-arrow="currentStage > 0" @click-left="prevStage" />
  <div @SetReceiverEvent="draftGem.receiverId = $event">
    <router-view class="select-friend" :name="currentStageName"></router-view>
  </div>
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
import { defineComponent, inject, ref } from "vue";
import { DROP_GEM } from "@/providers/GemProvider.vue";
import { NavBar } from "vant";
import { GemMessage } from "@/protobuf/gem_pb";

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
    dropMyGem() {
      console.log("called dropMyGem");
      const dropGem = inject(DROP_GEM, () => Promise.resolve(""));
      const response = ref<string>("");
      console.log(this.draftGem);
      dropGem(this.draftGem)
        .then((resp) => {
          response.value = resp;
        })
        .catch((err) => console.log(err, "Failed to drop gem"));
      console.log(response.value);
    },
  },
  components: {
    NavBar,
  },
});
</script>
