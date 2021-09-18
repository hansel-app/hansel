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
            class="profile-button"
          >
            Profile</CircleAvatar
          >
        </div>
      </div>
    </template>
    <template #fallback>
      <!-- TODO: add more useful loading / error message -->
      <div>Loading...</div>
    </template>
  </Suspense>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { PROFILE_ROUTE } from "@/constants";
import GemMap from "@/components/GemMap.vue";
import CircleAvatar from "@/components/CircleAvatar.vue";
import dayjs from "dayjs";

export default defineComponent({
  components: {
    CircleAvatar,
    GemMap,
  },

  setup() {
    // TODO: replace this with actual data source fetched from backend
    const gems = [
      {
        id: 1,
        position: { lat: 1.2966, lng: 103.7764 },
        dropper: "Bobby",
        dropTime: dayjs(),
      },
      {
        id: 2,
        position: { lat: 1.3483, lng: 103.6831 },
        dropper: "Bobby",
        dropTime: dayjs(),
      },
      {
        id: 3,
        position: { lat: 1.3644, lng: 103.9915 },
        dropper: "Bobby",
        dropTime: dayjs(),
      },
      {
        id: 4,
        position: { lat: 1.4382, lng: 103.7891 },
        dropper: "Bobby",
        dropTime: dayjs(),
      },
      {
        id: 5,
        position: { lat: 1.3109, lng: 103.7952 },
        dropper: "Bobby",
        dropTime: dayjs(),
      },
      {
        id: 6,
        position: { lat: 1.3309, lng: 103.8752 },
        dropper: "Bobby",
        dropTime: dayjs(),
      },
    ];

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
}
.container .profile-button {
  position: absolute;
  z-index: 10;
  top: 4px;
  left: 4px;
}
</style>
