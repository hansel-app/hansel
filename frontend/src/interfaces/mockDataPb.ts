// Contains mock data for protobuf models
// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore
import {
  Gem as GemPb,
  GemLogsWithFriend,
  GemColor,
} from "@/protobuf/gem_pb.d.ts";
// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore
import { User as UserPb } from "@/protobuf/user_pb.d.ts";

const buildGemPb = (gem: GemPb.AsObject): GemPb => {
  const gemPb = new GemPb();
  gemPb.setMessage(gem.message);
  gemPb.setColor(gem.color);
  return gemPb;
};

const buildUserPb = (user: UserPb.AsObject): UserPb => {
  const userPb = new UserPb();
  userPb.setDisplayName(user.displayName);
  userPb.setUsername(user.username);
  userPb.setAvatar(user.avatar);
  return userPb;
};

export const mockFriends: UserPb[] = [
  {
    userId: 1,
    displayName: "Groundskeeper",
    username: "goosehater123",
    avatar:
      "https://cdn.mos.cms.futurecdn.net/JycrJzD5tvbGHWgjtPrRZY-970-80.jpg.webp",
  },
  {
    userId: 2,
    displayName: "Wimp",
    username: "scared_of_geese",

    avatar:
      "https://cdn.mos.cms.futurecdn.net/JycrJzD5tvbGHWgjtPrRZY-970-80.jpg.webp",
  },
  {
    userId: 3,
    displayName: "Goose",
    username: "goose",

    avatar:
      "https://cdn.mos.cms.futurecdn.net/JycrJzD5tvbGHWgjtPrRZY-970-80.jpg.webp",
  },
  {
    userId: 4,
    displayName: "Pub owner",
    username: "abc",

    avatar:
      "https://cdn.mos.cms.futurecdn.net/JycrJzD5tvbGHWgjtPrRZY-970-80.jpg.webp",
  },
].map(buildUserPb);

export const mockSelfUser: UserPb = buildUserPb({
  userId: 100,
  displayName: "Maggie Mee",
  username: "maggiemee",
  avatar:
    "https://cdn.mos.cms.futurecdn.net/JycrJzD5tvbGHWgjtPrRZY-970-80.jpg.webp",
});

export const mockGems: GemPb[] = [
  {
    id: 1111,
    message: "Me gusta los camarones",
    latitude: 1.2966,
    longitude: 103.7764,
    createdAt: { seconds: 0, nanos: 0 },
    creatorId: mockFriends[0].getUserId(),
    receiverId: 1,
    receivedAt: { seconds: 0, nanos: 0 },
    color: GemColor.PURPLE,
    attachment: "",
  },
  {
    id: 222,
    message: "Me gusta los camarones",
    latitude: 1.2976,
    longitude: 103.7764,
    createdAt: { seconds: 0, nanos: 0 },
    creatorId: mockFriends[1].getUserId(),
    receiverId: 1,
    receivedAt: { seconds: 0, nanos: 0 },
    color: GemColor.BLACK,
    attachment: "",
  },
  {
    id: 333,
    message: "Me gusta los camarones",
    latitude: 1.2986,
    longitude: 103.7764,
    createdAt: { seconds: 0, nanos: 0 },
    creatorId: mockFriends[2].getUserId(),
    receiverId: 1,
    receivedAt: { seconds: 0, nanos: 0 },
    color: GemColor.BLUE,
    attachment: "",
  },
].map(buildGemPb);

export const mockGemLogsWithFriend = (): GemLogsWithFriend => {
  const mockGemLogsWithFriend: GemLogsWithFriend = new GemLogsWithFriend();
  mockGemLogsWithFriend.setFriend(mockFriends[0]);
  mockGemLogsWithFriend.setGemsList(mockGems);
  return mockGemLogsWithFriend;
};
