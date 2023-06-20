import nats
from api.models.events import Notification
from os import environ
import asyncio

NATS_HOST = environ.get('NATS_HOST', '127.0.0.1')
NATS_PORT = environ.get('NATS_PORT', '4222')

async def send(notif: Notification):
    nc = await nats.connect(f"nats://{NATS_HOST}:{NATS_PORT}")
    js = nc.jetstream()
    
    await js.publish(
        stream="EVENTS",
        subject=f"events.{notif.channel_id}", 
        payload=notif.json().encode(),
    )

if __name__ == "__main__":
    asyncio.run(send(Notification(channel_id="channel1", user_id="user2", text="Hello World!")))
