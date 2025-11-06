package tree

import (
	"fmt"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Node label helper: prefer rule name, fall back to token text/type.
func labelFor(node antlr.Tree, p antlr.Parser) string {
	switch n := node.(type) {
	case antlr.RuleNode:
		ruleIdx := n.GetRuleContext().GetRuleIndex()
		rules := p.GetRuleNames()
		if ruleIdx >= 0 && ruleIdx < len(rules) {
			return strings.ToUpper(rules[ruleIdx]) // 显示规则名（大写）
		}

	case antlr.TerminalNode:
		if tok := n.GetSymbol(); tok != nil {
			tt := tok.GetTokenType()

			// 新版 API：不再有 GetTokenNames()
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

	// 兜底：用 ANTLR 自带的字符串树
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

// RenderASCII renders the parse tree in a Git-like ASCII tree.
func RenderASCII(root antlr.Tree, p antlr.Parser) string {
	var b strings.Builder
	var walk func(node antlr.Tree, prefix string, isLast bool)
	walk = func(node antlr.Tree, prefix string, isLast bool) {
		conn := "├─ "
		nextPrefix := prefix + "│  "
		if isLast {
			conn = "└─ "
			nextPrefix = prefix + "   "
		}
		if prefix == "" { // root
			b.WriteString("ROOT\n")
		}
		if prefix != "" {
			b.WriteString(prefix + conn + labelFor(node, p) + "\n")
		} else {
			// After printing ROOT, print the first actual node
			b.WriteString(conn + labelFor(node, p) + "\n")
		}
		kids := childrenOf(node)
		for i, c := range kids {
			walk(c, nextPrefix, i == len(kids)-1)
		}
	}
	walk(root, "", true)
	return b.String()
}
