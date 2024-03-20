package races

import (
	"bufio"
	"fmt"
	"github.com/KageJinsei/CharacterCreation/internal/border"
	"github.com/KageJinsei/CharacterCreation/internal/classes"
	"github.com/KageJinsei/CharacterCreation/internal/global"
	"os"
)

func Racas() {
	fmt.Println("SELECIONE A SUA RAÇA:")
	fmt.Print("|1| Humano\n" +
		"|2| Elfo\n" +
		"|3| Fada\n" +
		"|4| Besta\n" +
		"|5| Vagante\n")

	for {
		fmt.Print("Qual raça deseja?: => ")

		scanner := bufio.NewScanner(os.Stdin)

		if scanner.Scan() {

			input := scanner.Text()

			switch input {
			case "1":
				global.Raca = "Humano"
				border.Borda()
				classes.Classes()
				return
			case "2":
				global.Raca = "Elfo"
				border.Borda()
				classes.Classes()
				return
			case "3":
				global.Raca = "Fada"
				border.Borda()
				classes.Classes()
				return
			case "4":
				global.Raca = "Besta"
				border.Borda()
				classes.Classes()
				return
			case "5":
				global.Raca = "Vagante"
				border.Borda()
				classes.Classes()
				return
			default:
				fmt.Println("Raça Inválida! Escolha uma das opções listadas.")
			}
		}
	}
}
