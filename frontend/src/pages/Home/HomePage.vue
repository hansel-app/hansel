<template>
  <Suspense>
    <template #default>
      <div>
      <button @click="goToProfile">Profile</button>
      <gem-map :gems="gems" />
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
import dayjs from "dayjs";

export default defineComponent({
  components: {
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
    }
  }
});
</script>
