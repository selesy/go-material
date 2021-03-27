package typescript

import (
	"fmt"
	"strings"
)

type Token struct {
	Row, Column int
	Bytes       []byte
	EOL         bool
	EOF         bool
}

func (t Token) String() string {
	return fmt.Sprintf("(Row: %d, Column: %d, EOL: %t, EOF: %t) %s", t.Row, t.Column, t.EOL, t.EOF, string(t.Bytes))
}

type Lexer struct {
	row, column int
	bytes       []byte
}

func NewLexer(bytes []byte) *Lexer {
	return &Lexer{
		row:    1,
		column: 1,
		bytes:  bytes,
	}
}

func (l *Lexer) Pop() (byte, bool) {
	if len(l.bytes) == 0 {
		return 0, true
	}

	b := l.bytes[0]
	l.bytes = l.bytes[1:]

	l.column++
	if b == '\n' {
		l.row++
		l.column = 1
	}

	return b, false
}

func (l *Lexer) Position() (row, column int) {
	return l.row, l.column
}

func (l *Lexer) Push(b byte) {
	l.bytes = append([]byte{b}, l.bytes...)
}

func (l *Lexer) Special() map[byte]bool {
	return map[byte]bool{'(': true, ')': true, '{': true, '}': true, '[': true, ']': true, '<': true, '>': true, ',': true}
}

func (l *Lexer) Token() *Token {
	row, column := l.Position()
	token := &Token{
		Row:    row,
		Column: column,
		Bytes:  nil,
		EOL:    false,
		EOF:    false,
	}

	b, eof := l.Pop()
	for !eof && !l.Whitespace()[b] {
		if l.Special()[b] && len(token.Bytes) != 0 {
			l.Push(b)

			break
		}

		token.Bytes = append(token.Bytes, b)

		if l.Special()[b] {
			break
		}

		b, eof = l.Pop()
	}

	if b == '\n' {
		token.EOL = true
	}

	token.EOF = eof

	token.Bytes = []byte(strings.TrimSpace(string(token.Bytes)))
	if !eof && len(token.Bytes) == 0 {
		return l.Token()
	}

	return token
}

func (l *Lexer) Whitespace() map[byte]bool {
	return map[byte]bool{'\n': true, '\t': true, ' ': true}
}
