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

func PrintPVOverlay(pv int) {
	if pv < 0 {
		pv = 0
	}
	if pv > 100 {
		pv = 100
	}
	filled := pv / 10
	bar := ""
	for index := 0; index < 10; index++ {
		if index < filled {
			bar += "#"
		} else {
			bar += "-"
		}
	}

	fmt.Print("\033[s")
	fmt.Printf("\033[22;75H\033[41m\033[37m PV [%s] %3d/100 \033[0m", bar, pv)
	fmt.Print("\033[u")
}
