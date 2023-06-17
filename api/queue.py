import asyncio
import nats

async def send(subscriptionId: str, msg: str):
    nc = await nats.connect("nats://127.0.0.1:4222")
    print(f"Sending message to 'events.{subscriptionId}' stream")
    js = nc.jetstream()
    streamid = f"events.{subscriptionId}"
    result = await js.find_stream_name_by_subject(streamid)
    print(f"Stream {streamid} exists: {result}")
    await js.publish(
        subject=streamid, 
        payload=msg.encode(),
        stream="EVENTS",
    )

if __name__ == '__main__':
    asyncio.run(send(subscriptionId="joe", msg="Hello World"))