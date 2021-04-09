package main

import "fmt"

type Coffee struct {
}

func (c *Coffee) PrepareRecipe() {
	c.boilWater()
	c.brewCoffeeGrinds()
	c.pourInCup()
	c.addSugarAndMilk()
}

func (c *Coffee) boilWater() {
	fmt.Println("물 끓이는 중")
}

func (c *Coffee) brewCoffeeGrinds() {
	fmt.Println("필터를 통해서 커피를 우려내는 중")
}

func (c *Coffee) pourInCup() {
	fmt.Println("컵에 따르는 중")
}

func (c *Coffee) addSugarAndMilk() {
	fmt.Println("설탕과 우유를 추가하는 중")
}
