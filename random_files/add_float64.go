package main

import (
	"fmt"
)

func add(x, y float64) float64{
	return x+y
}

func main(){
	var num1 float64 = 5.6
	var num2 float64 = 9.5

	fmt.Println(add(num1,num2))


	num3, num4 := 3.2, 4.6

	fmt.Println(add(num3,num4))


}