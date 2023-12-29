package main

import "fmt"

//Variables Decleration in Go
//Short Variable Decleration
// var_name := value
//long way
//var var_name string (nobody use this)

//funnctions in Go

/*
printf is a formatted printing function that allows you to specify a format string with placeholders for values.
It supports format specifiers like %d for integers, %s for strings, %f for floating-point numbers, etc.
It requires you to provide as many values as there are placeholders in the format string.
*/

/*
println is a variadic function that takes a series of values and prints them to the standard output, followed by a newline character.
It automatically inserts spaces between the values.
It does not support format specifiers like %d, %s, etc.
*/
func Variables() {
	congrats := "u Failed"
	fmt.Println(congrats)
}

func float() {
	penny := 2.0
	const money = 55
	//this will print the datatype of the variable
	fmt.Printf("Data type is %T\n",penny)
	fmt.Println(money)
}

func multi_variables(){
	fmt.Print("Decalring Multi Variables in single line \n")
	first_name,last_name := "Ritik","Agarwal"
	fmt.Println(first_name,last_name)
}

func conversion() {
	fmt.Print("Conversion of Data Types \n")
	number := 88.6;
	fmt.Printf("Before conversion %T\n",number)
	cnvt := int(number)
	fmt.Printf("After Conversion %T\n",cnvt)
	fmt.Println(cnvt)
}
func main() {
	//Variables();
	float()
	//multi_variables()
	//conversion()
}
 