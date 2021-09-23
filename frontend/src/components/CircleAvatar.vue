<template>
  <Image
    round
    :src="src"
    v-bind:style="dimensions"
    :show-loading="showLoading"
    :show-error="showError"
  />
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { Image } from "vant";

export default defineComponent({
  props: {
    avatarUrl: { type: String, default: () => "1" },
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
    src() {
      return base64AvatarToPlaceholderMapper(this.avatarUrl);
    },
  },
  components: {
    Image,
  },
});

enum DefaultAvatars {
  GIRL1 = "1",
  GIRL2 = "2",
  GUY1 = "3",
  GIRL3 = "4",
  GUY2 = "5",
}

// if User has not set an avatar, use one of our placeholders.
// Assumption is we reserve strings '1' to '5' for our placeholders haha
const base64AvatarToPlaceholderMapper = (base64: string): string => {
  if (!(Object as any).values(DefaultAvatars).includes(base64)) {
    return base64;
  }
  return require(`@/assets/avatars/avatar${base64}.png`);
};
</script>
