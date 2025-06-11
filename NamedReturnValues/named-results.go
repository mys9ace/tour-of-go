package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum - 1
	y = sum - x
	return // Named return values are returned implicitly
}

func main() {
	fmt.Println(split(17)) // This will print the values of x and y
}
