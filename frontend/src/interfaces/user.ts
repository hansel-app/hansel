import { Dayjs } from "dayjs";

export type Base64String = string;

export interface User {
  userId: number;
  username: string;
  displayName: string;
  avatar: Base64String;
}

export interface PendingFriendRequest {
  requester: User;
  createdAt: Dayjs;
}
