package classes

import (
	"bufio"
	"fmt"
  "os"

	"github.com/KageJinsei/CharacterCreation/internal/border"
	"github.com/KageJinsei/CharacterCreation/internal/global"
	"github.com/KageJinsei/CharacterCreation/internal/regions"
)

func Classes() {
	fmt.Println("SELECIONE A SUA CLASSE:")
	fmt.Print("|1| Guerreiro\n" +
		"|2| Mago\n" +
		"|3| Arqueiro\n" +
		"|4| Noviço\n" +
		"|5| Assassino\n" +
		"|6| Druida\n" +
    "|7| Paladino\n" +
    "|8| Necromante\n" +
    "|9| Monge\n")

	for {
		fmt.Print("Qual classe deseja?: => ")

		scanner := bufio.NewScanner(os.Stdin)

		if scanner.Scan() {

			input := scanner.Text()

			switch input {
			case "1":
				global.Classe = "Guerreiro"
				border.Borda()
				regions.Regioes()
				return
			case "2":
				global.Classe = "Mago"
				border.Borda()
				regions.Regioes()
				return
			case "3":
				global.Classe = "Arqueiro"
				border.Borda()
				regions.Regioes()
				return
			case "4":
				global.Classe = "Noviço"
				border.Borda()
				regions.Regioes()
				return
			case "5":
				global.Classe = "Assassino"
				border.Borda()
				regions.Regioes()
				return
      case "6":
				global.Classe = "Druida"
				border.Borda()
				regions.Regioes()
				return
      case "7":
				global.Classe = "Paladino"
				border.Borda()
				regions.Regioes()
				return
      case "8":
				global.Classe = "Necromante"
				border.Borda()
				regions.Regioes()
				return
      case "9":
				global.Classe = "Monge"
				border.Borda()
				regions.Regioes()
				return
			default:
				fmt.Println("Classe Inválida! Selecione uma das opções listadas.")
			}
		}
	}
}
