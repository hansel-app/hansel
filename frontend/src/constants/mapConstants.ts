import { hansel } from "@/interfaces";

export const SINGAPORE_CENTER: hansel.LatLng = Object.freeze({
  lat: 1.3521,
  lng: 103.8198,
});

// copy-pasted from https://snazzymaps.com/style/124771/google-maps-clean
const mapStyles: google.maps.MapTypeStyle[] = [
  {
    featureType: "administrative",
    elementType: "labels",
    stylers: [
      {
        visibility: "off",
      },
    ],
  },
  {
    featureType: "landscape",
    elementType: "all",
    stylers: [
      {
        visibility: "on",
      },
    ],
  },
  {
    featureType: "poi.attraction",
    elementType: "labels",
    stylers: [
      {
        visibility: "on",
      },
    ],
  },
  {
    featureType: "poi.business",
    elementType: "all",
    stylers: [
      {
        visibility: "on",
      },
    ],
  },
  {
    featureType: "poi.business",
    elementType: "labels",
    stylers: [
      {
        visibility: "on",
      },
    ],
  },
  {
    featureType: "poi.business",
    elementType: "labels.icon",
    stylers: [
      {
        visibility: "off",
      },
    ],
  },
  {
    featureType: "poi.government",
    elementType: "labels",
    stylers: [
      {
        visibility: "on",
      },
    ],
  },
  {
    featureType: "poi.school",
    elementType: "all",
    stylers: [
      {
        visibility: "on",
      },
    ],
  },
  {
    featureType: "poi.school",
    elementType: "labels",
    stylers: [
      {
        visibility: "off",
      },
    ],
  },
  {
    featureType: "road",
    elementType: "all",
    stylers: [
      {
        visibility: "on",
      },
    ],
  },
  {
    featureType: "road",
    elementType: "labels",
    stylers: [
      {
        visibility: "off",
      },
    ],
  },
];

// Note: if adding new fields to this object, do remember to
// pass it as a parameter to the GoogleMap component.
// (I don't know why the library doesn't want to just take in a config object
// which is what the original Google API does...)
export const DEFAULT_MAP_CONFIG = Object.freeze({
  center: SINGAPORE_CENTER,
  fullscreenControl: false,
  mapTypeControlOptions: {
    mapTypeIds: ["ROADMAP"],
  },
  maxZoom: 20,
  minZoom: 16,
  streetViewControl: false,
  styles: mapStyles,
  zoom: 18,
  zoomControl: false,
  clickableIcons: false,
});
