package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"interpreter/parser_src"
	"io"
	"os"
	"path/filepath"
	"strings"

	antlr "github.com/antlr4-go/antlr/v4" //
)

// ---- Error collector ----

type collectingErrorListener struct {
	*antlr.DefaultErrorListener
	Errors []string
}

func (l *collectingErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{},
	line, column int, msg string, e antlr.RecognitionException) {
	l.Errors = append(l.Errors, fmt.Sprintf("error at %d:%d: %s", line, column+1, msg))
}

// ---- Output structures ----

type TokenOut struct {
	Type   string `json:"type"`
	Lexeme string `json:"lexeme"`
	Line   int    `json:"line"`
	Col    int    `json:"col"`
}

type Result struct {
	Tokens []TokenOut `json:"tokens"`
	Errors []string   `json:"errors"`
}

// ---- Helpers ----

func outputFileName(arg string) string {
	if arg == "" || arg == "-" {
		return "stdin_output.txt"
	}
	base := filepath.Base(arg)
	base = strings.ReplaceAll(base, ".", "_")
	return base + "_output.txt"
}

func readInput() ([]byte, string, error) {
	if len(os.Args) > 1 && os.Args[1] != "-" {
		p := os.Args[1]
		b, err := os.ReadFile(p)
		return b, p, err
	}
	b, err := io.ReadAll(bufio.NewReader(os.Stdin))
	return b, "-", err
}

// ---- main ----

func main() {
	data, srcPath, err := readInput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "read error: %v\n", err)
		os.Exit(1)
	}

	// 1) LEX: collect tokens and lexer errors
	input := antlr.NewInputStream(string(data))
	lexer := parser_src.NewJLangLexer(input)

	lexErr := &collectingErrorListener{}
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexErr)

	tokStream := antlr.NewCommonTokenStream(lexer, 0)
	tokStream.Fill()

	var toks []TokenOut
	for _, tk := range tokStream.GetAllTokens() {
		tt := tk.GetTokenType()
		// token type name
		var name string
		if tt >= 0 && tt < len(lexer.SymbolicNames) {
			name = lexer.SymbolicNames[tt]
		}
		if name == "" {
			name = fmt.Sprintf("T_%d", tt)
		}
		toks = append(toks, TokenOut{
			Type:   name,
			Lexeme: tk.GetText(),
			Line:   tk.GetLine(),
			Col:    tk.GetColumn() + 1,
		})
	}

	// 2) PARSE: build parse tree and collect parser errors
	// 重置输入流
	input = antlr.NewInputStream(string(data))
	lexer = parser_src.NewJLangLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := parser_src.NewJLangParser(stream)

	parErr := &collectingErrorListener{}
	parser.RemoveErrorListeners()
	parser.AddErrorListener(parErr)

	_ = parser.Program()

	// 3) Aggregate results
	res := Result{
		Tokens: toks,
		Errors: append(lexErr.Errors, parErr.Errors...),
	}

	// 4) JSON → stdout
	js, _ := json.MarshalIndent(res, "", "  ")
	fmt.Println(string(js))

	// 5) JSON → file
	outPath := outputFileName(srcPath)
	if err := os.WriteFile(outPath, js, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "write output file error: %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stderr, "wrote %s\n", outPath)
}
