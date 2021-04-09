package main

import "fmt"

type Tea struct {
	CaffeineBeverage
}

func NewTea() *CaffeineBeverage {
	tea := Tea{}
	tea.CaffeineBeverage.Abstract = &tea
	return &tea.CaffeineBeverage
}

func (t *Tea) PrepareRecipe() {
	t.boilWater()
	t.steepTeaBag()
	t.addLemon()
	t.pourInCup()
}

func (t *Tea) steepTeaBag() {
	fmt.Println("차를 우려내는 중")
}

func (t *Tea) addLemon() {
	fmt.Println("레몬을 추가하는 중")
}
