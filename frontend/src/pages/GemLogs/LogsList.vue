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
import { defineComponent, reactive, onMounted } from "vue";
import Header from "@/components/Header.vue";
import Searchbar from "@/components/Searchbar.vue";
import BackSwipe from "@/components/BackSwipe.vue";
import { CellGroup } from "vant";
import { User, Gem, GemLogs } from "@/interfaces";
import { mapState, useStore } from "vuex";
import LogsPreview from "./LogsPreview.vue";

interface FriendAndMostRecentGemActivity {
  friend: User;
  mostRecentGem: Gem;
}

function findMostRecentGem(gemList: Gem[]): Gem {
  console.assert(
    gemList.length !== 0,
    "Friend must have at least one gem to appear in gem log"
  );

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

function parseGemLogs(gemLogs: GemLogs): FriendAndMostRecentGemActivity[] {
  const friends: FriendAndMostRecentGemActivity[] = [];

  for (const gemLog of gemLogs.gemLogsMap.values()) {
    const mostRecentGem = findMostRecentGem(gemLog.gems);

    friends.push({
      friend: gemLog.friend,
      mostRecentGem: mostRecentGem,
    });
  }
  return friends;
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
    const state = reactive({
      searchQuery: "",
      isCloseWindow: true,
      filteredFriendsAndActivities: [] as FriendAndMostRecentGemActivity[],
    });

    const store = useStore();
    const fetchGemsLogs = () => store.dispatch("getGemLogs");

    onMounted(() =>
      fetchGemsLogs().then(
        (gemLogs) =>
          (state.filteredFriendsAndActivities = parseGemLogs(gemLogs))
      )
    );

    return {
      state,
    };
  },
  computed: {
    ...mapState({
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      gemLogs: (state: any) => state.gems.gemLogs as GemLogs,
    }),
    allFriends(): FriendAndMostRecentGemActivity[] {
      return parseGemLogs(this.gemLogs);
    },
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
