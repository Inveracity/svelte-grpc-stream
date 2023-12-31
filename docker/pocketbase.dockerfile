FROM alpine:latest

ARG PB_VERSION=0.16.7

RUN apk add --no-cache \
  unzip \
  ca-certificates

# download and unzip PocketBase
ADD https://github.com/pocketbase/pocketbase/releases/download/v${PB_VERSION}/pocketbase_${PB_VERSION}_linux_amd64.zip /tmp/pb.zip
RUN unzip /tmp/pb.zip -d /pb/

COPY migrations /pb/pb_migrations

EXPOSE 8090

CMD [ "/pb/pocketbase", "serve", "--http=0.0.0.0:8090" ]

# docker build -t inveracity/pocketbase:0.16.7 -f docker/pocketbase.dockerfile .
# docker push inveracity/pocketbase:0.16.7
