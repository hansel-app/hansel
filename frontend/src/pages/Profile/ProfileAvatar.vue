<template>
  <CircleAvatar :avatarUrl="avatarUrl" radius="5" class="avatar-border" />
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
// TODO: abstract this out with friend-cell's gradient border maybe,
// but not sure how to abstract CSS
.avatar-border {
  position: relative;
  background: linear-gradient(to right, hsla(187, 89%, 52%, 0.6), #19d8f200),
    linear-gradient(to bottom, #c27df9, #c27df900),
    linear-gradient(to top, #e43d97, #e43d9700);
  background-clip: padding-box;
  padding: 8px;
  transform: translateX(-5%);
}
</style>
