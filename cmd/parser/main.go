package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	antlr "github.com/antlr4-go/antlr/v4"

	"interpreter/parser_src" // generated parser
	printer "interpreter/tree"
)

/* ---------------- Error Collector ---------------- */

type collectingErrorListener struct {
	*antlr.DefaultErrorListener
	Errors []string
}

func (l *collectingErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{},
	line, column int, msg string, e antlr.RecognitionException) {
	l.Errors = append(l.Errors, fmt.Sprintf("error at %d:%d: %s", line, column+1, msg))
}

/* ---------------- I/O Helpers ---------------- */

func readInput() ([]byte, string, error) {
	if len(os.Args) > 1 && os.Args[1] != "-" {
		p := os.Args[1]
		b, err := os.ReadFile(p)
		return b, p, err
	}
	b, err := io.ReadAll(bufio.NewReader(os.Stdin))
	return b, "-", err
}

func outName(srcPath, suffix string) string {
	if srcPath == "" || srcPath == "-" {
		return "stdin_" + suffix
	}
	base := filepath.Base(srcPath)
	base = strings.ReplaceAll(base, ".", "_")
	return base + "_" + suffix
}

/* ---------------- main ---------------- */

func main() {
	data, srcPath, err := readInput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "read error: %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stderr, "DEBUG: read %d bytes from %s\n", len(data), srcPath)

	input := antlr.NewInputStream(string(data))
	lexer := parser_src.NewJLangLexer(input)

	lexErr := &collectingErrorListener{}
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexErr)

	tokens := antlr.NewCommonTokenStream(lexer, 0)
	tokens.Fill()

	fmt.Println("TOKENS:")
	sym := lexer.GetSymbolicNames()
	lit := lexer.GetLiteralNames()
	for _, tk := range tokens.GetAllTokens() {
		tt := tk.GetTokenType()
		name := ""
		if tt >= 0 && tt < len(sym) && sym[tt] != "" {
			name = sym[tt]
		} else if tt >= 0 && tt < len(lit) && lit[tt] != "" {
			name = lit[tt]
		} else {
			name = fmt.Sprintf("T_%d", tt)
		}
		fmt.Printf("  %-18s %-12q @ %d:%d\n", name, tk.GetText(), tk.GetLine(), tk.GetColumn()+1)
	}

	input2 := antlr.NewInputStream(string(data))
	lexer2 := parser_src.NewJLangLexer(input2)

	lexErr2 := &collectingErrorListener{}
	lexer2.RemoveErrorListeners()
	lexer2.AddErrorListener(lexErr2)

	stream2 := antlr.NewCommonTokenStream(lexer2, 0)
	parser := parser_src.NewJLangParser(stream2)

	parErr := &collectingErrorListener{}
	parser.RemoveErrorListeners()
	parser.AddErrorListener(parErr)

	tree := parser.Program()

	ascii := printer.RenderASCII(tree, parser)
	fmt.Println("\nPARSE TREE:")
	fmt.Println(ascii)

	treeFile := outName(srcPath, "parsetree.txt")
	if err := os.WriteFile(treeFile, []byte(ascii), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "write tree error: %v\n", err)
	} else {
		fmt.Fprintf(os.Stderr, "wrote %s\n", treeFile)
	}

	if len(lexErr.Errors)+len(lexErr2.Errors)+len(parErr.Errors) > 0 {
		lines := append([]string{}, lexErr.Errors...)
		lines = append(lines, lexErr2.Errors...)
		lines = append(lines, parErr.Errors...)
		errText := strings.Join(lines, "\n") + "\n"
		errFile := outName(srcPath, "errors.txt")
		_ = os.WriteFile(errFile, []byte(errText), 0644)
		fmt.Fprintf(os.Stderr, "wrote %s\n", errFile)
	}
}
