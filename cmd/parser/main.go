package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	antlr "github.com/antlr4-go/antlr/v4"

	"interpreter/gen"
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

/* ---------------- Pretty Tree Printer ---------------- */

func labelFor(node antlr.Tree, p antlr.Parser) string {
	switch n := node.(type) {
	case antlr.RuleNode:
		ruleIdx := n.GetRuleContext().GetRuleIndex()
		rules := p.GetRuleNames()
		if ruleIdx >= 0 && ruleIdx < len(rules) {
			return strings.ToUpper(rules[ruleIdx]) // 规则名大写
		}
	case antlr.TerminalNode:
		if tok := n.GetSymbol(); tok != nil {
			tt := tok.GetTokenType()
			sym := p.GetSymbolicNames()
			lit := p.GetLiteralNames()

			name := ""
			if tt >= 0 && tt < len(sym) && sym[tt] != "" {
				name = sym[tt]
			} else if tt >= 0 && tt < len(lit) && lit[tt] != "" {
				name = lit[tt]
			} else {
				name = fmt.Sprintf("T_%d", tt)
			}

			txt := tok.GetText()
			if txt == "" {
				txt = "<ε>"
			}
			return fmt.Sprintf("%s %q", name, txt)
		}
	}
	// 兜底
	return antlr.TreesStringTree(node, nil, p)
}

func childrenOf(t antlr.Tree) []antlr.Tree {
	n := t.GetChildCount()
	out := make([]antlr.Tree, 0, n)
	for i := 0; i < n; i++ {
		out = append(out, t.GetChild(i))
	}
	return out
}

func renderASCII(root antlr.Tree, p antlr.Parser) string {
	var b strings.Builder
	var walk func(node antlr.Tree, prefix string, isLast bool)

	walk = func(node antlr.Tree, prefix string, isLast bool) {
		conn := "├─ "
		next := prefix + "│  "
		if isLast {
			conn = "└─ "
			next = prefix + "   "
		}
		if prefix == "" {
			b.WriteString("ROOT\n")
			b.WriteString(conn + labelFor(node, p) + "\n")
		} else {
			b.WriteString(prefix + conn + labelFor(node, p) + "\n")
		}
		kids := childrenOf(node)
		for i, c := range kids {
			walk(c, next, i == len(kids)-1)
		}
	}
	walk(root, "", true)
	return b.String()
}

/* ---------------- main ---------------- */

func main() {

	data, srcPath, err := readInput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "read error: %v\n", err)
		os.Exit(1)
	}

	// 1) 词法：构建 lexer，收集词法错误，填充 token 流
	input := antlr.NewInputStream(string(data))
	lexer := gen.NewJLangLexer(input) // ✅ 这里参数类型 antlr.CharStream 正确

	lexErr := &collectingErrorListener{}
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexErr)

	tokens := antlr.NewCommonTokenStream(lexer, 0)
	tokens.Fill()

	// 打印 token 列表（类型、词素、位置）
	fmt.Println("TOKENS:")

	// 正确获取 Token 名：使用 GetSymbolicNames / GetLiteralNames
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

	// 2) 语法：复用 token 流给 parser（重置指针即可）
	tokens.Reset()
	parser := gen.NewJLangParser(tokens)

	parErr := &collectingErrorListener{}
	parser.RemoveErrorListeners()
	parser.AddErrorListener(parErr)

	tree := parser.Program() // 触发解析

	/* 3) 输出 ASCII 解析树到终端与文件 */
	ascii := renderASCII(tree, parser)
	fmt.Println("\nPARSE TREE:")
	fmt.Println(ascii)

	treeFile := outName(srcPath, "parsetree.txt")
	if err := os.WriteFile(treeFile, []byte(ascii), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "write tree error: %v\n", err)
	} else {
		fmt.Fprintf(os.Stderr, "wrote %s\n", treeFile)
	}

	/* 4) 若有错误，写 *_errors.txt */
	if len(lexErr.Errors)+len(parErr.Errors) > 0 {
		lines := append([]string{}, lexErr.Errors...)
		lines = append(lines, parErr.Errors...)
		errText := strings.Join(lines, "\n") + "\n"
		errFile := outName(srcPath, "errors.txt")
		_ = os.WriteFile(errFile, []byte(errText), 0644)
		fmt.Fprintf(os.Stderr, "wrote %s\n", errFile)
	}
}
