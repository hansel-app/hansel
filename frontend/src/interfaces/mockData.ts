import { Gem, GemColor, User } from "@/interfaces";
import dayjs from "dayjs";

export const mockFriends: User[] = [
  {
    id: 1,
    displayName: "Groundskeeper",
    username: "goosehater123",
    avatar:
      "https://cdn.mos.cms.futurecdn.net/JycrJzD5tvbGHWgjtPrRZY-970-80.jpg.webp",
  },
  {
    id: 2,
    displayName: "Wimp",
    username: "scared_of_geese",

    avatar:
      "https://cdn.mos.cms.futurecdn.net/JycrJzD5tvbGHWgjtPrRZY-970-80.jpg.webp",
  },
  {
    id: 3,
    displayName: "Goose",
    username: "goose",

    avatar:
      "https://cdn.mos.cms.futurecdn.net/JycrJzD5tvbGHWgjtPrRZY-970-80.jpg.webp",
  },
  {
    id: 4,
    displayName: "Pub owner",
    username: "abc",

    avatar:
      "https://cdn.mos.cms.futurecdn.net/JycrJzD5tvbGHWgjtPrRZY-970-80.jpg.webp",
  },
];

export const mockSelfUser: User = {
  id: 100,
  displayName: "Maggie Mee",
  username: "maggiemee",
  avatar:
    "https://cdn.mos.cms.futurecdn.net/JycrJzD5tvbGHWgjtPrRZY-970-80.jpg.webp",
};

export const mockGems: Gem[] = [
  {
    id: 1111,
    message: "Me gusta los camarones",
    position: { lat: 1.2966, lng: 103.7764 },
    createdAt: dayjs(),
    createdBy: mockFriends[0],
    receiver: mockSelfUser,
    receivedAt: null,
    color: GemColor.PURPLE,
  },
  {
    id: 222,
    message: "Me gusta los camarones",
    position: { lat: 1.2976, lng: 103.7764 },
    createdAt: dayjs(),
    createdBy: mockFriends[1],
    receiver: mockSelfUser,
    receivedAt: null,
    color: GemColor.BLACK,
  },
  {
    id: 333,
    message: "Me gusta los camarones",
    position: { lat: 1.2986, lng: 103.7764 },
    createdAt: dayjs(),
    createdBy: mockFriends[2],
    receiver: mockSelfUser,
    receivedAt: null,
    color: GemColor.BLUE,
  },
];
