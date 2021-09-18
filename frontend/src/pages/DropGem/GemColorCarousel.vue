<template>
  <!-- TODO: make this into actual carousel -->
  <Carousel :settings="carouselSettings">
    <Slide v-for="gem in GemColor" :key="gem.color">
      <div class="carousel__item">
        <van-button round :color="gem" class="gem-color-button" />
      </div>
    </Slide>
  </Carousel>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { Carousel, Slide } from "vue3-carousel";
import { GemColor } from "@/interfaces";

import "vue3-carousel/dist/carousel.css";

export default defineComponent({
  data() {
    return {
      GemColor,
      carouselSettings: {
        itemsToShow: Object.keys(GemColor).length,
      },
    };
  },
  components: {
    Carousel,
    Slide,
  },
  inheritAttrs: false,
});
</script>

<style scoped>
.carousel__item {
  min-height: 48px;
  min-width: 48px;
  border-radius: 8px;
  display: flex;
  justify-content: center;
  align-items: center;
}

.carousel__slide > .carousel__item {
  transform: scale(1);
  transition: 0.5s;
}
.carousel__slide--visible > .carousel__item {
  transform: rotateY(0);
}
.carousel__slide--next > .carousel__item {
  transform: scale(0.8) translate(-10px);
}
.carousel__slide--prev > .carousel__item {
  transform: scale(0.8) translate(10px);
}
.carousel__slide--active > .carousel__item {
  transform: scale(1.5);
}

.gem-color-button {
  width: 100%;
}
</style>
