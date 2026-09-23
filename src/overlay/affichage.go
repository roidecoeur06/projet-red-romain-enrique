package overlay

import "fmt"

const startingCash = 100000

func PrintCashOverlay() {
	fmt.Print("\033[s")
	fmt.Printf("\033[18;75H\033[42m\033[30m ARGENT : %dK \033[0m", startingCash/1000)
	fmt.Print("\033[u")
}
