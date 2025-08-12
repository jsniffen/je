from editor import Editor
import pyray as pr

class Renderer:
    FONT_FILE = "c:\\windows\\fonts\\consola.ttf"

    def __init__(self) -> None:
        self.font = pr.load_font(self.FONT_FILE)

    def render(self, editor: Editor) -> None:
        pass
        # pr.begin_drawing()

        # pr.clear_background(pr.BLACK)
        # pr.draw_fps(0, 0)

        # x = y = 0
        # padding = 10
        # spacing = 5
        # font_size = 32
        # x += padding
        # y += padding

        # if len(editor.gap_buffer.read()) == 0:
        #     pr.draw_rectangle(x, y, 3, font_size, pr.WHITE)

        # for i, codepoint in enumerate(editor.gap_buffer.read()):
        #     if codepoint == '\n':
        #         x = padding
        #         y += font_size
        #     else:
        #         info = pr.get_glyph_info(self.font, codepoint)
        #         pr.draw_text_codepoint(self.font, codepoint, pr.Vector2(x, y), font_size, pr.WHITE)
        #         x += info.advanceX

        #     if i == editor.gap_buffer.start - 1:
        #         pr.draw_rectangle(x, y, 3, font_size, pr.WHITE)

        # pr.end_drawing()
