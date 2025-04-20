package example

import "fmt"

const Pi = 3.14

var Version = "1.0.0"

type Person struct {
	Name string
	Age int
}

func SayHello(name string) {
	fmt.Println("Hello", name)
}

func Add(a, b int) int {
	return a + b
}