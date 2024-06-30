package ctxcause

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "ctxcause reports the use of context.WithCancel, context.WithDeadline, and context.WithTimeout"

// Analyzer is the struct that reports the use of context.WithCancel, context.WithDeadline, and context.WithTimeout
var Analyzer = &analysis.Analyzer{
	Name: "ctxcause",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

var m = map[string]struct{}{
	"context.WithCancel":   {},
	"context.WithDeadline": {},
	"context.WithTimeout":  {},
}

func run(pass *analysis.Pass) (any, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.SelectorExpr)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.SelectorExpr:
			x, ok := n.X.(*ast.Ident)
			if !ok {
				return
			}
			name := fmt.Sprintf("%s.%s", x.Name, n.Sel.Name)
			if _, ok := m[name]; ok {
				pass.Reportf(n.Pos(), "%s should be replaced with %sCause", name, name)
			}
		}
	})

	return nil, nil
}
