package main

import (
	_ "headfirst.design.patterns/factory/pizza"

	pizzaStore "headfirst.design.patterns/factory/pizzastore"

)

func main() {
	pizzaStore := pizzaStore.Store{new(pizzaStore.NYPizzaStore)}
	pizzaStore.OrderPizza("cheese")
}

// func init() {
// 	fmt.Println("main의 init 함수")
// 	log.SetOutput(os.Stdout)
// }
