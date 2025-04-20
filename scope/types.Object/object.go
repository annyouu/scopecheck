package object

import (
	"fmt"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "identlist",
	Doc: "スコープ内の識別子を取得する",
	Run: run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	scope := pass.Pkg.Scope()
	for _, name := range scope.Names() {
		obj := scope.Lookup(name)
		fmt.Printf("Name %s, Object: %T, Type: %s, Pos: %d\n", 
		name, obj, obj.Type(), obj.Pos())
	}
	return nil, nil
}
