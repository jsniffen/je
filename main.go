package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const FONT_SIZE = 20

func drawGapBuffer(gb *GapBuffer, font rl.Font) {
	var x float32 = 10
	var y float32 = 10

	if gb.start == 0 {
		rl.DrawRectangle(int32(x), int32(y), 3, FONT_SIZE, rl.White)
	}

	for i, r := range gb.Read() {
		if r == '\n' {
			y += FONT_SIZE
			x = 10
		} else {
			cp := int32(r)
			info := rl.GetGlyphInfo(font, cp)
			rl.DrawTextCodepoint(font, cp, rl.Vector2{x, y}, FONT_SIZE, rl.White)
			x += float32(info.AdvanceX)
		}

		if i == gb.start-1 {
			rl.DrawRectangle(int32(x), int32(y), 3, FONT_SIZE, rl.White)
		}
	}
}

func main() {
	rl.InitWindow(800, 450, "je")
	defer rl.CloseWindow()

	rl.SetWindowState(rl.FlagVsyncHint)

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
			}
			editor.HandleEvent(event)
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		drawGapBuffer(editor.gb, font)

		rl.EndDrawing()
	}
}
