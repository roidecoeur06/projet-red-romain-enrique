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

var suits = []string{"Cœur", "Carreau", "Trèfle", "Pique"}
var ranks = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "V", "D", "R", "A"}

func createDeck(numDecks int) []Card {
	var deck []Card
	for i := 0; i < numDecks; i++ {
		for _, suit := range suits {
			color := "Noir"
			if suit == "Cœur" || suit == "Carreau" {
				color = "Rouge"
			}
			for _, rank := range ranks {
				val := 0
				if rank == "V" || rank == "D" || rank == "R" {
					val = 10
				} else if rank == "A" {
					val = 11
				} else {
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

func readInt(reader *bufio.Reader) int {
	input, _ := reader.ReadString('\n')
	val, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		return -1
	}
	return val
}

func Start(reader *bufio.Reader, jetons *int) {
	fmt.Print("\033[2J\033[3J\033[H")

	for {
		fmt.Println("\n====================================")
		fmt.Println("        TABLE DE BLACKJACK        ")
		fmt.Printf("        Vos jetons : %d\n", *jetons)
		fmt.Println("====================================")
		fmt.Println("1. Nouvelle partie")
		fmt.Println("2. Retour au menu")
		fmt.Println("====================================")
		fmt.Print("QUE VOULEZ VOUS FAIRE : ")

		choice := readInt(reader)

		switch choice {
		case 1:
			playBlackjack(reader, jetons)
		case 2:
			fmt.Print("\033[2J\033[3J\033[H")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func playBlackjack(reader *bufio.Reader, jetons *int) {
	if *jetons <= 0 {
		fmt.Println("Vous n'avez pas de jetons ! Allez en acheter au menu des Coins.")
		fmt.Println("Appuyez sur Entrée pour revenir.")
		reader.ReadString('\n')
		return
	}

	deck := createDeck(6)

	fmt.Printf("\nCombien de jetons voulez-vous miser ? (Max %d) : ", *jetons)
	mainBet := readInt(reader)
	if mainBet <= 0 || mainBet > *jetons {
		fmt.Println("Mise invalide. Retour au menu.")
		return
	}

	*jetons -= mainBet

	var playerHands []Hand
	playerHands = append(playerHands, Hand{Bet: mainBet, Cards: []Card{drawCard(&deck), drawCard(&deck)}})
	dealerCards := []Card{drawCard(&deck), drawCard(&deck)}

	fmt.Println("\n--- DISTRIBUTION ---")
	fmt.Printf("Main du croupier : %s et [Carte cachée]\n", printCard(dealerCards[0]))
	fmt.Printf("Votre main : %s %s (Score: %d)\n", printCard(playerHands[0].Cards[0]), printCard(playerHands[0].Cards[1]), calculateScore(playerHands[0].Cards))

	playerBJ := calculateScore(playerHands[0].Cards) == 21
	dealerBJ := calculateScore(dealerCards) == 21

	if playerBJ || dealerBJ {
		fmt.Printf("\nCarte cachée du croupier : %s (Score: %d)\n", printCard(dealerCards[1]), calculateScore(dealerCards))
		if playerBJ && !dealerBJ {
			fmt.Println("BLACKJACK NATUREL ! Vous gagnez 3:2 !")
			*jetons += int(float64(mainBet) * 2.5)
		} else if dealerBJ && !playerBJ {
			fmt.Println("Le Croupier a Blackjack. Vous perdez votre mise.")
		} else {
			fmt.Println("Égalité (Push) ! Les deux ont Blackjack.")
			*jetons += mainBet
		}
		fmt.Println("Appuyez sur Entrée pour continuer...")
		reader.ReadString('\n')
		return
	}

	fmt.Println("\n[La suite du jeu arrivera au prochain commit !]")
	fmt.Println("Pour l'instant, on vous rend votre mise de base.")
	*jetons += mainBet
	fmt.Println("Appuyez sur Entrée pour continuer...")
	reader.ReadString('\n')
}
