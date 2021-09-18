<template>
  <div>
    <GemHeader title="Drop a gem" />
    <GemColorCarousel />
    <h4>Choose media</h4>
    <Row>
      <Button>Camera</Button>
      <Button>Photos</Button>
    </Row>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { Button, Row } from "vant";
import GemHeader from "@/components/GemHeader.vue";
import GemColorCarousel from "./GemColorCarousel.vue";

export default defineComponent({
  components: {
    Button,
    GemHeader,
    GemColorCarousel,
    Row,
  },
});
</script>
