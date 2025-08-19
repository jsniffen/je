package main

import (
	"slices"
)

type GapBuffer struct {
	start int
	size  int
	buf   []rune
}

func NewGapBuffer() *GapBuffer {
	return &GapBuffer{
		start: 0,
		size:  1000,
		buf:   make([]rune, 1000, 1000),
	}
}

// If the size of the buffer is > 1MB
// then grow by a fixed size
func (gb *GapBuffer) grow() {
	newLen := 2 * len(gb.buf)
	if len(gb.buf) > 1_000_000 {
		newLen = len(gb.buf) + 1_000
	}

	buf := make([]rune, newLen)
	newSize := newLen - len(gb.buf) + gb.size

	copy(buf, gb.buf[:gb.start])
	copy(buf[gb.start+newSize:], gb.buf[gb.start+gb.size:])

	gb.size = newSize
	gb.buf = buf
}

func (gb *GapBuffer) Insert(r rune) {
	if gb.size <= 0 {
		gb.grow()
	}

	gb.buf[gb.start] = r
	gb.start += 1
	gb.size -= 1
}

func (gb *GapBuffer) Delete() {
	if gb.start <= 0 {
		return
	}

	gb.start -= 1
	gb.size += 1
}

func (gb *GapBuffer) Read() []rune {
	return slices.Concat(gb.buf[:gb.start], gb.buf[gb.start+gb.size:])
}

func (gb *GapBuffer) Up() {
	firstLine := true
	for i := 0; i < gb.start; i += 1 {
		if gb.buf[i] == '\n' {
			firstLine = false
			break
		}
	}

	if firstLine {
		return
	}

	moves := 0
	for gb.buf[gb.start-1] != '\n' {
		gb.Left()
		moves += 1
	}
	gb.Left()
	moves += 1

	lineLength := 0
	for i := gb.start - 1; i >= 0 && gb.buf[i] != '\n'; i -= 1 {
		lineLength += 1
	}

	for i := lineLength - moves; i >= 0; i -= 1 {
		gb.Left()
	}
}

func (gb *GapBuffer) Down() {
	lastLine := true
	for i := gb.start + gb.size; i < len(gb.buf); i += 1 {
		if gb.buf[i] == '\n' {
			lastLine = false
			break
		}
	}

	if lastLine {
		return
	}

	toLeftNewline := 0
	for i := gb.start - 1; i >= 0 && gb.buf[i] != '\n'; i -= 1 {
		toLeftNewline += 1

	}

	for gb.buf[gb.start+gb.size] != '\n' {
		gb.Right()
	}
	gb.Right()

	for i := gb.start + gb.size; i < len(gb.buf) && gb.buf[i] != '\n' && toLeftNewline > 0; i += 1 {
		gb.Right()
		toLeftNewline -= 1
	}
}

func (gb *GapBuffer) Left() {
	if gb.start > 0 {
		gb.buf[gb.start+gb.size-1] = gb.buf[gb.start-1]
		gb.start -= 1
	}
}

func (gb *GapBuffer) Right() {
	if gb.start+gb.size < len(gb.buf) {
		gb.buf[gb.start] = gb.buf[gb.start+gb.size]
		gb.start += 1
	}
}

// Read the space delimited runes under the cursor
func (gb *GapBuffer) ReadCursor() []rune {
	runes := gb.Read()
	pivot := gb.start - 1

	var start, end int

	for start = pivot; start > 0; start -= 1 {
		if runes[start] == '\n' ||
			runes[start] == ' ' ||
			runes[start] == '\t' {
			break
		}
	}

	if start > 0 {
		start += 1
	}

	for end = pivot; end < len(runes); end += 1 {
		if runes[end] == '\n' ||
			runes[end] == ' ' ||
			runes[end] == '\t' {
			break
		}
	}

	return runes[start:end]
}
