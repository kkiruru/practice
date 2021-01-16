package pizza

import (
	"fmt"
)

type Food interface {
	Prepare()
}

type Pizza struct {
	Food
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

type NYStyleCheesePizza struct {}

func (n *NYStyleCheesePizza) Prepare() {

}
