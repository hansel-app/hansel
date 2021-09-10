<template>
  <hello-world></hello-world>
  <gem-map></gem-map>
</template>

<script>
import HelloWorld from "./components/HelloWorld.vue";
import GemMap from "./components/GemMap.vue";

export default {
  name: "App",
  components: { HelloWorld, GemMap },
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
