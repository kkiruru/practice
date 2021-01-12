package main

import (
	"fmt"
	"log"
	"os"

	_ "headfirst.design.patterns/factory/ingredient"
)

func main() {
	fmt.Print("Hello world")
}

func init() {
	fmt.Println("main의 init 함수")
	log.SetOutput(os.Stdout)
}
