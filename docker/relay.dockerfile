FROM golang:1.19-alpine AS build
WORKDIR /app
COPY relay .

RUN go build -o /app/relay cmd/relay/main.go

FROM alpine:latest
WORKDIR /
COPY --from=build /app/relay /app/relay
EXPOSE 50051
CMD ["/app/relay"]
