package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is even")
		}
	}

	table := []string{"Golang", "Java", "Python", "PHP", "JavaScript", "Java"}
	for index, value := range table {
		fmt.Println("index --> ", index, "value --> ", value)
	}
}
