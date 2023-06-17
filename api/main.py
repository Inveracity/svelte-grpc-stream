from fastapi import FastAPI
from queue import send

app = FastAPI()


@app.get("/")
async def root():
    await send(subscriptionId="test", msg="Hello World")
    return {"message": "sent message to queue"}