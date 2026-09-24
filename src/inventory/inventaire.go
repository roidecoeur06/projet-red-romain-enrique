package inventory

import (
	"bufio"
	"fmt"
)

func Inventaire(reader *bufio.Reader, firstName string) {
	fmt.Print("\033[2J\033[3J\033[H")
	fmt.Println("INVENTAIRE")
	fmt.Printf("PRENOM : %s\n", firstName)
	fmt.Println("Votre inventaire est vide.")
	fmt.Print("Appuyez sur Entrée pour revenir au menu... ")

	_, _ = reader.ReadString('\n')
}
