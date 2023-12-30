package main

import "fmt"


// Go does not allow any unused variables
// In Go we can return multiple values from a function
/*
func add(x int , y int) int { // This returns only one value }

func add(x int , y int) (int,int) { // This returns two values , To return more than one value we need to use paranthesis
	return x+y,x-y
}

Some time we need to skip some values we use "_" to skip the values

*/
// types of function in go

func concat( s1 string , s2 string)  string  {
	return s1+s2
}

/* short cut to declare function s in Go
	fun concat (s1,s2 string) string {
		return s1+s2
	}

	this will work same as above
*/


func increment(x int) int{
	x++
	return x
}

func main(){
	//fmt.Println("Ritik","Agarwal")

	x := 5
	x =increment(x)

	fmt.Println(x)
}