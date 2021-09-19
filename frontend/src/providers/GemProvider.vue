<template><slot></slot></template>

<script lang="ts">
import {
  Gem as ProtoGem,
  GetPendingCollectionForUserRequest,
} from "@/protobuf/gem_pb";
import { defineComponent, InjectionKey, provide } from "vue";
import { Gem } from "@/interfaces";
import dayjs from "dayjs";
import services from "@/api/services";

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
    createdAt: dayjs(protoGem.getCreatedAt()?.toDate()),
    // TODO: replace this with actual user object
    createdBy: {
      id: 9999,
      username: "bobbythebobcat",
      displayName: "bobby",
      avatar:
        "https://cdn.mos.cms.futurecdn.net/JycrJzD5tvbGHWgjtPrRZY-970-80.jpg.webp",
    },
    receivedAt: dayjs(protoGem.getReceivedAt()?.toDate()),
  };
};

export default defineComponent({
  setup() {
    const client = services.gemsClient

    const getGemsPendingCollectionForUser = async (): Promise<Gem[]> => {
      const request = new GetPendingCollectionForUserRequest();

      try {
        return await client
          .getPendingCollectionForUser(request)
          .then((resp) => resp.getGemsList().map(protoGemToGemMapper));
      } catch (err) {
        return Promise.reject(err);
      }
    };

    provide(FETCH_GEMS_PENDING_COLLECTION, getGemsPendingCollectionForUser);
  },
});
</script>
