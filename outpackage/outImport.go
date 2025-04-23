package outImport

import (
	"fmt"
	// "go/types"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "objectchecker",
	Doc: "パッケージ外のオブジェクトを取得する",
	Run: run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, p := range pass.Pkg.Imports() {
		if p.Path() == "fmt" {
			obj := p.Scope().Lookup("Stringer")
			if obj != nil {
				fmt.Printf("fmt.Stringerが見つかりました %s (%T)\n", obj.Name(), obj)
			} else {
				fmt.Println("スコープ内でfmt.Stringerが見つかりませんでした")
			}
		}
	}
	return nil, nil
}