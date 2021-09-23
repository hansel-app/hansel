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
import Header from "@/components/Header.vue";
import Searchbar from "@/components/Searchbar.vue";
import BackSwipe from "@/components/BackSwipe.vue";
import { CellGroup } from "vant";

import LogsPreview from "./LogsPreview.vue";
import { User, Gem, GemLogs } from "@/interfaces";
import { useStore } from "vuex";

interface FriendAndMostRecentGemActivity {
  friend: User;
  mostRecentGem: Gem | null;
}

function findMostRecentGem(gemList: Gem[]): Gem | null {
  if (gemList.length === 0) {
    return null;
  }

  const reversedSorted = [...gemList]
    .map((gem) => {
      return {
        gem: gem,
        lastUpdated: gem.receivedAt || gem.createdAt,
      };
    })
    .sort((a, b) => b.lastUpdated.diff(a.lastUpdated));

  return reversedSorted[0].gem;
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
    const store = useStore();

    const gemLogs: GemLogs = store.state.gemLogs;

    const allFriends: FriendAndMostRecentGemActivity[] = [];

    for (const gemLog of gemLogs.gemLogsMap.values()) {
      const mostRecentGem = findMostRecentGem(gemLog.gems);

      allFriends.push({
        friend: gemLog.friend,
        mostRecentGem: mostRecentGem,
      });
    }

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
