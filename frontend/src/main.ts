import Vant from "vant";
import { createApp } from "vue";
import App from "./App.vue";
import "./registerServiceWorker";
import index from "@/router";
import store from "@/store";

createApp(App)
  .use(Vant)
  .use(index)
  .use(store)
  .mount("#app");
