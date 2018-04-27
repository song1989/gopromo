package main

import (
	"fmt"
)

func main() {
	start(NewB(C{}))
	start(NewB(D{}))
}

type A interface {
	what()
}

type B struct {
	A
}

type C struct {
}

func (b C) what() {
	fmt.Println("this is type C")
}

type D struct {
}

func (b D) what() {
	fmt.Println("this is type D")
}

func start(b B) {
	b.what()
}

func NewB(a A) B {
	return B{a}
}
