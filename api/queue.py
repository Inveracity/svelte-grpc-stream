import nats
from api.models.events import Event
from os import environ

NATS_HOST = environ.get('NATS_HOST', '127.0.0.1')
NATS_PORT = environ.get('NATS_PORT', '4222')

async def send(evt: Event):
    nc = await nats.connect(f"nats://{NATS_HOST}:{NATS_PORT}")
    js = nc.jetstream()
    
    await js.publish(
        stream="EVENTS",
        subject=f"events.{evt.subid}", 
        payload=evt.json().encode(),
    )