package pizzastore

import (
	pizza "headfirst.design.patterns/factory/pizza"
)


type PizzaStore interface {
	createPizza(pizzaType string) pizza.Pizza
}

type Store struct {
	PizzaStore
}


func (store *Store) OrderPizza(pizzaType string) pizza.Pizza {
	pizza := store.createPizza(pizzaType)

	pizza.Prepare();
	pizza.Bake();
	pizza.Cut();
	pizza.Box();

	return pizza
}


type NYPizzaStore struct {}

func (s *NYPizzaStore) createPizza(pizzaType string) pizza.Pizza {

	ny := new(pizza.NYStyleCheesePizza)

	p := pizza.Pizza{Abstract: ny}
	return p
}