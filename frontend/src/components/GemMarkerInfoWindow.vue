<template>
  <div>
    <p>{{ displayDistance }}</p>
    <p>From {{ dropperName }}</p>
    <p>{{ displayDropDateTime }}</p>
  </div>
</template>

<script lang="ts">
import { Dayjs } from "dayjs";
import { defineComponent } from "vue";

import { formatDistance } from "@/utils/geolocation";
import { formatDateTime } from "@/utils/date";

export default defineComponent({
  props: {
    distance: {
      type: Number,
      required: true,
    },
    dropperName: {
      type: String,
      required: true,
    },
    dropTime: {
      type: Dayjs,
      required: true,
    },
    dropperAvatar: {
      type: String, // TODO: replace this with actual avatar
      required: true,
    },
  },
  computed: {
    displayDistance() {
      return `${formatDistance(this.distance)} away`;
    },
    displayDropDateTime() {
      return `${formatDateTime(this.dropTime)}`;
    },
  },
});
</script>

<style scoped></style>
