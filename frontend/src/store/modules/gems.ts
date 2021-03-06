import services from "@/api/services";
import {
  DropGemFormState,
  Gem,
  GemColor,
  GemLogs,
  GemLogsWithFriend,
} from "@/interfaces";
import {
  DropRequest,
  DropResponse,
  GemLogs as ProtoGemLogs,
  GemMessage,
  GetPendingCollectionForUserResponse,
  PickUpRequest,
} from "@/protobuf/gem_pb";
import { RootState } from "@/store";
import { compressFileAndConvertToUint8Array } from "@/utils/attachment";
import {
  gemColorToProtoGemColorMapper,
  protoGemLogToGemLogMapper,
  protoGemToGemMapper,
} from "@/utils/mappers";
import { Empty } from "google-protobuf/google/protobuf/empty_pb";
import cloneDeep from "lodash/cloneDeep";
import { Module } from "vuex";
export interface GemsState {
  isGemMapFirstMount?: boolean;
  gemsPendingCollection: Gem[];
  dropGemFormState: DropGemFormState;
  lastPickedUpGem?: Gem;
  gemLogs?: GemLogs;
  selectedGemLog?: GemLogsWithFriend;
}

const initialDropGemFormState = { color: GemColor.PURPLE };
const initialState: GemsState = {
  isGemMapFirstMount: true,
  gemsPendingCollection: [],
  dropGemFormState: initialDropGemFormState,
  lastPickedUpGem: undefined,
  gemLogs: undefined,
  selectedGemLog: undefined,
};

const gemsModule: Module<GemsState, RootState> = {
  state: cloneDeep(initialState),
  mutations: {
    resetGemsStore(state) {
      Object.assign(state, cloneDeep(initialState));
    },
    setIsGemMapFirstMount(state, isFirstMount: boolean) {
      state.isGemMapFirstMount = isFirstMount;
    },
    setGemsPendingCollection(state, gems: Gem[]) {
      state.gemsPendingCollection = gems;
    },
    setGemLogs(state, gemLogs: GemLogs) {
      state.gemLogs = gemLogs;
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
    setLastPickedUpGem(state, gem: Gem) {
      state.lastPickedUpGem = gem;
    },
    setSelectedGemLog(state, gemLogWithFriend: GemLogsWithFriend) {
      state.selectedGemLog = gemLogWithFriend;
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
    getGemLogs({ commit }) {
      return new Promise((resolve, reject) => {
        services.gemsClient
          .getGemLogs(new Empty())
          .then((resp: ProtoGemLogs) => {
            const gemLogs = protoGemLogToGemLogMapper(resp);
            commit("setGemLogs", gemLogs);
            resolve(gemLogs);
          })
          .catch((err) => reject(err));
      });
    },
    async dropGem({ commit }) {
      const dropGemFormState = this.state.gems.dropGemFormState;
      console.assert(!dropGemFormState.id, "Form state was not reset properly");

      if (
        !dropGemFormState.message ||
        !dropGemFormState.receiverId ||
        !dropGemFormState.color ||
        !dropGemFormState.attachment
      ) {
        throw new Error("Some fields are missing");
      }

      const gemMessage = new GemMessage();
      gemMessage.setMessage(dropGemFormState.message);
      gemMessage.setReceiverId(dropGemFormState.receiverId);
      gemMessage.setColor(
        gemColorToProtoGemColorMapper(dropGemFormState.color)
      );
      gemMessage.setAttachment(
        await compressFileAndConvertToUint8Array(dropGemFormState.attachment)
      );
      gemMessage.setLatitude(this.state.user.currPosition.lat);
      gemMessage.setLongitude(this.state.user.currPosition.lng);

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
    pickUpGem({ commit, dispatch }, payload: { gem: Gem }) {
      const pickUpRequest = new PickUpRequest();
      pickUpRequest.setId(payload.gem.id);

      return new Promise((resolve, reject) => {
        services.gemsClient
          .pickUp(pickUpRequest)
          .then((resp) => {
            commit("setLastPickedUpGem", payload.gem);
            dispatch("getGemsPendingCollectionForUser");
            resolve(resp);
          })
          .catch((err) => reject(err));
      });
    },
  },
};

export default gemsModule;
