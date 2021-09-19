<template><slot></slot></template>

<script lang="ts">
import { GemServicePromiseClient } from "@/protobuf/gem_grpc_web_pb";
import {
  Gem as ProtoGem,
  GetPendingCollectionByUserRequest,
} from "@/protobuf/gem_pb";
import { defineComponent, InjectionKey, provide } from "vue";
import { Gem } from "@/interfaces";
import dayjs from "dayjs";

const SERVER_URL = process.env.VUE_APP_SERVER_URL;
const SERVER_PORT = process.env.VUE_APP_SERVER_PORT;

export const FETCH_GEMS_PENDING_COLLECTION: InjectionKey<(
  userId: number
) => Promise<Gem[]>> = Symbol("Fetch Gems Pending Collection");

const protoGemToGemMapper = (protoGem: ProtoGem): Gem => {
  return {
    id: protoGem.getId(),
    message: protoGem.getMessage(),
    position: {
      lat: protoGem.getLatitude(),
      lng: protoGem.getLongitude(),
    },
    createdAt: dayjs(protoGem.getCreatedAt().toDate()),
    // TODO: replace this with actual user object
    // createdBy: protoGem.getCreatorId(),
    createdBy: "Bobby",
    receivedAt: dayjs(protoGem.getReceivedAt()?.toDate()),
  };
};

export default defineComponent({
  setup() {
    const hostname = `${SERVER_URL}:${SERVER_PORT}`;
    const client = new GemServicePromiseClient(hostname, null, null);

    const getGemsPendingCollectionByUser = async (
      userId: number
    ): Promise<Gem[]> => {
      const request = new GetPendingCollectionByUserRequest();
      request.setUserid(userId);

      try {
        return await client
          .getPendingCollectionByUser(request)
          .then((resp) => resp.getGemsList().map(protoGemToGemMapper));
      } catch (err) {
        if (err instanceof Error) {
          return Promise.reject(err);
        }
        return Promise.reject(new Error("Unknown error type"));
      }
    };

    provide(FETCH_GEMS_PENDING_COLLECTION, getGemsPendingCollectionByUser);
  },
});
</script>
