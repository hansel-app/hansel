<template>
  <div>
    <h4>Hello Friendos</h4>
    Latitude: {{ currPosition.lat.toFixed(2) }}, Longitude:
    {{ currPosition.lng.toFixed(2) }}
  </div>

  <google-map-loader
    :currPosition="currPosition"
    :mapConfig="mapConfig"
    :apiKey="apiKey"
    :markersOptions="markersOptions"
  />
</template>

<script lang="ts">
import { computed, defineComponent, PropType } from "vue";
import { useGeolocation } from "../useGeolocation";
import GoogleMapLoader from "./GoogleMapLoader.vue";
import { DEFAULT_MAP_CONFIG } from "@/constants";
import { hansel } from "@/interfaces";

const GOOGLE_API_KEY = process.env.VUE_APP_GOOGLE_API_KEY;

// TODO: replaced this once Gem interface has been created
interface TempGem {
  position: hansel.LatLng;
}

export default defineComponent({
  props: {
    gems: {
      type: Array as PropType<TempGem[]>,
      // TODO: replace this with empty array once gems can be fetched
      // from an actual data source
      default: () => [
        { position: { lat: 1.2966, lng: 103.7764 } },
        { position: { lat: 1.3483, lng: 103.6831 } },
        { position: { lat: 1.3644, lng: 103.9915 } },
        { position: { lat: 1.4382, lng: 103.7891 } },
        { position: { lat: 1.3109, lng: 103.7952 } },
        { position: { lat: 1.3309, lng: 103.8752 } },
      ],
    },
  },
  components: {
    GoogleMapLoader,
  },

  setup() {
    const { coords } = useGeolocation();
    const currPosition = computed(
      () =>
        ({
          lat: coords.value.lat,
          lng: coords.value.lng,
        } as hansel.LatLng)
    );

    return {
      currPosition,
      apiKey: GOOGLE_API_KEY,
      mapConfig: { ...DEFAULT_MAP_CONFIG },
    };
  },

  computed: {
    markersOptions() {
      return this.gems.map((gem) => {
        return {
          ...gem,
        };
      });
    },
  },
});
</script>

<style scoped></style>
