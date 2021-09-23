<template>
  <div class="background-gradient">
    <BackSwipe />
    <Header title="Gem logs" />
    <div class="container">
      <Searchbar placeholder="Search a friend" @input="onSearchQueryInput" />
      <CellGroup>
        <LogsPreview
          v-for="friend in state.filteredFriends"
          :key="friend.userId"
          :friend="friend"
          :mostRecentGemActivity="mockSingleGem"
          :isClickable="true"
        />
      </CellGroup>
    </div>
  </div>
</template>
<script lang="ts">
import { defineComponent, reactive } from "vue";
import { mockFriends, mockGems } from "@/interfaces/mockData";
import Header from "@/components/Header.vue";
import Searchbar from "@/components/Searchbar.vue";
import BackSwipe from "@/components/BackSwipe.vue";
import { CellGroup } from "vant";

import LogsPreview from "./LogsPreview.vue";
import { User } from "@/interfaces";

export default defineComponent({
  components: {
    Header,
    Searchbar,
    LogsPreview,
    BackSwipe,
    CellGroup,
  },
  setup() {
    const allFriends = mockFriends;

    const state = reactive({
      searchQuery: "",
      isCloseWindow: true,
      // TODO: replace mockfriends and mockGems
      filteredFriends: allFriends as User[],
    });

    return {
      state,
      allFriends,
      mockSingleGem: mockGems[0],
    };
  },
  methods: {
    onSearchQueryInput(event: Event) {
      const target = event.target as HTMLInputElement;
      this.state.searchQuery = target.value;
      this.filterFriends();
    },
    filterFriends() {
      if (this.state.searchQuery.length === 0) {
        this.state.filteredFriends = this.allFriends;
        return;
      }

      this.state.filteredFriends = this.allFriends.filter((friend: User) => {
        const matchUsername = friend.username
          .toLowerCase()
          .includes(this.state.searchQuery.toLowerCase());
        const matchDisplayname = friend.displayName
          .toLowerCase()
          .includes(this.state.searchQuery.toLowerCase());
        return matchUsername || matchDisplayname;
      });
    },
  },
});
</script>
