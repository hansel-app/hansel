<template>
  <div>
    <Header title="ah" :backLink="GEM_LOGS_ROUTE" />
    <li v-for="gem in gems" :key="gem.id">
      {{ gem.getMessage() }}
    </li>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { GEM_LOGS_ROUTE } from "@/constants";
import { Gem, GemLogsWithFriend } from "@/protobuf/gem_pb";
import { mockGemLogsWithFriend } from "@/interfaces/mockDataPb";
import Header from "@/components/Header.vue";

export default defineComponent({
  components: {
    Header,
  },
  data() {
    return {
      GEM_LOGS_ROUTE,
    };
  },
  computed: {
    gemLogs(): GemLogsWithFriend {
      return mockGemLogsWithFriend();
    },
    gems(): Gem[] {
      console.log(this.gemLogs.getGemsList());
      return this.gemLogs.getGemsList();
    },
  },
});
</script>
