package main

import "strings"

type Window struct {
	Tag         *GapBuffer
	Body        *GapBuffer
	BodyFocused bool
	y0          int
	Weight      float32
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

	for _ = range bodyContent {
		body.Left()
	}

	return &Window{
		Tag:         tag,
		Body:        body,
		BodyFocused: false,
		y0:          0,
		Weight:      1.0,
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

func (w *Window) GetFileName() string {
	runes := w.Tag.Read()
	i := strings.IndexRune(string(runes), ' ')
	return string(runes[:i])
}
