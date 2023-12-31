package main

import "fmt"

type rectangle struct {
	height int
	weight int
}

func (r rectangle) area() int {
	return r.height * r.weight
}


func main() {
	r := rectangle{169, 35}

	fmt.Println("Area of rectangle is: ", r.area())
}