package interpreter

import (
	"bufio"
	"fmt"
	"io"
)

const tapeSize = 30_000

type Tape struct {
	input    string
	position int
	blocks   [tapeSize]byte
}

func New(input string) *Tape {
	return &Tape{
		input:    input,
		position: 0,
		blocks:   [tapeSize]byte{},
	}
}

func (t *Tape) Exec(in io.Reader, out io.Writer) {
	for i := 0; i < len(t.input); i++ {
		switch t.input[i] {
		case '>':
			t.position += 1
		case '<':
			t.position -= 1
		case '+':
			t.blocks[t.position] += 1
		case '-':
			t.blocks[t.position] -= 1
		case '.':
			fmt.Fprintln(out, string(t.blocks[t.position]))
		case ',':
			s := bufio.NewReader(in)

			b, err := s.ReadByte()

			if err != nil {
				fmt.Fprintln(out, "invalid input")

				return
			}

			t.blocks[t.position] = b
		case '[':
			if t.blocks[t.position] == 0 {
				for t.input[i] != ']' {
					i += 1
				}
			}
		case ']':
			if t.blocks[t.position] != 0 {
				for t.input[i] != '[' {
					i -= 1
				}
			}
		}
	}
}
