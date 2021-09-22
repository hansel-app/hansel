<template>
  <div>
    <Suspense>
      <template #default>
        <div>
          <HamburgerMenu ref="hamburgerMenu" />
          <div>
            <gem-map :gems="gems" />
            <div class="profile-button"></div>
          </div>
        </div>
      </template>
      <template #fallback>
        <!-- TODO: add more useful loading / error message -->
        <Loading size="24px">Loading...</Loading>
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
.container {
  position: relative;

  .profile-button {
    position: absolute;
    z-index: 10;
    top: 4px;
    left: 4px;
  }
}
</style>
