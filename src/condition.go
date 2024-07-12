package main

import "fmt"

func main() {
	var qty int = 89

	if qty >= 5 && qty < 6 {
		fmt.Println("--> Super")
	} else {
		fmt.Println("--> Très bien")
	}
}
