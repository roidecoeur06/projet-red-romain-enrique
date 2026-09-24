package table

import (
	"bufio"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Card struct {
	Rank  string
	Suit  string
	Value int
	Color string
}

type Hand struct {
	Cards       []Card
	Bet         int
	IsFinished  bool
	IsSurrender bool
	IsDoubled   bool
}

var suits = []string{"♥", "♦", "♣", "♠"}
var ranks = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}

func createDeck(numDecks int) []Card {
	var deck []Card
	for i := 0; i < numDecks; i++ {
		for _, suit := range suits {
			color := "Noir"
			if suit == "♥" || suit == "♦" {
				color = "Rouge"
			}
			for _, rank := range ranks {
				val := 0
				switch rank {
				case "J", "Q", "K":
					val = 10
				case "A":
					val = 11
				default:
					val, _ = strconv.Atoi(rank)
				}
				deck = append(deck, Card{Rank: rank, Suit: suit, Value: val, Color: color})
			}
		}
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })
	return deck
}

func drawCard(deck *[]Card) Card {
	card := (*deck)[0]
	*deck = (*deck)[1:]
	return card
}

func calculateScore(cards []Card) int {
	score := 0
	aces := 0
	for _, c := range cards {
		score += c.Value
		if c.Rank == "A" {
			aces++
		}
	}
	for score > 21 && aces > 0 {
		score -= 10
		aces--
	}
	return score
}

func printCard(c Card) string {
	return fmt.Sprintf("[%s de %s]", c.Rank, c.Suit)
}

func Start(reader *bufio.Reader) {
	fmt.Print("\033[2J\033[3J\033[H")

	for {
		fmt.Println()
		fmt.Println("====================================")
		fmt.Println("        TABLE DE BLACKJACK        ")
		fmt.Println("====================================")
		fmt.Println("1. Nouvelle partie")
		fmt.Println("2. Retour au menu")
		fmt.Println("====================================")
		fmt.Print("QUE VOULEZ VOUS FAIRE : ")

		input, err := reader.ReadString('\n')
		if err != nil && len(input) == 0 {
			return
		}

		choice, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			fmt.Println("Choix invalide.")
			continue
		}

		switch choice {
		case 1:
			fmt.Println("La partie de blackjack commence !")
			fmt.Println("Le jeu est en cours de développement, mais les structures de cartes sont prêtes en coulisses.")
			fmt.Println("Appuyez sur Entrée pour revenir au menu.")
			_, _ = reader.ReadString('\n')
		case 2:
			fmt.Print("\033[2J\033[3J\033[H")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
