<template>
  <CircleAvatar :avatarUrl="placeholderAvatarUrl" radius="5" />
</template>

<script lang="ts">
import { defineComponent } from "vue";
import CircleAvatar from "@/components/CircleAvatar.vue";

export default defineComponent({
  components: {
    CircleAvatar,
  },
  data() {
    return {
      // TODO: replace with user's own avatar
      placeholderAvatarUrl:
        "https://media.istockphoto.com/vectors/happy-young-woman-watching-into-rounded-frame-isolated-on-white-3d-vector-id1296058958?b=1&k=20&m=1296058958&s=170667a&w=0&h=6m2FU2hKv6emHjNtdNSqBJR1uMq64smptqwDAZNo6bg=",
    };
  },
});
</script>
