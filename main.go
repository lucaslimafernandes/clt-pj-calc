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

	c := utilities.ReadToml()

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

	handler(*salBrt, c)

}

func handler(salBrt float64, cfg *utilities.Cfg) {

	inssContrib := utilities.CalcINSS(salBrt)

	fmt.Println("CalcINSS: ", inssContrib)

	irpf := utilities.CalcIRPF(1, salBrt, inssContrib)

	fmt.Println("CalcIRPF: ", irpf)

	ferias := utilities.Round2(salBrt * 0.09)
	umTercFerias := utilities.Round2(salBrt * 0.03)
	decimoTerc := utilities.Round2(salBrt * 0.09)
	fgts := utilities.Round2(salBrt * 0.08)
	planoSaude := cfg.CustosFixos["planoSaude"]
	valeRefeicao := cfg.CustosFixos["valeRefeicao"]
	contabilidade := cfg.CustosFixos["contabilidade"]

	salLiq := utilities.Round2(salBrt - inssContrib - irpf)

	totalDespesas := ferias + umTercFerias + decimoTerc + fgts + planoSaude + valeRefeicao + contabilidade

	reservas := 0.0
	for _, v := range cfg.Reservas {
		if v <= 1.0 {
			reservas += salBrt * v
		} else {
			reservas += v
		}
	}

	// cfg.Reservas
	fat := salBrt + totalDespesas
	liqPJ := fat - totalDespesas
	impPJ := fat * 0.093

	relatorio(salBrt, salLiq, reservas, fat, liqPJ, totalDespesas, impPJ, cfg)

}

func relatorio(salCLT, liqCLT, reservas, fat, liqPJ, custosPJ, impostosPJ float64, cfg *utilities.Cfg) {
	fmt.Printf(
		`=== Comparativo CLT x PJ ===
CLT:
  Bruto:          R$ %.2f
  Líquido:        R$ %.2f
  Reservas eq.:   R$ %.2f  (férias %.2f%%, 1/3 férias %.2f%%, 13º %.2f%%, FGTS %.2f%%)
  Vale-refeiçãp   R$ %.2f
  Plano de saúde  R$ %.2f

PJ (alvo: manter Líquido CLT + Reservas):
  Pro-labore :    R$ %.2f
  Simples (≈):    %.2f%%

  Faturamento:    R$ %.2f
  Impostos:       R$ %.2f
  Custos fixos:   R$ %.2f
  Líquido PJ:     R$ %.2f

Check:
  Alvo (Liq CLT + Reservas):  R$ %.2f
  Diferença:                  R$ %.2f

Cálculo Valor/Hora
  160hrs		  R$ %.2f por hora
  168hrs		  R$ %.2f por hora
`,
		salCLT, liqCLT, reservas,
		cfg.Reservas["ferias"]*100, cfg.Reservas["adicionalferias"]*100, cfg.Reservas["decimoterceiro"]*100, cfg.Reservas["fgts"]*100,
		cfg.CustosFixos["valeRefeicao"], cfg.CustosFixos["planoSaude"],
		salCLT, cfg.PJ["simples"]*100,
		fat, impostosPJ, custosPJ, liqPJ,
		liqCLT+reservas, (liqPJ - (liqCLT + reservas)),
		fat/160, fat/168,
	)
}
