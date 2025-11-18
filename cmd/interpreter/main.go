package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"interpreter/eval"
	gen "interpreter/parser_src"

	antlr "github.com/antlr4-go/antlr/v4"
)

/*************** Error Listener ****************/

type syntaxErr struct {
	Line, Col int
	Msg       string
}

type collectingErrListener struct {
	*antlr.DefaultErrorListener
	errs []syntaxErr
}

func (l *collectingErrListener) SyntaxError(
	recognizer antlr.Recognizer,
	offendingSymbol interface{},
	line, column int,
	msg string,
	e antlr.RecognitionException,
) {
	l.errs = append(l.errs, syntaxErr{Line: line, Col: column, Msg: msg})
}

func (l *collectingErrListener) HasErrors() bool { return len(l.errs) > 0 }

func (l *collectingErrListener) Print(w io.Writer) {
	for _, e := range l.errs {
		fmt.Fprintf(w, "error at %d:%d: %s\n", e.Line, e.Col, e.Msg)
	}
}

/*************** Utilities ****************/

func readAll(r io.Reader) (string, error) {
	b := bufio.NewScanner(r)
	b.Buffer(make([]byte, 0, 64*1024), 16*1024*1024) // up to 16MB per line
	var s string
	for b.Scan() {
		s += b.Text()
		s += "\n"
	}
	return s, b.Err()
}

/*************** main ****************/

func main() {
	var (
		code string
		err  error
	)
	if len(os.Args) > 1 {
		f, e := os.Open(os.Args[1])
		if e != nil {
			fmt.Fprintf(os.Stderr, "open %s: %v\n", os.Args[1], e)
			os.Exit(1)
		}
		defer f.Close()
		code, err = readAll(f)
	} else {
		code, err = readAll(os.Stdin)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "read source: %v\n", err)
		os.Exit(1)
	}

	input := antlr.NewInputStream(code)
	lexer := gen.NewJLangLexer(input)

	lexer.RemoveErrorListeners()
	lexErr := &collectingErrListener{}
	lexer.AddErrorListener(lexErr)

	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := gen.NewJLangParser(tokens)
	parser.BuildParseTrees = true

	parser.RemoveErrorListeners()
	parErr := &collectingErrListener{}
	parser.AddErrorListener(parErr)

	prog := parser.Program()

	if lexErr.HasErrors() || parErr.HasErrors() {
		lexErr.Print(os.Stderr)
		parErr.Print(os.Stderr)
		os.Exit(2)
	}

	interp := eval.NewInterpreter(parser)
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "runtime panic: %v\n", r)
			os.Exit(3)
		}
	}()
	interp.Run(prog)
}
