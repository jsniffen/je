package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const FONT_SIZE = 20

func drawStatusBar(e *Editor, font rl.Font, x, y int) {
	sw := rl.GetScreenWidth()
	sh := rl.GetScreenHeight()

	rl.DrawRectangle(int32(x), int32(y), int32(sw), int32(sh), rl.Blue)
	pos := rl.Vector2{float32(x + 10), float32(y + 10)}
	if e.Mode == ModeNormal {
		rl.DrawTextCodepoints(font, []rune("NORMAL"), pos, FONT_SIZE, 0, rl.White)
	} else if e.Mode == ModeInsert {
		rl.DrawTextCodepoints(font, []rune("INSERT"), pos, FONT_SIZE, 0, rl.White)
	}
}

func drawWindow(w *Window, font rl.Font, x, y int) {
	sw := rl.GetScreenWidth()
	sh := rl.GetScreenHeight()

	rl.DrawRectangle(int32(x), int32(y), int32(sw), int32(sh), rl.Violet)
	drawGapBuffer(w.Tag, font, x, y, !w.BodyFocused)

	y += 2 * FONT_SIZE
	rl.DrawRectangle(int32(x), int32(y), int32(sw), int32(sh), rl.Gray)
	drawGapBuffer(w.Body, font, x, y, w.BodyFocused)
}

func drawGapBuffer(gb *GapBuffer, font rl.Font, x0, y0 int, focused bool) {
	var x float32 = float32(x0 + 10)
	var y float32 = float32(y0 + 10)

	if gb.start == 0 && focused {
		rl.DrawRectangle(int32(x), int32(y), 3, FONT_SIZE, rl.White)
	}

	for i, r := range gb.Read() {
		if r == '\n' {
			y += FONT_SIZE
			x = float32(x0 + 10)
		} else {
			cp := int32(r)
			info := rl.GetGlyphInfo(font, cp)
			rl.DrawTextCodepoint(font, cp, rl.Vector2{x, y}, FONT_SIZE, rl.White)
			x += float32(info.AdvanceX)
		}

		if i == gb.start-1 && focused {
			rl.DrawRectangle(int32(x), int32(y), 3, FONT_SIZE, rl.White)
		}
	}
}

func main() {
	rl.InitWindow(800, 450, "je")
	defer rl.CloseWindow()

	rl.SetWindowState(rl.FlagVsyncHint)
	rl.SetExitKey(0)

	font := rl.LoadFontEx("fonts/Tamzen10x20r.ttf", FONT_SIZE, nil, 0)
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

		screenHeight := rl.GetScreenHeight()

		drawWindow(editor.Window, font, 0, 0)
		drawStatusBar(editor, font, 0, screenHeight-2*FONT_SIZE)

		rl.EndDrawing()
	}
}
