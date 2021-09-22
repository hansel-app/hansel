<template>
  <div class="info-window">
    <van-row align="center">
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
    <van-row justify="end">
      <van-button
        v-if="shouldShowPickupButton"
        round
        class="pick-up-button"
        type="primary"
        icon-prefix="pick-up-button-icon"
        :icon="require('@/assets/icons/chevron-right-circle-white.svg')"
        icon-position="right"
      >
        Pick up
      </van-button>
    </van-row>
  </div>
</template>

<script lang="ts">
import { Dayjs } from "dayjs";
import { defineComponent, PropType } from "vue";
import { formatDistance } from "@/utils/geolocation";
import { formatDateTime } from "@/utils/date";
import { Row, Col, Button } from "vant";
import CircleAvatar from "./CircleAvatar.vue";
import { User } from "@/interfaces";

export default defineComponent({
  components: {
    "van-row": Row,
    "van-col": Col,
    "van-button": Button,
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
    shouldShowPickupButton: {
      type: Boolean,
      default: () => false,
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

<style lang="less">
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

  .pick-up-button {
    height: 30px;
    margin-bottom: 8px;
  }

  .pick-up-button-icon {
    display: flex;
  }
}
</style>
