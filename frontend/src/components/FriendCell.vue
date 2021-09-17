<template>
  <Cell center :title="friend.displayName" :is-link="isClickable">
    <template #icon>
      <!-- TODO: replace with user's own avatar -->
      <CircleAvatar :avatarUrl="placeholderAvatarUrl" />
    </template>
  </Cell>
</template>

<script>
import { defineComponent } from "vue";
import { Cell } from "vant";
import { User } from "@/interfaces";

import CircleAvatar from "./CircleAvatar.vue";

export default defineComponent({
  props: {
    friend: User,
    isClickable: {
      type: Boolean,
      default: () => false,
    }
  },
  data() {
    return {
      // TODO: replace with user's own avatar
      placeholderAvatarUrl:
        "https://cdn.mos.cms.futurecdn.net/JycrJzD5tvbGHWgjtPrRZY-970-80.jpg.webp",
    };
  },
  components: {
    Cell,
    CircleAvatar
  },
});
</script>
