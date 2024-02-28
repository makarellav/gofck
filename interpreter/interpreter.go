package interpreter

const tapeSize = 30_000

type Interpreter struct {
	input    string
	position int
	tape     [tapeSize]byte
}

func New(input string) *Interpreter {
	return &Interpreter{
		input:    input,
		position: 0,
		tape:     [30000]byte{},
	}
}
