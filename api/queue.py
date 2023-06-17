import nats
from api.models.events import Event

async def send(evt: Event):
    nc = await nats.connect("nats://nats:4222")
    
    print(f"Sending message to 'events.{evt.sub_id}' stream")
    
    js = nc.jetstream()
    streamid = f"events.{evt.sub_id}"
    result = await js.find_stream_name_by_subject(streamid)
    
    print(f"Stream {streamid} exists: {result}")
    
    await js.publish(
        subject=streamid, 
        payload=evt.text.encode(),
        stream="EVENTS",
    )