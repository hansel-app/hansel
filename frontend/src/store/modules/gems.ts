import services from "@/api/services";
import { Gem, GemColor } from "@/interfaces";
import { mockFriends, mockSelfUser } from "@/interfaces/mockData";
import {
  Gem as ProtoGem,
  GemColor as ProtoGemColor,
  GetPendingCollectionForUserRequest,
  GetPendingCollectionForUserResponse,
} from "@/protobuf/gem_pb";
import { RootState } from "@/store";
import dayjs from "dayjs";
import { Module } from "vuex";

export interface GemsState {
  gemsPendingCollection: Gem[];
}

const protoGemColorToGemColorMapper = (
  protoGemColor: ProtoGemColor
): GemColor => {
  switch (protoGemColor) {
    case ProtoGemColor.PURPLE:
      return GemColor.PURPLE;
    case ProtoGemColor.PINK:
      return GemColor.PINK;
    case ProtoGemColor.BLUE:
      return GemColor.BLUE;
    case ProtoGemColor.BLACK:
      return GemColor.BLACK;
    case ProtoGemColor.YELLOW:
      return GemColor.YELLOW;
    case ProtoGemColor.GREEN:
      return GemColor.GREEN;
    default:
      throw new Error("Unknown gem color received!");
  }
};

const protoGemToGemMapper = (protoGem: ProtoGem): Gem => {
  return {
    id: protoGem.getId(),
    message: protoGem.getMessage(),
    position: {
      lat: protoGem.getLatitude(),
      lng: protoGem.getLongitude(),
    },
    createdAt: dayjs(protoGem.getCreatedAt()?.toDate()),
    // TODO: replace these with actual user object
    createdBy: mockFriends[0],
    receiver: mockSelfUser,
    receivedAt: dayjs(protoGem.getReceivedAt()?.toDate()),
    color: protoGemColorToGemColorMapper(protoGem.getColor()),
  };
};

const gemsModule: Module<GemsState, RootState> = {
  state: {
    gemsPendingCollection: [],
  },
  mutations: {
    setGemsPendingCollection(state, gems: Gem[]) {
      state.gemsPendingCollection.splice(0, state.gemsPendingCollection.length);
      gems.forEach((gem) => state.gemsPendingCollection.push(gem));
    },
  },
  actions: {
    getGemsPendingCollectionForUser({ commit }) {
      const request = new GetPendingCollectionForUserRequest();

      return new Promise((resolve, reject) => {
        services.gemsClient
          .getPendingCollectionForUser(request)
          .then((resp: GetPendingCollectionForUserResponse) => {
            const gems: Gem[] = resp.getGemsList().map(protoGemToGemMapper);
            commit("setGemsPendingCollection", gems);
            resolve(resp);
          })
          .catch((err) => reject(err));
      });
    },
  },
};

export default gemsModule;
