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

	calc(r)

}

func calc(c *readf.Calculo) {

	var valorFinal float64
	var entradas float64
	// var descontos float64
	var obrigacoes float64

	// Entrada
	for _, val := range c.Entrada {
		entradas += val
	}

	fmt.Printf("Valor final: %v\n", valorFinal)

	// Impostos

	// Obrigacoes
	for _, val := range c.Impostos {
		if val > 1 {
			obrigacoes += val
		} else {
			_temp := entradas * val
			obrigacoes += _temp
		}
	}

	fmt.Printf("Valor obrigacoes: %v\n", obrigacoes)

	// Recomendados
	// Reservas

}
