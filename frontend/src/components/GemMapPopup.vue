<template>
  <van-cell>
    <template #value v-if="navigator && !navigator.onLine">
      <van-row>
        You are offline. Some functionalities may be unavailable.
      </van-row>
    </template>
    <template #value v-else-if="shouldShowTutorial">
      <van-row>
        Welcome to hansel. Try dropping your first gem!
      </van-row>
    </template>
    <template #value v-else-if="numberGemsPendingCollection > 0">
      <van-row>
        <van-col span="12">
          <van-row>
            <span class="number-label">
              {{ numberGemsPendingCollection }}
            </span>
          </van-row>
          <van-row>
            uncollected gem(s)
          </van-row>
        </van-col>

        <van-col span="12">
          <van-row>
            <span class="number-label">
              {{ displayNearestGemDistance }}
            </span>
          </van-row>
          <van-row>
            to the nearest gem
          </van-row>
        </van-col>
      </van-row>
    </template>

    <template #value v-else>
      <van-row>
        You currently have no uncollected gems~
      </van-row>
    </template>
  </van-cell>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { formatDistance } from "@/utils/geolocation";

export default defineComponent({
  props: {
    nearestGemDistance: {
      type: Number,
    },
    numberGemsPendingCollection: {
      type: Number,
      required: true,
    },
    shouldShowTutorial: {
      type: Boolean,
      required: true,
    },
  },
  computed: {
    displayNearestGemDistance() {
      if (!this.nearestGemDistance) {
        return null;
      }
      return formatDistance(this.nearestGemDistance);
    },
  },
});
</script>

<style scoped lang="less">
.number-label {
  font-size: @font-size-xl;
  font-weight: @font-weight-extra-bold;
}
.van-row {
  justify-content: center;
}
</style>
