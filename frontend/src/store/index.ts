import {createStore, Store} from "vuex";
import authModule, { AuthState } from "./modules/auth";

export interface RootState {
  auth: AuthState;
}

const store: Store<RootState> = createStore({
  modules: {
    auth: authModule
  }
});

export default store;
