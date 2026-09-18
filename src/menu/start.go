package menu

import (
	"bufio"
	"fmt"
	"os"
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

}
