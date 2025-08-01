class GapBuffer:
    def __init__(self, size):
        self.gap_start = 0
        self.gap_size = size
        self.size = size
        self.buffer = [None]*size

    def insert(self, char):
        self.buffer[self.gap_start] = char
        self.gap_start += 1
        self.gap_size -= 1 

    def delete(self):
        self.gap_start -= 1
        self.gap_size += 1

    def __str__(self):
        pre = self.buffer[:self.gap_start]
        post = self.buffer[self.gap_start + self.gap_size:]
        return "".join(pre+post)
