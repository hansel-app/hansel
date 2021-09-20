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
    <template v-if="Boolean(mapRef) && mapRef?.ready">
      <CustomControl class="right-top-controls" position="RIGHT_TOP">
        <van-row>
          <van-icon
            class="circle-button-icon-md"
            :name="require('@/assets/icons/compass.svg')"
            @click="centerMapOnCurrentLocation"
          />
        </van-row>
        <van-row>
          <van-icon
            class="circle-button-icon-md"
            :name="require('@/assets/icons/info.svg')"
            @click="showPopup"
          />
        </van-row>
      </CustomControl>

      <CustomControl class="bottom-controls" position="BOTTOM_CENTER">
        <van-row justify="center">
          <van-icon
            class="circle-button-icon-lg"
            :name="require('@/assets/images/purple_64.png')"
            @click="goToDropGem"
          />
        </van-row>
        Drop a Gem

        <van-cell>
          <van-row v-if="gems.length > 0" justify="space-around" align="center">
            <van-col>
              <van-icon
                class="circle-button-icon-xs"
                :name="require('@/assets/icons/chevron-left-circle.svg')"
                @click="prevGem"
              />
            </van-col>

            <van-col span="12" class="label van-multi-ellipsis--l2 ">
              You have {{ gems.length }} gems waiting to be collected.
            </van-col>

            <van-col>
              <van-icon
                class="circle-button-icon-xs"
                :name="require('@/assets/icons/chevron-right-circle.svg')"
                @click="nextGem"
              />
            </van-col>
          </van-row>
          <van-row v-else justify="space-around" align="center">
            <van-col class="label van-multi-ellipsis--l2 ">
              Great work, you have collected all your gems!
            </van-col>
          </van-row>

          <van-row justify="space-around">
            <van-button type="primary" size="large" round @click="goToDropGem">
              Drop a Gem
            </van-button>
          </van-row>
        </van-cell>
      </CustomControl>

      <Marker
        :ref="setGemMarkerRef"
        v-for="markerOptions in gemMarkerOptions"
        @click="onGemMarkerClick(markerOptions)"
        :key="markerOptions.id"
        :options="markerOptions"
      />
      <MapUserMarker
        :position="currPosition"
        :circle-radius="userCircleRadius"
      />
    </template>
  </GoogleMap>

  <GemMapPopup
    class="popup"
    v-show="shouldShowPopup"
    :number-gems-pending-collection="gems.length"
    :nearest-gem-distance="nearestGemDistance"
  />
</template>

<script lang="ts">
import { createApp, ref, defineComponent, PropType, onBeforeUpdate } from "vue";
import { useGeolocation } from "../useGeolocation";
import { GoogleMap, Marker, CustomControl } from "vue3-google-map";
import {
  DEFAULT_MAP_CONFIG,
  DROP_GEM_ROUTE,
  GEM_PICKUP_RADIUS_THRESHOLD,
  PICKUP_GEM_ROUTE,
} from "@/constants";
import { Gem, GemColor } from "@/interfaces";
import { formatDistance, getDistanceFromLatLonInKm } from "@/utils/geolocation";
import GemMarkerInfoWindow from "./GemMarkerInfoWindow.vue";
import MapUserMarker from "./MapUserMarker.vue";
import GemMapPopup from "./GemMapPopup.vue";

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
    MapUserMarker,
    GemMapPopup,
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

    const { coords: currPosition, getLocation } = useGeolocation();
    const initPos = await getLocation();
    const shouldShowPopup = ref<boolean>(false);

    return {
      mapRef,
      gemMarkerRefs,
      gemMarkerInfoWindowRef,
      setGemMarkerRef,
      currPosition,
      userCircleRadius: GEM_PICKUP_RADIUS_THRESHOLD,
      currGemIdx: null as null | number,
      apiKey: GOOGLE_API_KEY,
      mapConfig: { ...DEFAULT_MAP_CONFIG, center: initPos },
      shouldShowPopup,
    };
  },

  computed: {
    nearestGemDistance() {
      if (this.sortedGems.length === 0) {
        return undefined;
      }
      const nearest = this.sortedGems[0];
      return getDistanceFromLatLonInKm(nearest.position, this.currPosition);
    },
    shouldShowPickupButton() {
      return (
        this.nearestGemDistance &&
        this.nearestGemDistance <= GEM_PICKUP_RADIUS_THRESHOLD
      );
    },
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
  },

  methods: {
    centerMapOnCurrentLocation() {
      this.mapRef?.map?.panTo(this.currPosition);
    },

    getGemImageUrl(gem: Gem) {
      console.assert(Object.values(GemColor).includes(gem.color));
      return require(`@/assets/images/${gem.color}_64.png`);
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
    goToPickupGem() {
      // TODO: pass gem to new route
      this.$router.push(PICKUP_GEM_ROUTE);
    },

    onGemMarkerClick(markerOptions: GemMarkerOptions) {
      this.showGemMarkerInfoWindow(markerOptions);
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
        gemCreator: gem.createdBy,
        dropTime: gem.createdAt,
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

    showPopup() {
      this.shouldShowPopup = true;
    },
  },
});
</script>

<style scoped lang="less">
.google-map {
  width: 100%;
  height: 100vh;

  .right-top-controls {
    margin-top: 2vh;
    margin-right: 2vw;
  }

  .bottom-controls {
    width: 90vw;
    margin-bottom: 5vh;

    .van-cell {
      margin-top: 1.5vh;
      border-radius: 2rem;
      box-shadow: @shadow-medium;
    }

    .label {
      text-align: center;
    }

    .van-row {
      margin: 0.75rem 0;
    }
  }
}

.popup {
  position: absolute;
  z-index: 10;
  width: 80vw;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  border-radius: 1rem;
  box-shadow: @shadow-medium;

  .van-row {
    justify-content: center;
  }
}
</style>
