package main

import "fmt"

const boilingF = 212.0

var i, j, k int
var b, f, s = true, 2.3, "xd"

var names []string

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("ponto de ebulição = %g F \n ", f)
	fmt.Printf("ponto de ebulição em Celsius %g \n", c)

	fmt.Printf("%g F", ftoC(boilingF))
}

func ftoC(f float64) float64 {
	t := "lol"
	z, x := "massa", "agora sim"
	fmt.Printf(t)
	fmt.Printf(z)
	fmt.Printf(x)

	return (f - 32) * 5 / 9
}
