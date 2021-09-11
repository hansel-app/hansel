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
import { hansel } from "@/interfaces";

export default defineComponent({
  props: {
    currPosition: {
      type: Object as PropType<hansel.LatLng>,
    },
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

      // TODO: replace this with actual gem image / shape
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

  watch: {
    // TODO: replace this with a function that updates the
    // position of the marker representing the user
    // currPosition: function(newVal: hansel.LatLng) {
    // },
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
