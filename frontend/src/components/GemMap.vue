<template>
  <div>
    <h4>Hello Friendos</h4>
    Latitude: {{ currPosition.lat.toFixed(2) }}, Longitude:
    {{ currPosition.lng.toFixed(2) }}
  </div>
  <router-link :to="{ name: 'dropGem' }">Drop Gem</router-link>
  <GoogleMap
    ref="mapRef"
    class="google-map"
    :api-key="apiKey"
    :center="currPosition"
    :clickable-icons="mapConfig.clickableIcons"
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
      <CustomControl class="top-right-controls" position="TOP_RIGHT">
        <van-button
          round
          icon="aim"
          type="primary"
          @click="centerMapOnCurrentLocation"
        />
      </CustomControl>
      <CustomControl
        v-if="gems.length > 0"
        class="bottom-controls"
        position="BOTTOM_CENTER"
      >
        <van-button round icon="arrow-left" type="primary" @click="prevGem" />
        <van-button
          round
          icon="aim"
          type="primary"
          @click="centerMapOnCurrentLocation"
        />
        <van-button round icon="arrow" type="primary" @click="nextGem" />
      </CustomControl>

      <Marker
        :ref="setGemMarkerRef"
        v-for="markerOptions in gemMarkerOptions"
        @click="onGemMarkerClick(markerOptions)"
        :key="markerOptions.id"
        :options="markerOptions"
      />
      <Marker :options="userMarkerOptions" />
    </template>
  </GoogleMap>
</template>

<script lang="ts">
import { ref, computed, defineComponent, PropType, onBeforeUpdate } from "vue";
import { useGeolocation } from "../useGeolocation";
import { GoogleMap, Marker, CustomControl } from "vue3-google-map";
import { DEFAULT_MAP_CONFIG } from "@/constants";
import { hansel } from "@/interfaces";
import { formatDistance, getDistanceFromLatLonInKm } from "@/utils/geolocation";
import dayjs, { Dayjs } from "dayjs";
import { formatDateTime } from "@/utils/date";

const GOOGLE_API_KEY = process.env.VUE_APP_GOOGLE_API_KEY;

// TODO: replace this once Gem interface has been created
interface TempGem {
  id: number;
  position: hansel.LatLng;
  dropTime: Dayjs;
  dropper: string;
}

type GemMarkerOptions = TempGem &
  Pick<google.maps.MarkerOptions, "map" | "position">;

export default defineComponent({
  props: {
    gems: {
      type: Array as PropType<TempGem[]>,
      // TODO: replace this with empty array once gems can be fetched
      // from an actual data source
      default: () => [
        {
          id: 1,
          position: { lat: 1.2966, lng: 103.7764 },
          dropper: "Bobby",
          dropTime: dayjs(),
        },
        {
          id: 2,
          position: { lat: 1.3483, lng: 103.6831 },
          dropper: "Bobby",
          dropTime: dayjs(),
        },
        {
          id: 3,
          position: { lat: 1.3644, lng: 103.9915 },
          dropper: "Bobby",
          dropTime: dayjs(),
        },
        {
          id: 4,
          position: { lat: 1.4382, lng: 103.7891 },
          dropper: "Bobby",
          dropTime: dayjs(),
        },
        {
          id: 5,
          position: { lat: 1.3109, lng: 103.7952 },
          dropper: "Bobby",
          dropTime: dayjs(),
        },
        {
          id: 6,
          position: { lat: 1.3309, lng: 103.8752 },
          dropper: "Bobby",
          dropTime: dayjs(),
        },
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
    let gemMarkerRefs: Set<InstanceType<typeof Marker>> = new Set();

    const setGemMarkerRef = (el: InstanceType<typeof Marker>) => {
      if (el) {
        gemMarkerRefs.add(el);
      }
    };

    onBeforeUpdate(() => {
      gemMarkerRefs = new Set();
    });

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
      gemMarkerRefs,
      setGemMarkerRef,
      currPosition,
      currGemIdx: null as null | number,
      apiKey: GOOGLE_API_KEY,
      mapConfig: { ...DEFAULT_MAP_CONFIG, center: initPos },
    };
  },

  computed: {
    sortedGems() {
      return [...this.gems].sort((gem1, gem2) => {
        return (
          getDistanceFromLatLonInKm(gem1.position, this.currPosition) -
          getDistanceFromLatLonInKm(gem2.position, this.currPosition)
        );
      });
    },
    gemMarkerOptions() {
      return this.gems.map((gem) => {
        return {
          ...gem,
        };
      }) as GemMarkerOptions[];
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

    nextGem() {
      console.assert(this.sortedGems.length > 0);

      if (this.currGemIdx === null) {
        this.currGemIdx = 0;
      } else {
        this.currGemIdx = (this.currGemIdx + 1) % this.sortedGems.length;
      }
      const currGem = this.sortedGems[this.currGemIdx];
      this.mapRef?.map?.panTo(currGem.position);
    },
    prevGem() {
      console.assert(this.sortedGems.length > 0);

      if (this.currGemIdx === null || this.currGemIdx === 0) {
        this.currGemIdx = this.sortedGems.length - 1;
      } else {
        this.currGemIdx -= 1;
      }
      const currGem = this.sortedGems[this.currGemIdx];
      this.mapRef?.map?.panTo(currGem.position);
    },

    onGemMarkerClick(markerOptions: GemMarkerOptions) {
      const marker = Array.from(this.gemMarkerRefs).find(
        (marker) => marker.$props.options.position == markerOptions.position
      );

      const contentString =
        '<div id="bodyContent">' +
        `<p>${formatDistance(
          getDistanceFromLatLonInKm(
            markerOptions.position as hansel.LatLng,
            this.currPosition
          )
        )} away</p>` +
        `<p>From ${markerOptions.dropper}</p>` +
        `<p>${formatDateTime(markerOptions.dropTime)}</p>`;

      const gemMarkerInfoWindow = new google.maps.InfoWindow({
        content: contentString,
      });
      gemMarkerInfoWindow.open({
        anchor: marker?.marker.component.value,
        map: this.mapRef?.map,
        shouldFocus: false,
      });
    },
  },
});
</script>

<style scoped>
.google-map {
  width: 100%;
  height: 80vh;
}
.google-map .top-right-controls {
  margin: 10px;
}

.google-map .bottom-controls {
  width: 100%;
  display: flex;
  justify-content: space-around;
  margin-bottom: 20px;
}
</style>
