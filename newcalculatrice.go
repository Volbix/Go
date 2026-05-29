package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

// Struct pour un calcul individuel
type Calcul struct {
	A, B     float64
	Op       string
	Resultat float64
}

// Embedding : Afficheur est inclus dans Calculatrice
// Calculatrice hérite automatiquement de la méthode Titre()
type Afficheur struct{}

func (a Afficheur) Titre(msg string) {
	fmt.Printf("=== %s ===\n", msg)
}

type Calculatrice struct {
	Afficheur                 // embedding
	Historique []Calcul       // slice
	Stats      map[string]int // map op → nb utilisations
}

func nouvelleCalculatrice() *Calculatrice {
	return &Calculatrice{
		Historique: []Calcul{},
		Stats:      make(map[string]int),
	}
}

// Pointer receiver : modifie la vraie Calculatrice (pas une copie)
func (c *Calculatrice) enregistrer(calc Calcul) {
	c.Historique = append(c.Historique, calc)
	c.Stats[calc.Op]++
}

func (c *Calculatrice) afficherHistorique() {
	c.Titre("Historique")
	if len(c.Historique) == 0 {
		fmt.Println("Aucun calcul effectué.")
		return
	}
	for i, calc := range c.Historique {
		fmt.Printf("%d. %.2f %s %.2f = %.2f\n", i+1, calc.A, calc.Op, calc.B, calc.Resultat)
	}
}

func (c *Calculatrice) afficherStats() {
	c.Titre("Stats")
	for op, count := range c.Stats {
		fmt.Printf("%s : %d fois\n", op, count)
	}
}

func operer(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("division par zéro impossible")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("opération inconnue : %s", op)
	}
}

func creerOperation(op string) func(float64, float64) float64 {
	return func(a, b float64) float64 {
		resultat, _ := operer(a, b, op)
		return resultat
	}
}

func main() {
	calc := nouvelleCalculatrice()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Calculatrice (a b op | 'hist' | 'stats' | 'quit')")

	for {
		fmt.Print("> ")
		scanner.Scan()
		ligne := scanner.Text()

		switch ligne {
		case "quit":
			fmt.Println("Au revoir !")
			return
		case "hist":
			calc.afficherHistorique()
		case "stats":
			calc.afficherStats()
		default:
			var a, b float64
			var op string
			_, err := fmt.Sscanf(ligne, "%f %f %s", &a, &b, &op)
			if err != nil {
				fmt.Println("Format invalide (ex: 10 5 +)")
				continue
			}
			resultat, err := operer(a, b, op)
			if err != nil {
				fmt.Println("Erreur :", err)
			} else {
				fmt.Printf("%.2f %s %.2f = %.2f\n", a, op, b, resultat)
				calc.enregistrer(Calcul{a, b, op, resultat})
			}
		}
	}
}
