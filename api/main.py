from fastapi import FastAPI
from fastapi import BackgroundTasks
from pydantic import BaseModel

from api.queue import send

app = FastAPI()

class Message(BaseModel):
    destination: str
    msg: str

@app.get("/")
async def root():
    return {"message": "Hello World"}

@app.post("/send-notification/{subscriptionId}")
async def send_notification(subscriptionId: str, msg: Message, background_tasks: BackgroundTasks):
    background_tasks.add_task(send, subscriptionId=subscriptionId, msg="Hello World")
    send(subscriptionId=subscriptionId, msg=msg)
    return {"message": "notification sent"}


# curl -X POST "http://api.docker.localhost/send-notification/joe" -H "accept: application/json" -H "Content-Type: application/json" -d "{\"destination\":\"vito\",\"msg\":\"wassup\"}"