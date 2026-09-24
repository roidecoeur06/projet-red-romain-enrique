package overlay

import "fmt"

func PrintWalletOverlay(argent int, jetons int) {
	fmt.Print("\033[s")
	fmt.Printf("\033[18;75H\033[42m\033[30m ARGENT : %dK \033[0m", argent/1000)
	fmt.Printf("\033[19;75H\033[44m\033[30m JETONS : %d \033[0m", jetons)
	fmt.Print("\033[u")
}

func PrintFirstNameOverlay(firstName string) {
	fmt.Print("\033[s")
	fmt.Printf("\033[20;75H\033[46m\033[30m PRENOM : %-10s \033[0m", firstName)
	fmt.Print("\033[u")
}
