import { LatLng, User } from "@/interfaces";
import { Dayjs } from "dayjs";

export interface Gem {
  id: number;
  message: string;
  position: LatLng;
  createdAt: Dayjs;
  createdBy: User;
  receiver: User;
  receivedAt: Dayjs | null;
  color: GemColor;
}

export enum GemColor {
  PURPLE = "#cdb4db",
  PINK = "#ffc8dd",
  BLUE = "#a2d2ff",
  BLACK = "#080808",
  YELLOW = "#fffb8a",
  GREEN = "#a6ffb6",
}
