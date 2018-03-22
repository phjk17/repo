package main

import (
	"fmt"
)

func main() {
	var i,j int
	fmt.Scan(&i)
	for k := 0; k < i; k++ {
		fmt.Scan(&j)
		square := make([]int, j)
		for n := 0; n < j; n++ {
			fmt.Scan(&square[n])
		}
		fmt.Println(calculateSquareSums(square))
	}
}

func calculateSquareSums (values[]int)int {
	if(len(values)<=0){
		return 0
	}
	if(values[0] <= 0){
		return calculateSquareSums(values[1:len(values)])
	}
	return values[0]*values[0] + calculateSquareSums(values[1:len(values)])
}