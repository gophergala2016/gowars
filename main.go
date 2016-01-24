package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader
var player Player

func main() {

	reader = bufio.NewReader(os.Stdin)

	fmt.Println("Welcome young padawan. Please tell me your name: (Ex. Rey)")

	fmt.Print("Enter text: ")
	playerName, _ := reader.ReadString('\n')

	player = Player{
		Name: strings.TrimSpace(playerName),
		Bag: []Trade{
			{Name: "Light Saber", Quantity: 1},
		},
	}

	fmt.Println("Welcome to GoWars!")

	Travel()
}

func GetUserInput(question string) (selection int) {
	fmt.Println(question)

	strSelection, _ := reader.ReadString('\n')
	strSelection = strings.TrimSpace(strSelection)
	selection, _ = strconv.Atoi(strSelection)

	return selection
}
