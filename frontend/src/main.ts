import Vant from "vant";
import { createApp } from "vue";
import App from "./App.vue";
import "./registerServiceWorker";
import store from "@/store";
import router from "@/router";
import VueClickAway from "vue3-click-away";

// Load store before router.
createApp(App)
  .use(Vant)
  .use(store)
  .use(router)
  .use(VueClickAway)
  .mount("#app");
