package condimentdecorator

import (
	"headfirst.design.patterns/decorator/beverage"
)

type Mocha struct {
	beverage beverage.Beverage
}

func AddMocha(b beverage.Beverage) beverage.Beverage {
	c := Mocha{beverage: b}
	return c
}

func (e Mocha) Cost() (cost float32) {
	return 0.20 + e.beverage.Cost()
}

func (e Mocha) Description() string {
	return (e.beverage.Description() + ", 모카")
}

type Whip struct {
	beverage beverage.Beverage
}

func AddWhip(b beverage.Beverage) beverage.Beverage {
	c := Whip{beverage: b}
	return c
}

func (e Whip) Cost() (cost float32) {
	return 0.10 + e.beverage.Cost()
}

func (e Whip) Description() string {
	return (e.beverage.Description() + ", 휘핑 크림")
}

type Soy struct {
	beverage beverage.Beverage
}

func AddSoy(b beverage.Beverage) beverage.Beverage {
	c := Soy{beverage: b}
	return c
}

func (e Soy) Cost() (cost float32) {
	return 0.15 + e.beverage.Cost()
}

func (e Soy) Description() string {
	return (e.beverage.Description() + ", 두유")
}

type SteamMilk struct {
	beverage beverage.Beverage
}

func AddSteamMilk(b beverage.Beverage) beverage.Beverage {
	c := SteamMilk{beverage: b}
	return c
}

func (e SteamMilk) Cost() (cost float32) {
	return 0.10 + e.beverage.Cost()
}

func (e SteamMilk) Description() string {
	return (e.beverage.Description() + ", 스팀 밀크")
}
