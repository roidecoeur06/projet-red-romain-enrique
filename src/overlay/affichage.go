package overlay

import "fmt"

func formatMoney(argent int) string {
	millions := argent / 1000000
	thousands := (argent % 1000000) / 1000

	if millions > 0 {
		if thousands > 0 {
			return fmt.Sprintf("%dM %dK", millions, thousands)
		}
		return fmt.Sprintf("%dM", millions)
	}

	return fmt.Sprintf("%dK", argent/1000)
}

func PrintWalletOverlay(argent int, jetons int) {
	fmt.Print("\033[s")
	fmt.Printf("\033[18;75H\033[42m\033[30m ARGENT : %-10s \033[0m", formatMoney(argent))
	fmt.Printf("\033[19;75H\033[44m\033[30m JETONS : %d \033[0m", jetons)
	fmt.Print("\033[u")
}

func PrintFirstNameOverlay(firstName string) {
	fmt.Print("\033[s")
	fmt.Printf("\033[20;75H\033[46m\033[30m PRENOM : %-10s \033[0m", firstName)
	fmt.Print("\033[u")
}
