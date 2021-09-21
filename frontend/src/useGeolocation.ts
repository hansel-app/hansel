import { SINGAPORE_CENTER } from "@/constants";
import { LatLng } from "@/interfaces";
import { onMounted, onUnmounted, ref } from "vue";
import { useStore } from "vuex";

export function useGeolocation() {
  const coords = ref<LatLng>(SINGAPORE_CENTER);

  // check if geolocation api is supported
  const isSupported = "navigator" in window && "geolocation" in navigator;

  async function getLocation(): Promise<LatLng> {
    return new Promise((resolve, reject) => {
      if (!isSupported) {
        reject(new Error("Geolocation is not supported."));
      }

      navigator.geolocation.getCurrentPosition(
        (pos) => {
          resolve({ lat: pos.coords.latitude, lng: pos.coords.longitude });
        },
        (err) => reject(err)
      );
    });
  }

  let watcher: number | null = null;

  onMounted(() => {
    const store = useStore();
    if (isSupported) {
      watcher = navigator.geolocation.watchPosition((position) =>
        store.dispatch("updateCurrentPosition", position.coords)
      );
    }
  });
  onUnmounted(() => {
    if (watcher) navigator.geolocation.clearWatch(watcher);
  });

  return { coords, isSupported, getLocation };
}
