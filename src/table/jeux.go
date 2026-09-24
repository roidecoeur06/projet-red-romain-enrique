package table

import (
	"bufio"
	"fmt"
	"math/rand"
	"sort"
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
		waitAndClear(reader, "Appuyez sur Entrée pour revenir.")
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
	perfectPairsBet := askSideBet(reader, "Perfect Pairs", jetons)
	plusThreeBet := askSideBet(reader, "21+3", jetons)

	hand := Hand{Bet: mainBet, Cards: []Card{drawCard(&deck), drawCard(&deck)}}
	dealerCards := []Card{drawCard(&deck), drawCard(&deck)}

	fmt.Println("\n--- DISTRIBUTION ---")
	fmt.Printf("Main du croupier : %s et [Carte cachée]\n", printCard(dealerCards[0]))
	fmt.Printf("Votre main : %s %s (Score: %d)\n", printCard(hand.Cards[0]), printCard(hand.Cards[1]), calculateScore(hand.Cards))
	settleSideBets(hand.Cards, dealerCards[0], perfectPairsBet, plusThreeBet, jetons)

	playerBJ := calculateScore(hand.Cards) == 21
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
		waitAndClear(reader, "Appuyez sur Entrée pour continuer...")
		return
	}

	splitRequested := playHand(reader, &deck, &hand, jetons, true)
	hands := []Hand{hand}
	if splitRequested {
		*jetons -= hand.Bet
		hands = []Hand{
			{Bet: hand.Bet, Cards: []Card{hand.Cards[0], drawCard(&deck)}},
			{Bet: hand.Bet, Cards: []Card{hand.Cards[1], drawCard(&deck)}},
		}
		fmt.Println("\nLa main est divisée en deux.")
		for index := range hands {
			fmt.Printf("\n--- Main %d ---\n", index+1)
			playHand(reader, &deck, &hands[index], jetons, false)
		}
	}

	shouldPlayDealer := false
	for _, currentHand := range hands {
		if !currentHand.IsSurrender && calculateScore(currentHand.Cards) <= 21 {
			shouldPlayDealer = true
			break
		}
	}
	if shouldPlayDealer {
		for calculateScore(dealerCards) < 17 {
			dealerCards = append(dealerCards, drawCard(&deck))
		}
	}

	fmt.Printf("\nMain du croupier : ")
	for _, card := range dealerCards {
		fmt.Printf("%s ", printCard(card))
	}
	fmt.Printf("(Score: %d)\n", calculateScore(dealerCards))

	for index, currentHand := range hands {
		fmt.Printf("\nRésultat de la main %d :\n", index+1)
		settleHand(currentHand, dealerCards, jetons)
	}

	waitAndClear(reader, "Appuyez sur Entrée pour continuer...")
}

func settleHand(hand Hand, dealerCards []Card, jetons *int) {
	playerScore := calculateScore(hand.Cards)
	dealerScore := calculateScore(dealerCards)
	switch {
	case hand.IsSurrender:
		*jetons += hand.Bet / 2
		fmt.Printf("Abandon : vous récupérez %d jetons.\n", hand.Bet/2)
	case playerScore > 21:
		fmt.Println("Vous dépassez 21. Vous perdez votre mise.")
	case dealerScore > 21 || playerScore > dealerScore:
		*jetons += hand.Bet * 2
		fmt.Printf("Vous gagnez %d jetons !\n", hand.Bet)
	case playerScore == dealerScore:
		*jetons += hand.Bet
		fmt.Println("Égalité : votre mise est remboursée.")
	default:
		fmt.Println("Le croupier gagne. Vous perdez votre mise.")
	}
}

func waitAndClear(reader *bufio.Reader, message string) {
	fmt.Println(message)
	reader.ReadString('\n')
	fmt.Print("\033[2J\033[3J\033[H")
}

func askSideBet(reader *bufio.Reader, name string, jetons *int) int {
	if *jetons == 0 {
		return 0
	}

	fmt.Printf("Mise %s (0 pour passer, max %d) : ", name, *jetons)
	bet := readInt(reader)
	if bet < 0 || bet > *jetons {
		fmt.Println("Mise secondaire invalide, pari ignoré.")
		return 0
	}

	*jetons -= bet
	return bet
}

func settleSideBets(playerCards []Card, dealerVisible Card, perfectPairsBet int, plusThreeBet int, jetons *int) {
	if perfectPairsBet > 0 {
		name, multiplier := perfectPairsResult(playerCards[0], playerCards[1])
		if multiplier == 0 {
			fmt.Println("Perfect Pairs : perdu.")
		} else {
			*jetons += perfectPairsBet * (multiplier + 1)
			fmt.Printf("Perfect Pairs : %s, gain de %d jetons.\n", name, perfectPairsBet*multiplier)
		}
	}

	if plusThreeBet > 0 {
		name, multiplier := plusThreeResult(playerCards[0], playerCards[1], dealerVisible)
		if multiplier == 0 {
			fmt.Println("21+3 : perdu.")
		} else {
			*jetons += plusThreeBet * (multiplier + 1)
			fmt.Printf("21+3 : %s, gain de %d jetons.\n", name, plusThreeBet*multiplier)
		}
	}
}

func perfectPairsResult(first Card, second Card) (string, int) {
	if first.Value != second.Value {
		return "", 0
	}
	if first.Suit == second.Suit {
		return "paire parfaite", 25
	}
	if first.Color == second.Color {
		return "paire colorée", 12
	}
	return "paire mixte", 6
}

func plusThreeResult(first Card, second Card, dealer Card) (string, int) {
	cards := []Card{first, second, dealer}
	sameRank := cards[0].Rank == cards[1].Rank && cards[1].Rank == cards[2].Rank
	sameSuit := cards[0].Suit == cards[1].Suit && cards[1].Suit == cards[2].Suit
	straight := isStraight(cards)

	switch {
	case sameRank:
		return "brelan", 30
	case sameSuit && straight:
		return "quinte flush", 40
	case straight:
		return "suite", 10
	case sameSuit:
		return "couleur", 5
	default:
		return "", 0
	}
}

func isStraight(cards []Card) bool {
	values := make([]int, 0, len(cards))
	for _, card := range cards {
		value := card.Value
		switch card.Rank {
		case "V":
			value = 11
		case "D":
			value = 12
		case "R":
			value = 13
		case "A":
			value = 14
		}
		values = append(values, value)
	}

	sort.Ints(values)
	if values[0] == 2 && values[1] == 3 && values[2] == 14 {
		return true
	}
	return values[0]+1 == values[1] && values[1]+1 == values[2]
}

func playHand(reader *bufio.Reader, deck *[]Card, hand *Hand, jetons *int, allowSplit bool) bool {
	for {
		score := calculateScore(hand.Cards)
		if score > 21 {
			fmt.Printf("\nVotre score est de %d : vous dépassez 21.\n", score)
			return false
		}

		fmt.Printf("\nVotre main : ")
		for _, card := range hand.Cards {
			fmt.Printf("%s ", printCard(card))
		}
		fmt.Printf("(Score: %d)\n", score)
		fmt.Println("1. Tirer")
		fmt.Println("2. Rester")
		if !hand.IsDoubled {
			fmt.Println("3. Doubler")
			fmt.Println("4. Abandonner")
			if allowSplit && len(hand.Cards) == 2 && hand.Cards[0].Value == hand.Cards[1].Value {
				fmt.Println("5. Diviser (Split)")
			}
		}
		fmt.Print("Votre action : ")

		choice := readInt(reader)
		switch choice {
		case 1:
			hand.Cards = append(hand.Cards, drawCard(deck))
		case 2:
			return false
		case 3:
			if hand.IsDoubled {
				fmt.Println("Vous ne pouvez doubler qu'une seule fois.")
				continue
			}
			if *jetons < hand.Bet {
				fmt.Println("Vous n'avez pas assez de jetons pour doubler.")
				continue
			}
			*jetons -= hand.Bet
			hand.Bet *= 2
			hand.IsDoubled = true
			hand.Cards = append(hand.Cards, drawCard(deck))
			return false
		case 4:
			if len(hand.Cards) != 2 {
				fmt.Println("Vous ne pouvez abandonner qu'après la distribution initiale.")
				continue
			}
			hand.IsSurrender = true
			return false
		case 5:
			if !allowSplit || len(hand.Cards) != 2 || hand.Cards[0].Value != hand.Cards[1].Value {
				fmt.Println("Split impossible pour cette main.")
				continue
			}
			if *jetons < hand.Bet {
				fmt.Println("Vous n'avez pas assez de jetons pour diviser la main.")
				continue
			}
			return true
		default:
			fmt.Println("Action invalide.")
		}
	}
}
