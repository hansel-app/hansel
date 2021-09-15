import Vant from "vant";
import { createApp } from "vue";
import { createWebHistory, createRouter } from "vue-router";

import DropGem from "./components/DropGem.vue";
import { DROP_GEM_ROUTE, HOME_ROUTE } from "@/constants";
import HomePage from "./pages/Home/HomePage.vue";
import App from "./App.vue";
import "vant/lib/index.css";
import "./registerServiceWorker";

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: HOME_ROUTE, name: "home", component: HomePage },
    { path: DROP_GEM_ROUTE, name: "dropGem", component: DropGem },
  ],
});

createApp(App)
  .use(Vant)
  .use(router)
  .mount("#app");
