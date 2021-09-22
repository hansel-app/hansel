<template>
  <van-row class="info-window" type="flex" align="center">
    <van-col class="avatar-container">
      <CircleAvatar
        :avatarUrl="gemCreator.avatar"
        :showLoading="false"
        :showError="false"
        :radius="1.5"
      />
    </van-col>

    <van-col class="marker-text-info">
      <p class="distance-indicator">{{ displayDistance }}</p>
      <p>From: {{ gemCreator.displayName }}</p>
      <p>{{ displayDropDateTime }}</p>
    </van-col>
  </van-row>
</template>

<script lang="ts">
import { Dayjs } from "dayjs";
import { defineComponent, PropType } from "vue";
import { formatDistance } from "@/utils/geolocation";
import { formatDateTime } from "@/utils/date";
import { Row, Col } from "vant";
import CircleAvatar from "./CircleAvatar.vue";
import { User } from "@/interfaces";

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
    gemCreator: {
      type: Object as PropType<User>,
      required: true,
    },
    dropTime: {
      type: Dayjs,
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
.info-window {
  margin-left: 0.5rem;
  margin-right: 0.5rem;
  align-items: center;
  justify-self: center;

  .marker-text-info {
    p {
      text-align: left;
    }

    .distance-indicator {
      font-weight: @font-weight-bold;
      font-size: @font-size-md;
    }
  }

  .avatar-container {
    margin-right: 0.5em;
  }
}
</style>
