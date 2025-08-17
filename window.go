package main

type Window struct {
	Tag         *GapBuffer
	Body        *GapBuffer
	BodyFocused bool
	y0          int
}

func NewWindow(tagContent, bodyContent string) *Window {
	tag := NewGapBuffer()
	for _, c := range tagContent {
		tag.Insert(c)
	}
	body := NewGapBuffer()
	for _, c := range bodyContent {
		body.Insert(c)
	}

	return &Window{
		Tag:         tag,
		Body:        body,
		BodyFocused: true,
		y0:          0,
	}
}

func (w *Window) getActiveGapBuffer() *GapBuffer {
	if w.BodyFocused {
		return w.Body
	} else {
		return w.Tag
	}
}

func (w *Window) ReadCursor() string {
	gb := w.getActiveGapBuffer()
	return string(gb.ReadCursor())
}

func (w *Window) SwapFocus() {
	w.BodyFocused = !w.BodyFocused
}

func (w *Window) HandleEvent(ev Event) {
	gb := w.getActiveGapBuffer()

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
