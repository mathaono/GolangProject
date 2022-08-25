package main

import (
	"fmt"

	"github.com/GolangProject/assistant"
	"github.com/GolangProject/functions"
	"github.com/GolangProject/variables"
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("teste")
	assistant.Write()
	err := checkmail.ValidateFormat("aono.1304gmail.com.abc")
	fmt.Println(err)

	variables.Variable()

	sum := functions.Add(10, 20)
	fmt.Println(sum)

	functions.Funcs("Declarando uma variável do tipo func")
	sum, subtract := functions.Calculations(50, 25)
	fmt.Println(sum, subtract)
}
