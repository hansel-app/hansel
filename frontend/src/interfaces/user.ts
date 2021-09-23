export interface User {
  userId: number;
  username: string;
  displayName: string;
  avatar: string | Uint8Array; // TODO: replace this with bytes or something?
}
