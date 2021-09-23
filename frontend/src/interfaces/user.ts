import { User as ProtoUser } from "@/protobuf/user_pb";

export type User = ProtoUser.AsObject;
