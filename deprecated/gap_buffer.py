class GapBuffer:
    def __init__(self, size=10):
        self.start = 0
        self.size = size
        self.buffer = [None]*size

    def move(self, position):
        pass

    def insert(self, char):
        if self.size == 0:
            self.buffer = (
                self.buffer[:self.start] + 
                [None]*10 + 
                self.buffer[self.start:]
            )
            self.size = 10

        self.buffer[self.start] = char
        self.start += 1
        self.size -= 1 
        print(self.buffer)

    def delete(self):
        if self.start == 0:
            return

        self.start -= 1
        self.size += 1

    def read(self):
        return (
            self.buffer[:self.start] +
            self.buffer[self.start + self.size:]
        )

    def __str__(self):
        return "".join(
            self.buffer[:self.start] +
            self.buffer[self.start + self.size:]
        )
