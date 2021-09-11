import { SINGAPORE_CENTER } from "@/constants";
import { hansel } from "@/interfaces";
import { onMounted, onUnmounted, ref } from "vue";

export function useGeolocation() {
  const coords = ref<hansel.LatLng>(SINGAPORE_CENTER);

  // check if geolocation api is supported
  const isSupported = "navigator" in window && "geolocation" in navigator;

  let watcher: number | null = null;
  onMounted(() => {
    if (isSupported) {
      watcher = navigator.geolocation.watchPosition(
        (position) =>
          (coords.value = {
            lat: position.coords.latitude,
            lng: position.coords.longitude,
          })
      );
    }
  });
  onUnmounted(() => {
    if (watcher) navigator.geolocation.clearWatch(watcher);
  });
  return { coords, isSupported };
}
