package main

import "fmt"

var box_channel = make(chan int, 2)

const WORKER_NUM = 10

func main() {
	go nwServer()
	for result := range box_channel {
		fmt.Println(result)
		if result == 9 {
			close(box_channel)
		}
	}
	fmt.Println("I am Naruto")
}
func nwServer() {
	for i := 0; i < WORKER_NUM; i++ {
		go nwRoutine(i)
	}
}
func nwRoutine(i int) {
	box_channel <- i
	fmt.Println("I am Naruto too")
}
