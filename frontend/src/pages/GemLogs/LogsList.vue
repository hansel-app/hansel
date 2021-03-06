<template>
  <div class="background-gradient">
    <BackSwipe />
    <Header title="Gem logs" />
    <div class="container">
      <Searchbar v-model="searchQuery" placeholder="Search a friend" @input="filterFriends" />
      <div v-if="state.filteredFriendsAndActivities.length > 0">
        <CellGroup>
          <LogsPreview
            v-for="friendAndActivity in state.filteredFriendsAndActivities"
            :key="friendAndActivity.friend.userId"
            :friend="friendAndActivity.friend"
            :mostRecentGemActivity="friendAndActivity.mostRecentGem"
            :gemLogsWithFriend="
              gemLogs.gemLogsMap.get(friendAndActivity.friend.userId)
            "
            :isClickable="true"
          />
        </CellGroup>
      </div>
      <div v-else>
          <i>No results found</i> &#x1f4a9;
      </div>
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
  data() {
    return {
      searchQuery: "",
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
    filterFriends() {
      if (this.searchQuery.length === 0) {
        this.state.filteredFriendsAndActivities = this.allFriends;
        return;
      }

      this.state.filteredFriendsAndActivities = this.allFriends.filter(
        (friendAndGemActivity: FriendAndMostRecentGemActivity) => {
          const friend = friendAndGemActivity.friend;

          const matchUsername = friend.username
            .toLowerCase()
            .includes(this.searchQuery.toLowerCase());
          const matchDisplayname = friend.displayName
            .toLowerCase()
            .includes(this.searchQuery.toLowerCase());
          return matchUsername || matchDisplayname;
        }
      );
    },
  },
});
</script>
