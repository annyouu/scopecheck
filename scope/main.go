package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"
	// "analysis/scopetest1"
	// "analysis/scopetest"
	// "analysis/types.object"
	// "analysis/maxlen"
	"analysis/innserscope"
)

// func main() {
// 	singlechecker.Main(object.Analyzer)
// }


// func main() {
// 	singlechecker.Main(scopetest.Analyzer)
// }

// func main() {
// 	singlechecker.Main(limit.Analyzer)
// }

// inner
func main() {
	singlechecker.Main(inner.Analyzer)
}