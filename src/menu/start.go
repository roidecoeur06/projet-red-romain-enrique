package menu

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"blackjack/src/player"
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

	fmt.Printf("\033[32m%s\033[0m\n", title)
	fmt.Println()
	fmt.Println("1. \033[31mALLER A LA TABLE\033[0m")
	fmt.Println("2. \033[31mACHETER DES JETONS / VENDRE LES JETONS\033[0m")
	fmt.Println("3. \033[31mLES REGLES\033[0m")
	fmt.Println("4. \033[31mCHOIX DU PERSONNAGE\033[0m")
	fmt.Println("5. \033[31mQUITTER\033[0m")
	fmt.Println()
	fmt.Print("QUE VOULEZ VOUS FAIRE : ")

	input, err := reader.ReadString('\n')
	if err != nil && len(input) == 0 {
		return
	}

	choice, err := strconv.Atoi(strings.TrimSpace(input))
	if err == nil && choice == 4 {
		player.ChooseCharacter(reader)
	}
}
