import services from "@/api/services";
import { LatLng, User } from "@/interfaces";
import { SINGAPORE_CENTER } from "@/constants";
import {
  User as ProtoUser,
  GetOwnProfileResponse,
  GetFriendsResponse,
  GetFriendRequestsResponse,
  PendingFriendRequest,
  FriendRequest,
  SearchByUsernameRequest,
  SearchByUsernameResponse,
} from "@/protobuf/user_pb";
import { RootState } from "@/store";
import { Module } from "vuex";
import { Empty } from "google-protobuf/google/protobuf/empty_pb";

export interface UserState {
  friends: User[];

  // PendingFriendRequests includes User + datetime of request.
  friendRequests: PendingFriendRequest[];
  isSendFriendRequestSuccessful: boolean;
  self?: User;
  currPosition: LatLng;
}

const protoUserToUserMapper = (protoUser: ProtoUser): User => {
  return {
    id: protoUser.getUserId(),
    username: protoUser.getUsername(),
    displayName: protoUser.getDisplayName(),
    avatar: protoUser.getAvatar_asB64(),
  };
};

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
    setSelfInfo(state, self: User) {
      state.self = self;
    },
    setFriends(state, friends: User[]) {
      state.friends = friends;
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
    getOwnProfile({ commit }) {
      return new Promise((resolve, reject) => {
        services.userClient
          .getOwnProfile(new Empty())
          .then((resp: GetOwnProfileResponse) => {
            const friends: User[] = resp.getFriendsList().map(protoUserToUserMapper);
            commit("setFriends", friends);
            const self = resp.getInfo();
            commit("setSelfInfo", self);
            resolve(resp);
          })
          .catch((err) => reject(err));
      });
    },
    getFriends({ commit }) {
      return new Promise((resolve, reject) => {
        services.userClient
          .getFriends(new Empty())
          .then((resp: GetFriendsResponse) => {
            const friends: User[] = resp.getFriendsList().map(protoUserToUserMapper);
            commit("setFriends", friends);
            resolve(resp);
          })
          .catch((err) => reject(err));
      });
    },
    searchByUsername(_, searchQuery: string) {
      const request = new SearchByUsernameRequest();
      request.setSearchQuery(searchQuery);

      return new Promise((resolve, reject) => {
        services.userClient
          .searchByUsername(request)
          .then((resp: SearchByUsernameResponse) => {
            const matchedUsers: User[] = resp.getUsersList().map(protoUserToUserMapper);
            resolve(matchedUsers);
          })
          .catch((err) => reject(err));
      });
    },
    getFriendRequests({ commit }) {
      return new Promise((resolve, reject) => {
        services.userClient
          .getFriendRequests(new Empty())
          .then((resp: GetFriendRequestsResponse) => {
            const pendingFriends: PendingFriendRequest[] = resp.getFriendRequestsList();
            commit("setFriendRequests", pendingFriends);
            resolve(resp);
          })
          .catch((err) => reject(err));
      });
    },

    sendFriendRequest({ commit }, receiverId) {
      const request = new FriendRequest();
      request.setReceiverId(receiverId);

      return new Promise((resolve, reject) => {
        services.userClient
          .addFriendRequest(request)
          .then((resp) => {
            commit("updateSendFriendRequestStatus", true);
            resolve(resp);
          })
          .catch((err) => reject(err));
      });
    },

    acceptFriendRequest({ dispatch }, requesterId) {
      const request = new FriendRequest();
      request.setRequesterId(requesterId);

      return new Promise((resolve, reject) => {
        services.userClient
          .acceptFriendRequest(request)
          .then((resp) => {
            dispatch("getFriends");
            dispatch("getFriendRequests");
            resolve(resp);
          })
          .catch((err) => reject(err));
      });
    },

    declineFriendRequest({ dispatch }, requesterId) {
      const request = new FriendRequest();
      request.setRequesterId(requesterId);

      return new Promise((resolve, reject) => {
        services.userClient
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
