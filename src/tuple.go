// NOTE: Un tuples est une séquence non modifiable de données ordonnées

package main

import "fmt"

func main() {
	fmt.Println(calculeSquareAndCudeNumber(2))

	m, n := calculeSquareAndCudeNumber(3)
	fmt.Println(m, n)

	m, n = n, m
	fmt.Println(m, n)
}

func calculeSquareAndCudeNumber(x int) (int, int) {
	return x * x, x * x * x
}
