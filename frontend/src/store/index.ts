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
  plugins: [createPersistedState()],
});

export default store;
