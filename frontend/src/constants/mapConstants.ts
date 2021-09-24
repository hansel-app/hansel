import { LatLng } from "@/interfaces";

// Distance in km
export const GEM_PICKUP_RADIUS_THRESHOLD = 0.05;

export const SINGAPORE_CENTER: LatLng = Object.freeze({
  lat: 1.3521,
  lng: 103.8198,
});

// copy-pasted from https://snazzymaps.com/style/61/blue-essence
const mapStyles: google.maps.MapTypeStyle[] = [
  {
    featureType: "landscape.natural",
    elementType: "geometry.fill",
    stylers: [
      {
        color: "#e0efef",
      },
      {
        visibility: "on",
      },
    ],
  },
  {
    featureType: "poi",
    elementType: "geometry.fill",
    stylers: [
      {
        color: "#c0e8e8",
      },
      {
        hue: "#1900ff",
      },
      {
        visibility: "on",
      },
    ],
  },
  {
    featureType: "poi.business",
    stylers: [
      {
        visibility: "off",
      },
    ],
  },
  {
    featureType: "poi.park",
    elementType: "labels.text",
    stylers: [
      {
        visibility: "off",
      },
    ],
  },
  {
    featureType: "road",
    elementType: "geometry",
    stylers: [
      {
        lightness: 100,
      },
      {
        visibility: "simplified",
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
  {
    featureType: "road.arterial",
    elementType: "labels",
    stylers: [
      {
        visibility: "off",
      },
    ],
  },
  {
    featureType: "road.highway",
    elementType: "labels",
    stylers: [
      {
        visibility: "off",
      },
    ],
  },
  {
    featureType: "road.local",
    stylers: [
      {
        visibility: "off",
      },
    ],
  },
  {
    featureType: "transit.line",
    elementType: "geometry",
    stylers: [
      {
        lightness: 700,
      },
      {
        visibility: "on",
      },
    ],
  },
  {
    featureType: "water",
    stylers: [
      {
        color: "#7dcdcd",
      },
    ],
  },
];

// Note: when adding new fields to this object, do remember to
// pass it as a parameter to the GoogleMap component.
// (I don't know why the vue3-google-map library doesn't want to just take
// in a config object which is what the original Google API does...)
export const DEFAULT_MAP_CONFIG = Object.freeze({
  center: SINGAPORE_CENTER,
  fullscreenControl: false,
  mapTypeControlOptions: {
    mapTypeIds: ["ROADMAP"],
  },
  maxZoom: 20,
  minZoom: 12,
  streetViewControl: false,
  styles: mapStyles,
  zoom: 18,
  zoomControl: false,
  clickableIcons: false,
});
