import {createWebHistory, createRouter, RouteLocationNormalized, NavigationGuardNext} from "vue-router";
import {
  AttachMedia,
  AttachMessage,
  DropGem,
  SelectFriend,
} from "@/pages/DropGem/index";
import { PickupGem, FoundGem, MessageDisplay } from "@/pages/PickupGem/index";
import {
  ADD_FRIENDS_ROUTE,
  DROP_GEM_ROUTE,
  FRIEND_REQUESTS_ROUTE,
  GEM_LOGS_ROUTE,
  HOME_ROUTE,
  LOGIN_ROUTE,
  PROFILE_ROUTE,
  PICKUP_GEM_ROUTE,
  REGISTER_ROUTE,
} from "@/constants";
import HomePage from "@/pages/Home/HomePage.vue";
import ProfilePage from "@/pages/Profile/ProfilePage.vue";
import { AddFriendPage, FriendRequestsPage } from "@/pages/Friend";
import { LoginPage, RegisterPage } from "@/pages/LoginRegister";
import GemLogs from "@/pages/GemLogs/LogsList.vue"


// import store from "@/store";

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
    {
      path: LOGIN_ROUTE,
      component: LoginPage,
      meta: {
        isPublic: true
      }
    },
    {
      path: REGISTER_ROUTE,
      component: RegisterPage,
      meta: {
        isPublic: true
      }
    },
    {
      path: GEM_LOGS_ROUTE,
      component: GemLogs,
      meta: {
        isPublic: true
      }
    }
  ],
});

router.beforeEach((to: RouteLocationNormalized, _from: RouteLocationNormalized, next: NavigationGuardNext) => {
  if (to.matched.some(record => record.meta.isPublic)) {
    next();
    return;
  }
  // if (store.getters.isLoggedIn) {
    next();
    return;
  // }
  // next(LOGIN_ROUTE);
});

export default router;
