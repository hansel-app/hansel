export type Base64String = string;

export interface User {
  userId: number;
  username: string;
  displayName: string;
  avatar: Base64String;
}
