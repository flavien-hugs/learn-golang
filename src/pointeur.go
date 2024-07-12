package main

import "fmt"

func main() {
	m := 6
	decrement(&m)

	fmt.Println("m --> ", m)
	fmt.Println("afficher l'adresse mémoire de m (&m) --> ", &m)
}

func decrement(x *int) {
	*x--
}
