from fastapi import FastAPI
from fastapi import BackgroundTasks
from api.queue import send
from api.models.events import Event
from fastapi.middleware.cors import CORSMiddleware

app = FastAPI()

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

@app.get("/")
async def root():
    return {"message": "Hello World"}

@app.post("/send-notification")
async def send_notification(evt: Event, background_tasks: BackgroundTasks):
    background_tasks.add_task(send, evt=evt)
    return {"message": f"action {evt.action} sent to {evt.sub_id}"}

"""
curl -X POST "http://api.docker.localhost/send-notification" \
     -H "accept: application/json" \
     -H "Content-Type: application/json" \
     -d '{"text":"wassup","sub_id":"joe","action":"add"}' | jq
"""