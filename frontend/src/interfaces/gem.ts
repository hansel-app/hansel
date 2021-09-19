import { hansel } from "@/interfaces";
import { Dayjs } from "dayjs";

// TODO: replace with interface generated from protobuf?
export interface Gem {
  color: GemColor;
}

// TODO: replace with actual hex codes.
export enum GemColor {
  PURPLE = "#cdb4db",
  PINK = "#ffc8dd",
  BLUE = "#a2d2ff",
  YELLOW = "#fffb8a",
  GREEN = "#a6ffb6",
}

type HexCode = String;

export interface GemInfo {
  message: string;
  position: hansel.LatLng;
  dropper: string;
  receiver: string;
  dropTime: Dayjs;
  color: string;
}

export enum color {
  PURPLE = 'purple',
  PINK = 'pink',
  BLUE = 'blue',
  BLACK = 'black',
  YELLOW = 'yellow',
  GREEN = 'green',
}
