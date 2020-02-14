package main

import (
	"math"
	"fmt"
)

func main(){

	var num float64 = 24

	var result = math.Sqrt(num)
	var intResult = math.Floor(result) 
	fmt.Printf("The output is %.2f Thank you", intResult)
}
