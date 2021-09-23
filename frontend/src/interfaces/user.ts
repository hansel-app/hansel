import { User as ProtoUser, PendingFriendRequest as ProtoPendingFriendRequest } from "@/protobuf/user_pb";

export type User = ProtoUser.AsObject;
export type PendingFriendRequest = ProtoPendingFriendRequest.AsObject;
