package main

import "fmt"

type Tea struct {
	CaffeineBeverage
}

func NewTea() *Tea {
	tea := Tea{}
	tea.abstract = &tea
	return &tea
}

func (t *Tea) PrepareRecipe() {
	t.BoilWater()
	t.steepTeaBag()
	t.addLemon()
	t.PourInCup()
}

func (t *Tea) steepTeaBag() {
	fmt.Println("차를 우려내는 중")
}

func (t *Tea) addLemon() {
	fmt.Println("레몬을 추가하는 중")
}
