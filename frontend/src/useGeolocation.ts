import { onMounted, onUnmounted } from "vue";
import { useStore } from "vuex";

export function useGeolocation() {
  const store = useStore();

  // check if geolocation api is supported
  const isSupported = "navigator" in window && "geolocation" in navigator;

  async function getLocation(): Promise<void> {
    return new Promise((resolve, reject) => {
      if (!isSupported) {
        reject(new Error("Geolocation is not supported."));
      }

      navigator.geolocation.getCurrentPosition(
        (pos) => {
          console.log("got current position");
          store.dispatch("updateCurrentPosition", {
            lat: pos.coords.latitude,
            lng: pos.coords.longitude,
          });
          resolve();
        },
        (err) => reject(err)
      );
    });
  }

  let watcher: number | null = null;

  onMounted(() => {
    if (isSupported) {
      watcher = navigator.geolocation.watchPosition((position) =>
        store.dispatch("updateCurrentPosition", {
          lat: position.coords.latitude,
          lng: position.coords.longitude,
        })
      );
    }
  });
  onUnmounted(() => {
    if (watcher) navigator.geolocation.clearWatch(watcher);
  });

  return { isSupported, getLocation };
}
