<template>
  <van-row type="flex" align="center">
    <van-col span="6">
      <CircleAvatar
        :avatarUrl="dropperAvatar"
        :showLoading="false"
        :showError="false"
      />
    </van-col>
    <van-col class="marker-text-info" span="18">
      <p class="distance-indicator">{{ displayDistance }}</p>
      <p>From: {{ dropperName }}</p>
      <p>{{ displayDropDateTime }}</p>
    </van-col>
  </van-row>
</template>

<script lang="ts">
import { Dayjs } from "dayjs";
import { defineComponent } from "vue";
import { formatDistance } from "@/utils/geolocation";
import { formatDateTime } from "@/utils/date";
import { Row, Col } from "vant";
import CircleAvatar from "./CircleAvatar.vue";

export default defineComponent({
  components: {
    "van-row": Row,
    "van-col": Col,
    CircleAvatar,
  },
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
      type: String,
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

<style scoped lang="less">
.marker-text-info {
  padding-left: 2em;

  p {
    text-align: left;
  }

  .distance-indicator {
    font-weight: bold;
    font-size: large;
  }
}
</style>
