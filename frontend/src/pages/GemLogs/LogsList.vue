<template>
  <div class="background-gradient">
    <BackSwipe />
    <Header title="Gem logs" />
    <div class="container">
      <Searchbar placeholder="Search a friend" @input="onSearchQueryInput" />
      <CellGroup>
        <LogsPreview
          v-for="friendAndActivity in state.filteredFriendsAndActivities"
          :key="friendAndActivity.friend.userId"
          :friend="friendAndActivity.friend"
          :mostRecentGemActivity="friendAndActivity.mostRecentGem"
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
import { Gem, User } from "@/interfaces";

interface FriendAndMostRecentGemActivity {
  friend: User;
  mostRecentGem: Gem;
}

export default defineComponent({
  components: {
    Header,
    Searchbar,
    LogsPreview,
    BackSwipe,
    CellGroup,
  },
  setup() {
    // TODO: replace mockfriends and mockGems
    const mockSingleGem = mockGems[0];
    const allFriends: FriendAndMostRecentGemActivity[] = mockFriends.map(
      (friend) => {
        return {
          friend: friend,
          mostRecentGem: mockSingleGem,
        };
      }
    );

    const state = reactive({
      searchQuery: "",
      isCloseWindow: true,
      filteredFriendsAndActivities: allFriends as FriendAndMostRecentGemActivity[],
    });

    return {
      state,
      allFriends,
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
        this.state.filteredFriendsAndActivities = this.allFriends;
        return;
      }

      this.state.filteredFriendsAndActivities = this.allFriends.filter(
        (friendAndGemActivity: FriendAndMostRecentGemActivity) => {
          const friend = friendAndGemActivity.friend;

          const matchUsername = friend.username
            .toLowerCase()
            .includes(this.state.searchQuery.toLowerCase());
          const matchDisplayname = friend.displayName
            .toLowerCase()
            .includes(this.state.searchQuery.toLowerCase());
          return matchUsername || matchDisplayname;
        }
      );
    },
  },
});
</script>
