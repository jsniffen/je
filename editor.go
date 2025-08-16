package main

type Mode int

const (
	ModeNormal Mode = iota
	ModeInsert      = iota
)

type Editor struct {
	Window *Window
	Mode   Mode
}

func NewEditor() *Editor {
	return &Editor{
		Window: NewWindow(),
		Mode:   ModeInsert,
	}
}

func (e *Editor) HandleEvent(ev Event) {
	if e.Mode == ModeNormal {
		switch ev.Type {
		case EventRawKey:
			if ev.Rune == 'i' {
				e.Mode = ModeInsert
				return
			}
			if ev.Rune == 'o' {
				e.Mode = ModeInsert
				e.Window.SwapFocus()
				return
			}
		}
	}

	if e.Mode == ModeInsert {
		switch ev.Type {
		case EventKey:
			if ev.Key == KeyEscape {
				e.Mode = ModeNormal
				return
			}
			e.Window.HandleEvent(ev)
		case EventRawKey:
			e.Window.HandleEvent(ev)
		}
	}
}
