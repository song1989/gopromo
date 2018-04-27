package main

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once

func main() {
	onceDo()

	//for i, v := range make([]string, 10) {
	//	once.Do(onces)
	//	fmt.Println("count:", v, "---", i)
	//}
	//for i := 0; i < 10; i++ {
	//
	//	go func() {
	//		once.Do(onced)
	//		fmt.Println("213")
	//	}()
	//}
	//time.Sleep(4000)
}
func onces() {
	fmt.Println("onces")
}
func onced() {
	fmt.Println("onced")
}

func onceDo() {
	var num int
	sign := make(chan bool)
	var once sync.Once
	f := func(ii int) func() {
		return func() {
			num = (num + ii*2)
			sign <- true
		}
	}
	for i := 0; i < 3; i++ {
		fi := f(i + 1)
		go once.Do(fi)
	}
	for j := 0; j < 3; j++ {
		select {
		case <-sign:
			fmt.Println("Received a signal.")
		case <-time.After(100 * time.Millisecond):
			fmt.Println("Timeout!")
		}
	}
	fmt.Printf("Num: %d.\n", num)
}
