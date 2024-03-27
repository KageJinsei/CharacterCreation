package beliefs

import (
	"bufio"
  "fmt"
  "os"

	"github.com/KageJinsei/CharacterCreation/internal/border"
  "github.com/KageJinsei/CharacterCreation/internal/global"
  "github.com/KageJinsei/CharacterCreation/internal/token"
)

func Religioes() {
	fmt.Println("SELECIONE UMA CRENÇA:")
	fmt.Print("|1| Fé dos Sete\n" +
		"|2| Deuses Antigos\n" +
		"|3| Senhor da Luz\n" +
		"|4| Deus Afogado\n" +
    "|5| Deus de Muitas Faces\n")

	for {
		fmt.Print("Qual crença deseja?: => ")

		scanner := bufio.NewScanner(os.Stdin)

		if scanner.Scan() {

			input := scanner.Text()

			switch input {
			case "1":
				global.Religiao = "Fé dos Sete"
				border.Borda()
				token.Ficha()
				return
			case "2":
				global.Religiao = "Deuses Antigos"
				border.Borda()
				token.Ficha()
				return
			case "3":
				global.Religiao = "Senhor da Luz"
				border.Borda()
				token.Ficha()
				return
			case "4":
				global.Religiao = "Deus Afogado"
				border.Borda()
				token.Ficha()
				return
      case "5":
        global.Religiao = "Deus de Muitas Faces"
        border.Borda()
        token.Ficha()
        return
			default:
				fmt.Println("Religião Inválida! Selecione uma das opções listadas.")
			}
		}
	}
}
