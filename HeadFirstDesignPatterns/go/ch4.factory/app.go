package main

import (
	"fmt"
	"log"
	"os"

	_ "headfirst.design.patterns/factory/ingredient"

	pizza "headfirst.design.patterns/factory/pizza"

)

func main() {
	var nyPizza pizza.NYStyleCheesePizza
	nyPizza.Prepare()
}

func init() {
	fmt.Println("main의 init 함수")
	log.SetOutput(os.Stdout)
}
