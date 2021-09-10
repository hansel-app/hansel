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

interface MapConfig {
  center: {
    lat: number;
    lng: number;
  };
  zoom: number;
}

export default defineComponent({
  props: {
    mapConfig: {
      type: Object as PropType<MapConfig>,
      required: true,
    },
    apiKey: {
      type: String,
      default: "",
    },
  },

  data() {
    return {
      /* eslint-disable no-undef */
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
