import Vant from "vant";
import { createApp } from "vue";
import { createWebHistory, createRouter } from "vue-router";

import DropGem from "./components/DropGem.vue";
import GemMap from "./components/GemMap.vue";
import App from "./App.vue";
import "vant/lib/index.css";
import "./registerServiceWorker";

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: "/", name: "home", component: GemMap },
    { path: "/drop", name: "dropGem", component: DropGem },
  ],
});

createApp(App)
  .use(Vant)
  .use(router)
  .mount("#app");
