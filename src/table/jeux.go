package table

import (
	"bufio"
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"

	"blackjack/src/cards"
	"blackjack/src/inventory"
	"blackjack/src/overlay"
	"blackjack/src/player"
)

type Card struct {
	Rank   string
	Suit   string
	Value  int
	Color  string
	Visual cards.Card
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
				deck = append(deck, Card{
					Rank:  rank,
					Suit:  suit,
					Value: val,
					Color: color,
					Visual: cards.Card{
						Suit:  cardSuitSymbol(suit),
						Value: rank,
					},
				})
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

func cardSuitSymbol(suit string) string {
	switch suit {
	case "Cœur":
		return "♥"
	case "Carreau":
		return "♦"
	case "Trèfle":
		return "♣"
	default:
		return "♠"
	}
}

func renderCards(hand []Card) string {
	if len(hand) == 0 {
		return ""
	}

	lines := make([]string, len(hand[0].Visual.GetASCII()))
	for _, card := range hand {
		ascii := card.Visual.GetASCII()
		for index, line := range ascii {
			lines[index] += line + "  "
		}
	}
	return strings.Join(lines, "\n")
}

func renderHiddenCard() string {
	return strings.Join((cards.Card{}).GetBackASCII(), "\n")
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
		overlay.PrintPVOverlay(player.GetPV())
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
			if player.IsKO() {
				fmt.Println("Vous êtes KO : vos PV sont à 0, vous ne pouvez plus jouer.")
				waitAndClear(reader, "Appuyez sur Entrée pour revenir.")
				continue
			}
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
	fmt.Print("\033[2J\033[3J\033[H")
	fmt.Println("========== PARTIE EN COURS ==========")
	overlay.PrintPVOverlay(player.GetPV())
	fmt.Println("\n--- INVENTAIRE ---")
	inventory.Inventaire(reader, *jetons)

	dealerReduction, p3Active := inventory.StartRound()
	player.AdvancePowerCooldown()
	if dealerReduction > 0 {
		fmt.Println("Potion P2 activee : le score du croupier est reduit de 2 points pour ce tour.")
	}

	if *jetons <= 0 {
		fmt.Println("Vous n'avez pas de jetons ! Allez en acheter au menu des Coins.")
		waitAndClear(reader, "Appuyez sur Entrée pour revenir.")
		return
	}

	deck := createDeck(6)

	fmt.Printf("Jetons disponibles : %d\n", *jetons)
	fmt.Printf("Mise principale (max %d) : ", *jetons)
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
	doubleReward := false
	bonusReward := 0

	fmt.Println("\n--- DISTRIBUTION ---")
	fmt.Printf("CROUPIER :\n%s\n%s\n", renderCards(dealerCards[:1]), renderHiddenCard())
	if player.CanUsePower() {
		fmt.Printf("Pouvoir disponible (%s). Utiliser maintenant ? (oui/non) : ", characterPowerName(player.GetCharacter()))
		answer, _ := reader.ReadString('\n')
		if strings.ToLower(strings.TrimSpace(answer)) == "oui" && player.UsePower() {
			switch player.GetCharacter() {
			case 1:
				fmt.Printf("Pouvoir du Mentaliste : la carte cachee est %s.\n", printCard(dealerCards[1]))
			case 2:
				doubleReward = true
				fmt.Println("Pouvoir du Richissime : vos gains de cette manche sont doubles.")
			case 3:
				bonusReward = 10
				fmt.Println("Pouvoir du Joueur normal : une victoire rapporte 10 jetons supplementaires.")
			}
		}
	} else if player.GetCharacter() != 0 {
		fmt.Printf("Pouvoir en recharge : encore %d manche(s).\n", player.PowerCooldown())
	}
	settleSideBets(hand.Cards, dealerCards[0], perfectPairsBet, plusThreeBet, jetons)

	playerBJ := calculateScore(hand.Cards) == 21
	dealerBJ := dealerScore(dealerCards, dealerReduction) == 21

	if playerBJ || dealerBJ {
		fmt.Printf("\nCROUPIER :\n%s\nTOTAL : %d\n", renderCards(dealerCards), dealerScore(dealerCards, dealerReduction))
		if playerBJ && !dealerBJ {
			fmt.Println("BLACKJACK NATUREL ! Vous gagnez 3:2 !")
			*jetons += int(float64(mainBet) * 2.5)
			*jetons += bonusReward
			player.ApplyRoundResult(true)
		} else if dealerBJ && !playerBJ {
			fmt.Println("Le Croupier a Blackjack. Vous perdez votre mise.")
			player.ApplyRoundResult(false)
		} else {
			fmt.Println("Égalité (Push) ! Les deux ont Blackjack.")
			*jetons += mainBet
		}
		waitAndClear(reader, "Appuyez sur Entrée pour continuer...")
		return
	}

	splitRequested := playHand(reader, &deck, &hand, jetons, true, p3Active)
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
			playHand(reader, &deck, &hands[index], jetons, false, false)
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
		for dealerScore(dealerCards, dealerReduction) < 17 {
			dealerCards = append(dealerCards, drawCard(&deck))
		}
	}

	fmt.Printf("\nCROUPIER :\n%s\nTOTAL : %d\n", renderCards(dealerCards), dealerScore(dealerCards, dealerReduction))

	for index, currentHand := range hands {
		fmt.Printf("\nRésultat de la main %d :\n", index+1)
		settleHand(currentHand, dealerCards, dealerReduction, jetons, doubleReward, bonusReward)
	}

	waitAndClear(reader, "Appuyez sur Entrée pour continuer...")
}

func dealerScore(cards []Card, reduction int) int {
	score := calculateScore(cards) - reduction
	if score < 0 {
		return 0
	}
	return score
}

func settleHand(hand Hand, dealerCards []Card, dealerReduction int, jetons *int, doubleReward bool, bonusReward int) {
	playerScore := calculateScore(hand.Cards)
	dealerTotal := dealerScore(dealerCards, dealerReduction)
	switch {
	case hand.IsSurrender:
		*jetons += hand.Bet / 2
		player.ApplyRoundResult(false)
		fmt.Printf("Abandon : vous récupérez %d jetons.\n", hand.Bet/2)
	case playerScore > 21:
		fmt.Println("Vous dépassez 21. Vous perdez votre mise.")
		player.ApplyRoundResult(false)
	case dealerTotal > 21 || playerScore > dealerTotal:
		payout := hand.Bet * 2
		if doubleReward {
			payout *= 2
		}
		*jetons += payout + bonusReward
		player.ApplyRoundResult(true)
		fmt.Printf("Vous gagnez %d jetons !\n", payout+bonusReward)
	case playerScore == dealerTotal:
		*jetons += hand.Bet
		fmt.Println("Égalité : votre mise est remboursée.")
	default:
		fmt.Println("Le croupier gagne. Vous perdez votre mise.")
		player.ApplyRoundResult(false)
	}
}

func characterPowerName(character int) string {
	switch character {
	case 1:
		return "Mentaliste"
	case 2:
		return "Richissime"
	case 3:
		return "Joueur normal"
	default:
		return "Inconnu"
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

func playHand(reader *bufio.Reader, deck *[]Card, hand *Hand, jetons *int, allowSplit bool, p3Active bool) bool {
	for {
		score := calculateScore(hand.Cards)
		if score > 21 {
			fmt.Printf("\nVotre score est de %d : vous dépassez 21.\n", score)
			return false
		}
		if p3Active && score == 20 {
			hand.Cards = append(hand.Cards, aceCard())
			p3Active = false
			fmt.Println("Potion P3 activee : un As est ajoute. Votre score passe a 21.")
			continue
		}

		fmt.Printf("\nVOS CARTES :\n%s\nTOTAL : %d\n", renderCards(hand.Cards), score)
		fmt.Println("\nQUE FAIRE ?")
		fmt.Println("1. Tirer une carte")
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

func aceCard() Card {
	return Card{
		Rank:  "A",
		Suit:  "Pique",
		Value: 11,
		Color: "Noir",
		Visual: cards.Card{
			Suit:  "♠",
			Value: "A",
		},
	}
}
