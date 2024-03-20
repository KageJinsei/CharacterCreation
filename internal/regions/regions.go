package regions

import (
	"bufio"
	"fmt"
	"github.com/KageJinsei/CharacterCreation/internal/border"
	"github.com/KageJinsei/CharacterCreation/internal/global"
	"github.com/KageJinsei/CharacterCreation/internal/locations"
	"os"
)

func Regioes() {
	fmt.Println("SELECIONE UMA REGIÃO:")
	fmt.Print("|1| Westeros\n" +
		"|2| Essos\n" +
		"|3| Além do Mar Estreiro\n" +
		"|4| Além da Muralha\n" +
		"|0| Voltar\n")

	for {
		fmt.Print("Qual região deseja?: => ")

		scanner := bufio.NewScanner(os.Stdin)

		if scanner.Scan() {

			input := scanner.Text()

			switch input {
			case "1":
				global.Regiao = "Westeros"
				border.Borda()
				locations.LocaisWesteros()
				return
			case "2":
				global.Regiao = "Essos"
				border.Borda()
				locations.LocaisEssos()
				return
			case "3":
				global.Regiao = "Além do Mar Estreito"
				border.Borda()
				locations.LocaisMarEstreito()
				return
			case "4":
				global.Regiao = "Além da Muralha"
				border.Borda()
				locations.LocaisAlemMuralha()
				return
			default:
				fmt.Println("Região Inválida! Selecione uma das opções listadas.")
			}
		}
	}
}
