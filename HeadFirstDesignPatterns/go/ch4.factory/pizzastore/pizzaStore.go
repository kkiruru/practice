package pizzastore

import (
	"fmt"
	pizza "headfirst.design.patterns/factory/pizza"
)

func init() {
	fmt.Println("pizzastore package의 init 함수")
}


type PizzaStore interface {
	OrderPizza(pizzaType string) pizza.Pizza
	createPizza(pizzaType string) pizza.Pizza
}

type Store struct {
	PizzaStore
}


func (store Store) OrderPizza(pizzaType string) pizza.Pizza {
	pizza := store.createPizza(pizzaType)

	pizza.Prepare();
	pizza.Bake();
	pizza.Cut();
	pizza.Box();

	return pizza
}


type NYPizzaStore struct {
	Store
}

func (s NYPizzaStore) createPizza(pizzaType string) pizza.Pizza {
	pizza := pizza.NYStyleCheesePizza{}
	pizza.Name = "뉴욕피자"
	return pizza
}
