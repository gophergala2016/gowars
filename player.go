package main

import (
	"fmt"
)

type Player struct {
	Name  string
	Money float32
	Bag   Trades
}

func (p *Player) Print() {

	fmt.Println("Name:", p.Name)
	fmt.Printf("Money: %.2f\n", p.Money)
	fmt.Println("Bag:")

	p.Bag.Print()
}

func (p *Player) AddToBag(t Trade) {

	var found bool
	for i, item := range p.Bag {
		if item.Name == t.Name {

			p.Bag[i].Quantity += t.Quantity
			found = true
		}
	}

	if !found {
		c := make(Trades, len(p.Bag)+1)
		copy(c, p.Bag)
		c[len(p.Bag)] = t
		p.Bag = c
	}
}

func (p *Player) RemoveFromBag(t Trade) {

	for i, item := range p.Bag {
		if item.Name == t.Name {

			p.Bag[i].Quantity -= t.Quantity
		}
	}
}
