package main

import (
	"bufio"
	"fmt"
	"os"
)

var raca string
var classe string
var regiao string
var local string
var nome string

func main() {
	fmt.Println("RPG - IMAGINA UM NOME LEGAL AQUI!")

  borda()
	racas()
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
        borda()
        classes()
        return
		  case "2":
			  raca = "Elfo"
        borda()
        classes()
        return
		  case "3":
			  raca = "Fada"
        borda()
        classes()
        return
		  case "4":
			  raca = "Besta"
        borda()
        classes()
        return
		  case "5":
			  raca = "Vagante"
        borda()
        classes()
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
        borda()
        regioes()
        return
		  case "2":
			  classe = "Mago"
        borda()
        regioes()
        return
		  case "3":
			  classe = "Arqueiro"
        borda()
        regioes()
        return
		  case "4":
			  classe = "Noviço"
		    borda()
        regioes()
        return
      case "5":
			  classe = "Assassino"
        borda()
        regioes()
        return
      default:
        fmt.Println("Classe Inválida! Selecione uma das opções listadas.")
		  }
    }
	}
}

func regioes() {
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
        regiao = "Westeros"
        borda()
        locaisWesteros()
        return
      case "2":
        regiao = "Essos"
        borda()
        locaisEssos()
        return
      case "3":
        regiao = "Além do Mar Estreito"
        borda()
        locaisMarEstreito()
        return
      case "4":
        regiao = "Além da Muralha"
        borda()
        locaisAlemMuralha()
        return
      case "0":
        borda()
        classes()
      return
      default:
        fmt.Println("Região Inválida! Selecione uma das opções listadas.")
      }
    }
  }
}

func locaisWesteros() {
	fmt.Println("SELECIONE O LOCAL DE NASCIMENTO:")
	fmt.Print("|1| Winterfell\n" +
		        "|2| Highgarden\n" +
		        "|3| Carsterly Rock\n" +
		        "|4| Storm's End\n" +
		        "|5| Riverrun\n" +
            "|0| Voltar\n")
	  
  for {
    fmt.Print("Qual local deseja?: => ")

	  scanner := bufio.NewScanner(os.Stdin)

	  if scanner.Scan() {

		  input := scanner.Text()

		  switch input {
		  case "1":
			  local = "Winterfell"
        borda()
        ficha()
        return
		  case "2":
			  local = "Highgarden"
        borda()
        ficha()
        return
		  case "3":
			  local = "Casterly Rock"
        borda()
        ficha()
        return
		  case "4":
			  local = "Storm's End"
        borda()
        ficha()
        return
		  case "5":
			  local = "Red Keep"
        borda()
        ficha()
        return
      case "0":
        borda()
        regioes()
        return
      default:
        fmt.Println("Local Inválido! Selecione uma das opções listadas.")
		  }
    }
	}
}

func locaisEssos() {
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
			  local = "Braavos"
        borda()
        ficha()
        return
		  case "2":
			  local = "Pentos"
        borda()
        ficha()
        return
		  case "3":
			  local = "Volantis"
        borda()
        ficha()
        return
		  case "4":
			  local = "Meereen"
        borda()
        ficha()
        return
		  case "5":
			  local = "Astapor"
        borda()
        ficha()
        return
      case "0":
        borda()
        regioes()
        return
      default:
        fmt.Println("Local Inválido! Selecione uma das opções listadas.")
		  }
    }
	}
}

func locaisMarEstreito() {
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
			  local = "Asshai"
        borda()
        ficha()
        return
		  case "2":
			  local = "Shadow Lands"
        borda()
        ficha()
        return
		  case "3":
			  local = "Yi Ti"
        borda()
        ficha()
        return
		  case "4":
			  local = "Naath"
        borda()
        ficha()
        return
      case "5":
        local = "Valyria Peninsula"
        borda()
        ficha()
        return
      case "0":
        borda()
        regioes()
        return
      default:
        fmt.Println("Local Inválido! Selecione uma das opções listadas.")
		  }
    }
	}
}

func locaisAlemMuralha() {
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
			  local = "Castle Black"
        borda()
        ficha()
        return
		  case "2":
			  local = "Craster's Keep"
        borda()
        ficha()
        return
		  case "3":
			  local = "Hardhome"
        borda()
        ficha()
        return
		  case "4":
			  local = "Heart Tree"
        borda()
        ficha()
        return
		  case "5":
			  local = "Kingsroad"
        borda()
        ficha()
        return
      case "0":
        borda()
        regioes()
        return
      default:
        fmt.Println("Local Inválido! Selecione uma das opções listadas.")
		  }
    }
	}
}

func ficha() {
	fmt.Print("\nInsira o nome de personagem: => ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		nome = scanner.Text()
	}
	fmt.Println("")
	borda()
	fmt.Printf("Olá, %s! Seja bem-vindo(a) a %s!\n", nome, local)
  fmt.Printf("\nVeja as suas informações:\n-> Nome: %s\n-> Raça: %s\n-> Classe: %s\n-> Região: %s\n-> Local: %s\n", nome, raca, classe, regiao, local)
}

func borda() {
	linhas := ""

	for l := 0; l < 35; l++ {
		linhas += "="
	}

	fmt.Println(linhas)
}
