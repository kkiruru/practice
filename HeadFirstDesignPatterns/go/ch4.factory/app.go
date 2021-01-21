package main

import (
	pizzastore "headfirst.design.patterns/factory/pizzastore"
)


func main() {
	var pizzastore pizzastore.PizzaStore = pizzastore.PizzaStore{pizzastore.NYPizzaStore{}}
	pizzastore.OrderPizza("cheese")
}
