package main

import "fmt"

func for_loops() {
	fmt.Println("For Loop implementation")
	fmt.Println("------------------------")
	for i:= 0; i<5;i++{
		fmt.Println(i);
	} 
}

func while_loop(){
	fmt.Println("While Loop implementation")
	fmt.Println("------------------------")

	i := 0
	for i < 5 {
		fmt.Println(i);
		i++
	}
}

func range_loops(){
	fmt.Println("Range Loop implementation")
	fmt.Println("-------------------------")
	//Lets make array
	numbers := [] int{1,2,3,4,5,6,7,8,9,10}

	for index,value := range numbers {
		fmt.Println("Index is ",index,"Value is ",value)
	}


}

func main(){
	for_loops()
	while_loop()
	range_loops()
}
