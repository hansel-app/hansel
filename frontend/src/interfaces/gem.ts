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
  PURPLE = "purple",
  PINK = "pink",
  BLUE = "blue",
  BLACK = "black",
  YELLOW = "yellow",
  GREEN = "green",
}
