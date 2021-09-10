<template>
  <div>
    <div class="google-map" ref="googleMap"></div>
    <template v-if="Boolean(map)">
      <slot :map="map" />
    </template>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { Loader } from "@googlemaps/js-api-loader";

export default defineComponent({
  props: {
    mapConfig: Object,
    apiKey: String,
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
        apiKey: this.apiKey ?? "",
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
