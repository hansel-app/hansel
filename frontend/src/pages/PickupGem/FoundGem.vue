<template>
  <div class="bg-container">
    <video autoplay muted :poster="getPoster" id="gem-video">
      <source :src="getVideo" type="video/mp4" />
    </video>
    <div id="found-gem" class="big-header">
      You found a gem!
    </div>
    <div id="view-button">
      <van-button @click="$emit('next-stage')" size="large">
        View contents
      </van-button>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, PropType } from "vue";
import { Button } from "vant";
import { Gem, GemColor } from "@/interfaces";
import { getEnumKeyByEnumValue } from "@/utils/enum";

export default defineComponent({
  props: {
    gem: {
      type: Object as PropType<Gem>,
      required: true,
    },
  },
  components: {
    "van-button": Button,
  },
  computed: {
    getPoster() {
      const gemColorName = getEnumKeyByEnumValue(GemColor, this.gem.color);
      return require(`@/assets/images/gem-found-${gemColorName.toLowerCase()}.png`);
    },
    getVideo() {
      const gemColorName = getEnumKeyByEnumValue(GemColor, this.gem.color);
      return require(`@/assets/videos/gem-found-${gemColorName.toLowerCase()}.mp4`);
    },
  },
});
</script>

<style scoped>
.bg-container {
  font-size: 100%;
  text-align: center;
  height: 100vh;
  position: relative;
}

#found-gem {
  box-sizing: border-box;
  padding: 2em;
  display: inline-block;
  left: 0;
  position: absolute;
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

.big-header {
  color: black;
  margin-top: -1em;
}
</style>
