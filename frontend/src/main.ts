import Vant from "vant";
import { createApp } from "vue";
import { createWebHistory, createRouter } from "vue-router";

import {
  AttachMedia,
  AttachMessage,
  DropGem,
  SelectFriend,
} from "./pages/DropGem/index";
import { DROP_GEM_ROUTE, HOME_ROUTE, PICKUP_GEM_ROUTE } from "@/constants";
import HomePage from "./pages/Home/HomePage.vue";
import FoundGem from "./pages/PickupGem/FoundGem.vue";
import Temp from "./pages/PickupGem/MessageDisplay.vue";
import App from "./App.vue";
import "vant/lib/index.css";
import "./registerServiceWorker";

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: HOME_ROUTE, name: "home", component: HomePage },
    {
      path: DROP_GEM_ROUTE,
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
    { path: PICKUP_GEM_ROUTE, name: "pickupGem", component: FoundGem },
    { path: '/temp', name: "Temp", component: Temp },
  ],
});

createApp(App)
  .use(Vant)
  .use(router)
  .mount("#app");
