package main

import ("fmt")

func calc(x, y int) (out1, out2 int) {
	out1 = x + y
	out2 = x - y
	return
}
func main(){
	
	num1 := 5
	num2 := 4

	result1, result2 := calc(num1, num2)
	fmt.Println(result1, result2)
}

