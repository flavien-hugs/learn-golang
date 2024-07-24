package main

import (
	"fmt"
	"math"
)

func verifyNumber(n int) bool {
	// cette fonction vérifie si un nombre est premier ou non

	k := math.Floor(float64(n/2 + 1))
	for i := 1; i <= int(k); i++ {
		if (n % i) == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Le nombre 5 est-il un nombre premier ? ", verifyNumber(5))
}
