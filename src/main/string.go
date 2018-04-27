package main

import "fmt"

func main() {
	var a interface{}
	var b string
	a = "asdasdasdasd"
	b = a.(string)
	fmt.Println(a, b)
}
