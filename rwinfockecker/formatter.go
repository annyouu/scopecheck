package rwinfochecker

import (
	"fmt"
	"go/types"

	"github.com/gostaticanalysis/analysisutil"
	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "rwinfochecker",
	Doc: "io.ReadWriterの埋め込みインターフェースを表示する",
	Run: run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	// io.ReadWriterの名前付き型を取得
	typ := analysisutil.TypeOf(pass, "io", "ReadWriter")
	if typ == nil {
		fmt.Println("io.ReadWriterが見つかりませんでした。")
		return nil, nil
	}

	iface, ok := typ.Underlying().(*types.Interface)
	if !ok {
		fmt.Printf("%T\n", typ.Underlying())
		return nil, nil
	}

	for i := 0; i < iface.NumEmbeddeds(); i++ {
		emb := iface.EmbeddedType(i)
		fmt.Println(emb.String())
	}
	return nil, nil
}