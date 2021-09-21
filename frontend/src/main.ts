import router from "@/router";
import store from "@/store";
import Vant from "vant";
import { createApp } from "vue";
import VueGtag from "vue-gtag";
import VueClickAway from "vue3-click-away";
import App from "./App.vue";
import "./registerServiceWorker";

const GOOGLE_ANALYTICS_ID = window.env.VUE_APP_GOOGLE_ANALYTICS_ID;

// Load store before router.
createApp(App)
  .use(Vant)
  .use(store)
  .use(router)
  .use(
    VueGtag,
    {
      config: { id: GOOGLE_ANALYTICS_ID },
    },
    router
  )
  .use(VueClickAway)
  .mount("#app");
