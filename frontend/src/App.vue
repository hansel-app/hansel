<template>
  <router-view :key="$route.fullPath"></router-view>
</template>

<script>
import { defineComponent } from "vue";

export default defineComponent({
  name: "App",
});
</script>

<style>
#app {
  font-family: Circular-Std, Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
@font-face {
  font-family: "Circular-Std";
  src: local("CircularStd-Black"),
   url(./assets/fonts/CircularStd-Black.otf) format("truetype");
}
</style>
