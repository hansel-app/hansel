<template>
  <Swiper
    ref="swiperRef"
    :slides-per-view="4"
    :space-between="50"
    :centered-slides="true"
    :slide-to-clicked-slide="true"
  >
    <SwiperSlide v-for="color in colors" :key="color">
      <Button :color="color" />
    </SwiperSlide>
  </Swiper>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { Button } from "vant";
import { GemColor } from "@/interfaces";
import { Swiper, SwiperSlide } from "swiper/vue";

import "swiper/swiper-bundle.css";

export default defineComponent({
  components: {
    Button,
    Swiper,
    SwiperSlide,
  },

  setup() {
    return {
      colors: Object.values(GemColor),
    };
  },
});
</script>
