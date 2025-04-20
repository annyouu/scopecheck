package limit

// import (
// 	"fmt"
// 	"go/types"

// 	"golang.org/x/tools/go/analysis"
// 	"golang.org/x/tools/go/analysis/passes/inspect"
// )

// var Analyzer = &analysis.Analyzer{
// 	Name: "shortvar",
// 	Doc: "一文字変数を検出する",
// 	Run: run,
// 	Requires: []*analysis.Analyzer{
// 		inspect.Analyzer,
// 	},
// }

// func run(pass *analysis.Pass) (interface{}, error) {
// 	for _, n := range pass.Pkg.Scope().Names() {
// 		if len(n) != 1 {
// 			continue
// 		}
// 		obj, _ := pass.Pkg.Scope().Lookup(n).(*types.Var)
// 		if obj == nil {
// 			continue
// 		}
// 		fmt.Println(n, obj)
// 	}
// 	return nil, nil
// }

import (
	"flag"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

var (
	// 最大許容量をフラグで指定
	maxLen int
)

func init() {
	// Analyzer.Flagsでフラグを登録
	Analyzer.Flags.Init("shortvar", flag.ExitOnError)
	Analyzer.Flags.IntVar(&maxLen, "maxlen", 1, "report")
}

var Analyzer = &analysis.Analyzer{
	Name:     "shortvar",
	Doc:      "名前の短いパッケージ変数を探す",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	scope := pass.Pkg.Scope()

	for _, name := range scope.Names() {
		// 名前の長さがmaxLenより大きければ飛ばす
		if len(name) > maxLen {
			continue
		}

		// types.Varかどうかをチェックする
		if obj, ok := scope.Lookup(name).(*types.Var); ok {
			pass.Reportf(obj.Pos(), "パッケージ変数名%sが、max以下です: %d 長さ: %d ",
				name, maxLen, len(name),
			)
		}
	}
	return nil, nil
}
