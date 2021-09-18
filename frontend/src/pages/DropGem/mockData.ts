import { User, Gem } from "@/interfaces";

export const mockFriends: User[] = [
  { id: 1, username: "Groundskeeper" },
  { id: 2, username: "Wimp" },
  { id: 3, username: "Goose" },
  { id: 4, username: "Pub owner" },
];

// Color values are placeholders
export const mockGems: Gem[] = [
  {
    color: {
      selectionColor: "#cdb4db",
      tintColor: "#000000",
    },
  },
  {
    color: {
      selectionColor: "#ffc8dd",
      tintColor: "#000000",
    },
  },
  {
    color: {
      selectionColor: "#ffafcc",
      tintColor: "#000000",
    },
  },
  {
    color: {
      selectionColor: "#bde0fe",
      tintColor: "#000000",
    },
  },
  {
    color: {
      selectionColor: "#a2d2ff",
      tintColor: "#000000",
    },
  },
];