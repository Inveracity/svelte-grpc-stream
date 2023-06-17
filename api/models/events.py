from pydantic import BaseModel

from typing import Literal

class Event(BaseModel):
    text: str
    sub_id: str
    action: Literal["add", "delete", "update"] # TODO: remove this
