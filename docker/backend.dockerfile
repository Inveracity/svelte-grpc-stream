FROM golang:1.19-alpine AS build
WORKDIR /app
COPY internal internal
COPY main.go main.go
COPY go.mod go.mod
COPY go.sum go.sum

RUN go build -o /app/notifier

FROM alpine:latest
WORKDIR /
COPY --from=build /app/notifier /app/notifier
EXPOSE 50051
CMD ["/app/notifier"]
