<template>
  <div>
    <Header title="ah" :backLink="GEM_LOGS_ROUTE" />
    <FriendCell :friend="friend" />
    <GemMessage v-for="gem in gems" :key="gem.id" :message="gem.getMessage()" />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { GEM_LOGS_ROUTE } from "@/constants";
import { Gem, GemLogsWithFriend } from "@/protobuf/gem_pb";
import { User } from "@/interfaces/user";
import { mockGemLogsWithFriend } from "@/interfaces/mockDataPb";
import Header from "@/components/Header.vue";
import GemMessage from "./GemMessage.vue";
import FriendCell from "@/components/FriendCell.vue";

export default defineComponent({
  components: {
    Header,
    GemMessage,
    FriendCell,
  },
  data() {
    return {
      GEM_LOGS_ROUTE,
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
      return (this.gemLogs.getFriend()!).toObject();
    },
  },
});
</script>
