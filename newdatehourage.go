package main

import (
	"fmt"
	"time"
)

type Personne struct {
	Nom          string
	Anniversaire time.Time
}

// Value receiver : on ne modifie pas la personne
func (p Personne) Age() int {
	now := time.Now()
	age := now.Year() - p.Anniversaire.Year()
	if now.YearDay() < p.Anniversaire.YearDay() {
		age--
	}
	return age
}

func (p Personne) ProchainAnniversaire() time.Time {
	now := time.Now()
	prochain := time.Date(now.Year(), p.Anniversaire.Month(), p.Anniversaire.Day(), 0, 0, 0, 0, time.UTC)
	if prochain.Before(now) {
		prochain = prochain.AddDate(1, 0, 0)
	}
	return prochain
}

func (p Personne) Afficher() {
	fmt.Printf("%-10s | Né(e): %s | Âge: %d ans | Prochain anniv: %s\n",
		p.Nom,
		p.Anniversaire.Format("2006-01-02"),
		p.Age(),
		p.ProchainAnniversaire().Format("2006-01-02"),
	)
}

func main() {
	defer fmt.Println("Programme terminé.")

	now := time.Now()
	fmt.Printf("Date: %s | Heure: %s\n\n", now.Format("2006-01-02"), now.Format("15:04"))

	// slice de personnes
	personnes := []Personne{
		{"Francis", time.Date(1999, time.April, 18, 0, 0, 0, 0, time.UTC)},
		{"Alice", time.Date(2001, time.December, 3, 0, 0, 0, 0, time.UTC)},
		{"Bob", time.Date(1995, time.July, 22, 0, 0, 0, 0, time.UTC)},
	}

	// map mois → noms des personnes qui ont leur anniversaire ce mois
	parMois := make(map[time.Month][]string)
	for _, p := range personnes {
		mois := p.Anniversaire.Month()
		parMois[mois] = append(parMois[mois], p.Nom)
	}

	fmt.Println("=== Personnes ===")
	for _, p := range personnes {
		p.Afficher()
	}

	fmt.Println("\n=== Anniversaires par mois ===")
	for mois, noms := range parMois {
		fmt.Printf("%s : %v\n", mois, noms)
	}
}
