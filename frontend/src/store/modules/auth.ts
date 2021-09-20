import { Module } from "vuex";
import { RootState } from "@/store";
import services from "@/api/services";
import { LoginRequest, LoginResponse } from "@/protobuf/auth_pb";

export interface AuthState {
  accessToken: Nullable<string>;
}

const authModule: Module<AuthState, RootState> = {
  state: {
    accessToken: null
  },
  getters: {
    isLoggedIn(state): boolean {
      return state.accessToken != null;
    }
  },
  mutations: {
    setAccessToken(state, accessToken: string) {
      state.accessToken = accessToken;
    },
    clearAccessToken(state) {
      state.accessToken = null;
    }
  },
  actions: {
    login({ commit }, payload: { username: string, password: string }) {
      const request = new LoginRequest();
      request.setUsername(payload.username);
      request.setPassword(payload.password);

      return new Promise((resolve, reject) => {
        services.authClient
          .login(request)
          .then((response: LoginResponse) => {
            const token = response.getAccessToken();
            commit('setAccessToken', token);
            resolve(response);
          })
          .catch(err => reject(err));
      });
    }
  }
};

export default authModule;
