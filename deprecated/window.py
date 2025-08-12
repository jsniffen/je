from gap_buffer import GapBuffer

class Buffer:
    def __init__(self) -> None:
        self.gap_buffer = GapBuffer()

class Window:
    def __init__(self) -> None:
        self.tag = Buffer()
        self.body = Buffer()

    def handle_input(self) -> None:
        pass
