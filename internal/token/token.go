package token

import (
	"bufio"
	"fmt"
  "os"

	"github.com/KageJinsei/CharacterCreation/internal/border"
	"github.com/KageJinsei/CharacterCreation/internal/global"
)

func Ficha() {
	fmt.Print("\nInsira o nome de personagem: => ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		global.Nome = scanner.Text()
	}
	fmt.Println("")
	border.Borda()
	fmt.Printf("Olá, %s! Seja bem-vindo(a) a %s!\n", global.Nome, global.Local)
  fmt.Printf("\nVeja as suas informações:\n-> Nome: %s\n-> Raça: %s\n-> Classe: %s\n-> Região: %s\n-> Local: %s\n-> Crença: %s\n", global.Nome, global.Raca, global.Classe, global.Regiao, global.Local, global.Religiao)
}
