import nats
from api.models.events import Event

async def send(evt: Event):
    nc = await nats.connect("nats://nats:4222")
    js = nc.jetstream()
    
    await js.publish(
        stream="EVENTS",
        subject=f"events.{evt.subid}", 
        payload=evt.json().encode(),
    )