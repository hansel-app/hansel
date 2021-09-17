import { User, Gem } from "@/interfaces";

export const mockFriends: User[] = [
  { id: 1, displayName: "Groundskeeper", username: "abc" },
  { id: 2, displayName: "Wimp", username: "abc" },
  { id: 3, displayName: "Goose", username: "abc" },
  { id: 4, displayName: "Pub owner", username: "abc" },
];

export const mockSelfUser: User = {
  id: 100, displayName: "Maggie Mee", username: "maggiemee"
}

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