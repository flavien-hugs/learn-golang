package main

import "fmt"

func main() {
	var cities []string
	cities = append(cities, "Django")
	cities = append(cities, "FastAPI")
	cities = append(cities, "Flask")

	fmt.Println("cities --> ", cities)
	fmt.Println("\n update cities")
	updateCities(cities)
	fmt.Println()
	fmt.Println(cities)
}

func updateCities(cities []string) {
	cities[0] = "AA"
	cities[1] = "BB"
	cities[2] = "CC"
}
