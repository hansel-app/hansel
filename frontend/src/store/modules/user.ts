import services from "@/api/services";
import { User, LatLng } from "@/interfaces";
import { SINGAPORE_CENTER } from "@/constants";
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
  isSendFriendRequestSuccessful: boolean;
  self: User;
  currPosition: LatLng;
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
    isSendFriendRequestSuccessful: false,
    currPosition: SINGAPORE_CENTER,
  },
  mutations: {
    setCurrPosition(state, newPosition: LatLng) {
      state.currPosition = newPosition;
    },
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
    // Used an action here instead of directly committing as
    // actions are asynchronous.
    updateCurrentPosition({ commit }, newPosition) {
      commit("setCurrPosition", newPosition);
    },
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
            commit("updateSendFriendRequestStatus", true);
            resolve(resp);
          })
          .catch((err) => reject(err));
      });
    },

    acceptFriendRequest({ dispatch }, requester_id) {
      const request = new FriendRequest();
      // Note: receiver_id (self) will be retrieved
      // from context keys in the backend
      request.setRequesterId(requester_id);

      return new Promise((resolve, reject) => {
        services.friendsClient
          .accept(request)
          .then((resp) => {
            dispatch("getFriends");
            dispatch("getFriendRequests");
            resolve(resp);
          })
          .catch((err) => reject(err));
      });
    },

    declineFriendRequest({ dispatch }, requester_id) {
      const request = new FriendRequest();
      // Note: receiver_id (self) will be retrieved
      // from context keys in the backend
      request.setRequesterId(requester_id);

      return new Promise((resolve, reject) => {
        services.friendsClient
          .decline(request)
          .then((resp) => {
            dispatch("getFriendRequests");
            resolve(resp);
          })
          .catch((err) => reject(err));
      });
    },
  },
};

export default userModule;
