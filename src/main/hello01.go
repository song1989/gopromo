package main

import "fmt"

type Car struct {
	WheelCount int
}

func (car *Car) numberOfWheels() int {
	return car.WheelCount
}

type Ferrari struct {
	Car
}

func main() {
	f := Ferrari{Car{4}}
	fmt.Println(f.numberOfWheels())
}
