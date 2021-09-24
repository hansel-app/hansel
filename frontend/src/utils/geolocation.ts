import { LatLng } from "@/interfaces";
import axios from "axios";

type Kilometer = number;

const GOOGLE_API_KEY = window.env.VUE_APP_GOOGLE_API_KEY;

// Adapted from https://stackoverflow.com/a/27943
// Returns distance between 2 points (as indicated by their LatLng values) in km
export function getDistanceFromLatLonInKm(
  latlng1: LatLng,
  latlng2: LatLng
): Kilometer {
  const R = 6371; // Radius of the earth in km
  const dLat = degreesToRadian(latlng2.lat - latlng1.lat);
  const dLon = degreesToRadian(latlng2.lng - latlng1.lng);

  const a =
    Math.sin(dLat / 2) * Math.sin(dLat / 2) +
    Math.cos(degreesToRadian(latlng1.lat)) *
      Math.cos(degreesToRadian(latlng2.lat)) *
      Math.sin(dLon / 2) *
      Math.sin(dLon / 2);
  const c = 2 * Math.atan2(Math.sqrt(a), Math.sqrt(1 - a));
  const d = R * c;

  return d;
}

export function formatDistance(km: number): string {
  if (km >= 1) {
    return `${km.toFixed(3)} km`;
  } else {
    return `${(km * 1000).toFixed(0)} m`;
  }
}

function degreesToRadian(deg: number): number {
  return deg * (Math.PI / 180);
}

export async function getLocationNameFromLatLng(
  coords: LatLng
): Promise<string> {
  return new Promise((resolve, reject) => {
    axios
      .get(
        "https://maps.googleapis.com/maps/api/geocode/json?latlng=" +
          coords.lat +
          "," +
          coords.lng +
          "&key=" +
          GOOGLE_API_KEY
      )
      .then((response) => {
        const address = response.data.results[1];
        if (address !== undefined) resolve(address.formatted_address);
      })
      .catch((error) => reject(error));
  });
}
