<template>
  <Swiper
    ref="swiperRef"
    class="gem-color-carousel"
    :slides-per-view="4"
    :space-between="25"
    :centered-slides="true"
    :slide-to-clicked-slide="true"
    :initial-slide="initialIndex"
    @real-index-change="onRealIndexChange"
  >
    <SwiperSlide v-for="color in colors" :key="color">
      <Button round class="gem-color-button" :color="color" />
    </SwiperSlide>
  </Swiper>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { useStore } from "vuex";
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
    const store = useStore();
    const initialIndex = store.state.gems.dropGemFormState.colorIndex;

    return {
      store,
      colors: Object.values(GemColor),
      initialIndex,
    };
  },
  methods: {
    onRealIndexChange(swiper: SwiperType) {
      const selectedColor: GemColor = this.colors[swiper.realIndex];
      this.store.commit("updateDropGemFormState", {
        colorIndex: swiper.realIndex,
      });
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

.gem-color-carousel {
  margin-top: -2vh;
  margin-bottom: -1vh;
}
</style>
