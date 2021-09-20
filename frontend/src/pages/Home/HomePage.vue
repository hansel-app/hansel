<template>
  <Suspense>
    <template #default>
      <div>
        <HamburgerMenu ref="hamburgerMenu" />
        <div class="container">
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
</template>

<script lang="ts">
import { onMounted, defineComponent, inject, ref } from "vue";
import { PROFILE_ROUTE } from "@/constants";
import GemMap from "@/components/GemMap.vue";
import { FETCH_GEMS_PENDING_COLLECTION } from "@/providers/GemProvider.vue";
import { Gem } from "@/interfaces";
import { Loading } from "vant";
import HamburgerMenu from "@/components/HamburgerMenu.vue";

export default defineComponent({
  components: {
    GemMap,
    HamburgerMenu,
    Loading,
  },
  setup() {
    const fetchGems = inject(FETCH_GEMS_PENDING_COLLECTION, () =>
      Promise.resolve([])
    );
    const gems = ref<Gem[]>([]);

    onMounted(() => {
      // TODO: remove hardcoded userid
      fetchGems(2)
        .then((resp) => {
          gems.value = resp;
        })
        // TODO: better error handling
        .catch((err) => console.log(err, "Failed to fetch gems"));
    });

    return {
      gems,
    };
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
