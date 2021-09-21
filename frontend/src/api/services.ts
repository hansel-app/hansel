import { AuthServicePromiseClient } from "@/protobuf/auth_grpc_web_pb";
import { GemServicePromiseClient } from "@/protobuf/gem_grpc_web_pb";
import AuthInterceptor from "@/api/interceptors/auth";

const SERVER_URL = window.env.VUE_APP_SERVER_URL;
const SERVER_PORT = window.env.VUE_APP_SERVER_PORT;
const hostname = `${SERVER_URL}:${SERVER_PORT}`;

const authInterceptor = new AuthInterceptor();

// Note that the equivalent stream authentication interceptor is not defined.
// If we need to protect stream RPC calls, then it will have to be implemented.
const options = {
  unaryInterceptors: [authInterceptor]
};

const services = Object.freeze({
  authClient: new AuthServicePromiseClient(hostname, null, options),
  gemsClient: new GemServicePromiseClient(hostname, null, options)
});

export default services;
