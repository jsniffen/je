package main

type Editor struct {
	gb *GapBuffer
}

func NewEditor() *Editor {
	return &Editor{
		gb: NewGapBuffer(),
	}
}

func (e *Editor) HandleEvent(ev Event) {
	switch ev.Type {
	case EventRawKey:
		e.gb.Insert(ev.Rune)
	case EventKey:
		switch ev.Key {
		case KeyBackspace:
			e.gb.Delete()
		case KeyEnter:
			e.gb.Insert('\n')
		case KeyLeft:
			e.gb.Left()
		case KeyRight:
			e.gb.Right()
		}
	}
}
