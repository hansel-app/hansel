import * as Cookies from "js-cookie";
import { createStore, Store } from "vuex";
import createPersistedState from "vuex-persistedstate";
import authModule, { AuthState } from "./modules/auth";
import gemsModule, { GemsState } from "./modules/gems";
import userModule, { UserState } from "./modules/user";
export interface RootState {
  auth: AuthState;
  gems: GemsState;
  user: UserState;
}

const store: Store<RootState> = createStore({
  modules: {
    auth: authModule,
    gems: gemsModule,
    user: userModule,
  },

  plugins: [
    createPersistedState({
      getState: (key) => Cookies.getJSON(key),
      setState: (key, state) =>
        Cookies.set(key, state, { expires: 3, secure: true }),
    }),
  ],
});

export default store;
