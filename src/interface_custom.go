// CUSTOM INTERFACE: Calculer de l'air et du volume de l'aire de forme géométrique

package main

import (
	"fmt"
	"math"
)

type volumeAreaInterface interface {
	aire() float64
	volume() float64
}

type cubeStructure struct {
	length float64
}

type sphereStructure struct {
	radius float64
}

func (cbs cubeStructure) aire() float64 {
	return cbs.length * cbs.length * 6
}

func (cbs cubeStructure) volume() float64 {
	return cbs.length * cbs.length * cbs.length
}

func (spr sphereStructure) aire() float64 {
	// Pour trouver l'aire d'un cercle, il faut multiplier le carré du rayon (r) par pi (π)
	return spr.radius * spr.radius * 4 * math.Pi
}

func (spr sphereStructure) volume() float64 {
	return spr.radius * spr.radius * 4 * math.Pi / 3
}

func mesure(vlm volumeAreaInterface) {
	fmt.Println(vlm)
	fmt.Println("L'aire est égale à => ", vlm.aire())
	fmt.Println("Le volume est égale à => ", vlm.volume())
}

func main() {
	cude := cubeStructure{length: 10}
	sphere := sphereStructure{radius: 4}

	fmt.Println("Mesure cude : ")
	mesure(cude)
	fmt.Println("Mesure sphére : ")
	mesure(sphere)
}
