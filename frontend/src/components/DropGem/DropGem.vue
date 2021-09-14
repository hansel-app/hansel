<template>
  <router-view class="select-friend" v-bind:name="currentStage"></router-view>
  <div>
    <button v-if="currentStageIdx > 0" v-on:click="prevStage">Previous</button>
    <button v-if="currentStageIdx < stages.length - 1" v-on:click="nextStage">
      Next
    </button>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  data() {
    return {
      stages: ["friend", "media", "message"],
      currentStageIdx: 0,
    };
  },
  computed: {
    currentStage() {
      return this.stages[this.currentStageIdx];
    },
  },
  methods: {
    nextStage() {
      if (this.currentStageIdx < this.stages.length - 1) {
        this.currentStageIdx += 1;
      }
    },
    prevStage() {
      if (this.currentStageIdx > 0) {
        this.currentStageIdx -= 1;
      }
    },
  },
});
</script>
