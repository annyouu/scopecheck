package fmttypeofchecker

import (
	"fmt"

	"github.com/gostaticanalysis/analysisutil"
	"golang.org/x/tools/go/analysis"
)
// Type.ofを取得する

// var Analyzer = &analysis.Analyzer{
// 	Name: "fmttypeofchecker",
// 	Doc: "TypeOfでfmtimportパスと名前で型を探す",
// 	Run: run,
// }

// func run(pass *analysis.Pass) (interface{}, error) {
// 	typ := analysisutil.TypeOf(pass, "fmt", "Stringer")
// 	if typ != nil {
// 		fmt.Printf("Found fmt.Stringer: %v\n", typ)
// 	} else {
// 		fmt.Println("fmt.Stringerが見つかりませんでした。")
// 	}
// 	return nil, nil
// }


// Obejctofを取得する

var Analyzer = &analysis.Analyzer{
	Name: "printcheck",
	Doc: "fmt.Printlnで、Objectofを見つける",
	Run: run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	obj := analysisutil.ObjectOf(pass, "fmt", "Println")
	if obj != nil {
		fmt.Printf("Found: %s: %T\n", obj.String(), obj)
	} else {
		fmt.Println("fmt.Println not found")
	}
	return nil, nil
}