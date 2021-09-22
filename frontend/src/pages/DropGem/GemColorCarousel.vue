<template>
  <Swiper
    ref="swiperRef"
    :slides-per-view="4"
    :space-between="50"
    :centered-slides="true"
    :slide-to-clicked-slide="true"
    @real-index-change="onRealIndexChange"
  >
    <SwiperSlide v-for="color in colors" :key="color">
      <Button round class="gem-color-button" :color="color" />
    </SwiperSlide>
  </Swiper>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { Button } from "vant";
import { GemColor } from "@/interfaces";
import { Swiper, SwiperSlide } from "swiper/vue";
import { Swiper as SwiperType } from "swiper";

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
  methods: {
    onRealIndexChange(swiper: SwiperType) {
      const selectedColor: GemColor = this.colors[swiper.realIndex];
      this.$emit("selected-color-changed", selectedColor);
    },
  },
});
</script>

<style scoped lang="less">
.gem-color-button {
  width: 3rem;
  height: 3rem;
}
</style>
