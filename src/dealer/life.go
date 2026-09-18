package dealer

import (
	"fmt"
)

type Dealer struct {
	Hand  []string
	Score int
}

func NewDealer() *Dealer {
	return &Dealer{
		Hand:  []string{},
		Score: 0,
	}
}

func (d *Dealer) CalculateScore(hand []int) {
	score := 0
	aces := 0

	for _, cardValue := range hand {
		if cardValue == 11 { // As
			aces++
		}
		score += cardValue
	}
	for score > 21 && aces > 0 {
		score -= 10
		aces--
	}

	d.Score = score
}

func (d *Dealer) PlayTurn(drawCardFunc func() int) {
	fmt.Println("\n--- Tour du Croupier ---")

	currentHandValues := []int{}

	for d.Score < 17 {
		newCard := drawCardFunc()
		currentHandValues = append(currentHandValues, newCard)
		d.CalculateScore(currentHandValues)
		fmt.Printf("Le croupier a tiré une carte. Nouveau score du croupier : %d\n", d.Score)
	}

	if d.Score > 21 {
		fmt.Println("Le croupier a dépassé 21 (Bust) ! Le joueur gagne.")
	} else {
		fmt.Printf("Le croupier s'arrête avec un score de : %d\n", d.Score)
	}
}
