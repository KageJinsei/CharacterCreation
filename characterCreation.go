package main

import "fmt"

var raca string
var classe string
var local string
var nome string

func main() {
  fmt.Println("RPG - IMAGINA UM NOME FODA AQUI!\n")
  
  borda()
  racas()
  borda()
  classes()
  borda()
  locais()
  borda()

  fmt.Print("\nInsira o nome da sua personagem: => ")
  fmt.Scan(&nome)
  fmt.Print("\n")
  
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
  fmt.Print("Qual raça deseja?: => ") 
  fmt.Scan(&raca)

  switch {
  case raca == "1":
    raca = "Humano"
  
  case raca == "2":
    raca = "Elfo"

  case raca == "3":
    raca = "Fada"

  case raca == "4":
    raca = "Besta"

  case raca == "5":
    raca = "Vagante"
  }
}

func classes() {
  fmt.Println("SELECIONE A SUA CLASSE:")
  fmt.Print("|1| Guerreiro\n" +
            "|2| Mago\n" +
            "|3| Arqueiro\n" +
            "|4| Noviço\n" +
            "|5| Assassino\n")
  fmt.Print("Qual classe deseja?: => ")
  fmt.Scan(&classe)

  switch {
  case classe == "1":
    classe = "Guerreiro"

  case classe == "2":
    classe = "Mago"

  case classe == "3":
    classe = "Arqueiro"

  case classe == "4":
    classe = "Noviço"

  case classe == "5":
    classe = "Assassino"
  }
}

func locais() {
  fmt.Println("SELECIONE O LOCAL DE NASCIMENTO:")
  fmt.Print("|1| Winterfell\n" +
            "|2| Highgarden\n" + 
            "|3| Carsterly Rock\n" +
            "|4| Storm's End\n" +
            "|5| Red Keep\n")
  fmt.Print("Qual local deseja?: => ")
  fmt.Scan(&local)
  
  switch {
  case local == "1":
    local = "Winterfell"

  case local == "2":
    local = "Highgarden"

  case local == "3":
    local = "Casterly Rock"

  case local == "4":
    local = "Storm's End"

  case local == "5":
    local = "Red Keep"
  }
}

func ficha() {
  fmt.Printf("Olá, %s! Seja bem-vindo(a) a %s!\n", nome, local)
  fmt.Printf("\nVeja as suas informações:\n-> Nome: %s\n-> Raça: %s\n-> Classe: %s\n-> Local: %s\n", nome, raca, classe, local)
}

func borda() {
  linhas := ""

  for l := 0; l < 35; l++ {
    linhas += "="
  }

  fmt.Println(linhas)
}
