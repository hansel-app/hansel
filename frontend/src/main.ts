import Vant from "vant";
import { createApp } from "vue";
import { createWebHistory, createRouter } from "vue-router";

import {
  AttachMedia,
  AttachMessage,
  DropGem,
  SelectFriend,
} from "./pages/DropGem/index";
import GemMap from "./components/GemMap.vue";
import App from "./App.vue";
import "vant/lib/index.css";
import "./registerServiceWorker";

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: "/", name: "home", component: GemMap },
    {
      path: "/drop",
      name: "dropGem",
      component: DropGem,
      children: [
        {
          path: "",
          components: {
            friend: SelectFriend,
            media: AttachMedia,
            message: AttachMessage,
          },
        },
      ],
    },
  ],
});

createApp(App)
  .use(Vant)
  .use(router)
  .mount("#app");
