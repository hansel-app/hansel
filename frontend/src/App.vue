<template>
  <Provider>
    <router-view :key="$route.fullPath"></router-view>
  </Provider>
</template>

<script>
import { defineComponent } from "vue";
import Provider from "./providers/Provider.vue";

export default defineComponent({
  name: "App",
  components: { Provider },
});
</script>

<style lang="less">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

@font-face {
  font-family: "CircularStd";
  src: local("CircularStd-Black"),
    url(./assets/fonts/CircularStd-Black.otf) format("truetype");
  font-weight: bolder;
  font-style: normal;
}

@font-face {
  font-family: "CircularStd";
  src: local("CircularStd-Bold"),
    url(./assets/fonts/CircularStd-Bold.otf) format("truetype");
  font-weight: bold;
  font-style: normal;
}

@font-face {
  font-family: "CircularStd";
  src: local("CircularStd-Medium"),
    url(./assets/fonts/CircularStd-Medium.otf) format("truetype");
  font-weight: normal;
  font-style: normal;
}

@font-face {
  font-family: "CircularStd";
  src: local("CircularStd-MediumItalic"),
    url(./assets/fonts/CircularStd-MediumItalic.otf) format("truetype");
  font-weight: normal;
  font-style: italic;
}
</style>
