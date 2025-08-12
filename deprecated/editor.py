from window import Window
import pyray as pr

class Editor:
    def __init__(self) -> None:
        self.windows = [Window()]

    def handle_input(self) -> None:
        pass
        # while char := pr.get_char_pressed():
        #     if char >= 32 and char <= 125:
        #         self.gap_buffer.insert(char)

        # if pr.is_key_pressed(pr.KeyboardKey.KEY_ENTER):
        #     self.gap_buffer.insert('\n')

        # if pr.is_key_pressed(pr.KeyboardKey.KEY_BACKSPACE):
        #     self.gap_buffer.delete()
