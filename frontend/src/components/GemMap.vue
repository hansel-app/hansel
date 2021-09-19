<template>
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
        <van-button round icon="plus" type="primary" @click="goToDropGem" />
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
import {
  createApp,
  ref,
  computed,
  defineComponent,
  PropType,
  onBeforeUpdate,
} from "vue";
import { useGeolocation } from "../useGeolocation";
import { GoogleMap, Marker, CustomControl } from "vue3-google-map";
import {
  DEFAULT_MAP_CONFIG,
  DROP_GEM_ROUTE,
  GEM_PICKUP_RADIUS_THRESHOLD,
} from "@/constants";
import { Gem, LatLng } from "@/interfaces";
import { getDistanceFromLatLonInKm } from "@/utils/geolocation";
import GemMarkerInfoWindow from "./GemMarkerInfoWindow.vue";

const GOOGLE_API_KEY = process.env.VUE_APP_GOOGLE_API_KEY;

type GemMarkerOptions = Gem &
  Pick<google.maps.MarkerOptions, "map" | "position" | "icon">;

export default defineComponent({
  props: {
    gems: {
      type: Array as PropType<Gem[]>,
      default: () => [],
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
    const gemMarkerInfoWindowRef = ref<google.maps.InfoWindow | null>(null);

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
        } as LatLng)
    );

    const initPos = await getLocation();

    return {
      mapRef,
      gemMarkerRefs,
      gemMarkerInfoWindowRef,
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
        const gemImageUrl = this.getGemImageUrl(gem);
        return {
          ...gem,
          icon: gemImageUrl,
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

    getGemImageUrl(gem: Gem) {
      // TODO: replace with if-else block based on gem colour
      console.log(gem);
      return require("@/assets/blue_gem_64.png");
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
      this.showGemMarkerInfoWindow(currGem);
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
      this.showGemMarkerInfoWindow(currGem);
    },
    goToDropGem() {
      this.$router.push(DROP_GEM_ROUTE);
    },

    onGemMarkerClick(markerOptions: GemMarkerOptions) {
      const distFromSelf = getDistanceFromLatLonInKm(
        markerOptions.position,
        this.currPosition
      );

      if (distFromSelf >= GEM_PICKUP_RADIUS_THRESHOLD) {
        this.showGemMarkerInfoWindow(markerOptions);
      } else {
        // TODO: pick up gem
      }
    },

    showGemMarkerInfoWindow(gem: Gem) {
      const marker = Array.from(this.gemMarkerRefs).find(
        (marker) => marker.$props.options.position == gem.position
      );

      const distFromSelf = getDistanceFromLatLonInKm(
        gem.position,
        this.currPosition
      );

      const infoWindow = createApp(GemMarkerInfoWindow, {
        distance: distFromSelf,
        dropperName: gem.createdBy,
        dropTime: gem.createdAt,
        // TODO: replace this with actual avatar
        dropperAvatar:
          "https://cdn.mos.cms.futurecdn.net/JycrJzD5tvbGHWgjtPrRZY-970-80.jpg",
      });
      const el = document.createElement("div");
      const mountedApp = infoWindow.mount(el);
      const contentString = mountedApp.$el.outerHTML;

      if (!this.gemMarkerInfoWindowRef) {
        this.gemMarkerInfoWindowRef = new google.maps.InfoWindow();
      }
      // Close the existing window (if any)
      this.gemMarkerInfoWindowRef.close();
      this.gemMarkerInfoWindowRef.setContent(contentString);
      this.gemMarkerInfoWindowRef.open({
        anchor: marker?.marker.component.value,
        map: this.mapRef?.map,
        shouldFocus: false,
      });

      this.mapRef?.map?.addListener("click", () => {
        this.gemMarkerInfoWindowRef?.close();
      });
    },
  },
});
</script>

<style scoped lang="less">
.google-map {
  width: 100%;
  height: 100vh;
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
