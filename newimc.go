package main

import "fmt"

const (
	IMCMaigreur = 18.5
	IMCNormal   = 25.0
	IMCSurpoids = 30.0
)

type Personne struct {
	Nom    string
	Poids  float64
	Taille float64
}

// Value receiver : calcul sans modifier la personne
func (p Personne) IMC() float64 {
	return p.Poids / (p.Taille * p.Taille)
}

func (p Personne) Categorie() string {
	// switch sans condition = if/else if en plus propre
	switch imc := p.IMC(); {
	case imc < IMCMaigreur:
		return "Maigreur"
	case imc < IMCNormal:
		return "Normal"
	case imc < IMCSurpoids:
		return "Surpoids"
	default:
		return "Obésité"
	}
}

func (p Personne) Afficher() {
	fmt.Printf("%-10s | IMC: %.2f | Catégorie: %s\n", p.Nom, p.IMC(), p.Categorie())
}

// Pointer receiver : modifie le poids de la personne
func (p *Personne) MettreAJourPoids(nouveauPoids float64) {
	p.Poids = nouveauPoids
}

func main() {
	defer fmt.Println("Programme terminé.")

	// slice de personnes
	personnes := []Personne{
		{"Francis", 95.5, 1.78},
		{"Alice", 55.0, 1.65},
		{"Bob", 110.0, 1.80},
	}

	// map catégorie → liste de noms
	parCategorie := make(map[string][]string)
	for _, p := range personnes {
		cat := p.Categorie()
		parCategorie[cat] = append(parCategorie[cat], p.Nom)
	}

	fmt.Println("=== IMC individuel ===")
	for _, p := range personnes {
		p.Afficher()
	}

	fmt.Println("\n=== Par catégorie ===")
	for cat, noms := range parCategorie {
		fmt.Printf("%s : %v\n", cat, noms)
	}

	// pointer receiver en action
	fmt.Println("\n=== Mise à jour poids Francis ===")
	personnes[0].MettreAJourPoids(80.0)
	personnes[0].Afficher()
}
