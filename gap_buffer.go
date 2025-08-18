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
	//todo
}

func (gb *GapBuffer) Down() {
	//todo
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
