# github.com/inveracity/ssh-tunnel
tunnel {
    user = "root"
    name = "Nomad"
    local {
        port = 4646
        cmd = ["wslview", "http://localhost:4646"]
    }
    remote {
        host = "linode:22"
        port = 4646
    }
}

tunnel {
    name = "Traefik"
    user = "root"
    local {
        port = 8080
        cmd = ["wslview", "http://localhost:8080"]
    }
    remote {
        host = "linode:22"
        port = 8080
    }
}
