package table

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func Start(reader *bufio.Reader) {
	fmt.Print("\033[2J\033[3J\033[H")

	for {
		fmt.Println()
		fmt.Println("====================================")
		fmt.Println("         TABLE DE BLACKJACK         ")
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
			fmt.Println("Le jeu est en cours de developpement, mais la table est bien accessible.")
			fmt.Println("Appuyez sur Entree pour revenir au menu.")
			_, _ = reader.ReadString('\n')
			return
		case 2:
			fmt.Print("\033[2J\033[3J\033[H")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
