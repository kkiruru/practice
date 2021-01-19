package main

import (
	_ "headfirst.design.patterns/factory/pizza"
	pizzaStore "headfirst.design.patterns/factory/pizzastore"
)


func main() {
	pizzaStore := pizzaStore.Store{PizzaStore: new(pizzaStore.NYPizzaStore)}
	pizzaStore.OrderPizza("cheese")
}

