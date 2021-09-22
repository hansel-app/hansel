import { Gem as ProtoGem, GemColor as ProtoGemColor } from "@/protobuf/gem_pb";
import { mockFriends, mockSelfUser } from "@/interfaces/mockData";
import dayjs from "dayjs";
import { Gem, GemColor } from "@/interfaces";

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
  return {
    id: protoGem.getId(),
    message: protoGem.getMessage(),
    position: {
      lat: protoGem.getLatitude(),
      lng: protoGem.getLongitude(),
    },
    createdAt: dayjs(protoGem.getCreatedAt()?.toDate()),
    // TODO: replace these with actual user object
    createdBy: mockFriends[0],
    receiver: mockSelfUser,
    receivedAt: dayjs(protoGem.getReceivedAt()?.toDate()),
    color: protoGemColorToGemColorMapper(protoGem.getColor()),
  };
};
