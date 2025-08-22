package main

import (
	"flag"
	"fmt"

	"github.com/lucaslimafernandes/clt-pj-calc/internal/utilities"
)

func main() {

	help := flag.Bool("help", false, "Show available commands")
	generate_cfg := flag.Bool("gen-config", false, "Exporta o arquivo de configuração")
	salBrt := flag.Float64("salario", 0.0, "Digite o salário bruto, fmt 1000.10")

	flag.Parse()

	if *help {
		fmt.Printf(`[project]
name = "clt-pj-calc"
version = "v0.0.1"
description = "Calculadora de salário."

[repository]
url = "https://github.com/lucaslimafernandes/clt-pj-calc"

`)
		return
	}

	if *generate_cfg {
		utilities.CreateToml()
		return
	}

	fmt.Println(*salBrt)
	// r := utilities.ReadToml()

	// fmt.Println(*r)

	// fmt.Println(r.Reservas)
	// fmt.Println(r.Reservas)
	// fmt.Println(r.Reservas["emergencia"])

	// fmt.Println(r.Obrigacoes["prolabore"])

	// for i, c := range r.Reservas {
	// 	fmt.Println("i", i)
	// 	fmt.Println("c", c)
	// }

	// calc(r)

	handler(*salBrt)

}

func handler(salBrt float64) {

	inssContrib := utilities.CalcINSS(salBrt)

	fmt.Println("CalcINSS: ", inssContrib)

	irpf := utilities.CalcIRPF(1, salBrt, inssContrib)

	fmt.Println("CalcIRPF: ", irpf)

	ferias := utilities.Round2(salBrt * 0.09)
	umTercFerias := utilities.Round2(salBrt * 0.03)
	decimoTerc := utilities.Round2(salBrt * 0.09)
	fgts := utilities.Round2(salBrt * 0.08)
	unimed := 1500.00
	VA := 1000.00
	contabilidade := 300.00

	fmt.Println("Férias: ", ferias)
	fmt.Println("1/3 férias: ", umTercFerias)
	fmt.Println("Décimo Terc: ", decimoTerc)
	fmt.Println("FGTS: ", fgts)

	fmt.Println("Salário Liq: ", utilities.Round2(salBrt-inssContrib-irpf))

	totalDespesas := ferias + umTercFerias + decimoTerc + fgts + unimed + VA + contabilidade

	fmt.Println("Total de despesas PJ: ", totalDespesas)
	fmt.Println("Fat: ", salBrt+totalDespesas)

}

func Calc(c *utilities.Calculo) {

	// var valorFinal float64
	var entradas float64
	var impostos float64
	// var descontos float64
	var obrigacoes float64

	// Entrada
	for _, val := range c.Entrada {
		entradas += val
	}

	// Impostos
	for _, val := range c.Impostos {
		impostos += val * entradas
	}

	fmt.Printf("Valor de impostos: %v\n", impostos)

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
