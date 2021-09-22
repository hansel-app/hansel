<template>
  <Suspense>
    <template #default>
      <div>
        <HamburgerMenu ref="hamburgerMenu" />
        <div>
          <gem-map :gems="gems" />
        </div>
      </div>
    </template>
    <template #fallback>
      <van-row class="loading-container" justify="center" align="center">
        <Loading class="loading" size="36px">Loading...</Loading>
      </van-row>
    </template>
  </Suspense>
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

    onMounted(() => {
      fetchGems();
    });
  },
  computed: {
    ...mapState({
      gems: (state: any) => state.gems.gemsPendingCollection as Gem[],
    }),
  },
});
</script>

<style scoped lang="less">
.loading-container {
  height: 100vh;
}
</style>
