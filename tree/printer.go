package tree

import (
	"fmt"
	"strings"

	antlr "github.com/antlr4-go/antlr/v4"
)

func labelFor(node antlr.Tree, p antlr.Parser) string {
	switch n := node.(type) {
	case antlr.RuleNode:
		ruleIdx := n.GetRuleContext().GetRuleIndex()
		rules := p.GetRuleNames()
		if ruleIdx >= 0 && ruleIdx < len(rules) {
			return strings.ToUpper(rules[ruleIdx])
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

func RenderASCII(root antlr.Tree, p antlr.Parser) string {
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
