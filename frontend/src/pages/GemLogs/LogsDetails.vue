<template>
  <div>
    <Header :backLink="GEM_LOGS_ROUTE">
      <p>{{ friend.displayName }}</p>
      <CircleAvatar :avatarUrl="placeholderAvatarUrl" />
    </Header>

    <GemMessage v-for="gem in gems" :key="gem.id" :message="gem.getMessage()" />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { GEM_LOGS_ROUTE } from "@/constants";
import { Gem, GemLogsWithFriend } from "@/protobuf/gem_pb";
import { User } from "@/interfaces/user";
import { mockGemLogsWithFriend } from "@/interfaces/mockDataPb";
import { protoUserToUserMapper } from "@/utils/mappers";
import CircleAvatar from "@/components/CircleAvatar.vue";
import Header from "@/components/Header.vue";
import GemMessage from "./GemMessage.vue";

export default defineComponent({
  components: {
    Header,
    CircleAvatar,
    GemMessage,
  },
  data() {
    return {
      GEM_LOGS_ROUTE,
      placeholderAvatarUrl:
        "https://cdn.mos.cms.futurecdn.net/JycrJzD5tvbGHWgjtPrRZY-970-80.jpg.webp",
    };
  },
  computed: {
    gemLogs(): GemLogsWithFriend {
      // TODO: replace with gem logs from store.
      return mockGemLogsWithFriend();
    },
    gems(): Gem[] {
      console.log(this.gemLogs.getGemsList());
      return this.gemLogs.getGemsList();
    },
    friend(): User {
      // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
      return protoUserToUserMapper(this.gemLogs.getFriend()!);
    },
  },
});
</script>
