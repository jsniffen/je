from pyray import *
from gap_buffer import GapBuffer

gap_buffer = GapBuffer(10)

init_window(800, 450, "Je")

x = 190
y = 200

while not window_should_close():
    while key := get_key_pressed():
        if key == KEY_BACKSPACE:
            gap_buffer.delete()
        else:
            gap_buffer.insert(chr(key))

    begin_drawing()
    clear_background(WHITE)
    draw_text(str(gap_buffer), x, y, 20, VIOLET)
    end_drawing()

close_window()
