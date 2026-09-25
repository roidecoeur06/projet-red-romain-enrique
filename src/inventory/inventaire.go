package inventory

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var backpackCapacity = 10
var potionP1 = 3
var potionP2Cooldown = 0
var potionP2Pending = false
var potionP3Charges = 2
var potionP3Cooldown = 0
var potionP3Pending = false
var backpackTimer *time.Timer
var roundNumber = 0

func Inventaire(reader *bufio.Reader, jetons int) {
	for {
		printInventory(jetons)
		fmt.Print("Voulez-vous utiliser une potion ? (oui/non) : ")
		answer, _ := reader.ReadString('\n')
		if strings.ToLower(strings.TrimSpace(answer)) != "oui" {
			return
		}

		fmt.Print("Quelle potion voulez-vous utiliser ? (1, 2 ou 3) : ")
		potionInput, _ := reader.ReadString('\n')
		potion, err := strconv.Atoi(strings.TrimSpace(potionInput))
		if err != nil || potion < 1 || potion > 3 {
			fmt.Println("Potion invalide. Appuyez sur Entrée pour continuer.")
			_, _ = reader.ReadString('\n')
			continue
		}

		usePotion(potion)
		fmt.Println("Appuyez sur Entrée pour continuer.")
		_, _ = reader.ReadString('\n')
	}
}

func printInventory(jetons int) {
	fmt.Print("\033[2J\033[3J\033[H")
	fmt.Println("+------------------------------------------------------+")
	fmt.Println("|              == INVENTAIRE BLACKJACK ==              |")
	fmt.Println("+------------------------------------------------------+")
	fmt.Printf("| JETONS: %-5d   MANCHE: %-3d SAC: 3/%-2d             |\n", jetons, roundNumber, backpackCapacity)
	fmt.Println("+------------------------------------------------------+")
	fmt.Printf("|                 SAC A DOS (%-2d slots)                 |\n", backpackCapacity)
	fmt.Println("|                                                      |")
	fmt.Println("|   +------+------+------+------+------+               |")
	fmt.Println("|   |  P1  |  P2  |  P3  |      |      |               |")
	fmt.Println("|   +------+------+------+------+------+               |")
	fmt.Println("|   |      |      |      |      |      |               |")
	fmt.Println("|   +------+------+------+------+------+               |")
	fmt.Println("+--------------------+------------------------+--------+")
	fmt.Println("| POTION             | EFFET                  | ETAT   |")
	fmt.Println("+--------------------+------------------------+--------+")
	fmt.Printf("| P1 Potion x%-3d     | Sac agrandi a 30 slots  | %-6s |\n", potionP1, potionState(potionP1))
	fmt.Printf("| P2 Potion Malus x%-1d   | Croupier -2 pendant 1 tour | %-4s |\n", potionP2Quantity(), potionP2State())
	fmt.Printf("| P3 Potion As x%-4d   | 20 devient 21          | %-6s |\n", potionP3Quantity(), potionP3State())
	fmt.Println("+--------------------+------------------------+--------+")
	fmt.Println("| Recharge : 5 manches apres chaque usage              |")
	fmt.Println("+------------------------------------------------------+")
}

func potionState(quantity int) string {
	if quantity == 0 {
		return "VIDE"
	}
	return "PRET"
}

func potionP2Quantity() int {
	if potionP2Cooldown > 0 {
		return 0
	}
	return 1
}

func potionP2State() string {
	if potionP2Cooldown > 0 {
		return "RECHARGE"
	}
	return "PRET"
}

func potionP3Quantity() int {
	if potionP3Cooldown > 0 {
		return 0
	}
	return potionP3Charges
}

func potionP3State() string {
	if potionP3Charges == 0 {
		return "VIDE"
	}
	if potionP3Cooldown > 0 {
		return "RECHARGE"
	}
	return "PRET"
}

func GetRoundNumber() int {
	return roundNumber
}

// StartRound advances potion cooldowns and returns the effects for this round.
func StartRound() (int, bool) {
	roundNumber++

	if potionP2Cooldown > 0 {
		potionP2Cooldown--
	}
	if potionP3Cooldown > 0 {
		potionP3Cooldown--
	}

	dealerReduction := 0
	if potionP2Pending {
		potionP2Pending = false
		dealerReduction = 2
	}

	p3Active := potionP3Pending
	potionP3Pending = false
	return dealerReduction, p3Active
}

func usePotion(potion int) {
	switch potion {
	case 1:
		if potionP1 == 0 {
			fmt.Println("Vous n'avez plus de Potion P1.")
			return
		}
		potionP1--
		backpackCapacity = 30
		if backpackTimer != nil {
			backpackTimer.Stop()
		}
		backpackTimer = time.AfterFunc(5*time.Minute, func() {
			backpackCapacity = 10
		})
		fmt.Println("Potion P1 utilisee : votre sac a dos contient maintenant 30 slots pendant 5 minutes.")
	case 2:
		if potionP2Cooldown > 0 {
			fmt.Printf("La Potion P2 est en recharge. Encore %d tours.\n", potionP2Cooldown)
			return
		}
		potionP2Pending = true
		potionP2Cooldown = 25
		fmt.Println("Potion P2 utilisee : le croupier aura -2 points pendant le prochain tour.")
	case 3:
		if potionP3Charges == 0 {
			fmt.Println("Vous n'avez plus de charges pour la Potion P3.")
			return
		}
		if potionP3Cooldown > 0 {
			fmt.Printf("La Potion P3 est en recharge. Encore %d tours.\n", potionP3Cooldown)
			return
		}
		potionP3Charges--
		potionP3Pending = true
		potionP3Cooldown = 30
		fmt.Println("Potion P3 utilisee : si votre main atteint 20, un As sera ajoute pour faire 21.")
	}
}
