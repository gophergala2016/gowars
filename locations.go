package main

import "fmt"

var locations = []string{"Tatoine", "Forest moon of Endor", "Yodas Hut", "Death Star", "Exit"}

func LocationSelection() (selection int) {

	var position int

	for _, location := range locations {
		position++
		fmt.Printf("%d. %s\n", position, location)
	}

	selection = GetUserInput("Where would you like to go?")

	fmt.Println("Light speed ahead!")

	return selection
}

func Visit(position int) {
	fmt.Printf("Welcome to %s. What are you trading?\n", locations[position-1])
	fmt.Println("Todays prices:")
	stuff := GenerateGoods()

	for true {
		stuff.Print()
		selection := GetUserInput("Do you want to (1) Buy : (2) Sell : (3) Leave?")

		switch selection {
		case 1:
			Buy(stuff)
		case 2:
			Sell(stuff)
		default:
			return
		}
	}
}

func Buy(g Goods) {
	selection := GetUserInput("What do you want?") - 1

	if selection > 3 || selection < 0 {
		fmt.Println("No selection")
		return
	}

	quantity := GetUserInput("How many?")

	if quantity > g[selection].Quantity || quantity < 1 {
		fmt.Println("Too much")
		return
	}

	if g[selection].Price*float32(quantity) > player.Money {
		fmt.Println("Not enough money")
		return
	}

	g[selection].Quantity -= quantity
	player.Money -= g[selection].Price * float32(quantity)

	player.AddToBag(Trade{g[selection].Trade.Name, quantity})

}

func Sell(g Goods) {
	player.Print()

	selection := GetUserInput("What are you selling?") - 1

	if selection > len(player.Bag) || selection < 0 {
		fmt.Println("No selection")
		return
	}

	quantity := GetUserInput("How many?")

	if quantity > player.Bag[selection].Quantity || quantity < 1 {
		fmt.Println("Too much")
		return
	}

	player.Money += g.Add(player.Bag[selection].Name, quantity)
	player.Bag[selection].Quantity -= quantity
}
