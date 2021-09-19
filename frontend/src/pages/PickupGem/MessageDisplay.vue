<template>
  <div class="bg-container" :style="addBackground">
    <BottomSheet :gem="gem" />
  </div>
</template>
<script lang="ts">
import { defineComponent, PropType } from "vue";
import { Gem } from "@/interfaces";
import BottomSheet from "@/components/BottomSheet.vue";

export default defineComponent({
  props: {
    gem: {
      type: Object as PropType<Gem>,
      required: true,
    },
  },
  components: {
    BottomSheet,
  },
  data() {
    return {
      show: true,
      addBackground: {
        // TODO: get and chuck image here
        backgroundImage: `url(${require("@/assets/images/gem-found-yellow.png")})`,
      },
    };
  },
  methods: {
    showPopup() {
      this.show = true;
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
  background-size: cover;
  background-position: 50% 50%;
  -webkit-background-size: cover;
  -moz-background-size: cover;
  -o-background-size: cover;
  background-size: cover;
}
</style>
