<template>
  <Cell
    center
    :title="friend.displayName"
    :label="label"
    :is-link="isClickable"
    class="content remove-padding"
    v-bind:class="purpleSelect"
  >
    <template #icon>
      <CircleAvatar
        :avatarUrl="friend.avatar"
        class="avatar-border"
        :radius="1.5"
      />
    </template>
    <template #right-icon>
      <slot />
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
    friend: {
      type: User,
      required: true,
    },
    isClickable: {
      type: Boolean,
      default: () => false,
    },
    shouldDisplayUsername: {
      type: Boolean,
      default: () => false,
    },
    labelOverride: {
      type: String,
    },
  },
  computed: {
    label() {
      if (this.labelOverride) {
        return this.labelOverride;
      }
      return this.shouldDisplayUsername ? "@" + this.friend.username : "";
    },
    purpleSelect() {
      return this.isClickable ? 'hover-change' : '';
    }
  },
  components: {
    Cell,
    CircleAvatar,
  },
});
</script>

<style scoped lang="less">
.avatar-border {
  position: relative;
  background: linear-gradient(to right, hsla(187, 89%, 52%, 0.6), #19d8f200),
    linear-gradient(to bottom, #c27df9, #c27df900),
    linear-gradient(to top, #e43d97, #e43d9700);
  background-clip: padding-box;
  padding: 4px;
  margin-right: 12px;
  margin-left: -4px;
  padding-left: -4px;
  transform: translateX(-5%);
}

.remove-padding {
  margin-left: -8px;
  width: 104%;
}

.van-cell {
  border-radius: 100em;
}

.hover-change:hover {
  background: @primary-color;
  color: white;
}
</style>
