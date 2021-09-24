import { LatLng } from "@/interfaces";

// Distance in km
export const GEM_PICKUP_RADIUS_THRESHOLD = 0.05;

export const SINGAPORE_CENTER: LatLng = Object.freeze({
  lat: 1.3521,
  lng: 103.8198,
});

// copy-pasted from https://snazzymaps.com/style/104349/pastel-cold
const mapStyles: google.maps.MapTypeStyle[] = [
  {
    featureType: "administrative",
    elementType: "labels.text.fill",
    stylers: [
      {
        color: "#53608f",
      },
    ],
  },
  {
    featureType: "administrative.country",
    elementType: "geometry.stroke",
    stylers: [
      {
        saturation: "22",
      },
      {
        lightness: "7",
      },
    ],
  },
  {
    featureType: "administrative.neighborhood",
    elementType: "labels.text.fill",
    stylers: [
      {
        color: "#53608f",
      },
      {
        lightness: "24",
      },
    ],
  },
  {
    featureType: "landscape",
    elementType: "labels.text.fill",
    stylers: [
      {
        color: "#53608f",
      },
    ],
  },
  {
    featureType: "landscape.man_made",
    elementType: "geometry.fill",
    stylers: [
      {
        color: "#e9ecf5",
      },
    ],
  },
  {
    featureType: "landscape.man_made",
    elementType: "geometry.stroke",
    stylers: [
      {
        color: "#edeff3",
      },
      {
        lightness: "-6",
      },
    ],
  },
  {
    featureType: "poi",
    elementType: "all",
    stylers: [
      {
        saturation: "32",
      },
      {
        color: "#e3e5ed",
      },
    ],
  },
  {
    featureType: "poi",
    elementType: "labels.text.stroke",
    stylers: [
      {
        color: "#ffffff",
      },
    ],
  },
  {
    featureType: "poi.attraction",
    elementType: "all",
    stylers: [
      {
        visibility: "off",
      },
    ],
  },
  {
    featureType: "poi.business",
    elementType: "all",
    stylers: [
      {
        visibility: "off",
      },
    ],
  },
  {
    featureType: "poi.government",
    elementType: "all",
    stylers: [
      {
        visibility: "off",
      },
    ],
  },
  {
    featureType: "poi.medical",
    elementType: "all",
    stylers: [
      {
        visibility: "off",
      },
    ],
  },
  {
    featureType: "poi.park",
    elementType: "geometry.fill",
    stylers: [
      {
        color: "#c0deab",
      },
    ],
  },
  {
    featureType: "poi.park",
    elementType: "labels.text.fill",
    stylers: [
      {
        color: "#919ab7",
      },
    ],
  },
  {
    featureType: "poi.place_of_worship",
    elementType: "all",
    stylers: [
      {
        visibility: "off",
      },
    ],
  },
  {
    featureType: "poi.school",
    elementType: "all",
    stylers: [
      {
        visibility: "off",
      },
    ],
  },
  {
    featureType: "poi.sports_complex",
    elementType: "all",
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
      {
        saturation: "-11",
      },
    ],
  },
  {
    featureType: "road",
    elementType: "geometry.stroke",
    stylers: [
      {
        visibility: "on",
      },
    ],
  },
  {
    featureType: "road",
    elementType: "labels.text.fill",
    stylers: [
      {
        saturation: "0",
      },
      {
        lightness: "0",
      },
      {
        color: "#919ab7",
      },
    ],
  },
  {
    featureType: "road",
    elementType: "labels.icon",
    stylers: [
      {
        saturation: "-24",
      },
      {
        visibility: "on",
      },
      {
        hue: "#e9ff00",
      },
      {
        lightness: "16",
      },
    ],
  },
  {
    featureType: "road.highway",
    elementType: "all",
    stylers: [
      {
        saturation: "0",
      },
      {
        lightness: "0",
      },
    ],
  },
  {
    featureType: "road.highway",
    elementType: "geometry.fill",
    stylers: [
      {
        color: "#d8cafb",
      },
    ],
  },
  {
    featureType: "road.highway",
    elementType: "geometry.stroke",
    stylers: [
      {
        color: "#b296f8",
      },
    ],
  },
  {
    featureType: "road.highway",
    elementType: "labels.icon",
    stylers: [
      {
        saturation: "55",
      },
      {
        lightness: "0",
      },
      {
        weight: "1",
      },
    ],
  },
  {
    featureType: "road.highway.controlled_access",
    elementType: "all",
    stylers: [
      {
        visibility: "on",
      },
      {
        saturation: "0",
      },
      {
        lightness: "16",
      },
    ],
  },
  {
    featureType: "road.highway.controlled_access",
    elementType: "geometry.fill",
    stylers: [
      {
        color: "#b296f8",
      },
    ],
  },
  {
    featureType: "road.highway.controlled_access",
    elementType: "geometry.stroke",
    stylers: [
      {
        color: "#865bf4",
      },
    ],
  },
  {
    featureType: "road.highway.controlled_access",
    elementType: "labels.icon",
    stylers: [
      {
        visibility: "on",
      },
      {
        lightness: "0",
      },
      {
        saturation: "-17",
      },
      {
        hue: "#5c00ff",
      },
      {
        weight: "1.00",
      },
      {
        gamma: "1.00",
      },
    ],
  },
  {
    featureType: "road.arterial",
    elementType: "labels.icon",
    stylers: [
      {
        visibility: "off",
      },
    ],
  },
  {
    featureType: "road.local",
    elementType: "geometry.fill",
    stylers: [
      {
        color: "#ffffff",
      },
    ],
  },
  {
    featureType: "transit.station",
    elementType: "geometry.fill",
    stylers: [
      {
        visibility: "on",
      },
      {
        hue: "#0000ff",
      },
      {
        saturation: "20",
      },
    ],
  },
  {
    featureType: "water",
    elementType: "all",
    stylers: [
      {
        color: "#e4e5ff",
      },
      {
        visibility: "on",
      },
    ],
  },
  {
    featureType: "water",
    elementType: "geometry.fill",
    stylers: [
      {
        color: "#a8e0ff",
      },
    ],
  },
  {
    featureType: "water",
    elementType: "labels.text.fill",
    stylers: [
      {
        color: "#7ba8ff",
      },
      {
        visibility: "off",
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
