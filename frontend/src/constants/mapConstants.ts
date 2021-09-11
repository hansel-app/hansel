import { hansel } from "@/interfaces";

export const SINGAPORE_CENTER: hansel.LatLng = Object.freeze({
  lat: 1.3521,
  lng: 103.8198,
});

export const DEFAULT_MAP_CONFIG: google.maps.MapOptions = Object.freeze({
  mapTypeControlOptions: {
    mapTypeIds: ["ROADMAP"],
  },
  fullscreenControl: false,
  streetViewControl: false,
  zoomControl: false,
  center: SINGAPORE_CENTER,
  zoom: 18,
});
