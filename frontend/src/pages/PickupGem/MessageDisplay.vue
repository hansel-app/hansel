<template>
  <div class="bg-container" :style="addBackground">
    <BottomDrawer/>
  </div>
</template>
<script lang="ts">
import { defineComponent } from "vue";
import BottomDrawer from "@/components/MessageBottomSheet.vue";

export default defineComponent({
  components: {
    BottomDrawer
  },
  data() {
    return {
      show: true,
      addBackground: {
        // TODO: get and chuck image here
        backgroundImage: `url(${require('@/assets/images/purple-gem.png')})`
      },
    };
  },
  methods: {
    showPopup() {
      this.show = true;
    }
  }
});
</script>
<style scoped>
.bg-container {
  font-size: 100%;
  text-align: center;
  height:100vh;
  position: relative;
  background-size: cover;
  background-position: 50% 50%;
  -webkit-background-size: cover;
  -moz-background-size: cover;
  -o-background-size: cover;
  background-size: cover;
}
</style>