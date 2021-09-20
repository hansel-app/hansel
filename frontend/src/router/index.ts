import { createWebHistory, createRouter } from "vue-router";
import {
  AttachMedia,
  AttachMessage,
  DropGem,
  SelectFriend,
} from "@/pages/DropGem/index";
import { PickupGem, FoundGem, MessageDisplay } from "@/pages/PickupGem/index";
import {
  DROP_GEM_ROUTE,
  ADD_FRIENDS_ROUTE,
  FRIEND_REQUESTS_ROUTE,
  HOME_ROUTE,
  PROFILE_ROUTE,
  LOGIN_ROUTE,
  REGISTER_ROUTE,
  PICKUP_GEM_ROUTE,
} from "@/constants";
import HomePage from "@/pages/Home/HomePage.vue";
import ProfilePage from "@/pages/Profile/ProfilePage.vue";
import { AddFriendPage, FriendRequestsPage } from "@/pages/Friend";
import { LoginPage, RegisterPage } from "@/pages/LoginRegister";

const index = createRouter({
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
    {
      path: PICKUP_GEM_ROUTE,
      name: "pickupGem",
      component: PickupGem,
      children: [
        {
          path: "",
          components: {
            FoundGem: FoundGem,
            MessageDisplay: MessageDisplay,
          },
        },
      ],
    },
    { path: PROFILE_ROUTE, component: ProfilePage },
    { path: ADD_FRIENDS_ROUTE, component: AddFriendPage },
    { path: FRIEND_REQUESTS_ROUTE, component: FriendRequestsPage },
    { path: LOGIN_ROUTE, component: LoginPage },
    { path: REGISTER_ROUTE, component: RegisterPage },
  ],
});

export default index;