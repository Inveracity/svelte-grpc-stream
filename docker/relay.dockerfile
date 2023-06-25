FROM golang:1.20.4-alpine3.18 AS build
WORKDIR /app

# Cache dependencies before building
COPY relay/go.mod .
COPY relay/go.sum .
RUN go mod download

# Build after caching should prevent installing dependencies on every build
COPY relay .

RUN go build -o /app/relay cmd/relay/main.go
CMD ["/app/relay"]

# FROM scratch
# WORKDIR /
# COPY --from=build /app/relay /app/relay
# EXPOSE 50051
# CMD ["/app/relay"]
