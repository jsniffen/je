package main

import (
	"fmt"
	"io/ioutil"
)

type Mode int

const (
	ModeNormal Mode = iota
	ModeInsert      = iota
)

type Column struct {
	x0      int
	windows []*Window
}

type Editor struct {
	cols []*Column
	Mode Mode
}

func NewEditor() *Editor {
	return &Editor{
		cols: []*Column{{
			x0:      0,
			windows: []*Window{NewWindow("scratch", "")},
		}},
		Mode: ModeInsert,
	}
}

func (e *Editor) Execute(s string) {
	e.OpenFile(s)
}

func (e *Editor) HandleEvent(ev Event) {
	if e.Mode == ModeNormal {
		switch ev.Type {
		case EventKey:
			if ev.Key == KeyEnter {
				s := e.getActiveWindow().ReadCursor()
				e.Execute(s)
			}

		case EventRawKey:
			if ev.Rune == 'i' {
				e.Mode = ModeInsert
				return
			}
			if ev.Rune == 'o' {
				e.Mode = ModeInsert
				e.getActiveWindow().SwapFocus()
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
			e.getActiveWindow().HandleEvent(ev)
		case EventRawKey:
			e.getActiveWindow().HandleEvent(ev)
		}
	}
}

func (e *Editor) OpenFile(s string) {
	b, err := ioutil.ReadFile(s)
	if err != nil {
		fmt.Println(err)
		return
	}

	w := NewWindow(s, string(b))
	e.addWindow(w)
}

func (e *Editor) addWindow(w *Window) {
	e.cols[0].windows = append(e.cols[0].windows, w)
}

func (e *Editor) getActiveWindow() *Window {
	return e.cols[0].windows[0]
}
