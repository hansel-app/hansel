<template>
  <Cell center :title="friend.username" is-link>
    <template #icon>
      <!-- TODO: replace with user's own avatar -->
      <Image round :src="placeholderAvatarUrl" class="circle-avatar" />
    </template>
  </Cell>
</template>

<script>
import { defineComponent } from "vue";
import { Cell, Image } from "vant";
import { User } from "@/interfaces";

export default defineComponent({
  props: {
    friend: User,
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
    Image,
  },
});
</script>

<style scoped>
.circle-avatar {
  width: 4rem;
  height: 4rem;
}
</style>
