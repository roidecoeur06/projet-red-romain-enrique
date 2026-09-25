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

type character struct {
	title     string
	firstName string
}

var currentFirstName = "JOUEUR"
var currentCharacter = 0
var powerCooldown = 0
var currentPV = 100
var winStreak = 0

func ChooseCharacter(reader *bufio.Reader) {
	characters := []character{
		{title: "LE MENTALISTE", firstName: "ALEX"},
		{title: "LE RICHISSIME", firstName: "VICTOR"},
		{title: "LE JOUEUR NORMAL", firstName: "LOUIS"},
	}

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
		selectedCharacter := characters[choice-1]
		fmt.Printf("Vous avez choisi %s.\n", selectedCharacter.title)
		printCharacterCharacteristics(choice)

		fmt.Println()
		fmt.Printf("Le prenom actuel est %s.\n", selectedCharacter.firstName)
		fmt.Print("VOULEZ-VOUS CHANGER LE PRENOM DU PERSONNAGE ? (oui/non) : ")
		changeName, err := reader.ReadString('\n')
		if err != nil && len(changeName) == 0 {
			return
		}

		selectedFirstName := selectedCharacter.firstName
		if strings.ToLower(strings.TrimSpace(changeName)) == "oui" {
			fmt.Print("ECRIVEZ LE PRENOM DE VOTRE PERSONNAGE : ")
			newFirstName, err := reader.ReadString('\n')
			if err != nil && len(newFirstName) == 0 {
				return
			}

			newFirstName = strings.TrimSpace(newFirstName)
			if newFirstName == "" {
				fmt.Println("Le prenom ne peut pas etre vide.")
				continue
			}
			selectedFirstName = newFirstName
		}

		fmt.Println()
		fmt.Print("VOULEZ-VOUS CHANGER DE PERSONNAGE ? (oui/non) : ")
		change, err := reader.ReadString('\n')
		if err != nil && len(change) == 0 {
			return
		}

		if strings.ToLower(strings.TrimSpace(change)) == "oui" {
			continue
		}

		currentFirstName = selectedFirstName
		currentCharacter = choice
		currentPV = 100
		winStreak = 0
		powerCooldown = 0
		fmt.Println("Vous allez etre redirige vers le menu principal.")
		clearTerminal()
		return
	}
}

func GetFirstName() string {
	return currentFirstName
}

func GetCharacter() int {
	return currentCharacter
}

func AdvancePowerCooldown() {
	if powerCooldown > 0 {
		powerCooldown--
	}
}

func CanUsePower() bool {
	return currentCharacter != 0 && powerCooldown == 0
}

func UsePower() bool {
	if !CanUsePower() {
		return false
	}
	powerCooldown = 10
	return true
}

func PowerCooldown() int {
	return powerCooldown
}

func GetPV() int {
	return currentPV
}

func IsKO() bool {
	return currentPV == 0
}

func ApplyRoundResult(won bool) {
	if won {
		winStreak++
		gain := 10
		if currentCharacter == 2 {
			gain *= 2
		}
		currentPV += gain
		if currentPV > 100 {
			currentPV = 100
		}
		if currentCharacter == 1 && winStreak >= 5 {
			currentPV -= 30
			winStreak = 0
		}
		return
	}

	winStreak = 0
	loss := 10
	if currentCharacter == 2 {
		loss *= 2
	}
	currentPV -= loss
	if currentPV < 0 {
		currentPV = 0
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
