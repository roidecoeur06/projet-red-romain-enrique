package menu

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	coins "blackjack/src/jeton"
	"blackjack/src/overlay"
	"blackjack/src/player"
	"blackjack/src/regles"
	table "blackjack/src/table"
)

const title = `████████    ██            ██████      ██████    ██      ██      ██████    ██████      ██████    ██      ██
████████    ██            ██████      ██████    ██      ██      ██████    ██████      ██████    ██      ██
██      ██  ██          ██      ██  ██          ██    ██          ██    ██      ██  ██          ██    ██
██      ██  ██          ██      ██  ██          ██    ██          ██    ██      ██  ██          ██    ██
████████    ██          ██████████  ██          ██████            ██    ██████████  ██          ██████
████████    ██          ██████████  ██          ██████            ██    ██████████  ██          ██████
██      ██  ██          ██      ██  ██          ██    ██    ██    ██    ██      ██  ██          ██    ██
██      ██  ██          ██      ██  ██          ██    ██    ██    ██    ██      ██  ██          ██    ██
████████    ██████████  ██      ██    ██████    ██      ██    ████      ██      ██    ██████    ██      ██
████████    ██████████  ██      ██    ██████    ██      ██    ████      ██      ██    ██████    ██      ██`

func Start() {
	reader := bufio.NewReader(os.Stdin)

	for {
		clearTerminal()
		fmt.Printf("\033[32m%s\033[0m\n", title)
		fmt.Println()
		fmt.Println("1. \033[31mALLER A LA TABLE\033[0m")
		fmt.Println("2. \033[31mACHETER DES JETONS / VENDRE LES JETONS\033[0m")
		fmt.Println("3. \033[31mLES REGLES\033[0m")
		fmt.Println("4. \033[31mCHOIX DU PERSONNAGE\033[0m")
		fmt.Println("5. \033[31mQUITTER\033[0m")
		fmt.Println()
		wallet := coins.GetWalletPtr()
		overlay.PrintWalletOverlay(wallet.Argent, wallet.Jetons)
		overlay.PrintFirstNameOverlay(player.GetFirstName())
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
			table.Start(reader, &wallet.Jetons)
		case 2:
			coins.Start(reader)
		case 3:
			regles.ShowRules(reader)
		case 4:
			player.ChooseCharacter(reader)
		case 5:
			return
		default:
			fmt.Println("Cette fonctionnalite arrive bientot.")
		}
	}
}

func clearTerminal() {
	fmt.Print("\033[2J\033[3J\033[H")
}
