package pizza

import (
	"fmt"
)

func init() {
	fmt.Println("pizza package의 init 함수")
}

type Pizza interface {
	Bake()
	Cut()
	Box()
	Prepare()
}

type property struct {
	Name string
	Price float64
}

type parent struct {
	Pizza
	property
}


func (p parent) Bake() {
	fmt.Println("Bake for 25 minutes at 350")
}


func (p parent) Cut() {
	fmt.Println("Cutting the pizza into diagonal slices")
}


func (p parent) Box() {
	fmt.Println("Place pizza in official PizzaStore box")
}

