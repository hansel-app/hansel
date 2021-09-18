import { hansel } from "@/interfaces";
import { Dayjs } from "dayjs";

// TODO: replace with interface generated from protobuf?
export interface Gem {
  color: GemColor;
}

export interface GemColor {
  selectionColor: HexCode;
  tintColor: HexCode;
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