package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var player Player

func main() {

	Intro()

	fmt.Println("Welcome young padawan. Please tell me your name: (Ex. Rey)")

	playerName := GetUserString("Enter text: ")

	player = Player{
		Name: strings.TrimSpace(playerName),
		Bag: []Trade{
			{Name: "Light Saber", Quantity: 1},
		},
	}

	fmt.Println("Welcome to GoWars!")

	Travel()
}

func GetUserString(question string) (selection string){
	
	fmt.Println(question)
	reader := bufio.NewReader(os.Stdin)	
	
	strSelection, _ := reader.ReadString('\n')
	selection = strings.TrimSpace(strSelection)
	
	return selection
}

func GetUserInput(question string) (selection int) {
		
	selection, _ = strconv.Atoi(GetUserString(question))

	return selection
}
