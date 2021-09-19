import { GemServicePromiseClient } from "@/protobuf/gem_grpc_web_pb";
import AuthInterceptor from "@/api/interceptors/auth";

const SERVER_URL = process.env.VUE_APP_SERVER_URL;
const SERVER_PORT = process.env.VUE_APP_SERVER_PORT;
const hostname = `${SERVER_URL}:${SERVER_PORT}`;

// TODO: Inject token.
const authInterceptor = new AuthInterceptor();

// Note that the equivalent stream authentication interceptor is not defined.
// If we need to protect stream RPC calls, then it will have to be implemented.
const options = {
  unaryInterceptors: [authInterceptor]
};

const services = Object.freeze({
  gemsClient: new GemServicePromiseClient(hostname, null, options)
});

export default services;
