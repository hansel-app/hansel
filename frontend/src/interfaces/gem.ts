import { LatLng } from "@/interfaces";
import { Dayjs } from "dayjs";

// TODO: replace with interface generated from protobuf?
export interface Gem {
  id: number;
  message: string;
  position: LatLng;
  createdAt: Dayjs;
  createdBy: string; // TODO: replace this with user object
  receivedAt: Dayjs;
  // color: GemColor; // TODO: replace this with enum
}

export interface GemColor {
  selectionColor: HexCode;
  tintColor: HexCode;
}

type HexCode = string;

// TODO: remove this and replace it directly with Gem
export interface GemInfo {
  message: string;
  position: LatLng;
  dropper: string;
  receiver: string;
  dropTime: Dayjs;
  color: string;
}

export enum color {
  PURPLE = "purple",
  PINK = "pink",
  BLUE = "blue",
  BLACK = "black",
  YELLOW = "yellow",
  GREEN = "green",
}
