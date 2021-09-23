import { User as ProtoUser } from "@/protobuf/user_pb";
import { Dayjs } from "dayjs";

export type User = ProtoUser.AsObject;

export interface PendingFriendRequest {
  requester: User;
  createdAt: Dayjs;
}
