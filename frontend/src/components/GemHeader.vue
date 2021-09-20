<template>
  <div class="stack-container">
    <p class="big-header overlay">{{ title }}</p>
    <div class="square" />
    <div class="semi-circle" />
    <Image :src="getGemImageUrl" class="overlay" />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { Image } from "vant";

export default defineComponent({
  props: {
    title: String,
  },
  methods: {
    gemPath() {
      return require("@/assets/images/blue_64.png");
    },
  },
  data() {
    return {
      // TODO: replace with our own gem asset
      placeholderGemUrl:
        "https://images.squarespace-cdn.com/content/v1/5a773a9f8fd4d27532408347/1530563687072-GRIQPL3SZQDPZ7ZY7JEC/butt+lightened.jpg?format=1000w",
    };
  },
  components: { Image },
});
</script>
<style scoped>
.stack-container {
  position: relative;
  height: 325px;
}

.square {
  height: 250px;
  width: 100%;
  background-color: #e4bfbf;
}

.overlay {
  position: absolute;
  left: 0;
  right: 0;
  margin-left: auto;
  margin-right: auto;
}

.semi-circle {
  position: absolute;
  transform: translate(0%, -50%);
  height: 150px;
  width: 100%;
  border-radius: 0 0 50% 50%;
  background-color: #e4bfbf;
}

.gem-display {
  width: 100px;
  height: 100px;
}
</style>
