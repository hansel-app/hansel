import { LatLng, User } from "@/interfaces";
import { Dayjs } from "dayjs";

export interface Gem {
  id: number;
  message: string;
  position: LatLng;
  createdAt: Dayjs;
  createdBy: User;
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
