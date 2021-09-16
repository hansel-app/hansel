<template>
  <van-row type="flex" align="center">
    <van-col span="6">
      <van-image
        round
        fit="cover"
        src="https://cdn.mos.cms.futurecdn.net/JycrJzD5tvbGHWgjtPrRZY-970-80.jpg.webp"
        :show-loading="false"
        :show-error="false"
      />
    </van-col>
    <van-col span="18">
      <p>{{ displayDistance }}</p>
      <p>From {{ dropperName }}</p>
      <p>{{ displayDropDateTime }}</p>
    </van-col>
  </van-row>
</template>

<script lang="ts">
import { Dayjs } from "dayjs";
import { defineComponent } from "vue";
import { formatDistance } from "@/utils/geolocation";
import { formatDateTime } from "@/utils/date";
import { Row, Col, Image } from "vant";

export default defineComponent({
  components: { "van-image": Image, "van-row": Row, "van-col": Col },
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
