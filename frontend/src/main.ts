import Vant from "vant";
import "vant/lib/index.css";
import { createApp } from "vue";
import App from "./App.vue";
import "./registerServiceWorker";

createApp(App)
  .use(Vant)
  .mount("#app");
