package main

import "fmt"

func main() {
	var qty uint64 = 15
	var price = 19.65

	var cost = float32(float64(qty) * price)
	fmt.Println("PRODUCT COST --> ", cost)
}
