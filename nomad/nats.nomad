job "nats" {
  datacenters = ["linode"]
  type        = "service"
  namespace   = "default"

  group "nats" {
    count = 1

    network {
      port "nats" {
        static = 4222
      }
    }

    service {
      name     = "nats"
      provider = "nomad"
      port     = "nats"
      tags = [
        "traefik.enable=false",
      ]
    }

    task "server" {
      driver = "docker"
      config {
        image = "nats:latest"
        ports = ["nats"]
      }
    }
  }
}
