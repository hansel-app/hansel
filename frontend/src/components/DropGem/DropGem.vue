<template>
  <NavBar v-bind:left-arrow="currentStageIdx > 0" @click-left="prevStage" />
  <router-view class="select-friend" v-bind:name="currentStage"></router-view>
  <div>
    <button v-if="currentStageIdx < stages.length - 1" v-on:click="nextStage">
      Next
    </button>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { NavBar } from "vant";

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
  components: {
    NavBar,
  },
});
</script>
