package main

type Window struct {
	Tag         *GapBuffer
	Body        *GapBuffer
	BodyFocused bool
}

func NewWindow() *Window {
	return &Window{
		Tag:         NewGapBuffer(),
		Body:        NewGapBuffer(),
		BodyFocused: true,
	}
}

func (w *Window) SwapFocus() {
	w.BodyFocused = !w.BodyFocused
}

func (w *Window) HandleEvent(ev Event) {
	var gb *GapBuffer
	if w.BodyFocused {
		gb = w.Body
	} else {
		gb = w.Tag
	}

	switch ev.Type {
	case EventRawKey:
		gb.Insert(ev.Rune)
	case EventKey:
		switch ev.Key {
		case KeyBackspace:
			gb.Delete()
		case KeyEnter:
			gb.Insert('\n')
		case KeyLeft:
			gb.Left()
		case KeyRight:
			gb.Right()
		}
	}
}
