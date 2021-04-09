package main

import "fmt"

type Tea struct {
}

func (t *Tea) PrepareRecipe() {
	t.boilWater()
	t.steepTeaBag()
	t.addLemon()
	t.pourInCup()
}

func (t *Tea) boilWater() {
	fmt.Println("물 끓이는 중")
}

func (t *Tea) steepTeaBag() {
	fmt.Println("차를 우려내는 중")
}

func (t *Tea) addLemon() {
	fmt.Println("레몬을 추가하는 중")
}

func (t *Tea) pourInCup() {
	fmt.Println("컵에 따르는 중")
}
