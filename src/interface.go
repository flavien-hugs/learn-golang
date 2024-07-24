// les interfaces sont un  ensembles de méthodes et un type utilsé
// permettant de définir le comportement d'autres types.
// Elle permet d'optimiser le code en réduisant la duplication, découpler les différences parties du code
// et optimiser les tests unitaires.

package main

import (
	"fmt"
)

// Définition d'une structure emploi
type emploi struct {
	ID    string
	Titre string
}

// Cette fonction prend en argument la structure 'emploi' la méthode 'String()' permet de récupérer
// permet les éléments de la structure 'emploi'
func (emp emploi) String() string {
	return fmt.Sprintf("L'emploi '%s' a pour identifiant '%s' .", emp.Titre, emp.ID)
}

func main() {
	emp := emploi{
		ID:    "DEV-GOLANG",
		Titre: "Développeur GO",
	}

	fmt.Println(emp.String())
}
