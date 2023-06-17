from pydantic import BaseModel

class Event(BaseModel):
    subid: str
    text: str
    sender: str
