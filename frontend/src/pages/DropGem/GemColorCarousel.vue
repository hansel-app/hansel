<template>
  <Swiper
    ref="swiperRef"
    :slides-per-view="4"
    :space-between="50"
    :centered-slides="true"
    :slide-to-clicked-slide="true"
  >
    <SwiperSlide v-for="color in colors" :key="color">
      <Button
        round
        class="gem-color-button"
        :color="color"
        @click="this.$emit('SetGemColorEvent', $event.target.value)"
      />
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

<style scoped lang="less">
.gem-color-button {
  width: 3rem;
  height: 3rem;
}
</style>
