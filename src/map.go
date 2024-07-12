package main

import "fmt"

func main() {
	roads := make(map[int]string)
	fmt.Println("EMPTY MAP --> ", roads)

	// add data to map
	roads[1] = "Bingerville"
	roads[5] = "Cocody"

	fmt.Println("DISPLAY MAP WITH TWO VALUES --> ", roads)
	fmt.Println("MAP LENGHT --> ", len(roads))

	for key, value := range roads {
		fmt.Println("KEY ROAD: ", key, "VALUE ROAD: ", value)
	}

	// OR USING
	cites := map[int]string{
		1: "Yopougon",
		2: "Rivera 2",
	}
	fmt.Println("\nDISPLAY MAP WITH TWO VALUES --> ", cites)

	for key, value := range cites {
		fmt.Println("KEY CITIE: ", key, "VALUE CITIE: ", value)
	}
}
