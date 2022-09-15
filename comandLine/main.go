package main

import (
	"fmt"
	"log"
	"os"

	"github.com/GolangProject/comandLine/app"
)

func main() {

	fmt.Println("Ponto de partida")

	app := app.ToGenerate()
	error := app.Run(os.Args)
	if error != nil {
		log.Fatal(error)
	}

}
