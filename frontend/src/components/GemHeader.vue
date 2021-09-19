<template>
  <span class="big-header">{{ title }}</span>
  <Image :src="placeholderGemUrl" />
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { Image } from "vant";

export default defineComponent({
  props: {
    title: String,
  },
  data() {
    return {
      // TODO: replace with our own gem asset
      placeholderGemUrl:
        "https://images.squarespace-cdn.com/content/v1/5a773a9f8fd4d27532408347/1530563687072-GRIQPL3SZQDPZ7ZY7JEC/butt+lightened.jpg?format=1000w",
    };
  },
  components: { Image },
});
</script>

