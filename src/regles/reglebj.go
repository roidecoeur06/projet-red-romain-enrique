package regles

import (
	"bufio"
	"fmt"
)

const reglesBlackjack = `
        ████  █████  ███  █     █████    ████  █   █    ████  █      ███   ███    ███  ███   ███  █   █         
        █   █ █     █     █     █        █   █ █   █    █   █ █     █   █ █        █  █   █ █     █  █          
████    ████  ████  █  ██ █     ████     █   █ █   █    ████  █     █████ █        █  █████ █     ███      ████ 
        █  █  █     █   █ █     █        █   █ █   █    █   █ █     █   █ █     █  █  █   █ █     █  █          
        █   █ █████  ███  █████ █████    ████   ███     ████  █████ █   █  ███   ██   █   █  ███  █   █         
 
But du jeu
Obtenir une main dont la valeur est la plus proche possible de 21, sans la dépasser, et battre la main du croupier.
 
Valeur des cartes
- Les cartes de 2 à 10 valent leur valeur affichée
- Le Valet, la Dame et le Roi valent 10
- L'As vaut 11, ou 1 si 11 fait dépasser 21
 
Déroulement
1. Chaque joueur mise
2. Le croupier distribue deux cartes à chaque joueur et deux cartes à lui-même (une visible, une cachée)
3. Si les deux premières cartes forment un As + une carte de valeur 10, c'est un Blackjack naturel, payé 3 contre 2
4. Chaque joueur joue son tour, puis le croupier joue le sien
 
Actions du joueur
- Tirer : demander une carte supplémentaire
- Rester : garder sa main telle quelle
- Doubler : doubler sa mise, tirer une seule carte, puis rester automatiquement
- Diviser (Split) : si les deux cartes ont la même valeur, séparer en deux mains distinctes avec une mise égale sur chacune
- Abandonner (Surrender) : abandonner la main dès le départ en récupérant la moitié de la mise
- Assurance : si le croupier montre un As, parier jusqu'à la moitié de la mise qu'il a un Blackjack, payé 2 contre 1 si c'est le cas
 
Règle du croupier
Le croupier tire tant que sa main vaut 16 ou moins, et reste dès qu'elle atteint 17 ou plus.
 
Résultats
- Si le joueur dépasse 21, il perd immédiatement
- Si le croupier dépasse 21, tous les joueurs encore en jeu gagnent
- Si la main du joueur est plus proche de 21 que celle du croupier, il gagne
- En cas d'égalité, la mise est remboursée
- Si le croupier a une main plus forte, le joueur perd sa mise
 
 
SIDE BET : PERFECT PAIRS
 
Ce pari porte uniquement sur les deux premières cartes du joueur. Il gagne si elles forment une paire :
 
- Paire parfaite (même valeur, même enseigne) : payée 25 contre 1
- Paire colorée (même valeur, même couleur, enseigne différente) : payée 12 contre 1
- Paire mixte (même valeur, couleurs différentes) : payée 6 contre 1
 
 
SIDE BET : 21+3
 
Ce pari combine les deux cartes du joueur et la carte visible du croupier, évaluées comme une main de poker à trois cartes :
 
- Brelan (trois cartes de même valeur) : payé 30 contre 1
- Quinte flush (trois cartes consécutives de la même enseigne) : payée 40 contre 1
- Suite (trois cartes consécutives, enseignes différentes) : payée 10 contre 1
- Couleur (trois cartes de la même enseigne, non consécutives) : payée 5 contre 1
`

func ShowRules(reader *bufio.Reader) {
	fmt.Print("\033[2J\033[3J\033[H")
	fmt.Print(reglesBlackjack)
	fmt.Println()
	fmt.Print("Appuyez sur Entrée pour revenir au menu... ")

	_, err := reader.ReadString('\n')
	if err != nil {
		return
	}

	fmt.Print("\033[2J\033[3J\033[H")
}
