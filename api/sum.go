package model

import "fmt"

//Sum sum of two integer
func Sum(x, y int) int {
	fmt.Println("Received Digits :", x, y)
	return x + y
}