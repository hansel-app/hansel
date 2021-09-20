import services from "@/api/services";
import { GetFriendsRequest, GetFriendsResponse } from "@/protobuf/friend_pb";
import { RootState } from "@/store";
import { Module } from "vuex";

export interface FriendsState {
  // List of friend's ids
  friends: number[];
}

const friendsModule: Module<FriendsState, RootState> = {
  state: {
    friends: [],
  },
  mutations: {
    setFriends(state, friends: number[]) {
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
            const friends: number[] = resp.getFriendsList();
            commit("setFriends", friends);
            resolve(resp);
          })
          .catch((err) => reject(err));
      });
    },
  },
};

export default friendsModule;
