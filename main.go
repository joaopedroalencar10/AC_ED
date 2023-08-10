package main

import (
	"fmt"
)

func main() {
  e_primo()
  semana()
  bisexto()
}
func semana(){

var dia = "1"

	switch dia {
	  case "1":
		fmt.Println("domingo")
	   case "2":
		fmt.Println("segunda")
    case "3":
		fmt.Println("terça")
    case "4":
		fmt.Println("quarta")
    case "5":
		fmt.Println("quinta")
    case "6":
		fmt.Println("sexta")
    case "7":
		fmt.Println("sábado")
	default:
		fmt.Println("Erro, dia inválido.")
	}
}

func e_primo(){
  var n = 10
 
  for i := 2; i < n-1; i++ {
  if n % i == 0{
    fmt.Println("numero não é primo")
    break
  } else {
     fmt.Println("numero é primo")
    
  }
  }
}

func bisexto(){
  var ano = 1984
  if ano % 4 == 0 {
    fmt.Println("ano é bissexto")
  }else {
    fmt.Println("ano não é bissexto")
  }
}