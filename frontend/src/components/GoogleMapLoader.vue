<template>
  <div>
    <div class="google-map" ref="googleMap"></div>
    <template v-if="Boolean(map)">
      <slot :map="map" />
    </template>
  </div>
</template>

<script lang="ts">
/* eslint-disable no-undef */
import { defineComponent, PropType } from "vue";
import { Loader } from "@googlemaps/js-api-loader";

export default defineComponent({
  props: {
    mapConfig: {
      type: Object as PropType<google.maps.MapOptions>,
      required: true,
    },
    apiKey: {
      type: String,
      default: "",
    },
    markersOptions: {
      type: Array as PropType<google.maps.MarkerOptions[]>,
      default: () => [],
    },
  },

  data() {
    return {
      map: null as google.maps.Map | null,
    };
  },

  async mounted() {
    await this.initializeMap();
  },

  methods: {
    async initializeMap() {
      const mapContainer = this.$refs.googleMap as Element;
      const api = await new Loader({
        apiKey: this.apiKey,
      }).load();

      this.map = new api.maps.Map(mapContainer, this.mapConfig);

      const shape = {
        coords: [1, 1, 1, 20, 18, 20, 18, 1],
        type: "poly",
      };

      this.markersOptions.forEach((options) => {
        new api.maps.Marker({
          ...options,
          map: this.map,
          shape,
        });
      });
    },
  },
});
/* eslint-enable no-undef */
</script>

<style scoped>
.google-map {
  width: 100%;
  height: 80vh;
}
</style>
