import { Gem, GemColor, User } from "@/interfaces";
import dayjs from "dayjs";
import { placeholderBase64Avatar } from "./placeholderBase64Avatar";

export const mockFriends: User[] = [
  {
    userId: 1,
    displayName: "Groundskeeper",
    username: "goosehater123",
    avatar: "",
  },
  {
    userId: 2,
    displayName: "Wimp",
    username: "scared_of_geese",
    avatar: "2",
  },
  {
    userId: 3,
    displayName: "Goose",
    username: "goose",
    avatar: "3",
  },
  {
    userId: 4,
    displayName: "Pub owner",
    username: "abc",
    avatar: "4",
  },
];

export const mockSelfUser: User = {
  userId: 100,
  displayName: "Maggie Mee",
  username: "maggiemee",
  avatar: placeholderBase64Avatar,
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
    attachment: "",
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
    attachment: "",
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
    attachment: "",
  },
];
