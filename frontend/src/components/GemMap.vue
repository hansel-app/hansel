<template>
  <div>
    <h4>Hello Friendos</h4>
    Latitude: {{ currPosition.lat.toFixed(2) }}, Longitude:
    {{ currPosition.lng.toFixed(2) }}
  </div>

  <GoogleMap
    ref="mapRef"
    class="google-map"
    :api-key="apiKey"
    :center="currPosition"
    :fullscreen-control="mapConfig.fullscreenControl"
    :map-type-control-options="mapConfig.mapTypeControlOptions"
    :max-zoom="mapConfig.maxZoom"
    :min-zoom="mapConfig.minZoom"
    :street-view-control="mapConfig.streetViewControl"
    :styles="mapConfig.styles"
    :zoom="mapConfig.zoom"
    :zoom-control="mapConfig.zoomControl"
  >
    <template v-if="Boolean(mapRef)">
      <CustomControl class="center-map-control" position="TOP_RIGHT">
        <van-button
          round
          icon="aim"
          type="primary"
          @click="centerMapOnCurrentLocation"
        />
      </CustomControl>

      <Marker
        v-for="markerOptions in markersOptions"
        :key="markerOptions.id"
        :options="markerOptions"
      />
      <Marker :options="userMarkerOptions" />
    </template>
  </GoogleMap>
</template>

<script lang="ts">
import { ref, computed, defineComponent, PropType } from "vue";
import { useGeolocation } from "../useGeolocation";
import { GoogleMap, Marker, CustomControl } from "vue3-google-map";
import { DEFAULT_MAP_CONFIG } from "@/constants";
import { hansel } from "@/interfaces";

const GOOGLE_API_KEY = process.env.VUE_APP_GOOGLE_API_KEY;

// TODO: replace this once Gem interface has been created
interface TempGem {
  id: number;
  position: hansel.LatLng;
}

export default defineComponent({
  props: {
    gems: {
      type: Array as PropType<TempGem[]>,
      // TODO: replace this with empty array once gems can be fetched
      // from an actual data source
      default: () => [
        { id: 1, position: { lat: 1.2966, lng: 103.7764 } },
        { id: 2, position: { lat: 1.3483, lng: 103.6831 } },
        { id: 3, position: { lat: 1.3644, lng: 103.9915 } },
        { id: 4, position: { lat: 1.4382, lng: 103.7891 } },
        { id: 5, position: { lat: 1.3109, lng: 103.7952 } },
        { id: 6, position: { lat: 1.3309, lng: 103.8752 } },
      ],
    },
  },
  components: {
    GoogleMap,
    Marker,
    CustomControl,
  },

  async setup() {
    const mapRef = ref<InstanceType<typeof GoogleMap> | null>(null);

    const { coords, getLocation } = useGeolocation();
    const currPosition = computed(
      () =>
        ({
          lat: coords.value.lat,
          lng: coords.value.lng,
        } as hansel.LatLng)
    );
    const initPos = await getLocation();

    return {
      mapRef,
      currPosition,
      apiKey: GOOGLE_API_KEY,
      mapConfig: { ...DEFAULT_MAP_CONFIG, center: initPos },
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
    userMarkerOptions() {
      if (!this.mapRef?.ready) {
        return {};
      }

      const svgMarker: google.maps.Symbol = {
        path: google.maps.SymbolPath.CIRCLE,
        fillColor: "#3989fc",
        fillOpacity: 0.8,
        strokeWeight: 2,
        strokeColor: "#fff",
        scale: 12,
      };

      return {
        position: this.currPosition,
        icon: svgMarker,
      };
    },
  },

  methods: {
    centerMapOnCurrentLocation() {
      this.mapRef?.map?.panTo(this.currPosition);
    },
  },
});
</script>

<style scoped>
.google-map {
  width: 100%;
  height: 80vh;
}

.google-map .center-map-control {
  margin: 10px;
}
</style>
