<template>
  <div class="background-gradient">
    <BackSwipe />
    <Header title="Gem logs" />
    <div class="container">
      <Searchbar placeholder="Search a friend" />
      <CellGroup>
        <LogsPreview
          v-for="friend in mockFriends"
          :key="friend.id"
          :friend="friend"
          :mostRecentGemActivity="mockSingleGem"
          :isClickable="true"
        />
      </CellGroup>
    </div>
  </div>
</template>
<script lang="ts">
import { defineComponent } from "vue";
import { mockFriends, mockGems } from "@/interfaces/mockData.ts";
import Header from "@/components/Header.vue";
import Searchbar from "@/components/Searchbar.vue";
import BackSwipe from "@/components/BackSwipe.vue";
import { CellGroup } from "vant";

import LogsPreview from "./LogsPreview.vue";

export default defineComponent({
  components: {
    Header,
    Searchbar,
    LogsPreview,
    BackSwipe,
    CellGroup,
  },
  data() {
    return {
      isCloseWindow: true,
      mockFriends,
      mockSingleGem: mockGems[0],
    };
  },
});
</script>
