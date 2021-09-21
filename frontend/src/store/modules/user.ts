import services from "@/api/services";
import {
  FriendInfo,
  GetFriendsRequest,
  GetFriendsResponse,
} from "@/protobuf/friend_pb";
import { RootState } from "@/store";
import { Module } from "vuex";

export interface UserState {
  // User's friends.
  // FriendInfo includes their id, username, displayName and avatar (TODO).
  friends: FriendInfo[];

  // TODO: Signed in user
  // TODO: user's current geolocation
}

const userModule: Module<UserState, RootState> = {
  state: {
    friends: [],
  },
  mutations: {
    setFriends(state, friends: FriendInfo[]) {
      state.friends.splice(0, state.friends.length);
      friends.forEach((friend) => state.friends.push(friend));
    },
  },
  actions: {
    getFriends({ commit }) {
      const request = new GetFriendsRequest();

      return new Promise((resolve, reject) => {
        services.friendsClient
          .getFriends(request)
          .then((resp: GetFriendsResponse) => {
            const friends: FriendInfo[] = resp.getFriendsList();
            commit("setFriends", friends);
            resolve(resp);
          })
          .catch((err) => reject(err));
      });
    },
  },
};

export default userModule;
