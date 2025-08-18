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
	e.OpenFile(s)
}

func (e *Editor) up() {
	if e.WindowFocus > 0 {
		e.WindowFocus -= 1
	}
}

func (e *Editor) down() {
	if e.WindowFocus < len(e.Columns[e.ColumnFocus].Windows)-1 {
		e.WindowFocus += 1
	}
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
			if ev.Rune == 'k' {
				e.up()
			}
			if ev.Rune == 'j' {
				e.down()
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
	e.Columns[0].Windows = append(e.Columns[0].Windows, w)
}

func (e *Editor) getActiveWindow() *Window {
	return e.Columns[e.ColumnFocus].Windows[e.WindowFocus]
}
