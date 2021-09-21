import services from "@/api/services";
import { LatLng } from "@/interfaces";
import { SINGAPORE_CENTER } from "@/constants";
import {
  PersonInfo,
  ProfileRequest,
  ProfileResponse,
  GetFriendsRequest,
  GetFriendsResponse,
  GetPendingFriendRequestsRequest,
  GetPendingFriendRequestsResponse,
  PendingFriendRequest,
} from "@/protobuf/user_pb";
import { FriendRequest } from "@/protobuf/friend_pb";
import { RootState } from "@/store";
import { Module } from "vuex";

export interface UserState {
  // User's friends.
  // PersonInfo includes their id, username, displayName and avatar (TODO).
  friends: PersonInfo[];

  // PendingFriendRequests includes PersonInfo + datetime of request.
  friendRequests: PendingFriendRequest[];
  isSendFriendRequestSuccessful: boolean;
  self: PersonInfo | undefined;
  currPosition: LatLng;
}

const userModule: Module<UserState, RootState> = {
  state: {
    friends: [],
    friendRequests: [],
    self: undefined,
    isSendFriendRequestSuccessful: false,
    currPosition: SINGAPORE_CENTER,
  },
  mutations: {
    setCurrPosition(state, newPosition: LatLng) {
      state.currPosition = newPosition;
    },
    setSelfInfo(state, info: PersonInfo) {
      state.self = info;
    },
    setFriends(state, friends: PersonInfo[]) {
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
      console.log(newPosition);
      commit("setCurrPosition", newPosition);
    },
    getMyProfile({ commit }) {
      const request = new ProfileRequest();

      return new Promise((resolve, reject) => {
        services.userClient
          .getProfile(request)
          .then((resp: ProfileResponse) => {
            const friends: PersonInfo[] = resp.getFriendsList();
            commit("setFriends", friends);
            const self = resp.getInfo();
            commit("setSelfInfo", self);
            resolve(resp);
          })
          .catch((err) => reject(err));
      });
    },
    getFriends({ commit }) {
      const request = new GetFriendsRequest();

      return new Promise((resolve, reject) => {
        services.userClient
          .getFriends(request)
          .then((resp: GetFriendsResponse) => {
            const friends: PersonInfo[] = resp.getFriendsList();
            commit("setFriends", friends);
            resolve(resp);
          })
          .catch((err) => reject(err));
      });
    },
    getFriendRequests({ commit }) {
      const request = new GetPendingFriendRequestsRequest();

      return new Promise((resolve, reject) => {
        services.userClient
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
          .addFriendRequest(request)
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
          .acceptFriendRequest(request)
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
          .declineFriendRequest(request)
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
