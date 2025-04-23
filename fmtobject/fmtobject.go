package fmtobject

import (
	"fmt"
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name: "fmtobjectchecker",
	Doc: "スコープの比較",
	Run: run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	// fmtパッケージを探す
	var fmtpkg *types.Package
	for _, p := range pass.Pkg.Imports() {
		if p.Path() == "fmt" {
			fmtpkg = p
		}
	}
	if fmtpkg == nil {
		return nil, nil
	}

	// fmtパッケージのスコープとの比較
	nodeFilter := []ast.Node{(*ast.Ident)(nil)}
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		id := n.(*ast.Ident)
		obj := pass.TypesInfo.ObjectOf(id)
		if obj != nil && obj.Parent() == fmtpkg.Scope() {
			fmt.Printf("fmt package object used: %s (%T)\n", obj.Name(), obj)
		}
	})
	return nil, nil
}