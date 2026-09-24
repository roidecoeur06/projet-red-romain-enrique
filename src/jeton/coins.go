package coins

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

const StartingMoney = 100000

var wallet = Wallet{Argent: StartingMoney, Jetons: 0}

type Wallet struct {
	Argent int
	Jetons int
}

func GetWallet() Wallet {
	return wallet
}

func (w *Wallet) AcheterJetons(nombre int) bool {
	if nombre <= 0 {
		return false
	}
	if w.Argent < nombre {
		return false
	}
	w.Argent -= nombre
	w.Jetons += nombre
	return true
}

func (w *Wallet) VendreJetons(nombre int) bool {
	if nombre <= 0 {
		return false
	}
	if w.Jetons < nombre {
		return false
	}
	w.Jetons -= nombre
	w.Argent += nombre
	return true
}

func afficherResultatVente() {
	if wallet.Argent < StartingMoney {
		fmt.Println("\033[31mDEFAITE : le casino a gagne, votre argent est inferieur a 100K.\033[0m")
		return
	}

	if wallet.Argent > StartingMoney {
		fmt.Println("\033[33mVICTOIRE : vous avez gagne contre le casino, votre argent depasse 100K.\033[0m")
		return
	}

	fmt.Println("EGALITE : vous avez exactement 100K.")
}

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
		fmt.Printf("ARGENT : %d | JETONS : %d\n", wallet.Argent, wallet.Jetons)
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
			fmt.Print("Combien de jetons souhaitez-vous acheter ? : ")
			amountInput, err := reader.ReadString('\n')
			if err != nil {
				return
			}
			amount, err := strconv.Atoi(strings.TrimSpace(amountInput))
			if err != nil || amount <= 0 {
				fmt.Println("Montant invalide.")
				continue
			}
			if wallet.AcheterJetons(amount) {
				fmt.Printf("Achat de %d jetons effectue.\n", amount)
			} else {
				fmt.Println("Solde insuffisant pour acheter ces jetons.")
			}
			fmt.Println("Appuyez sur Entrée pour revenir au menu.")
			_, _ = reader.ReadString('\n')
			return
		case 2:
			fmt.Print("Combien de jetons souhaitez-vous vendre ? : ")
			amountInput, err := reader.ReadString('\n')
			if err != nil {
				return
			}
			amount, err := strconv.Atoi(strings.TrimSpace(amountInput))
			if err != nil || amount <= 0 {
				fmt.Println("Montant invalide.")
				continue
			}
			if wallet.VendreJetons(amount) {
				fmt.Printf("Vente de %d jetons effectuee.\n", amount)
				fmt.Println()
				afficherResultatVente()
			} else {
				fmt.Println("Vous n'avez pas assez de jetons pour vendre cette quantite.")
			}
			fmt.Println()
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
