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

// 'any' annotations is to support accessing the key given the enum value
// https://stackoverflow.com/questions/62215454/how-to-get-enum-key-by-value-in-typescript
export enum GemColor {
  PURPLE = <any>"#cdb4db",
  PINK = <any>"#ffc8dd",
  BLUE = <any>"#a2d2ff",
  BLACK = <any>"#080808",
  YELLOW = <any>"#fffb8a",
  GREEN = <any>"#a6ffb6",
}
