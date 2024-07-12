package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Appel struct {
	Name      string  `json:"name,omitempty"`
	Telephone []Phone `json:"telephone,omitempty"`
}

type Phone struct {
	Portable bool   `json:"portable,omitempty"`
	Number   string `json:"number,omitempty"`
}

func main() {
	AppelsList := Appel{
		Name: "Flavien HUGS",
		Telephone: []Phone{
			{Portable: true, Number: "(225) 01 51 57 13 96"},
			{Portable: true, Number: "(225) 07 77 27 48 86"},
			{Portable: false, Number: "(225) 05 03 34 96"},
		},
	}

	directory := "data"

	// Create folder if not exist
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		err := os.Mkdir(directory, 0755)
		if err != nil {
			fmt.Println("Error --> ", err)
			return
		}
	}

	// Create json file
	file, err := json.MarshalIndent(AppelsList, "", " ")
	if err != nil {
		fmt.Println("Error --> ", err)
		return
	}
	filePath := filepath.Join(directory, "appels.json")

	err = os.WriteFile(filePath, file, 0755)
	if err != nil {
		fmt.Println("Error --> ", err)
		return
	}

	fmt.Println("Write data into file ", filePath, "successfully")

	fmt.Println("\nRead data into file ")
	diplayData(filePath)
}

func diplayData(filePath string) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error --> ", err)
		return
	}

	var appelsList Appel
	err = json.Unmarshal(file, &appelsList)
	if err != nil {
		fmt.Println("Error --> ", err)
		return
	}

	fmt.Printf("Content JSON file:\n%+v\n\n", appelsList)
}
