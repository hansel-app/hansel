<template>
  <div>
    <van-image />
    <p>{{ message }}</p>
  </div>
</template>

<script lang="ts">
import { PropType, defineComponent } from "vue";
import { HexCode } from "@/interfaces";

export default defineComponent({
  props: {
    color: String as PropType<HexCode>,
    message: String,
  },
});
</script>
