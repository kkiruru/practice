package pizza

import (
	"fmt"
)

type Abstract interface {
	Prepare()
}

type Property struct {
	Name string
}

type Pizza struct {
	Abstract
}


func (p *Pizza) Bake() {
	fmt.Println("Bake for 25 minutes at 350")
}


func (p *Pizza) Cut() {
	fmt.Println("Cutting the pizza into diagonal slices")
}


func (p *Pizza) Box() {
	fmt.Println("Place pizza in official PizzaStore box")
}

type NYStyleCheesePizza struct {
	Property
}

func (n *NYStyleCheesePizza) Prepare() {
	fmt.Println("Prepare NYStyleCheesePizza")

	// n.Name = "뉴욕 피자"
}

func (n *NYStyleCheesePizza) init() {
	fmt.Println("init NYStyleCheesePizza")
}