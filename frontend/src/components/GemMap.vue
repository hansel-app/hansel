<template>
  <div>
    <h4>Hello Friendos</h4>
    Latitude: {{ currPosition.lat.toFixed(2) }}, Longitude:
    {{ currPosition.lng.toFixed(2) }}
  </div>
  <google-map-loader :mapConfig="mapConfig" apiKey="">
    <template v-slot="{ map }">
      {{ map }}
    </template>
  </google-map-loader>
</template>

<script lang="ts">
import { setup, Options, Vue } from "vue-class-component";
import { computed } from "vue";
import { useGeolocation } from "../useGeolocation";
import GoogleMapLoader from "./GoogleMapLoader.vue";

const DEFAULT_MAP_SETTINGS = {
  center: { lat: 1.3521, lng: 103.8198 },
  zoom: 13,
};

@Options({
  props: {},
  components: { GoogleMapLoader },
})
export default class GemMap extends Vue {
  currPosition = setup(() => {
    const { coords } = useGeolocation();
    const currPosition = computed(() => ({
      lat: coords.value.latitude,
      lng: coords.value.longitude,
    }));

    return currPosition;
  });

  get mapConfig() {
    return {
      ...DEFAULT_MAP_SETTINGS,
    };
  }
}
</script>

<style scoped></style>
