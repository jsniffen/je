from pyray import *
from gap_buffer import GapBuffer

FONT_FILE = "c:\\windows\\fonts\\consola.ttf"

class Editor:
    def __init__(self):
        self.gap_buffer = GapBuffer(10)
        self.font = load_font(FONT_FILE)
        print(is_font_valid(self.font))

    def handle_input(self):
        while char := get_char_pressed():
            if char >= 32 and char <= 125:
                self.gap_buffer.insert(char)

        if is_key_pressed(KEY_ENTER):
            self.gap_buffer.insert('\n')

        if is_key_pressed(KEY_BACKSPACE):
            self.gap_buffer.delete()

    def render(self):
        x = y = 0
        padding = 10
        spacing = 5
        font_size = 32
        x += padding
        y += padding

        if len(self.gap_buffer.read()) == 0:
            draw_rectangle(x, y, 3, font_size, WHITE)

        for i, codepoint in enumerate(self.gap_buffer.read()):
            if codepoint == '\n':
                x = padding
                y += font_size
            else:
                info = get_glyph_info(self.font, codepoint)
                draw_text_codepoint(self.font, codepoint, Vector2(x, y), font_size, WHITE)
                x += info.advanceX

            if i == self.gap_buffer.start - 1:
                draw_rectangle(x, y, 3, font_size, WHITE)

init_window(800, 450, "Je")
editor = Editor()

while not window_should_close():
    editor.handle_input()

    begin_drawing()
    clear_background(BLACK)
    editor.render()
    end_drawing()

close_window()
