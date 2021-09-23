import {
  ADD_FRIENDS_ROUTE,
  DROP_GEM_ROUTE,
  EDIT_PROFILE_ROUTE,
  FRIEND_REQUESTS_ROUTE,
  FRIEND_LIST_ROUTE,
  GEM_LOGS_DETAILS_ROUTE,
  GEM_LOGS_ROUTE,
  HOME_ROUTE,
  LOGIN_ROUTE,
  PICKUP_GEM_ROUTE,
  PROFILE_ROUTE,
  REGISTER_ROUTE,
} from "@/constants";
import {
  AttachMedia,
  AttachMessage,
  DropGem,
  SelectFriend,
} from "@/pages/DropGem/index";
import {
  AddFriendPage,
  FriendRequestsPage,
  FriendListPage,
} from "@/pages/Friend";
import { GemLogsDetailsPage, GemLogsPage } from "@/pages/GemLogs/index";
import HomePage from "@/pages/Home/HomePage.vue";
import { LoginPage, RegisterPage } from "@/pages/LoginRegister";
import { FoundGem, MessageDisplay, PickupGem } from "@/pages/PickupGem/index";
import { EditProfilePage, ProfilePage } from "@/pages/Profile/index";
import store from "@/store";
import "animate.css";
import {
  createRouter,
  createWebHistory,
  NavigationGuardNext,
  RouteLocationNormalized,
} from "vue-router";

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
            Friend: SelectFriend,
            Media: AttachMedia,
            Message: AttachMessage,
          },
        },
      ],
      meta: {
        enterClass: "animate__animated animate__slideInRight",
        leaveClass: "animate__animated animate__slideOutLeft",
      },
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
      meta: {
        enterClass: "animate__animated animate__slideInRight",
        leaveClass: "animate__animated animate__slideOutLeft",
      },
    },
    { path: PROFILE_ROUTE, component: ProfilePage },
    { path: EDIT_PROFILE_ROUTE, component: EditProfilePage },
    { path: ADD_FRIENDS_ROUTE, component: AddFriendPage },
    { path: FRIEND_REQUESTS_ROUTE, component: FriendRequestsPage },
    { path: FRIEND_LIST_ROUTE, component: FriendListPage },
    {
      path: LOGIN_ROUTE,
      component: LoginPage,
      meta: {
        isPublic: true,
      },
    },
    {
      path: REGISTER_ROUTE,
      component: RegisterPage,
      meta: {
        isPublic: true,
      },
    },
    {
      path: GEM_LOGS_ROUTE,
      component: GemLogsPage,
      meta: {
        enterClass: "animate__animated animate__fadeInRight",
        leaveClass: "animate__animated animate__fadeOutLeft",
      },
    },
    {
      path: GEM_LOGS_DETAILS_ROUTE,
      component: GemLogsDetailsPage,
    },
  ],
});

router.beforeEach(
  (
    to: RouteLocationNormalized,
    _from: RouteLocationNormalized,
    next: NavigationGuardNext
  ) => {
    if (to.matched.some((record) => record.meta.isPublic)) {
      next();
      return;
    }
    if (store.getters.isLoggedIn) {
      next();
      return;
    }
    next(LOGIN_ROUTE);
  }
);

export default router;
