<template>
  <div>
    <HamburgerMenu
      ref="hamburgerMenu"
      :requestCount="pendingFriendRequestCount"
    />
    <Suspense>
      <template #default>
        <div>
          <div>
            <gem-map :gems="gems" />
          </div>
        </div>
      </template>
      <template #fallback>
        <van-row class="loading-container" justify="center" align="center">
          <Loading class="loading" size="36px"
            >Loading map & stacking gems...</Loading
          >
        </van-row>
      </template>
    </Suspense>
  </div>
</template>

<script lang="ts">
import { onMounted, defineComponent } from "vue";
import GemMap from "@/components/GemMap.vue";
import { Gem } from "@/interfaces";
import { Loading } from "vant";
import HamburgerMenu from "@/components/HamburgerMenu.vue";
import { mapState, useStore } from "vuex";
import { PendingFriendRequest } from "@/protobuf/user_pb";
import { mockFriends, mockGems, mockSelfUser } from "@/interfaces/mockData.ts"

export default defineComponent({
  components: {
    GemMap,
    HamburgerMenu,
    Loading,
  },
  setup() {
    const store = useStore();
    const fetchGems = () => store.dispatch("getGemsPendingCollectionForUser");
    const fetchFriendRequests = () => store.dispatch("getFriendRequests");

    store.commit("setSelfInfo", mockSelfUser);
    store.commit("setFriends", mockFriends);
    store.commit("setGemsPendingCollection", mockGems);

    onMounted(() => {
      fetchGems();
      fetchFriendRequests();
    });
  },
  computed: {
    ...mapState({
      gems: (state: any) => state.gems.gemsPendingCollection as Gem[],
      friendRequests: (state: any) =>
        state.user.friendRequests as PendingFriendRequest[],
    }),
    pendingFriendRequestCount(): number {
      return this.friendRequests.length;
    },
  },
});
</script>

<style scoped lang="less">
.loading-container {
  height: 100vh;
}
</style>
