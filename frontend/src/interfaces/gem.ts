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

export type HexCode = string;

export enum GemColor {
  PURPLE = "#BC88B8",
  PINK = "#ffc8dd",
  BLUE = "#99CEE0",
  BLACK = "#8B8B8B",
  YELLOW = "#F0D072",
  GREEN = "#99BC88",
}

export interface DropGemFormState {
  id?: number;
  message?: string;
  receiverId?: number;
  color: GemColor;
  attachment?: File;
}
