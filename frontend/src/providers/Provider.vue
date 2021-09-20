<template>
  <GemProvider>
    <slot></slot>
  </GemProvider>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import GemProvider from "@/providers/GemProvider.vue";

export default defineComponent({
  components: {
    GemProvider,
  },
});
</script>
