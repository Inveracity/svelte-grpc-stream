FROM golang:1.20.4-alpine3.18 AS build
WORKDIR /app
COPY relay .

RUN go build -o /app/relay cmd/relay/main.go
CMD ["/app/relay"]
# FROM alpine:latest
# WORKDIR /
# COPY --from=build /app/relay /app/relay
# EXPOSE 50051
# CMD ["/app/relay"]
