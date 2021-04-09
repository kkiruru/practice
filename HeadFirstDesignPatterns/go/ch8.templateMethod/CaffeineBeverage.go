package main

import "fmt"

type abstract interface {
	PrepareRecipe()
	BoilWater()
	PourInCup()
}

type CaffeineBeverage struct {
	abstract
}

func (c *CaffeineBeverage) BoilWater() {
	fmt.Println("물 끓이는 중..")
}

func (c *CaffeineBeverage) PourInCup() {
	fmt.Println("컵에 따르는 중..")
}
