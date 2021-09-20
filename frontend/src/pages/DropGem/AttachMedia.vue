<template>
  <div>
    <GemHeader title="Drop a gem" />
    <GemColorCarousel />
    <div class="container">
      <p class="sub-header">Choose media</p>
      <Row>
        <Button>Camera</Button>
        <Button>Photos</Button>
      </Row>
    </div>
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
