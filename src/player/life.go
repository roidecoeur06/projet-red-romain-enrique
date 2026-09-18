package player

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func ChooseCharacter(reader *bufio.Reader) {
	clearTerminal()
	fmt.Println("CHOIX DU PERSONNAGE")
	fmt.Println()
	fmt.Println("1. LE STRATEGE")
	fmt.Println("2. LE BLUFFEUR")
	fmt.Println("3. LE RISK-TAKER")
	fmt.Println()
	fmt.Print("CHOISISSEZ VOTRE PERSONNAGE : ")

	input, err := reader.ReadString('\n')
	if err != nil && len(input) == 0 {
		return
	}

	choice, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil || choice < 1 || choice > 3 {
		fmt.Println("Choix invalide.")
		return
	}

	characters := []string{"LE STRATEGE", "LE BLUFFEUR", "LE RISK-TAKER"}
	fmt.Printf("Vous avez choisi %s.\n", characters[choice-1])
}

func clearTerminal() {
	fmt.Print("\033[2J\033[3J\033[H")
}
