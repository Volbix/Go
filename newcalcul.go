package main

import (
	"fmt"
	"math"
)

type ResultatPuissance struct {
	Base     uint64
	Exposant int
	Resultat uint64
}

type Calculateur struct {
	Historique []ResultatPuissance // slice
	Cache      map[string]uint64   // map "base^exp" → résultat
}

func nouveauCalculateur() *Calculateur {
	return &Calculateur{
		Historique: []ResultatPuissance{},
		Cache:      make(map[string]uint64),
	}
}

// Pointer receiver : modifie le vrai Calculateur
func (c *Calculateur) puissance(base uint64, exp int) uint64 {
	cle := fmt.Sprintf("%d^%d", base, exp)

	// idiome map : val, ok := m[k]
	if val, ok := c.Cache[cle]; ok {
		fmt.Printf("(cache) %s = %d\n", cle, val)
		return val
	}

	resultat := uint64(math.Pow(float64(base), float64(exp)))
	c.Cache[cle] = resultat
	c.Historique = append(c.Historique, ResultatPuissance{base, exp, resultat})
	return resultat
}

func (c *Calculateur) afficherHistorique() {
	fmt.Println("=== Historique ===")
	for _, r := range c.Historique {
		fmt.Printf("%d^%d = %d\n", r.Base, r.Exposant, r.Resultat)
	}
}

// closure : fixe la base, retourne une fonction qui calcule base^exp
func creerPuissanceDe(base uint64) func(int) uint64 {
	return func(exp int) uint64 {
		return uint64(math.Pow(float64(base), float64(exp)))
	}
}

func main() {
	defer fmt.Println("Programme terminé.")

	calc := nouveauCalculateur()

	fmt.Println(calc.puissance(64, 8))
	fmt.Println(calc.puissance(2, 10))
	fmt.Println(calc.puissance(64, 8)) // depuis le cache

	// closure : puissances de 2
	puissanceDe2 := creerPuissanceDe(2)
	fmt.Println("=== Puissances de 2 ===")
	for i := 1; i <= 8; i++ {
		fmt.Printf("2^%d = %d\n", i, puissanceDe2(i))
	}

	calc.afficherHistorique()
}
