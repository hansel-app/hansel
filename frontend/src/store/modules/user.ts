import services from "@/api/services";
import { SINGAPORE_CENTER } from "@/constants";
import { LatLng, PendingFriendRequest, User } from "@/interfaces";
import {
  EditProfileRequest,
  FriendRequest,
  GetFriendRequestsResponse,
  GetFriendsResponse,
  GetOwnProfileResponse,
  SearchByUsernameRequest,
  SearchByUsernameResponse,
} from "@/protobuf/user_pb";
import { RootState } from "@/store";
import { blobToUint8Array } from "@/utils/attachment";
import {
  protoPendingFriendRequestToPendingFriendRequestMapper,
  protoUserToUserMapper,
} from "@/utils/mappers";
import { Empty } from "google-protobuf/google/protobuf/empty_pb";
import cloneDeep from "lodash.clonedeep";
import { Module } from "vuex";

export interface UserState {
  friends: User[];

  // PendingFriendRequests includes User + datetime of request.
  friendRequests: PendingFriendRequest[];
  self?: User;
  currPosition: LatLng;
}

const initialState: UserState = {
  friends: [],
  friendRequests: [],
  self: undefined,
  currPosition: SINGAPORE_CENTER,
};

const userModule: Module<UserState, RootState> = {
  state: cloneDeep(initialState),
  mutations: {
    resetUserStore(state) {
      Object.assign(state, cloneDeep(initialState));
    },
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
      state.friendRequests = pendingFriends;
    },
  },
  actions: {
    // Used an action here instead of directly committing as
    // actions are asynchronous.
    updateCurrentPosition({ commit }, newPosition) {
      commit("setCurrPosition", newPosition);
    },
    getOwnProfile({ commit }) {
      return new Promise((resolve, reject) => {
        services.userClient
          .getOwnProfile(new Empty())
          .then((resp: GetOwnProfileResponse) => {
            const friends: User[] = resp
              .getFriendsList()
              .map(protoUserToUserMapper);
            commit("setFriends", friends);

            const self = resp.getInfo();
            if (!self) {
              throw new Error(
                "Failed to retrieve information about current logged in user"
              );
            }
            commit("setSelfInfo", protoUserToUserMapper(self));

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
            const friends: User[] = resp
              .getFriendsList()
              .map(protoUserToUserMapper);
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
            const matchedUsers: User[] = resp
              .getUsersList()
              .map(protoUserToUserMapper);
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
            const pendingFriendRequests: PendingFriendRequest[] = resp
              .getFriendRequestsList()
              .map(protoPendingFriendRequestToPendingFriendRequestMapper);
            commit("setFriendRequests", pendingFriendRequests);
            resolve(resp);
          })
          .catch((err) => reject(err));
      });
    },

    sendFriendRequest(_, receiverId) {
      const request = new FriendRequest();
      request.setReceiverId(receiverId);

      return new Promise((resolve, reject) => {
        services.userClient
          .addFriendRequest(request)
          .then((resp) => resolve(resp))
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

    async editOwnProfile(
      _,
      payload: { newDisplayName?: string; newAvatar?: File }
    ) {
      const request = new EditProfileRequest();
      if (payload.newDisplayName)
        request.setNewDisplayName(payload.newDisplayName);
      if (payload.newAvatar)
        request.setNewAvatar(await blobToUint8Array(payload.newAvatar));

      return new Promise((resolve, reject) => {
        services.userClient
          .editOwnProfile(request)
          .then((resp) => {
            // Ideally there is a separate RPC call that doesn't
            // fetch friends here.
            this.dispatch("getOwnProfile");
            resolve(resp);
          })
          .catch((err) => reject(err));
      });
    },
  },
};

export default userModule;
