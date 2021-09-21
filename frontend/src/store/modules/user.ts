import services from "@/api/services";
import { User } from "@/interfaces";
import {
  FriendInfo,
  GetFriendsRequest,
  GetFriendsResponse,
  GetPendingFriendRequestsRequest,
  GetPendingFriendRequestsResponse,
} from "@/protobuf/friend_pb";
import { RootState } from "@/store";
import { Module } from "vuex";

export interface UserState {
  // User's friends.
  // FriendInfo includes their id, username, displayName and avatar (TODO).
  friends: FriendInfo[];
  friendRequests: FriendInfo[];
  self: User;

  // TODO: user's current geolocation
}

const userModule: Module<UserState, RootState> = {
  state: {
    friends: [],
    friendRequests: [],
    // TODO: replace hardcode
    self: {
      id: 100,
      username: "maggieMee",
      displayName: "Maggie Mee",
      avatar: "avatarrrr",
    },
  },
  mutations: {
    setFriends(state, friends: FriendInfo[]) {
      state.friends.splice(0, state.friends.length);
      friends.forEach((friend) => state.friends.push(friend));
    },
    setFriendRequests(state, pendingFriends: FriendInfo[]) {
      state.friends.splice(0, state.friends.length);
      pendingFriends.forEach((friendRequest) =>
        state.friendRequests.push(friendRequest)
      );
    },
  },
  actions: {
    getFriends({ commit }) {
      const request = new GetFriendsRequest();
      request.setUserId(this.state.user.self.id);

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
    getFriendRequests({ commit }) {
      const request = new GetPendingFriendRequestsRequest();
      request.setUserId(this.state.user.self.id);

      return new Promise((resolve, reject) => {
        services.friendsClient
          .getFriendRequests(request)
          .then((resp: GetPendingFriendRequestsResponse) => {
            const pendingFriends: FriendInfo[] = resp.getFriendsList();
            commit("setFriendRequests", pendingFriends);
            resolve(resp);
          })
          .catch((err) => reject(err));
      });
    },
  },
};

export default userModule;
