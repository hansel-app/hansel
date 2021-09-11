<template>
  <Suspense>
    <template #default>
      <gem-map></gem-map>
    </template>
    <template #fallback>
      <div>Loading...</div>
    </template>
  </Suspense>
</template>

<script>
import { defineComponent } from "vue";
import GemMap from "./components/GemMap.vue";

export default defineComponent({
  name: "App",
  components: { GemMap },
});
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
