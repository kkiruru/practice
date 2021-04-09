package main

import "fmt"

type Coffee struct {
	CaffeineBeverage
}

func NewCoffee() *CaffeineBeverage {
	coffee := Coffee{}
	coffee.CaffeineBeverage.Abstract = &coffee
	return &coffee.CaffeineBeverage
}

func (c *Coffee) PrepareRecipe() {
	c.boilWater()
	c.brewCoffeeGrinds()
	c.pourInCup()
	c.addSugarAndMilk()
}

func (c *Coffee) brewCoffeeGrinds() {
	fmt.Println("필터를 통해서 커피를 우려내는 중")
}

func (c *Coffee) addSugarAndMilk() {
	fmt.Println("설탕과 우유를 추가하는 중")
}
