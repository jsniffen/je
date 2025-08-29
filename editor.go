package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"slices"
)

type Mode int

const (
	ModeNormal Mode = iota
	ModeInsert
	ModeVisual
	ModeVisualLine
)

type Column struct {
	x0      int
	Windows []*Window
}

type Editor struct {
	Columns     []*Column
	ColumnFocus int
	WindowFocus int
	Mode        Mode
}

func NewEditor() *Editor {
	return &Editor{
		Columns: []*Column{{
			x0:      0,
			Windows: []*Window{NewWindow("scratch", "")},
		}},
		Mode: ModeInsert,
	}
}

func (e *Editor) Execute(s string) {
	switch s {
	case "put":
		e.Put()
	default:
		err := e.OpenFile(s)
		if err == nil {
			return
		}
		cmd := exec.Command("bash", "-c", s)
		output, err := cmd.CombinedOutput()
		w := NewWindow(s+"+", string(output))
		e.AddWindow(w)
	}
}

func (e *Editor) Up() {
	if e.WindowFocus > 0 {
		e.WindowFocus -= 1
	}
}

func (e *Editor) Down() {
	if e.WindowFocus < len(e.Columns[e.ColumnFocus].Windows)-1 {
		e.WindowFocus += 1
	}
}

func (e *Editor) HandleEvent(ev Event) {
	gb := e.GetActiveWindow().GetActiveGapBuffer()

	switch e.Mode {
	case ModeNormal:
		switch ev.Type {
		case EventKey:
			if ev.Key == KeyEnter {
				s := e.GetActiveWindow().ReadCursor()
				e.Execute(s)
			}

		case EventRawKey:
			switch ev.Rune {
			case 'Q':
				e.DeleteActiveWindow()
			case 'I':
				e.Mode = ModeInsert
			case 'O':
				e.Mode = ModeInsert
				e.GetActiveWindow().SwapFocus()
			case 'H':
				gb.Left()
			case 'L':
				gb.Right()
			case 'K':
				if ev.CtrlPressed {
					e.Up()
				} else {
					gb.Up()
				}
			case 'J':
				if ev.CtrlPressed {
					e.Down()
				} else {
					gb.Down()
				}
			case 'V':
				if ev.ShiftPressed {
					e.Mode = ModeVisualLine
				} else {
					e.Mode = ModeVisual
					e.GetActiveWindow().GetActiveGapBuffer().SelectBegin()
				}
			}
		}
	case ModeInsert:
		switch ev.Type {
		case EventKey:
			switch ev.Key {
			case KeyEscape:
				e.Mode = ModeNormal
			case KeyTab:
				gb.Insert('\t')
			case KeyBackspace:
				gb.Delete()
			case KeyEnter:
				gb.Insert('\n')
			case KeyUp:
				gb.Up()
			case KeyDown:
				gb.Down()
			case KeyLeft:
				gb.Left()
			case KeyRight:
				gb.Right()
			}
		case EventRawKey:
			r := ev.TranslateRawKey()
			gb.Insert(r)
		}
	case ModeVisual:
		switch ev.Type {
		case EventKey:
			switch ev.Key {
			case KeyEscape:
				e.Mode = ModeNormal
			case KeyEnter:
				s := string(e.GetActiveWindow().GetActiveGapBuffer().ReadSelection())
				e.Execute(s)
				e.Mode = ModeNormal
			}
		case EventRawKey:
			switch ev.Rune {
			case 'H':
				gb.Left()
			case 'L':
				gb.Right()
			case 'K':
				gb.Up()
			case 'J':
				gb.Down()
			}
		}
	case ModeVisualLine:
		switch ev.Type {
		case EventKey:
			switch ev.Key {
			case KeyEscape:
				e.Mode = ModeNormal
			}
		}
	}
}

func (e *Editor) OpenFile(s string) error {
	b, err := ioutil.ReadFile(s)
	if err != nil {
		return err
	}

	w := NewWindow(s, string(b))
	e.AddWindow(w)
	return nil
}

func (e *Editor) AddWindow(w *Window) {
	e.Columns[e.ColumnFocus].Windows = append(e.Columns[0].Windows, w)
}

func (e *Editor) DeleteActiveWindow() {
	if len(e.Columns[e.ColumnFocus].Windows) <= 1 {
		return
	}

	e.Columns[e.ColumnFocus].Windows = slices.Concat(
		e.Columns[e.ColumnFocus].Windows[:e.WindowFocus],
		e.Columns[e.ColumnFocus].Windows[e.WindowFocus+1:],
	)

	if e.WindowFocus > 0 {
		e.WindowFocus -= 1
	}
}

func (e *Editor) GetActiveWindow() *Window {
	return e.Columns[e.ColumnFocus].Windows[e.WindowFocus]
}

func (e *Editor) Put() error {
	w := e.GetActiveWindow()
	name := w.GetFileName()
	data := []byte(string(e.GetActiveWindow().Body.Read()))
	return os.WriteFile(name, data, 0666)
}
