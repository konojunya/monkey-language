package repl

import (
	"bufio"
	"fmt"
	"github.com/konojunya/monkey-language/lexer"
	"github.com/konojunya/monkey-language/token"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for t := l.NextToken(); t.Type != token.EOF; t = l.NextToken() {
			fmt.Printf("%+v\n", t)
		}
	}
}
