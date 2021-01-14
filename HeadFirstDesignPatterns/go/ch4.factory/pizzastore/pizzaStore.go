package pizzastore

import (
	pizza "headfirst.design.patterns/factory/pizza"
)


type Store interface {
	createPizza(pizzaType string) pizza.Pizza
}


type PizzaStore struct {
	Store
}


func (s *PizzaStore) OrderPizza(pizzaType string) pizza.Pizza {
	pizza := s.createPizza(pizzaType)
	return pizza
}


type NYPizzaStore struct {}

func (s *NYPizzaStore) createPizza(pizzaType string) pizza.Pizza {
	var p pizza.Pizza
	return p
}