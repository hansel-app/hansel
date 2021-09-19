import { User, Gem, GemColor } from "@/interfaces";

export const mockFriends: User[] = [
  { id: 1, displayName: "Groundskeeper", username: "goosehater123" },
  { id: 2, displayName: "Wimp", username: "scared_of_geese" },
  { id: 3, displayName: "Goose", username: "goose" },
  { id: 4, displayName: "Pub owner", username: "abc" },
];

export const mockSelfUser: User = {
  id: 100,
  displayName: "Maggie Mee",
  username: "maggiemee",
};

export const mockGems: Gem[] = [
  {
    color: GemColor.PURPLE,
  },
  {
    color: GemColor.BLUE,
  },
  {
    color: GemColor.GREEN,
  },
  {
    color: GemColor.YELLOW,
  },
  {
    color: GemColor.PINK,
  },
];