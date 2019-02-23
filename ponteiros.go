package main

import (
	"flag"
	"fmt"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	new_in()
	x := 1
	p := &x // "endereço de X" ponteiro para int
	z := &x
	s := &z

	// ponteiros podem ser comparados;
	if p == z {
		fmt.Println("ok")
	}
	// *p fornece o valor da variavel;
	fmt.Println(*p)
	fmt.Println(*s)
}

// incrementa o valor da variavel, mais não a altera;
func increment(p *int) int {
	*p++
	return *p
}

// cria uma variavel sem nome do tipo T;
// inicializa com o valor default 0 ; e devolve seu endereço
// que é um valor do tipo *T;
func new_in() {
	p := new(int)
	fmt.Println(*p)
	*p = 2
	fmt.Println(*p)
}
