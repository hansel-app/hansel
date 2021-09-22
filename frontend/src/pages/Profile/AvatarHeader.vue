<template>
  <div class="overlay">
    <van-row>
      <CircleAvatar
        :avatarUrl="placeholderAvatarUrl"
        radius="5"
        class="overlay"
      />
    </van-row>
    <van-row>
      <h2 class="overlay">{{ displayName }}</h2>
    </van-row>
    <h4 class="overlay">{{ "@" + username }}</h4>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import CircleAvatar from "@/components/CircleAvatar.vue";

export default defineComponent({
  props: {
    avatarUrl: String,
    displayName: String,
    username: String,
  },
  data() {
    return {
      // TODO: replace with user's own avatar
      placeholderAvatarUrl:
        "https://media.istockphoto.com/vectors/happy-young-woman-watching-into-rounded-frame-isolated-on-white-3d-vector-id1296058958?b=1&k=20&m=1296058958&s=170667a&w=0&h=6m2FU2hKv6emHjNtdNSqBJR1uMq64smptqwDAZNo6bg=",
    };
  },
  components: {
    CircleAvatar,
  },
});
</script>
