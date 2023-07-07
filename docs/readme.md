# Project

This document describes how this project works.

> Note: This is all subject to change

## Idea

The idea of this project is quite similar to the idea of a realtime chat application.

1. User A performs an action (sends a message).
2. User B receives a notification about the action (receives the message).

## Event flow

Whenever a user subscribes to realtime notifications, it creates a queue in NATS and the relay opens a stream back to the browser.

```mermaid
sequenceDiagram
autonumber
actor Alice
actor Bob
participant Relay
participant Redis
participant PocketBase
participant Nats


Alice      ->>  PocketBase : Authenticate
Bob        ->>  PocketBase : Authenticate
Bob        ->>+ Relay      : Sends msg to "channel1"
Relay      ->>+ PocketBase : Verify Token
PocketBase ->>- Relay      : OK
Relay      ->>  Redis      : Cache message in Redis
Relay      ->>- Nats       : Publish msg to "channel1"
Note       over Nats       : Since there are no subscribers the msg is dropped
Alice      ->>+ Relay      : Subscribe to "channel1"
Relay      ->>+ PocketBase : Verify Token
PocketBase ->>- Relay      : OK
Relay      ->>+ Redis      : Get "channel1" history from last timestamp
Redis      ->>- Relay      : Bob's message
Relay      ->>  Alice      : Bob's message arrives at Alice
Relay      ->>- Nats       : Subscribe to "channel1"
Relay      ->>  Alice      : Open server stream

loop Server Stream
    Nats  -->+ Relay : Subscribe
    Bob   ->>  Relay : New message
    Relay ->>  Redis : Cache
    Relay ->>  Nats  : Publish
    Nats  -->> Relay : Msg
    Note  over Relay : Nats -> gRPC
    Relay -->> Alice : gRPC server stream
    Relay ->>- Nats  : Ack msg
end
```

## Data

The message payload has the following structure

```yaml
channelid: the NATS queue ID is the same as a channel
userid: where the message came from
text: the actual message
ts: timestamp in Unix Nano seconds
```

The relay creates a hardcoded group called `events` and each user that presses the subscribe button is the `subject` on that group.
So for user `bob` the queue is named `events.bob`.

Anyone can subscribe to that queue but since it's a queue only one instance will pick up the latest message.
