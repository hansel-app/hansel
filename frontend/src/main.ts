import Vant from "vant";
import { createApp } from "vue";
import { createWebHistory, createRouter } from "vue-router";

import {
  AttachMedia,
  AttachMessage,
  DropGem,
  SelectFriend,
} from "./pages/DropGem/index";
import {
  DROP_GEM_ROUTE,
  ADD_FRIENDS_ROUTE,
  FRIEND_REQUESTS_ROUTE,
  HOME_ROUTE,
  PROFILE_ROUTE,
} from "@/constants";
import HomePage from "./pages/Home/HomePage.vue";
import ProfilePage from "./pages/Profile/ProfilePage.vue";
import { AddFriendPage, FriendRequestsPage } from "./pages/Friend";
import App from "./App.vue";
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
    { path: PROFILE_ROUTE, component: ProfilePage },
    { path: ADD_FRIENDS_ROUTE, component: AddFriendPage },
    { path: FRIEND_REQUESTS_ROUTE, component: FriendRequestsPage },
  ],
});

createApp(App)
  .use(Vant)
  .use(router)
  .mount("#app");
