FROM golang:1.19-alpine AS build
WORKDIR /app
COPY internal internal
COPY cmd cmd
COPY go.mod go.mod
COPY go.sum go.sum

RUN go build -o /app/server cmd/server/main.go 

FROM alpine:latest
WORKDIR /
COPY --from=build /app/server /app/server
EXPOSE 50051
CMD ["/app/server"]
