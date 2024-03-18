package main

import (
	"bufio"
	"fmt"
	"os"
)

var raca string
var classe string
var local string
var nome string

func main() {
	fmt.Println("RPG - IMAGINA UM NOME LEGAL AQUI!")

	borda()
	racas()
	borda()
	classes()
	borda()
	locais()
	borda()
	ficha()
	borda()
}

func racas() {
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
			  raca = "Humano"
        return
		  case "2":
			  raca = "Elfo"
        return
		  case "3":
			  raca = "Fada"
        return
		  case "4":
			  raca = "Besta"
        return
		  case "5":
			  raca = "Vagante"
        return
      default:
        fmt.Println("Raça Inválida! Escolha uma das opções listadas.")
		  }
	  }
  }
}

func classes() {
	fmt.Println("SELECIONE A SUA CLASSE:")
	fmt.Print("|1| Guerreiro\n" +
		"|2| Mago\n" +
		"|3| Arqueiro\n" +
		"|4| Noviço\n" +
		"|5| Assassino\n")
	
  for {
  fmt.Print("Qual classe deseja?: => ")

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {

		input := scanner.Text()

		switch input {
		case "1":
			classe = "Guerreiro"
      return
		case "2":
			classe = "Mago"
      return
		case "3":
			classe = "Arqueiro"
      return
		case "4":
			classe = "Noviço"
		  return
    case "5":
			classe = "Assassino"
      return
    default:
      fmt.Println("Classe Inválida! Selecione uma das opções listadas.")
		  }
    }
	}
}

func locais() {
	fmt.Println("SELECIONE O LOCAL DE NASCIMENTO:")
	fmt.Print("|1| Winterfell\n" +
		"|2| Highgarden\n" +
		"|3| Carsterly Rock\n" +
		"|4| Storm's End\n" +
		"|5| Red Keep\n")
	
  for {
  fmt.Print("Qual local deseja?: => ")

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {

		input := scanner.Text()

		switch input {
		case "1":
			local = "Winterfell"
      return
		case "2":
			local = "Highgarden"
      return
		case "3":
			local = "Casterly Rock"
      return
		case "4":
			local = "Storm's End"
      return
		case "5":
			local = "Red Keep"
      return
    default:
      fmt.Println("Local Inválido! Selecione uma das opções listadas.")
		  }
    }
	}
}

func ficha() {

	fmt.Print("\nInsira o nome da sua personagem: => ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		nome = scanner.Text()
	}
	fmt.Println("")
	borda()
	fmt.Printf("\nOlá, %s! Seja bem-vindo(a) a %s!\n", nome, local)
	fmt.Printf("\nVeja as suas informações:\n-> Nome: %s\n-> Raça: %s\n-> Classe: %s\n-> Local: %s\n", nome, raca, classe, local)
}

func borda() {
	linhas := ""

	for l := 0; l < 35; l++ {
		linhas += "="
	}

	fmt.Println(linhas)
}
