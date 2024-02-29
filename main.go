package main

import (
	"bufio"
	"fmt"
	"github.com/makarellav/gofck/interpreter"
	"os"
)

const PROMPT = "BrainFuck>> "

func main() {
	s := bufio.NewScanner(os.Stdin)

	for {
		fmt.Fprint(os.Stdout, PROMPT)

		scanned := s.Scan()

		if !scanned {
			return
		}

		input := s.Text()

		i := interpreter.New(input)

		i.Exec(os.Stdin, os.Stdout)
	}
}
