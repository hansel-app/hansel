<template>
  <van-cell>
    <template v-if="numberGemsPendingCollection > 0">
      <van-row>
        <van-col span="12">
          <van-row>
            {{ numberGemsPendingCollection }}
          </van-row>
          <van-row>
            uncollected gems
          </van-row>
        </van-col>

        <van-col span="12">
          <van-row>
            {{ displayNearestGemDistance }}
          </van-row>
          <van-row>
            to the nearest gem
          </van-row>
        </van-col>
      </van-row>
    </template>

    <template v-else>
      <van-row>
        Great work, you have collected all your gems!
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
