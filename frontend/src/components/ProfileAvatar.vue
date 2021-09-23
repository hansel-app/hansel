<template>
  <CircleAvatar :avatarUrl="avatarUrl" :radius="5" class="avatar-border" />
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
  height: 100;
  position: relative;
  background: @profile-gradient;
  background-clip: padding-box;
  padding: 6px;
  margin-top: -10%;
}
</style>
