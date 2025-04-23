package fmtlookupchecker

import (
	"fmt"

	"github.com/gostaticanalysis/analysisutil"
	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "fmtlookupchecker",
	Doc: "インポートされたパッケージからの特定のオブジェクトを探すアナライザー",
	Run: run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	// 指定したimport path("fmt")とオブジェクト名("Stringer")を探し出す
	obj := analysisutil.LookupFromImports(pass.Pkg.Imports(), "fmt", "Stringer")
	if obj != nil {
		// 見つかったオブジェクトの名前と型を出力
		fmt.Printf("Found: %s: %s (%T)\n", "fmt.Stringer", obj.Name(), obj)
	} else {
		fmt.Println("fmt.Stringerが見つかりませんでした。")
	}
	return nil, nil
}