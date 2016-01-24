package main

import (
	"fmt"
	"os"
)

func Travel() {

	for true {
		Divider()
		player.Print()
		selection := LocationSelection()

		if selection == 5 {
			fmt.Println("Thanks for playing!")

			player.Print()

			os.Exit(0)
		}

		Visit(selection)
	}
}
