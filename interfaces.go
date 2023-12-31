package main

import (
	"fmt"
	"math"
)

type shape interface{
	area() float64
}

type rectangle struct{
	width float64
	height float64
}

type circle struct{
	radius float64
}

func (r rectangle) area() float64{
	return r.width * r.height
}

func (c circle) area() float64{
	return math.Pi * c.radius * c.radius	
}

func main(){
	r := rectangle{ 10,  20}
	c := circle{69}

	fmt.Println("Area of rectangle is: ", r.area())
	fmt.Println("Area of circle is: ", c.area())
}