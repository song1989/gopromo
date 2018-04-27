package main

import (
	"fmt"

	"github.com/speps/go-hashids"
)

func main() {
	hdata := NewData()
	hd := hashids.NewData()
	hd.Salt = "this is my salt"
	h := hashids.NewWithData(hd)
	id, _ := h.Encode([]int{1, 2, 3})
	numbers, _ := h.DecodeWithError(id)
	fmt.Println(numbers)
}
