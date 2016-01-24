package main

import (
	"fmt"
	"math/rand"
)

type Trade struct {
	Name     string
	Quantity int
}

type Sellable struct {
	Trade
	Price float32
}

type Trades []Trade

func (t *Trade) Print() {

	fmt.Printf("%s : (%d)\n", t.Name, t.Quantity)
}

func (t Trades) Print() {

	position := 0
	for _, trade := range t {
		position++
		fmt.Printf("	%d. ", position)
		trade.Print()
	}
}

type Goods []Sellable

func (s *Sellable) Print() {

	fmt.Printf("%s : %.2f (%d)\n", s.Name, s.Price, s.Quantity)
}

func (g Goods) Add(name string, quantity int) (payment float32) {

	for i, item := range g {
		if item.Name == name {
			g[i].Quantity += quantity
			payment = float32(quantity) * item.Price
		}
	}

	return payment
}

func GenerateGoods() Goods {

	return Goods{
		{Trade: Trade{Name: "Light Saber", Quantity: rand.Intn(5)}, Price: (rand.Float32() * 20) + 20},
		{Trade: Trade{Name: "Phaser Gun", Quantity: rand.Intn(5)}, Price: (rand.Float32() * 15) + 15},
		{Trade: Trade{Name: "Food Packet", Quantity: rand.Intn(5)}, Price: (rand.Float32() * 5) + 5},
		{Trade: Trade{Name: "Thremal Detonator", Quantity: rand.Intn(5)}, Price: (rand.Float32() * 50) + 50},
	}
}

func (g Goods) Print() {

	position := 0
	for _, trade := range g {
		position++
		fmt.Printf("	%d. ", position)
		trade.Print()
	}
}
