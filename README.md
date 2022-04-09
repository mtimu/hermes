# Hermes
Signaling server over EMQX for WebRTC. It uses MongoDB to manage rooms and EMQX for sending(receiving) events to(from) clients.

## WebRTC
WebRTC is direct peer-to-peer connection which is done using WebRTC API in browsers and mobile applications. But for starting a group call, it requires a mechnaism to share peers address and connectivity information with each other. This mechanism is called Signaling. It keeps an open connection with each peer to exchange peers' address, connectivity info, and room events.

## Run Locally
Run command `make up` to run Hermes and dependencies locally. to kill the services run `make down`.

## Deployment
Deploy EMQX and MongoDB services and pass their configurations to hermes `values.yaml`. Use Helm to deploy Hermes's Chart. corresponding command is `make deploy`. and to bring down the Hermes run `make deployd`.

## Other Useful Commands
- `make croom`: create a test room
- `make jroom`: join a user to a room. (pass room_id parameter)
- `make image`: build image and push to registry


## Warning
Please update parameters in `Makefile` and configuration in `deployment/hermes/values.yaml` and `internal/config/default.go` for the project to run properly. If necessary, you can create a `config.yaml` file in root directory to pass configuration.
