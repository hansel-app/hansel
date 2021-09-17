<template>
  <div class="bg-container" :style="addBackground">
    <van-popup v-model="show" position="bottom" :style="{ width: '100%', height: '100%' }">
      <div>
        <h1>内容s</h1>
      </div>
    </van-popup>
  </div>
</template>
<script lang="ts">
import { defineComponent } from "vue";
import { Popup } from "vant";

export default defineComponent({
  components: {
    'van-popup': Popup,
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