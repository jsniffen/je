package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const FONT_SIZE = 24
const PADDING = FONT_SIZE / 2
const BAR_HEIGHT = 2*PADDING + FONT_SIZE

func drawCursor(mode Mode, x, y int) {
	switch mode {
	case ModeNormal:
		rl.DrawRectangle(int32(x), int32(y), 10, FONT_SIZE, rl.Orange)
	case ModeInsert:
		rl.DrawRectangle(int32(x), int32(y), 3, FONT_SIZE, rl.White)
	case ModeVisual:
		rl.DrawRectangle(int32(x), int32(y), 10, FONT_SIZE, rl.Violet)
	case ModeVisualLine:
		rl.DrawRectangle(int32(x), int32(y), 10, FONT_SIZE, rl.Violet)
	default:
		rl.DrawRectangle(int32(x), int32(y), 3, FONT_SIZE, rl.White)
	}
}

func drawWindow(w *Window, font rl.Font, x, y int, focused bool, mode Mode) {
	sw := rl.GetScreenWidth()
	sh := rl.GetScreenHeight()

	rl.DrawRectangle(int32(x), int32(y), int32(sw), int32(sh), rl.Blue)
	y += drawGapBuffer(w.Tag, font, x, y, !w.BodyFocused && focused, mode)
	y += FONT_SIZE + PADDING
	rl.DrawRectangle(int32(x), int32(y), int32(sw), int32(sh), rl.Black)
	drawGapBuffer(w.Body, font, x, y, w.BodyFocused && focused, mode)
}

func drawGapBuffer(gb *GapBuffer, font rl.Font, x0, y0 int, focused bool, mode Mode) int {
	x := x0 + PADDING
	y := y0 + PADDING

	if gb.start == 0 && focused {
		drawCursor(mode, x, y)
	}

	for i, r := range gb.Read() {
		if r == '\r' {
			continue
		}

		if r == '\n' {
			y += FONT_SIZE
			x = x0 + PADDING
		} else if r == '\t' {
			info := rl.GetGlyphInfo(font, ' ')
			w := int(4 * info.AdvanceX)
			if mode == ModeVisual && gb.InSelectionRange(i) && focused {
				rl.DrawRectangle(int32(x), int32(y), int32(w), FONT_SIZE, rl.Violet)
			}
			x += w
		} else {
			cp := int32(r)
			info := rl.GetGlyphInfo(font, cp)
			w := int(info.AdvanceX)
			if mode == ModeVisual && gb.InSelectionRange(i) && focused {
				rl.DrawRectangle(int32(x), int32(y), int32(w), FONT_SIZE, rl.Violet)
			}
			rl.DrawTextCodepoint(font, cp, rl.Vector2{float32(x), float32(y)}, FONT_SIZE, rl.White)
			x += w
		}

		if i == gb.start-1 && focused {
			drawCursor(mode, x, y)
		}
	}

	return y - y0
}

func main() {
	rl.SetWindowState(rl.FlagWindowHighdpi)

	rl.InitWindow(800, 450, "je")
	defer rl.CloseWindow()

	rl.SetWindowState(rl.FlagVsyncHint | rl.FlagWindowResizable)
	rl.SetExitKey(0)

	font := rl.LoadFontEx("fonts/Go-Mono.ttf", FONT_SIZE, nil, 0)
	if !rl.IsFontValid(font) {
		fmt.Println("Font failed to load")
		return
	}

	editor := NewEditor()
	for !rl.WindowShouldClose() {
		ctrlPressed := rl.IsKeyDown(rl.KeyLeftControl) || rl.IsKeyDown(rl.KeyRightControl)
		shiftPressed := rl.IsKeyDown(rl.KeyLeftShift) || rl.IsKeyDown(rl.KeyRightShift)
		for key := rl.KeyNull; key <= rl.KeyKpEqual; key += 1 {
			if !rl.IsKeyPressed(int32(key)) && !rl.IsKeyPressedRepeat(int32(key)) {
				continue
			}

			event := Event{
				CtrlPressed:  ctrlPressed,
				ShiftPressed: shiftPressed,
			}

			if key >= 32 && key <= 127 {
				event.Type = EventRawKey
				event.Rune = rune(key)
			} else {
				event.Type = EventKey
			}

			switch key {
			case rl.KeyBackspace:
				event.Key = KeyBackspace
			case rl.KeyEnter:
				event.Key = KeyEnter
			case rl.KeyUp:
				event.Key = KeyUp
			case rl.KeyDown:
				event.Key = KeyDown
			case rl.KeyLeft:
				event.Key = KeyLeft
			case rl.KeyRight:
				event.Key = KeyRight
			case rl.KeyEscape:
				event.Key = KeyEscape
			case rl.KeyTab:
				event.Key = KeyTab
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

		rl.EndDrawing()
	}
}
