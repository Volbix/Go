package main

import "fmt"

func main() {
	const (
		Nom          = "Francis"
		IMCMaigreur  = 18.5
		IMCNormal    = 25.0
		IMCSurpoids  = 30.0
	)

	var poids float64 = 95.5
	var taille float64 = 1.78

	imc := poids / (taille * taille)

	fmt.Printf("Bonjour %s !\n", Nom)
	fmt.Printf("IMC : %.2f\n", imc)

	var categorie string
	if imc < IMCMaigreur {
		categorie = "Maigreur"
	} else if imc < IMCNormal {
		categorie = "Normal"
	} else if imc < IMCSurpoids {
		categorie = "Surpoids"
	} else {
		categorie = "Obésité"
	}

	fmt.Printf("Catégorie : %s\n", categorie)
}
