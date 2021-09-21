<template>
  <router-view :key="$route.fullPath"></router-view>
</template>

<script>
import { defineComponent } from "vue";

export default defineComponent({
  name: "App",
});
</script>

<style lang="less">
#app {
  font-family: CircularStd, Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

@font-face {
  font-family: "CircularStd";
  src: local("CircularStd-Black"),
    url(./assets/fonts/CircularStd-Black.otf) format("truetype");
  font-weight: bold;
  font-style: normal;
}

@font-face {
  font-family: "CircularStd-Bold";
  src: local("CircularStd-Bold"),
    url(./assets/fonts/CircularStd-Bold.otf) format("truetype");
  font-weight: normal;
  font-style: normal;
}

@font-face {
  font-family: "CircularStd";
  src: local("CircularStd-Medium"),
    url(./assets/fonts/CircularStd-Medium.otf) format("truetype");
  font-weight: lighter;
  font-style: normal;
}

@font-face {
  font-family: "CircularStd";
  src: local("CircularStd-MediumItalic"),
    url(./assets/fonts/CircularStd-MediumItalic.otf) format("truetype");
  font-weight: lighter;
  font-style: italic;
}
</style>
