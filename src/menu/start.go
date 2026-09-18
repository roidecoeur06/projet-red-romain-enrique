package menu

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	fmt.Println(title)
	fmt.Println()
	fmt.Println()
	fmt.Println("1. ALLER A LA TABLE")
	fmt.Println("2. ACHETER DES JETONS / VENDRE LES JETONS")
	fmt.Println("3. LES REGLES")
	fmt.Println("4. QUITTER")
	fmt.Println()
	fmt.Print("QUE VOULEZ VOUS FAIRE : ")

	input, err := reader.ReadString('\n')
	if err != nil && len(input) == 0 {
		return
	}

	choice, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		fmt.Println("Choix invalide.")
		return
	}

	switch choice {
	case 1:
		fmt.Println("La table arrive bientôt.")
	case 2:
		fmt.Println("La gestion des jetons arrive bientôt.")
	case 3:
		fmt.Println("Les règles arrivent bientôt.")
	case 4:
		fmt.Println("A bientôt !")
	default:
		fmt.Println("Choix invalide.")
	}
}
