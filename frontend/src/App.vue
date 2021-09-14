<template>
  <Suspense>
    <template #default>
      <router-view :key="$route.fullPath"></router-view>
    </template>
    <template #fallback>
      <!-- TODO: add more useful loading / error message -->
      <div>Loading...</div>
    </template>
  </Suspense>
</template>

<script>
import { defineComponent } from "vue";

export default defineComponent({
  name: "App",
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
