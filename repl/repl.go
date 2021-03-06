package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/tehmantra/monkey/evaluator"
	"github.com/tehmantra/monkey/lexer"
	"github.com/tehmantra/monkey/object"
	"github.com/tehmantra/monkey/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) > 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			fmt.Fprint(out, evaluated.Inspect())
		}

		fmt.Fprintln(out)
	}
}

func printParserErrors(out io.Writer, errors []string) {
	fmt.Fprintln(out, "Woops! We ran into some monkey business here!")
	fmt.Fprintln(out, "Parser errors:")
	for _, msg := range errors {
		fmt.Fprintf(out, "\t%s\n", msg)
	}
}
