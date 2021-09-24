import services from "@/api/services";
import { LOGIN_ROUTE } from "@/constants";
import {
  LoginRequest,
  LoginResponse,
  RegisterRequest,
} from "@/protobuf/auth_pb";
import router from "@/router";
import { RootState } from "@/store";
import Cookies from "js-cookie";
import { Module } from "vuex";

const ACCESS_TOKEN_COOKIE = "access-token";

export interface AuthState {
  accessToken: Nullable<string>;
}

const authModule: Module<AuthState, RootState> = {
  state: {
    accessToken: Cookies.get(ACCESS_TOKEN_COOKIE) ?? null,
  },
  getters: {
    isLoggedIn(state): boolean {
      return state.accessToken != null;
    },
  },
  mutations: {
    setAccessToken(state, accessToken: string) {
      state.accessToken = accessToken;
      Cookies.set(ACCESS_TOKEN_COOKIE, accessToken, {
        expires: 7,
        secure: true,
        sameSite: "strict",
      });
    },
    clearAccessToken(state) {
      state.accessToken = null;
      Cookies.remove(ACCESS_TOKEN_COOKIE);
    },
  },
  actions: {
    login(
      { commit, dispatch },
      payload: { username: string; password: string }
    ) {
      const request = new LoginRequest();
      request.setUsername(payload.username);
      request.setPassword(payload.password);

      return new Promise((resolve, reject) => {
        services.authClient
          .login(request)
          .then((response: LoginResponse) => {
            const token = response.getAccessToken();
            commit("setAccessToken", token);
            dispatch("getOwnProfile");
            resolve(response);
          })
          .catch((err) => reject(err));
      });
    },
    register(
      { commit, dispatch },
      payload: { displayName: string; username: string; password: string }
    ) {
      const request = new RegisterRequest();
      request.setDisplayName(payload.displayName);
      request.setUsername(payload.username);
      request.setPassword(payload.password);

      return new Promise<void>((resolve, reject) => {
        services.authClient
          .register(request)
          .then(() => {
            commit("setShouldShowTutorial", true);
            dispatch("login", payload).then(() => resolve());
          })
          .catch((err) => reject(err));
      });
    },
    logout({ commit }) {
      commit("clearAccessToken");
      router.push(LOGIN_ROUTE).then(() => {
        commit("resetGemsStore");
        commit("resetUserStore");
      });
    },
  },
};

export default authModule;
