package main

import (
	"fmt"
	"go/types"
	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "scopetracker",
	Doc: "子スコープを見つける",
	Run: run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	allScopes(pass.Pkg.Scope())
	return nil, nil
}

func allScopes(s *types.Scope) {
	fmt.Printf("Scope: %s, Objects: %d, Children: %s\n", s.String(), s.Len(), s.NumChildren())
	for i := 0; i < s.NumChildren(); i++ {
		allScopes(s.Child(i))
	}
}