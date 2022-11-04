package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	auxiliar.Escrever()
	fmt.Println("Guilherme Go")
	erro := checkmail.ValidateFormat("guilherme@gmail.com")
	fmt.Println(erro)
}
