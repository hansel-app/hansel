<template>
  <div>
    <h4>Hello Friendos</h4>
    Latitude: {{ currPosition.lat.toFixed(2) }}, Longitude:
    {{ currPosition.lng.toFixed(2) }}
  </div>

  <google-map-loader :mapConfig="mapConfig" :apiKey="apiKey">
    <template v-slot="{ map }">
      {{ map }}
    </template>
  </google-map-loader>
</template>

<script lang="ts">
import { computed, defineComponent } from "vue";
import { useGeolocation } from "../useGeolocation";
import GoogleMapLoader from "./GoogleMapLoader.vue";

const DEFAULT_MAP_CONFIG = {
  center: { lat: 1.3521, lng: 103.8198 },
  zoom: 13,
};

const GOOGLE_API_KEY = process.env.VUE_APP_GOOGLE_API_KEY;

export default defineComponent({
  components: {
    GoogleMapLoader,
  },

  setup() {
    const { coords } = useGeolocation();
    const currPosition = computed(() => ({
      lat: coords.value.latitude,
      lng: coords.value.longitude,
    }));

    return { currPosition };
  },

  computed: {
    mapConfig() {
      return { ...DEFAULT_MAP_CONFIG };
    },
    apiKey() {
      return GOOGLE_API_KEY;
    },
  },
});
</script>

<style scoped></style>
