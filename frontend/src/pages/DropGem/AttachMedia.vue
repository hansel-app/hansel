<template>
  <div>
    <GemHeader title="Drop a gem" :color="color" />
    <GemColorCarousel @selected-color-changed="onSelectedColorChanged" />
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
import { GemColor } from "@/interfaces";
import { mapState } from "vuex";

export default defineComponent({
  components: {
    Button,
    GemHeader,
    GemColorCarousel,
    Row,
  },
  computed: {
    ...mapState({
      color: (state: any) => state.gems.dropGemFormState.color as GemColor,
    }),
  },
  methods: {
    onSelectedColorChanged(selectedColor: GemColor) {
      this.$emit("set-gem-color-event", selectedColor);
    },
  },
});
</script>
