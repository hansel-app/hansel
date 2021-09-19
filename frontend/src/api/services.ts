import { GemServicePromiseClient } from "@/protobuf/gem_grpc_web_pb";

const SERVER_URL = process.env.VUE_APP_SERVER_URL;
const SERVER_PORT = process.env.VUE_APP_SERVER_PORT;
const hostname = `${SERVER_URL}:${SERVER_PORT}`;

const services = Object.freeze({
  gemsClient: new GemServicePromiseClient(hostname, null, null)
});

export default services;
