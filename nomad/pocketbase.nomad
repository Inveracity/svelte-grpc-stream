job "pocketbase" {
  datacenters = ["linode"]
  type        = "service"
  namespace   = "default"

  group "pocketbase" {
    count = 1

    network {
      port "http" {
        static = 8090
      }
    }

    volume "pb-data" {
        type      = "host"
        read_only = false
        source    = "pb-data"
    }

    service {
      name     = "pocketbase"
      provider = "nomad"
      port     = "http"
      tags = [
        "traefik.enable=true",
        # Router settings
        "traefik.http.routers.pocketbase.rule=Host(`chat.christopherbaklid.com`) && PathPrefix(`/pocketbase`)",
        "traefik.http.routers.pocketbase.entrypoints=websecure",
        "traefik.http.routers.pocketbase.tls.certresolver=myresolver",
        "traefik.http.routers.pocketbase.tls=true",
        "traefik.http.routers.pocketbase.middlewares=pb-stripprefix",
        # Middlewares
        "traefik.http.middlewares.pb-stripprefix.stripprefix.prefixes=/pocketbase",
      ]
    }

    task "server" {
      driver = "docker"

      template {
        destination = "${NOMAD_SECRETS_DIR}/env.vars"
        env = true
        data = <<-EOF
        {{ with nomadVar "nomad/jobs" }}
        PB_ADMIN_USER={{ .pb_admin_email }}
        PB_ADMIN_PASS={{ .pb_admin_password }}
        {{ end }}
        EOF
      }

      config {
        image = "inveracity/pocketbase:0.16.7"
        ports = ["http"]
      }
      volume_mount {
        volume      = "pb-data"
        destination = "/pb/pb_data"
      }
    }
  }
}
