<template>
  <AuthProvider>
    <GemProvider>
      <slot></slot>
    </GemProvider>
  </AuthProvider>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import AuthProvider from "@/providers/AuthProvider.vue";
import GemProvider from "@/providers/GemProvider.vue";

export default defineComponent({
  components: {
    AuthProvider,
    GemProvider,
  },
});
</script>
