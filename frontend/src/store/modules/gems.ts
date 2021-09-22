import services from "@/api/services";
import { DropGemFormState, Gem, GemColor } from "@/interfaces";
import { mockFriends, mockSelfUser } from "@/interfaces/mockData";
import {
  DropRequest,
  DropResponse,
  Gem as ProtoGem,
  GemColor as ProtoGemColor,
  GemMessage,
  GetPendingCollectionForUserResponse,
} from "@/protobuf/gem_pb";
import { RootState } from "@/store";
import dayjs from "dayjs";
import { Module } from "vuex";
import { Empty } from "google-protobuf/google/protobuf/empty_pb";

export interface GemsState {
  gemsPendingCollection: Gem[];
  dropGemFormState: DropGemFormState;
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

const initialDropGemFormState = Object.freeze({ color: GemColor.PURPLE });

const gemsModule: Module<GemsState, RootState> = {
  state: {
    gemsPendingCollection: [],
    dropGemFormState: initialDropGemFormState,
  },
  mutations: {
    setGemsPendingCollection(state, gems: Gem[]) {
      state.gemsPendingCollection = gems;
    },
    updateDropGemFormState(state, dropGemFormState: Partial<DropGemFormState>) {
      state.dropGemFormState = {
        ...state.dropGemFormState,
        ...dropGemFormState,
      };
    },
    clearDropGemFormState(state) {
      state.dropGemFormState = initialDropGemFormState;
    },
  },
  actions: {
    getGemsPendingCollectionForUser({ commit }) {
      return new Promise((resolve, reject) => {
        services.gemsClient
          .getPendingCollectionForUser(new Empty())
          .then((resp: GetPendingCollectionForUserResponse) => {
            const gems: Gem[] = resp.getGemsList().map(protoGemToGemMapper);
            commit("setGemsPendingCollection", gems);
            resolve(resp);
          })
          .catch((err) => reject(err));
      });
    },
    dropGem({ commit }) {
      const dropGemFormState = this.state.gems.dropGemFormState;
      console.assert(!dropGemFormState.id, "Form state was not reset properly");

      if (
        !dropGemFormState.message ||
        !dropGemFormState.receiverId ||
        !dropGemFormState.color
      ) {
        throw new Error("Missing fields when trying to drop a gem");
      }

      const gemMessage = new GemMessage();
      gemMessage.setMessage(dropGemFormState.message);
      gemMessage.setReceiverId(dropGemFormState.receiverId);
      gemMessage.setColor(dropGemFormState.color);
      // TODO:  populate latlng from store

      const request = new DropRequest();
      request.setGemMessage(gemMessage);

      return new Promise((resolve, reject) => {
        services.gemsClient
          .drop(request)
          .then((resp: DropResponse) => {
            const droppedGemId: number = resp.getDroppedGemId();
            commit("updateDropGemFormState", { id: droppedGemId });
            resolve(resp);
          })
          .catch((err) => reject(err));
      });
    },
  },
};

export default gemsModule;
