<template>
  <div class="bg-container">
    <video autoplay muted :poster="getPoster" id="gem-video">
      <source :src="getVideo" type="video/mp4">
    </video>
    <div id="found-gem">
      <h1>You found a gem! {{ gemPackage.color }}</h1>
    </div>
    <div id="view-button">
      <van-button @click="$emit('nextStage')">View contents</van-button>
    </div>
  </div>
</template>
<script lang="ts">
import { defineComponent, PropType } from "vue";
import { Button } from "vant";
import { GemInfo, color } from "@/interfaces";

export default defineComponent({
  props: {
    gemPackage: {
      type: Object as PropType<GemInfo>,
      required: true,
    }
  },
  components: { 
    "van-button": Button,
  },
  computed: {
    getPoster() {
      // is this bad code, should i use switch instead :-)
      if (Object.values(color).includes(this.gemPackage.color as color)) {
        return require('@/assets/images/gem-' + this.gemPackage.color +'-found.png');
      }
      return require('@/assets/images/gem-purple-found.png');
    },
    getVideo() {
      if (Object.values(color).includes(this.gemPackage.color as color)) {
        return require('@/assets/videos/gem-' + this.gemPackage.color +'-found.mp4');
      }
      return require('@/assets/videos/gem-purple-found.mp4');
    },
  }
});
</script>
<style scoped>
.bg-container {
  font-size: 100%;
  text-align: center;
  height:100vh;
  position: relative;
}

#found-gem {
  box-sizing: border-box;
  padding: 2em;
  display: inline-block;
  left: 0;
  position:absolute;
  height: 80%;
  width: 100%;
}

#gem-video {
  object-fit: cover;
  width: 100vw;
  height: 100vh;
  position: fixed;
  top: 0;
  left: 0;
}

.van-button {
  padding: 1em 2em;
  background-color: black;
  border: none;
  border-radius: 2em;
  color: white;
}

#view-button {
  position: absolute; 
  bottom: 20%; 
  left: 50%; 
  -webkit-transform: translateX(-50%);
  transform: translateX(-50%);
}

h1 {
  font-family: CircularStd-Black;
  font-size: 3em;
  color: rgb(87, 40, 80);
}
</style>
