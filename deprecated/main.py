from editor import Editor
from event import Event
import pyray as pr

pr.init_window(800, 450, "Je")
pr.set_window_state(pr.ConfigFlags.FLAG_WINDOW_RESIZABLE | pr.ConfigFlags.FLAG_VSYNC_HINT)

if __name__ == "__main__":
    editor = Editor()

    while not pr.window_should_close():
        event = None

        while char := pr.get_char_pressed():
            if char >= 32 and char <= 125:
                event = Event(

        if pr.is_key_pressed(pr.KeyboardKey.KEY_ENTER):
            self.gap_buffer.insert('\n')

        if pr.is_key_pressed(pr.KeyboardKey.KEY_BACKSPACE):
            self.gap_buffer.delete()

    pr.close_window()
