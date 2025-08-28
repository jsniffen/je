package main

type EventType int
type Key int

const (
	EventNone EventType = iota
	EventRawKey
	EventKey
)

const (
	KeyNone Key = iota
	KeyBackspace
	KeyEnter
	KeyLeft
	KeyRight
	KeyUp
	KeyDown
	KeyEscape
	KeyTab
)

type Event struct {
	Repeat       bool
	Rune         rune
	Key          Key
	Type         EventType
	CtrlPressed  bool
	ShiftPressed bool
}

func (e Event) TranslateRawKey() rune {
	if e.Rune >= 65 && e.Rune <= 90 && !e.ShiftPressed {
		return e.Rune + 32
	}

	if e.ShiftPressed {
		switch e.Rune {
		case '1':
			return '!'
		case '2':
			return '@'
		case '3':
			return '#'
		case '4':
			return '$'
		case '5':
			return '%'
		case '6':
			return '^'
		case '7':
			return '&'
		case '8':
			return '*'
		case '9':
			return '('
		case '0':
			return ')'
		case '-':
			return '_'
		case '=':
			return '+'
		case '[':
			return '{'
		case ']':
			return '}'
		case '\\':
			return '|'
		case ';':
			return ':'
		case '\'':
			return '"'
		case ',':
			return '<'
		case '.':
			return '>'
		case '/':
			return '?'
		default:
			return e.Rune
		}
	}

	return e.Rune
}
