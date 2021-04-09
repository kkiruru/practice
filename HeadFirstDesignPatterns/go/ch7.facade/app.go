package main

import "fmt"

func main() {
	a := GetAmplifierInstance("Top-O-Line Amplifier")
	t := GetTunerInstance("Top-O-Line AM/FM Tuner", a)

	h := GetHomeTheaterInstance(a, t)

	h.WatchMovie("Raiders of the Lost Ark")
	h.EndMovie()

	b()
}

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}
