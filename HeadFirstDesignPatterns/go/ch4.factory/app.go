package main

import (
	"fmt"
	"log"
	"os"

	_ "headfirst.design.patterns/factory/pizza"

	pizzaStore "headfirst.design.patterns/factory/pizzastore"

)

func main() {
	pizzaStore := pizzaStore.PizzaStore{new(pizzaStore.NYPizzaStore)}
	pizza := pizzaStore.OrderPizza("cheese")

	fmt.Println("Ethan ordered a " + pizza.Name );
}

func init() {
	fmt.Println("main의 init 함수")
	log.SetOutput(os.Stdout)
}
