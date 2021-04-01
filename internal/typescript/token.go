package typescript

import "fmt"

type Token struct {
	Row, Column int
	Bytes       []byte
	EOL         bool
	EOF         bool
}

func (t Token) String() string {
	return fmt.Sprintf("(Row: %d, Column: %d, EOL: %t, EOF: %t) %s", t.Row, t.Column, t.EOL, t.EOF, string(t.Bytes))
}
