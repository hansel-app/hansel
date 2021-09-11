<template>
  <div>
    <div class="google-map" ref="googleMap"></div>
    <template v-if="Boolean(map)">
      <slot :map="map" />
    </template>
  </div>
</template>

<script lang="ts">
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
      userMarker: null as google.maps.Marker | null,
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

      this.markersOptions.forEach((options) => {
        new api.maps.Marker({
          ...options,
          map: this.map,
          // TODO: replace this with gem icon
          // icon:
        });
      });

      const svgMarker: google.maps.Symbol = {
        path: google.maps.SymbolPath.CIRCLE,
        fillColor: "#3989fc",
        fillOpacity: 0.8,
        strokeWeight: 2,
        strokeColor: "#fff",
        scale: 12,
      };

      this.userMarker = new api.maps.Marker({
        position: this.mapConfig.center,
        map: this.map,
        icon: svgMarker,
      });
    },
  },

  watch: {
    // TODO: replace this with a function that updates the
    // position of the marker representing the user
    currPosition: function(newVal: hansel.LatLng) {
      this.userMarker?.setPosition(newVal);
    },
  },
});
</script>

<style scoped>
.google-map {
  width: 100%;
  height: 80vh;
}
</style>
