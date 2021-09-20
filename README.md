# Hansel

## Getting Started

1. Install Protocol Buffer using the instructions [here](https://grpc.io/docs/protoc-installation/).
1. Follow the instructions in the respective frontend and backend `README`s.
1. Install Docker, Docker Compose, and run `docker-compose up` in the root directory to run an Envoy proxy.

   - If you are facing any difficulties getting the proxy to work, you may need replace `address: 127.0.0.1` with `address: host.docker.internal` or `address: docker.for.mac.localhost` in `envoy/envoy.yaml`
