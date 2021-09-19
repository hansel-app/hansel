<template>
  <Suspense>
    <template #default>
      <div class="container">
        <gem-map :gems="gems" />
        <div class="profile-button">
          <CircleAvatar
            :avatarUrl="placeholderAvatarUrl"
            :radius="3"
            @click="goToProfile"
          >
            Profile
          </CircleAvatar>
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
import { defineComponent, inject, ref } from "vue";
import { PROFILE_ROUTE } from "@/constants";
import GemMap from "@/components/GemMap.vue";
import CircleAvatar from "@/components/CircleAvatar.vue";
import { FETCH_GEMS_PENDING_COLLECTION } from "@/providers/GemProvider.vue";
import { Gem } from "@/interfaces";
import { Loading } from "vant";

export default defineComponent({
  components: {
    CircleAvatar,
    GemMap,
    Loading,
  },

  setup() {
    const fetchGems = inject(FETCH_GEMS_PENDING_COLLECTION, () =>
      Promise.resolve([])
    );
    const gems = ref<Gem[]>([]);
    // TODO: remove hardcoded userid
    fetchGems(2)
      .then((resp) => {
        gems.value = resp;
      })
      // TODO: better error handling
      .catch((err) => console.log(err, "Failed to fetch gems"));

    return {
      gems,
    };
  },
  methods: {
    goToProfile() {
      this.$router.push(PROFILE_ROUTE);
    },
  },
  data() {
    return {
      // TODO: replace with user's own avatar
      placeholderAvatarUrl:
        "https://media.istockphoto.com/vectors/happy-young-woman-watching-into-rounded-frame-isolated-on-white-3d-vector-id1296058958?b=1&k=20&m=1296058958&s=170667a&w=0&h=6m2FU2hKv6emHjNtdNSqBJR1uMq64smptqwDAZNo6bg=",
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
