package coins

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func Start(reader *bufio.Reader) {
	for {
		fmt.Print("\033[2J\033[3J\033[H")
		fmt.Println(`
        █   █ █████ █   █ █   █      ███ █████ █████  ███  █   █         
        ██ ██ █     ██  █ █   █       █  █       █   █   █ ██  █         
████    █ █ █ ████  █ █ █ █   █       █  ████    █   █   █ █ █ █    ████ 
        █   █ █     █  ██ █   █    █  █  █       █   █   █ █  ██         
        █   █ █████ █   █  ███      ██   █████   █    ███  █   █         `)
		fmt.Println()
		fmt.Println("====================================")
		fmt.Println("1. Acheter des jetons")
		fmt.Println("2. Vendre des jetons")
		fmt.Println("3. Retour au menu")
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
			fmt.Println("Vous avez choisi d'acheter des jetons.")
			fmt.Println("Appuyez sur Entrée pour revenir au menu.")
			_, _ = reader.ReadString('\n')
			return
		case 2:
			fmt.Println("Vous avez choisi de vendre des jetons.")
			fmt.Println("Appuyez sur Entrée pour revenir au menu.")
			_, _ = reader.ReadString('\n')
			return
		case 3:
			fmt.Print("\033[2J\033[3J\033[H")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
