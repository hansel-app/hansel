<template>
  <!-- TODO: make this into actual carousel -->
  <Row>
    <Button
      v-for="gem in GemColor"
      :key="gem"
      :color="getGemColorHex(gem)"
      @click="this.$emit('SetGemColorEvent', $event.target.value)"
    />
  </Row>
</template>

<script>
import { defineComponent } from "vue";
import { Button, Row } from "vant";
import { GemColor } from "@/interfaces";

export default defineComponent({
  data() {
    return {
      GemColor,
    };
  },
  methods: {
    getGemColorHex(color) {
      switch (color) {
        case GemColor.BLACK:
          return "#8d8d8d";
        case GemColor.PURPLE:
          return "#cdb4db";
        case GemColor.PINK:
          return "#ffc8dd";
        case GemColor.BLUE:
          return "#a2d2ff";
        case GemColor.YELLOW:
          return "#fffb8a";
        case GemColor.GREEN:
          return "#a6ffb6";
        default:
          return "#FFFFFF";
      }
    },
  },
  components: {
    Button,
    Row,
  },
  inheritAttrs: false,
});
</script>
