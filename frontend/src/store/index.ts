import { createStore, Store } from "vuex";
import authModule, { AuthState } from "./modules/auth";
import gemsModule, { GemsState } from "./modules/gems";

export interface RootState {
  auth: AuthState;
  gems: GemsState;
}

const store: Store<RootState> = createStore({
  modules: {
    auth: authModule,
    gems: gemsModule,
  },
});

export default store;
