<template><slot></slot></template>

<script lang="ts">
import {
  Gem as ProtoGem,
  GemColor as ProtoGemColor,
  GetPendingCollectionForUserRequest,
  DropRequest,
  GemMessage
} from "@/protobuf/gem_pb";
import { defineComponent, InjectionKey, provide } from "vue";
import { Gem, GemColor } from "@/interfaces";
import dayjs from "dayjs";
import services from "@/api/services";
import { mockSelfUser } from "@/interfaces/mockData";

export const FETCH_GEMS_PENDING_COLLECTION: InjectionKey<(
  userId: number
) => Promise<Gem[]>> = Symbol("Fetch Gems Pending Collection");

export const DROP_GEM: InjectionKey<(
  gemMessage: GemMessage
) => Promise<string>> = Symbol("Drop Gem");

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
    // TODO: replace this with actual user object
    createdBy: {
      id: 9999,
      username: "bobbythebobcat",
      displayName: "bobby",
      avatar:
        "https://cdn.mos.cms.futurecdn.net/JycrJzD5tvbGHWgjtPrRZY-970-80.jpg.webp",
    },
    receiver: mockSelfUser,
    receivedAt: dayjs(protoGem.getReceivedAt()?.toDate()),
    color: protoGemColorToGemColorMapper(protoGem.getColor()),
  };
};


export default defineComponent({
  setup() {
    const client = services.gemsClient;
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

    const dropGem = async (gemMessage : GemMessage) : Promise<string> => {
      const request = new DropRequest();
      request.setGemMessage(gemMessage);

      try {
        return await client
        .drop(request)
        .then((resp) => resp.getMessage())
      } catch (err) {
        return Promise.reject(err);
      }
    };

    provide(FETCH_GEMS_PENDING_COLLECTION, getGemsPendingCollectionForUser);
    provide(DROP_GEM, dropGem);
  },
});
</script>
