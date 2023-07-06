job "relay" {
  datacenters = ["linode"]
  type        = "service"
  namespace   = "default"

  group "relay" {
    count = 1

    network {
      port "grpc" {
        to = 50051
      }
    }

    service {
      name     = "relay-grpc"
      provider = "nomad"
      port     = "grpc"
      tags = [
        "traefik.enable=true",
        "traefik.http.routers.relay.rule=Host(`chat.christopherbaklid.com`) && PathPrefix(`/proto.chat.v1.ChatService`)",
        "traefik.http.routers.relay.entrypoints=websecure",
        "traefik.http.routers.relay.tls=true",
        "traefik.http.routers.relay.tls.certresolver=myresolver",
        "traefik.http.services.relay.loadbalancer.server.scheme=h2c",
        "traefik.http.middlewares.relay-grpc.grpcWeb.allowOrigins=*",
        "traefik.http.routers.relay.middlewares=relay-grpc",
      ]
    }

    task "server" {
      driver = "docker"

      config {
        image = "inveracity/chat-relay:latest"
        ports = ["grpc"]

        args = [
          "-port=${NOMAD_PORT_grpc}",
          "-redis=127.0.0.1:6379",
          "-nats=127.0.0.1:4222",
        ]
      }
    }
  }
}
