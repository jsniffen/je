from enum import Enum, auto

class Key(Enum):
    BACKSPACE = auto()
    ENTER = auto()
    LITERAL = auto()




@dataclass
class Event:
    key: Key | None = None
    raw: int | None = None
