<template>
  <!-- TODO: make this into actual carousel -->
  <Row>
    <Button
      v-for="gem in mockGems"
      :key="gem.color"
      :color="gem.color.selectionColor"
    />
  </Row>
</template>

<script>
import { defineComponent } from "vue";
import { Button, Row } from "vant";
import { mockGems } from "@/interfaces/mockData";

export default defineComponent({
  data() {
    return {
      mockGems,
    };
  },
  components: {
    Button,
    Row,
  },
  inheritAttrs: false,
});
</script>
