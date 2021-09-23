<template>
  <div>
    <HamburgerMenu ref="hamburgerMenu" :requestCount="pendingFriendRequestCount"/>
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
          <Loading class="loading" size="36px">Loading map & stacking gems...</Loading>
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

    onMounted(() => {
      fetchGems();
      fetchFriendRequests();
    });
  },
  computed: {
    ...mapState({
      gems: (state: any) => state.gems.gemsPendingCollection as Gem[],
    }),
    pendingFriendRequestCount() {
      // return (state: any) => state.user.friendRequests.length;
      return 10;
    },
  },
});
</script>

<style scoped lang="less">
.loading-container {
  height: 100vh;
}
</style>
