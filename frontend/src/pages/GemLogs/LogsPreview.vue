<template>
  <FriendCell :friend="friend" :isClickable="true" labelOverride="test" />
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { User, Gem } from "@/interfaces";
import FriendCell from "@/components/FriendCell.vue";

export default defineComponent({
  props: {
    friend: {
      type: User,
      required: true,
    },
    mostRecentGemActivity: {
      type: Gem,
      required: true,
    },
  },
  components: {
    FriendCell,
  },
});
</script>
