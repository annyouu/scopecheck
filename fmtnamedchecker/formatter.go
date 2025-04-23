package fmtnamedchecker

import (
	"fmt"

	"github.com/gostaticanalysis/analysisutil"
	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "fmtnamedchecker",
	Doc: "名前付き型と基底型を出力",
	Run: run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	typ := analysisutil.TypeOf(pass, "fmt", "Stringer")
	if typ == nil {
		fmt.Println("fmt.Stringerが見つかりませんでした")
		return nil, nil
	}

	fmt.Printf("Named type: %T\n", typ)
	fmt.Printf("Underlying(): %T\n", typ.Underlying())

	return nil, nil
}