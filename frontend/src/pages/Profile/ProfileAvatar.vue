<template>
  <CircleAvatar :avatarUrl="avatarUrl" :radius="4" class="avatar-border" />
</template>

<script lang="ts">
import { defineComponent } from "vue";
import CircleAvatar from "@/components/CircleAvatar.vue";

export default defineComponent({
  props: {
    avatarUrl: String,
  },
  components: {
    CircleAvatar,
  },
});
</script>

<style scoped lang="less">
.avatar-border {
  position: relative;
  background: @profile-gradient;
  background-clip: padding-box;
  padding: 8px;
}
</style>
