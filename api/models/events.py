from pydantic import BaseModel

class Notification(BaseModel):
    channel_id: str
    user_id: str
    text: str
