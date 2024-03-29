package locations

import (
	"bufio"
	"fmt"
	"os"

	"github.com/KageJinsei/CharacterCreation/internal/beliefs"
	"github.com/KageJinsei/CharacterCreation/internal/border"
	"github.com/KageJinsei/CharacterCreation/internal/global"
)

func LocaisWesteros() {
	fmt.Println("SELECIONE O LOCAL DE NASCIMENTO:")
	fmt.Print("|1| Winterfell\n" +
		"|2| Highgarden\n" +
		"|3| Carsterly Rock\n" +
		"|4| Storm's End\n" +
		"|5| Riverrun\n")

	for {
		fmt.Print("Qual local deseja?: => ")

		scanner := bufio.NewScanner(os.Stdin)

		if scanner.Scan() {

			input := scanner.Text()

			switch input {
			case "1":
				global.Local = "Winterfell"
				border.Borda()
				beliefs.Religioes()
				return
			case "2":
				global.Local = "Highgarden"
				border.Borda()
				beliefs.Religioes()
				return
			case "3":
				global.Local = "Casterly Rock"
				border.Borda()
				beliefs.Religioes()
				return
			case "4":
				global.Local = "Storm's End"
				border.Borda()
				beliefs.Religioes()
				return
			case "5":
				global.Local = "Riverrun"
				border.Borda()
				beliefs.Religioes()
				return
			default:
				fmt.Println("Local Inválido! Selecione uma das opções listadas.")
			}
		}
	}
}

func LocaisEssos() {
	fmt.Println("SELECIONE O LOCAL DE NASCIMENTO:")
	fmt.Print("|1| Braavos\n" +
		"|2| Pentos\n" +
		"|3| Volantis\n" +
		"|4| Meereen\n" +
		"|5| Astapor\n" +
		"|0| Voltar\n")

	for {
		fmt.Print("Qual local deseja?: => ")

		scanner := bufio.NewScanner(os.Stdin)

		if scanner.Scan() {

			input := scanner.Text()

			switch input {
			case "1":
				global.Local = "Braavos"
				border.Borda()
				beliefs.Religioes()
				return
			case "2":
				global.Local = "Pentos"
				border.Borda()
				beliefs.Religioes()
				return
			case "3":
				global.Local = "Volantis"
				border.Borda()
				beliefs.Religioes()
				return
			case "4":
				global.Local = "Meereen"
				border.Borda()
				beliefs.Religioes()
				return
			case "5":
				global.Local = "Astapor"
				border.Borda()
				beliefs.Religioes()
				return
			default:
				fmt.Println("Local Inválido! Selecione uma das opções listadas.")
			}
		}
	}
}

func LocaisMarEstreito() {
	fmt.Println("SELECIONE O LOCAL DE NASCIMENTO:")
	fmt.Print("|1| Asshai\n" +
		"|2| Shadow Lands\n" +
		"|3| Yi Ti\n" +
		"|4| Naath\n" +
		"|5| Valyria Peninsula\n" +
		"|0| Voltar\n")

	for {
		fmt.Print("Qual local deseja?: => ")

		scanner := bufio.NewScanner(os.Stdin)

		if scanner.Scan() {

			input := scanner.Text()

			switch input {
			case "1":
				global.Local = "Asshai"
				border.Borda()
				beliefs.Religioes()
				return
			case "2":
				global.Local = "Shadow Lands"
				border.Borda()
				beliefs.Religioes()
				return
			case "3":
				global.Local = "Yi Ti"
				border.Borda()
				beliefs.Religioes()
				return
			case "4":
				global.Local = "Naath"
				border.Borda()
				beliefs.Religioes()
				return
			case "5":
				global.Local = "Valyria Peninsula"
				border.Borda()
				beliefs.Religioes()
				return
			default:
				fmt.Println("Local Inválido! Selecione uma das opções listadas.")
			}
		}
	}
}

func LocaisAlemMuralha() {
	fmt.Println("SELECIONE O LOCAL DE NASCIMENTO:")
	fmt.Print("|1| Castle Black\n" +
		"|2| Craster's Keep\n" +
		"|3| Hardhome\n" +
		"|4| Heart Tree\n" +
		"|5| Kingsroad\n" +
		"|0| Voltar\n")

	for {
		fmt.Print("Qual local deseja?: => ")

		scanner := bufio.NewScanner(os.Stdin)

		if scanner.Scan() {

			input := scanner.Text()

			switch input {
			case "1":
				global.Local = "Castle Black"
				border.Borda()
				beliefs.Religioes()
				return
			case "2":
				global.Local = "Craster's Keep"
				border.Borda()
				beliefs.Religioes()
				return
			case "3":
				global.Local = "Hardhome"
				border.Borda()
				beliefs.Religioes()
				return
			case "4":
				global.Local = "Heart Tree"
				border.Borda()
				beliefs.Religioes()
				return
			case "5":
				global.Local = "Kingsroad"
				border.Borda()
				beliefs.Religioes()
				return
			default:
				fmt.Println("Local Inválido! Selecione uma das opções listadas.")
			}
		}
	}
}
