package main

import (
	"fmt"
)

func main() {
	// 型宣言によりfmt.Stringerがインポートされ、LookupFromImportsで検出される
	var _ fmt.Stringer
	fmt.Println("Hello world!")
}