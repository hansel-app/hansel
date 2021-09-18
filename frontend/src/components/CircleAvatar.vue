<template>
  <Image
    round
    :src="avatarUrl"
    :style="dimensions"
    :show-loading="showLoading"
    :show-error="showError"
  />
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { Image } from "vant";

export default defineComponent({
  props: {
    avatarUrl: String,
    radius: {
      type: Number,
      default: () => 2,
    },
    showLoading: {
      type: Boolean,
      default: () => true,
    },
    showError: {
      type: Boolean,
      default: () => true,
    },
  },
  computed: {
    dimensions() {
      return {
        width: this.radius * 2 + "rem",
        height: this.radius * 2 + "rem",
      };
    },
  },
  components: {
    Image,
  },
});
</script>
