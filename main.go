package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const FONT_SIZE = 24
const PADDING = FONT_SIZE / 2
const BAR_HEIGHT = 2*PADDING + FONT_SIZE

func drawCursor(mode Mode, x, y int) {
	if mode == ModeNormal {
		rl.DrawRectangle(int32(x), int32(y), 10, FONT_SIZE, rl.Orange)
	} else {
		rl.DrawRectangle(int32(x), int32(y), 3, FONT_SIZE, rl.White)
	}
}

func drawStatusBar(e *Editor, font rl.Font, x, y int) {
	sw := rl.GetScreenWidth()
	sh := rl.GetScreenHeight()

	rl.DrawRectangle(int32(x), int32(y), int32(sw), int32(sh), rl.Blue)
	pos := rl.Vector2{float32(x + PADDING), float32(y + PADDING)}
	if e.Mode == ModeNormal {
		rl.DrawTextCodepoints(font, []rune("NORMAL"), pos, FONT_SIZE, 0, rl.Orange)
	} else if e.Mode == ModeInsert {
		rl.DrawTextCodepoints(font, []rune("INSERT"), pos, FONT_SIZE, 0, rl.White)
	}
}

func drawWindow(w *Window, font rl.Font, x, y int, focused bool, mode Mode) {
	sw := rl.GetScreenWidth()
	sh := rl.GetScreenHeight()

	rl.DrawRectangle(int32(x), int32(y), int32(sw), int32(sh), rl.Violet)
	y += drawGapBuffer(w.Tag, font, x, y, !w.BodyFocused && focused, mode)
	y += FONT_SIZE + PADDING
	rl.DrawRectangle(int32(x), int32(y), int32(sw), int32(sh), rl.Gray)
	drawGapBuffer(w.Body, font, x, y, w.BodyFocused && focused, mode)
}

func drawGapBuffer(gb *GapBuffer, font rl.Font, x0, y0 int, focused bool, mode Mode) int {
	x := x0 + PADDING
	y := y0 + PADDING

	if gb.start == 0 && focused {
		drawCursor(mode, x, y)
	}

	for i, r := range gb.Read() {
		if r == '\n' {
			y += FONT_SIZE
			x = x0 + PADDING
		} else if r == '\t' {
			info := rl.GetGlyphInfo(font, ' ')
			x += int(4 * info.AdvanceX)
		} else {
			cp := int32(r)
			info := rl.GetGlyphInfo(font, cp)
			rl.DrawTextCodepoint(font, cp, rl.Vector2{float32(x), float32(y)}, FONT_SIZE, rl.White)
			x += int(info.AdvanceX)
		}

		if i == gb.start-1 && focused {
			drawCursor(mode, x, y)
		}
	}

	return y - y0
}

func main() {
	rl.InitWindow(800, 450, "je")
	defer rl.CloseWindow()

	rl.SetWindowState(rl.FlagVsyncHint | rl.FlagWindowResizable)
	rl.SetExitKey(0)

	font := rl.LoadFontEx("c:/windows/fonts/Go-Mono.ttf", FONT_SIZE, nil, 0)
	if !rl.IsFontValid(font) {
		fmt.Println("Font failed to load")
		return
	}

	editor := NewEditor()
	for !rl.WindowShouldClose() {
		for c := rl.GetCharPressed(); c != 0; c = rl.GetCharPressed() {
			event := Event{
				Type: EventRawKey,
				Rune: rune(c),
			}
			editor.HandleEvent(event)
		}

		for k := rl.GetKeyPressed(); k != 0; k = rl.GetKeyPressed() {
			var event Event
			switch k {
			case rl.KeyBackspace:
				event.Type = EventKey
				event.Key = KeyBackspace
			case rl.KeyEnter:
				event.Type = EventKey
				event.Key = KeyEnter
			case rl.KeyLeft:
				event.Type = EventKey
				event.Key = KeyLeft
			case rl.KeyRight:
				event.Type = EventKey
				event.Key = KeyRight
			case rl.KeyEscape:
				event.Type = EventKey
				event.Key = KeyEscape
			}
			editor.HandleEvent(event)
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		x, y := 0, 0

		screenHeight := rl.GetScreenHeight()
		colHeight := screenHeight - 2*BAR_HEIGHT

		for i, col := range editor.Columns {
			numWindows := len(col.Windows)
			winHeight := colHeight / numWindows

			for j, win := range col.Windows {
				focused := editor.ColumnFocus == i && editor.WindowFocus == j
				drawWindow(win, font, x, y, focused, editor.Mode)
				y += winHeight
			}
		}
		drawStatusBar(editor, font, 0, screenHeight-(2*PADDING+FONT_SIZE))

		rl.EndDrawing()
	}
}
