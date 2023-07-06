job "redis" {
  datacenters = ["linode"]
  type        = "service"
  namespace   = "default"

  group "redis" {
    count = 1

    network {
      port "redis" {
        static = 6379
      }
    }

    service {
      name     = "redis"
      provider = "nomad"
      port     = "redis"
      tags = [
        "traefik.enable=false",
      ]
    }

    task "server" {
      driver = "docker"
      config {
        image = "redis:7-alpine"
        ports = ["redis"]
        args = [
          "--appendonly", "yes",
        ]
      }
    }
  }
}
