<template>
  <div v-cloak>
    <router-view v-slot="{ Component, route }" :key="$route.fullPath">
      <transition
        :enter-active-class="route.meta.enterClass"
        :leave-active-class="route.meta.leaveClass"
      >
        <div>
          <component :is="Component"></component>
        </div>
      </transition>
    </router-view>
  </div>
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

.fade-enter-from,
.fade-leave-to,
.fade-enter {
  opacity: 0;
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 2s ease-out;
}
</style>
