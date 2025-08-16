package main

type EventType int
type Key int

const (
	EventNone EventType = iota
	EventRawKey
	EventKey
)

const (
	KeyBackspace Key = iota
	KeyEnter
	KeyLeft
	KeyRight
	KeyUp
	KeyDown
	KeyEscape
)

type Event struct {
	Repeat bool
	Rune   rune
	Key    Key
	Type   EventType
}
