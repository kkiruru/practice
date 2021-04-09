package main

import "fmt"

type Abstract interface {
	PrepareRecipe()

	boilWater()
	pourInCup()
}

type CaffeineBeverage struct {
	Abstract
}

func (c *CaffeineBeverage) boilWater() {
	fmt.Println("물 끓이는 중..")
}

func (c *CaffeineBeverage) pourInCup() {
	fmt.Println("컵에 따르는 중..")
}
