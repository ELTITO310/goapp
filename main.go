package main

import (
	"flag"
	"log"
	"strconv"

	"github.com/ELTITO310/goapp/commands"
	"github.com/fatih/color"
)

func main() {

	var expenses []float32
	var export string

	flag.StringVar(&export, "export", "", "Exports the expenses to file .txt")

	flag.Parse()

	for {
		input, err := commands.GetInput()
		if err != nil {
			log.Fatal(err)
		}

		if input == "cls" {
			break
		}

		expense, err := strconv.ParseFloat(input, 32)
		if err != nil {
			d := color.New(color.FgRed, color.Bold)
			d.Printf("Type a number float32\n")
			continue
		}

		expenses = append(expenses, float32(expense))
	}

	if export == "" {
		commands.ShowInConsole(expenses)

		d := color.New(color.FgRed, color.Bold)
		d.Printf("Closing...")
	} else {
		commands.Export(export, expenses)
	}

}
