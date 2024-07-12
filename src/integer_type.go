package main

import (
	"fmt"
	"time"
)

func main() {
	year := time.Now().Year()
	birthyear := 1993

	age := year - birthyear
	fmt.Println("YOUR YEAR OLD: ", age, "ans")
}
