<template>
  <Circle :options="circleOptions" />
  <Marker :options="userMarkerOptions" />
</template>

<script lang="ts">
import { defineComponent, PropType } from "vue";
import { Marker, Circle } from "vue3-google-map";
import { LatLng } from "@/interfaces";

export default defineComponent({
  props: {
    position: {
      type: Object as PropType<LatLng>,
      required: true,
    },
    circleRadius: {
      type: Number,
      required: true,
    },
  },
  components: {
    Marker,
    Circle,
  },

  computed: {
    userMarkerOptions() {
      const svgMarker: google.maps.Symbol = {
        path: google.maps.SymbolPath.CIRCLE,
        fillColor: "#3989fc",
        fillOpacity: 0.8,
        strokeWeight: 2,
        strokeColor: "#fff",
        scale: 12,
      };

      return {
        position: this.position,
        icon: svgMarker,
      };
    },
    circleOptions() {
      return {
        strokeColor: "#3989fc",
        strokeOpacity: 0.8,
        strokeWeight: 2,
        fillColor: "#3989fc",
        fillOpacity: 0.35,
        center: this.position,
        // convert km to m
        radius: this.circleRadius * 1000,
      };
    },
  },
});
</script>
