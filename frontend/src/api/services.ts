import { AuthServicePromiseClient } from "@/protobuf/auth_grpc_web_pb";
import { GemServicePromiseClient } from "@/protobuf/gem_grpc_web_pb";
import { FriendServicePromiseClient } from "@/protobuf/friend_grpc_web_pb";
import AuthInterceptor from "@/api/interceptors/auth";
import { UserServicePromiseClient } from "@/protobuf/user_grpc_web_pb";

const SERVER_URL = window.env.VUE_APP_SERVER_URL;
const SERVER_PORT = window.env.VUE_APP_SERVER_PORT;
const hostname = `${SERVER_URL}:${SERVER_PORT}`;

const authInterceptor = new AuthInterceptor();

// Note that the equivalent stream authentication interceptor is not defined.
// If we need to protect stream RPC calls, then it will have to be implemented.
const options = {
  unaryInterceptors: [authInterceptor],
};

const services = Object.freeze({
  authClient: new AuthServicePromiseClient(hostname, null, options),
  gemsClient: new GemServicePromiseClient(hostname, null, options),
  userClient: new UserServicePromiseClient(hostname, null, options),
  friendsClient: new FriendServicePromiseClient(hostname, null, options),
});

export default services;
