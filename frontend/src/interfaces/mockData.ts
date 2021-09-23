import { Gem, GemColor, GemLogs, GemLogsWithFriend, User } from "@/interfaces";
import dayjs from "dayjs";
import { placeholderBase64Avatar } from "./placeholderBase64Avatar";
import { utownBase64, socBase64 } from "./mockBase64Pictures";

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
    id: 111,
    message: "Yo did you oversleep for tutorial again",
    position: { lat: 1.2986, lng: 103.7764 },
    createdAt: dayjs(),
    createdBy: mockSelfUser,
    receiver: mockFriends[0],
    receivedAt: dayjs(),
    color: GemColor.BLUE,
    attachment: socBase64,
  },
  {
    id: 222,
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
    id: 333,
    message: "I saw ashy here today!",
    position: { lat: 1.2976, lng: 103.7764 },
    createdAt: dayjs(),
    createdBy: mockSelfUser,
    receiver: mockFriends[0],
    receivedAt: dayjs(),
    color: GemColor.BLACK,
    attachment: utownBase64,
  },
  {
    id: 444,
    message: "Me gusta los camarones",
    position: { lat: 1.2986, lng: 103.7764 },
    createdAt: dayjs(),
    createdBy: mockFriends[0],
    receiver: mockSelfUser,
    receivedAt: null,
    color: GemColor.PINK,
    attachment: "",
  },
];

export const mockGemLogs: GemLogs = {
  gemLogsMap: new Map([
    [1, { friend: mockFriends[0], gems: mockGems } as GemLogsWithFriend],
    [2, { friend: mockFriends[1], gems: mockGems } as GemLogsWithFriend],
    [3, { friend: mockFriends[2], gems: mockGems } as GemLogsWithFriend],
    [4, { friend: mockFriends[3], gems: mockGems } as GemLogsWithFriend],
  ]),
};
