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
            class="circle-button-icon-md filter-to-primary-color"
            :name="require('@/assets/icons/compass.svg')"
            @click="centerMapOnCurrentLocation"
          />
        </van-row>
        <van-row>
          <van-icon
            class="circle-button-icon-md filter-to-primary-color"
            :name="require('@/assets/icons/info.svg')"
            @click="togglePopup"
          />
        </van-row>
      </CustomControl>

      <CustomControl class="bottom-controls" position="BOTTOM_CENTER">
        <van-row justify="end" v-if="shouldShowPickupGemButton">
          <van-button
            class="pick-up-button"
            round
            plain
            type="primary"
            icon-prefix="pick-up-button-icon"
            :icon="require('@/assets/icons/chevron-right-circle.svg')"
            icon-position="right"
            @click="pickUpGem"
          >
            Pick up
          </van-button>
        </van-row>

        <van-row justify="space-around" align="center">
          <van-col v-if="gems.length > 0">
            <van-icon
              class="circle-button-icon-md prev-gem filter-to-primary-color"
              :name="require('@/assets/icons/chevron-left-circle.svg')"
              color="@primary-color"
              @click="() => prevGem()"
            />
          </van-col>

          <van-col>
            <van-icon
              class="circle-button-icon-lg"
              :name="require('@/assets/images/drop-gem-button.png')"
              @click="goToDropGem"
              size="100"
            />
          </van-col>

          <van-col v-if="gems.length > 0">
            <van-icon
              class="circle-button-icon-md next-gem filter-to-primary-color"
              :name="require('@/assets/icons/chevron-right-circle.svg')"
              @click="() => nextGem()"
            />
          </van-col>
        </van-row>
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
    class="map-popup"
    v-show="shouldShowPopup"
    v-click-away="hidePopup"
    :number-gems-pending-collection="gems.length"
    :nearest-gem-distance="nearestGemDistance"
    :should-show-tutorial="shouldShowTutorial"
  />
</template>

<script lang="ts">
import { createApp, ref, defineComponent, PropType, watchEffect } from "vue";
import { useGeolocation } from "@/useGeolocation";
import { GoogleMap, Marker, CustomControl } from "vue3-google-map";
import {
  DEFAULT_MAP_CONFIG,
  DROP_GEM_ROUTE,
  GEM_PICKUP_RADIUS_THRESHOLD,
  PICKUP_GEM_ROUTE,
} from "@/constants";
import { Gem, GemColor, LatLng } from "@/interfaces";
import { getDistanceFromLatLonInKm } from "@/utils/geolocation";
import GemMarkerInfoWindow from "./GemMarkerInfoWindow.vue";
import MapUserMarker from "./MapUserMarker.vue";
import GemMapPopup from "./GemMapPopup.vue";
import { mapState, useStore } from "vuex";
import { getEnumKeyByEnumValue } from "@/utils/enum";
import { Toast } from "vant";

const GOOGLE_API_KEY = window.env.VUE_APP_GOOGLE_API_KEY;

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
    const gemMarkerInfoWindowRef = ref<google.maps.InfoWindow | null>(null);
    const openedInfoWindowGem = ref<Gem | null>(null);
    const store = useStore();
    const isGemMapFirstMount = store.state.gems.isGemMapFirstMount;

    const shouldShowPopup = ref<boolean>(isGemMapFirstMount);

    const { getLocation } = useGeolocation();

    watchEffect(
      () => {
        const popupCloseEvents = ["mousedown", "dragstart"];
        popupCloseEvents.forEach((event) => {
          mapRef?.value?.map?.addListener(event, () => {
            shouldShowPopup.value = false;
          });
        });
      },
      {
        flush: "post",
      }
    );

    await getLocation();

    store.commit("setIsGemMapFirstMount", false);
    return {
      mapRef,
      gemMarkerInfoWindowRef,
      userCircleRadius: GEM_PICKUP_RADIUS_THRESHOLD,
      currGemIdx: null as null | number,
      apiKey: GOOGLE_API_KEY,
      shouldShowPopup,
      openedInfoWindowGem,
      store,
    };
  },

  beforeUpdate() {
    this.gemMarkerRefs = new Set();
  },
  data() {
    return {
      gemMarkerRefs: new Set() as Set<InstanceType<typeof Marker>>,
    };
  },

  computed: {
    ...mapState({
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      currPosition: (state: any) => state.user.currPosition as LatLng,
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      shouldShowTutorial: (state: any) =>
        state.user.shouldShowTutorial as boolean,
    }),
    mapConfig() {
      return {
        ...DEFAULT_MAP_CONFIG,
        currPosition: this.currPosition,
      };
    },
    nearestGemDistance() {
      if (this.sortedGems.length === 0) {
        return undefined;
      }
      const nearest = this.sortedGems[0];
      return getDistanceFromLatLonInKm(nearest.position, this.currPosition);
    },
    sortedGems(): Gem[] {
      return [...this.gems].sort((gem1, gem2) => {
        return (
          getDistanceFromLatLonInKm(gem1.position, this.currPosition) -
          getDistanceFromLatLonInKm(gem2.position, this.currPosition)
        );
      });
    },
    gemMarkerOptions(): GemMarkerOptions[] {
      return this.gems.map((gem) => {
        const gemImageUrl = this.getGemImageUrl(gem);
        return {
          ...gem,
          icon: gemImageUrl,
        };
      }) as GemMarkerOptions[];
    },
    shouldShowPickupGemButton(): boolean {
      if (this.openedInfoWindowGem === null) {
        return false;
      }
      const distToSelf = getDistanceFromLatLonInKm(
        this.openedInfoWindowGem.position,
        this.currPosition
      );

      return distToSelf <= GEM_PICKUP_RADIUS_THRESHOLD;
    },
  },

  methods: {
    setGemMarkerRef(el: InstanceType<typeof Marker>) {
      if (el) {
        this.gemMarkerRefs.add(el);
      }
    },
    centerMapOnCurrentLocation() {
      this.mapRef?.map?.panTo(this.currPosition);
      Toast.loading({
        duration: 1000,
        forbidClick: true,
        loadingType: "spinner",
        message: "Centering map to your location",
      });
    },

    getGemImageUrl(gem: Gem) {
      const gemColorName = getEnumKeyByEnumValue(GemColor, gem.color);
      return require(`@/assets/images/${gemColorName.toLowerCase()}_64.png`);
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

      const gemInfoWindowCloseEvents = ["click", "mousedown"];
      gemInfoWindowCloseEvents.forEach((event) => {
        this.mapRef?.map?.addListener(event, () => {
          this.gemMarkerInfoWindowRef?.close();
          this.openedInfoWindowGem = null;
        });
      });

      this.openedInfoWindowGem = gem;
    },

    togglePopup() {
      this.shouldShowPopup = !this.shouldShowPopup;
    },
    hidePopup() {
      this.shouldShowPopup = false;
      this.store.commit("setShouldShowTutorial", false);
    },

    pickUpGem() {
      if (!this.openedInfoWindowGem) {
        console.error("No gem selected for picking up");
        return;
      }
      this.store
        .dispatch("pickUpGem", { gem: this.openedInfoWindowGem })
        .then(() => this.$router.push(PICKUP_GEM_ROUTE))
        .catch((err) => Toast.fail(`Failed to pick up gem: ${err}`));
    },
  },
});
</script>

<style lang="less">
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

  .pick-up-button {
    margin-bottom: 8px;
  }

  .pick-up-button-icon {
    display: flex;
  }
}

.map-popup {
  position: absolute;
  z-index: 10;
  width: 90vw;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  border-radius: 1rem;
  box-shadow: @shadow-medium;
  background: @white !important;
}

// HSV of og blue color (226, 100, 45)
// HSV of desired primary color purple (252, 57, 68)
.filter-to-primary-color {
  filter: hue-rotate(26deg);
}
</style>
