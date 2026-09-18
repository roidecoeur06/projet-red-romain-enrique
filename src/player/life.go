package player

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

const characterTitle = `
         ███  █   █  ███  ███ █   █    ████  █   █    ████  █████ ████   ████  ███  █   █ █   █  ███   ███  █████         
        █     █   █ █   █  █   █ █     █   █ █   █    █   █ █     █   █ █     █   █ ██  █ ██  █ █   █ █     █             
████    █     █████ █   █  █    █      █   █ █   █    ████  ████  ████   ███  █   █ █ █ █ █ █ █ █████ █  ██ ████     ████ 
        █     █   █ █   █  █   █ █     █   █ █   █    █     █     █  █      █ █   █ █  ██ █  ██ █   █ █   █ █             
         ███  █   █  ███  ███ █   █    ████   ███     █     █████ █   █ ████   ███  █   █ █   █ █   █  ███  █████`

func ChooseCharacter(reader *bufio.Reader) {
	characters := []string{"LE MENTALISTE", "LE RICHISSIME", "LE JOUEUR NORMAL"}

	for {
		clearTerminal()
		fmt.Println(characterTitle)
		fmt.Println()
		fmt.Println()
		fmt.Println("1. LE MENTALISTE")
		fmt.Println("2. LE RICHISSIME")
		fmt.Println("3. LE JOUEUR NORMAL")
		fmt.Println()
		fmt.Print("CHOISISSEZ VOTRE PERSONNAGE : ")

		input, err := reader.ReadString('\n')
		if err != nil && len(input) == 0 {
			return
		}

		choice, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil || choice < 1 || choice > 3 {
			fmt.Println("Choix invalide.")
			continue
		}

		fmt.Println()
		fmt.Printf("Vous avez choisi %s.\n", characters[choice-1])
		printCharacterCharacteristics(choice)

		fmt.Println()
		fmt.Print("VOULEZ-VOUS CHANGER DE PERSONNAGE ? (oui/non) : ")
		change, err := reader.ReadString('\n')
		if err != nil && len(change) == 0 {
			return
		}

		if strings.ToLower(strings.TrimSpace(change)) == "oui" {
			continue
		}

		fmt.Println("Vous allez etre redirige vers le menu principal.")
		clearTerminal()
		return
	}
}

func printCharacterCharacteristics(choice int) {
	fmt.Println()
	fmt.Println("CARACTERISTIQUES DU PERSONNAGE")
	fmt.Println()

	switch choice {
	case 1:
		fmt.Println("Mecanique principale : Detection du casino")
		fmt.Println("Capacite : Compte les cartes pour anticiper les tirages.")
		fmt.Println("Impact sur les PV : Apres 5 victoires d'affilee, le casino le soupconne de triche ATTENTION !.")
		fmt.Println("Malus : Controle du casino et perte de 30 PV.")
	case 2:
		fmt.Println("Mecanique principale : Quitte ou Double")
		fmt.Println("Gains : Quand il gagne une manche, il double ses gains de PV.")
		fmt.Println("Pertes : Quand il perd, il perd 2 fois plus de PV.")
	case 3:
		fmt.Println("Mecanique principale : Sante mentale / Dopamine")
		fmt.Println("Victoires (Dopamine) : Chaque victoire lui fait gagner 10 PV.")
		fmt.Println("Pertes (Stress) : S'il perd trop de coins ou de jetons, ses PV chutent.")
	}
}

func clearTerminal() {
	fmt.Print("\033[2J\033[3J\033[H")
}
