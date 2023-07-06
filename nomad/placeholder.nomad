job "whoami" {
  datacenters = ["linode"]
  type        = "service"
  namespace   = "default"

  group "whoami" {
    count = 1

    network {
      port "http" {
        to = 80
      }
    }

    service {
      name     = "whoami-http"
      provider = "nomad"
      port     = "http"
      tags = [
        "traefik.enable=true",
        "traefik.http.routers.whoami.rule=Host(`chat.christopherbaklid.com`)",
        "traefik.http.routers.whoami.entrypoints=websecure",
        "traefik.http.routers.whoami.tls=true",
        "traefik.http.routers.whoami.tls.certresolver=myresolver",
        "traefik.http.routers.whoami.middlewares=whoami-https",
        "traefik.http.middlewares.whoami-https.redirectscheme.scheme=https",
        "traefik.http.middlewares.whoami-https.redirectscheme.permanent=true",
      ]
    }

    task "server" {
      driver = "docker"
      env {
        TEST = "test"
      }
      config {
        image = "traefik/whoami:latest"
        ports = ["http"]
      }
    }
  }
}
