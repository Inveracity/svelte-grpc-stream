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
        # Router settings
        "traefik.http.routers.relay.rule=Host(`chat.christopherbaklid.com`) && PathPrefix(`/relay/proto.chat.v1.ChatService`)",
        "traefik.http.routers.relay.entrypoints=websecure",
        "traefik.http.routers.relay.tls=true",
        "traefik.http.routers.relay.tls.certresolver=myresolver",
        "traefik.http.routers.relay.middlewares=relay-grpc,relay-stripprefix",
        # Middlewares
        "traefik.http.middlewares.relay-grpc.grpcWeb.allowOrigins=*",
        "traefik.http.middlewares.relay-stripprefix.stripprefix.prefixes=/relay",
        # Services
        "traefik.http.services.relay.loadbalancer.server.scheme=h2c",
      ]
    }

    task "server" {
      driver = "docker"

      template {
        destination = "${NOMAD_SECRETS_DIR}/env.vars"
        env = true
        data = <<-EOF
        {{ with nomadVar "nomad/jobs" }}
        PB_PASS={{ .pb_admin_password }}
        PB_USER={{ .pb_admin_email }}
        {{ end }}
        EOF
      }

      config {
        image = "inveracity/chat-relay:latest"
        ports = ["grpc"]
        args = [
          "-port=${NOMAD_PORT_grpc}",
          "-redis=${meta.public_ip}:6379",
          "-nats=${meta.public_ip}:4222",
          "-pocketbase=http://${meta.public_ip}:8090",
        ]
      }
    }
  }
}
