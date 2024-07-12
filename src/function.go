// anonyme function in go

package main

import "fmt"

func main() {
	// Fonction anonyme ou fermeture: Ce sont des fonctions qui ne necissite pas d'être nommer.
	// Elle est dite fonction définie inline

	funcAnonym := func() {
		fmt.Println("Hello, je suis une fonction anonyme !")
	}

	funcAnonym()

	// example:
	addition := func(x int, y int) int {
		return x + y
	}

	for n := 0; n < 3; n++ {
		fmt.Println("result --> ", addition(10, n))
	}
}
