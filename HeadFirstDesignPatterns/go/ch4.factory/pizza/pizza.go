package pizza

import (
	"fmt"
)


type Pizza interface {
	Bake()
	Cut()
	Box()
	Prepare()
}

type property struct {
	Name string
}

type Parent struct {
	Pizza
	property
}

type NYStyleCheesePizza struct {
	Parent
}


func (p Parent) Bake() {
	fmt.Println("Bake for 25 minutes at 350")
}


func (p Parent) Cut() {
	fmt.Println("Cutting the pizza into diagonal slices")
}


func (p Parent) Box() {
	fmt.Println("Place pizza in official PizzaStore box")
}


func (n NYStyleCheesePizza) Prepare() {
	fmt.Println("Prepare NYStyleCheesePizza")
	n.Name = "뉴욕 피자"
}
