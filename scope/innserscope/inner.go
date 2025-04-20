package inner

import (
	"fmt"
	"go/ast"
	// "go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer {
	Name: "innerscope",
	Doc: "指定識別子の内側スコープを取得する",
	Run: run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	ins := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{&ast.Ident{}}

	ins.Preorder(nodeFilter, func(n ast.Node) {
		id, ok := n.(*ast.Ident)
		if !ok || id.Name != "gopher" {
			return
		}

		scope := pass.Pkg.Scope().Innermost(id.Pos())
		if scope == nil {
			return
		}

		fmt.Printf("Found gopher at %d\n", id.Pos())
		for _, name := range scope.Names() {
			obj := scope.Lookup(name)
			fmt.Printf("Name: %s, Type: %s\n", name, obj.Type())
		}
	})
	return nil, nil
}