package main

import (
	"fmt"

	"github.com/lucaslimafernandes/clt-pj-calc/internal/readf"
)

func main() {

	r := readf.ReadToml()

	fmt.Println(*r)

	fmt.Println(r.Reservas)
	fmt.Println(r.Reservas)
	fmt.Println(r.Reservas["emergencia"])

	for i, c := range r.Reservas {
		fmt.Println("i", i)
		fmt.Println("c", c)
	}

}
