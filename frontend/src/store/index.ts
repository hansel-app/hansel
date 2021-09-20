import { createStore, Store } from "vuex";
import authModule, { AuthState } from "./modules/auth";
import gemModule, { GemState } from "./modules/gem";

export interface RootState {
  auth: AuthState;
  gems: GemState;
}

const store: Store<RootState> = createStore({
  modules: {
    auth: authModule,
    gems: gemModule,
  },
});

export default store;
