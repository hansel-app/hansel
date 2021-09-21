import services from "@/api/services";
import { User } from "@/interfaces";
import {
  FriendInfo,
  FriendRequest,
  GetFriendsRequest,
  GetFriendsResponse,
  GetPendingFriendRequestsRequest,
  GetPendingFriendRequestsResponse,
  PendingFriendRequest,
} from "@/protobuf/friend_pb";
import { RootState } from "@/store";
import { Module } from "vuex";

export interface UserState {
  // User's friends.
  // FriendInfo includes their id, username, displayName and avatar (TODO).
  friends: FriendInfo[];

  // PendingFriendRequests includes FriendInfo + datetime of request.
  friendRequests: PendingFriendRequest[];
  self: User;

  // TODO: user's current geolocation

  isSendFriendRequestSuccessful: boolean;
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
    setFriendRequests(state, pendingFriends: PendingFriendRequest[]) {
      state.friends.splice(0, state.friends.length);
      pendingFriends.forEach((friendRequest) =>
        state.friendRequests.push(friendRequest)
      );
    },
    updateSendFriendRequestStatus(state, isSuccess: boolean) {
      state.isSendFriendRequestSuccessful = isSuccess;
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
    getFriendRequests({ commit }) {
      const request = new GetPendingFriendRequestsRequest();

      return new Promise((resolve, reject) => {
        services.friendsClient
          .getFriendRequests(request)
          .then((resp: GetPendingFriendRequestsResponse) => {
            const pendingFriends: PendingFriendRequest[] = resp.getFriendRequestsList();
            commit("setFriendRequests", pendingFriends);
            resolve(resp);
          })
          .catch((err) => reject(err));
      });
    },

    sendFriendRequest({ commit }, receiver_id) {
      const request = new FriendRequest();
      // Note: requester_id (self) will be retrieved
      // from context keys in the backend
      request.setReceiverId(receiver_id);

      return new Promise((resolve, reject) => {
        services.friendsClient
          .add(request)
          .then((resp) => {
            // Is it possible to call getFriends here?
            commit("updateSendFriendRequestStatus", true);
            resolve(resp);
          })
          .catch((err) => reject(err));
      });
    },

    acceptFriendRequest({ commit },requester_id) {
      const request = new FriendRequest();
      // Note: receiver_id (self) will be retrieved
      // from context keys in the backend
      request.setRequesterId(requester_id);

      return new Promise((resolve, reject) => {
        services.friendsClient
          .accept(request)
          .then((resp) => {
            // Ditto with above, getFriends to retrieve updated list.
            resolve(resp);
          })
          .catch((err) => reject(err));
      });
    },

    declineFriendRequest({ commit }, requester_id) {
      const request = new FriendRequest();
      // Note: receiver_id (self) will be retrieved
      // from context keys in the backend
      request.setRequesterId(requester_id);

      return new Promise((resolve, reject) => {
        services.friendsClient
          .decline(request)
          .then((resp) => {
            resolve(resp);
          })
          .catch((err) => reject(err));
      });
    },
  },
};

export default userModule;
