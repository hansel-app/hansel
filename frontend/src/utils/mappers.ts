import {
  Gem,
  GemColor,
  GemLogs,
  GemLogsWithFriend,
  PendingFriendRequest,
  User,
} from "@/interfaces";
import {
  Gem as ProtoGem,
  GemColor as ProtoGemColor,
  GemLogs as ProtoGemLogs,
  GemLogsWithFriend as ProtoGemLogsWithFriend,
} from "@/protobuf/gem_pb";
import {
  PendingFriendRequest as ProtoPendingFriendRequest,
  User as ProtoUser,
} from "@/protobuf/user_pb";
import dayjs from "dayjs";

export const protoGemColorToGemColorMapper = (
  protoGemColor: ProtoGemColor
): GemColor => {
  switch (protoGemColor) {
    case ProtoGemColor.PURPLE:
      return GemColor.PURPLE;
    case ProtoGemColor.PINK:
      return GemColor.PINK;
    case ProtoGemColor.BLUE:
      return GemColor.BLUE;
    case ProtoGemColor.BLACK:
      return GemColor.BLACK;
    case ProtoGemColor.YELLOW:
      return GemColor.YELLOW;
    case ProtoGemColor.GREEN:
      return GemColor.GREEN;
    default:
      throw new Error("Unknown gem color received!");
  }
};

export const gemColorToProtoGemColorMapper = (
  gemColor: GemColor
): ProtoGemColor => {
  switch (gemColor) {
    case GemColor.PURPLE:
      return ProtoGemColor.PURPLE;
    case GemColor.PINK:
      return ProtoGemColor.PINK;
    case GemColor.BLUE:
      return ProtoGemColor.BLUE;
    case GemColor.BLACK:
      return ProtoGemColor.BLACK;
    case GemColor.YELLOW:
      return ProtoGemColor.YELLOW;
    case GemColor.GREEN:
      return ProtoGemColor.GREEN;
    default:
      throw new Error("Unknown gem color received!");
  }
};

export const protoGemToGemMapper = (protoGem: ProtoGem): Gem => {
  const creator = protoGem.getCreator();
  const receiver = protoGem.getReceiver();
  if (!creator || !receiver) {
    throw new Error("Gem must have a creator and receiver");
  }

  return {
    id: protoGem.getId(),
    message: protoGem.getMessage(),
    position: {
      lat: protoGem.getLatitude(),
      lng: protoGem.getLongitude(),
    },
    createdAt: dayjs(protoGem.getCreatedAt()?.toDate()),
    createdBy: protoUserToUserMapper(creator),
    receivedAt: protoGem.getReceivedAt()
      ? dayjs(protoGem.getReceivedAt()?.toDate())
      : null,
    receiver: protoUserToUserMapper(receiver),
    color: protoGemColorToGemColorMapper(protoGem.getColor()),
    attachment: protoGem.getAttachment_asB64(),
  };
};

export const protoGemLogWithFriendsToGemLogWithFriendsMapper = (
  protoGemLogWithFriends: ProtoGemLogsWithFriend
): GemLogsWithFriend => {
  const protoFriend = protoGemLogWithFriends.getFriend();
  const protoGems = protoGemLogWithFriends.getGemsList();

  if (!protoFriend) {
    throw new Error("Gem log must be associated with a friend");
  }

  return {
    friend: protoUserToUserMapper(protoFriend),
    gems: protoGems.map(protoGemToGemMapper),
  };
};

export const protoGemLogToGemLogMapper = (
  protoGemLog: ProtoGemLogs
): GemLogs => {
  const gemLogs = protoGemLog.getGemLogsMap();

  const mappedMap = new Map<number, GemLogsWithFriend>();

  for (const [friendId, val] of gemLogs.entries()) {
    const gemLogsWithFriends = protoGemLogWithFriendsToGemLogWithFriendsMapper(
      val
    );
    mappedMap.set(friendId, gemLogsWithFriends);
  }

  return {
    gemLogsMap: mappedMap,
  };
};

export const protoUserToUserMapper = (protoUser: ProtoUser): User => {
  return {
    userId: protoUser.getUserId(),
    username: protoUser.getUsername(),
    displayName: protoUser.getDisplayName(),
    avatar: protoUser.getAvatar_asB64(),
  };
};

export const protoPendingFriendRequestToPendingFriendRequestMapper = (
  protoPendingFriendRequest: ProtoPendingFriendRequest
): PendingFriendRequest => {
  const requester = protoPendingFriendRequest.getRequester();
  if (!requester) {
    throw new Error("Pending friend request must have a requester");
  }

  return {
    requester: protoUserToUserMapper(requester),
    createdAt: dayjs(protoPendingFriendRequest.getCreatedAt()?.toDate()),
  };
};
