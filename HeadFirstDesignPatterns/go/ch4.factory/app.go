package main

import (
	_ "headfirst.design.patterns/factory/pizza"
	pizzastore "headfirst.design.patterns/factory/pizzastore"
)


func main() {
	var pizzastore pizzastore.PizzaStore = pizzastore.Store{PizzaStore: pizzastore.NYPizzaStore{}}
	pizzastore.OrderPizza("cheese")
}
