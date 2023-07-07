
job "frontend" {
  datacenters = ["linode"]
  type        = "service"
  namespace   = "default"

  group "frontend" {
    count = 1

    network {
      port "http" {
        to = 3000
      }
    }

    service {
      name     = "frontend"
      provider = "nomad"
      port     = "http"
      tags = [
        "traefik.enable=true",
        # Router settings
        "traefik.http.routers.frontend.rule=Host(`chat.christopherbaklid.com`)",
        "traefik.http.routers.frontend.entrypoints=websecure",
        "traefik.http.routers.frontend.tls.certresolver=myresolver",
        "traefik.http.routers.frontend.tls=true",
      ]
    }

    task "server" {
      driver = "docker"

      env {
        PUBLIC_POCKETBASE_URL = "https://chat.christopherbaklid.com/pocketbase"
        PUBLIC_RELAY_URL = "https://chat.christopherbaklid.com/relay"
      }

      config {
        image = "inveracity/chat-frontend:latest"
        ports = ["http"]
      }
    }
  }
}
