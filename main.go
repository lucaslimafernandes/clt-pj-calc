package main

import (
	"flag"
	"fmt"

	"github.com/lucaslimafernandes/clt-pj-calc/internal/utilities"
)

type Relatorio struct {
	salarioCLT    float64
	contribINSS   float64
	ImpostoIRPF   float64
	liquidoCLT    float64
	reservas      float64
	faturamentoPJ float64
	liquidoPJ     float64
	totalDespesas float64
	custosPJ      float64
	impostosPJ    float64
	cfg           *utilities.Cfg
}

func main() {

	help := flag.Bool("help", false, "Show available commands")
	generate_cfg := flag.Bool("gen-config", false, "Exporta o arquivo de configuração")
	salBrt := flag.Float64("salario", 0.0, "Digite o salário bruto, fmt 1000.10")

	flag.Parse()

	if *help {
		fmt.Printf(`[project]
name = "clt-pj-calc"
version = "0.1.0"
description = "Calculadora de salário."

-help			Exibe esta ajuda
-gen-config		Cria um arquivo de configuração .toml
-salario		Passe o valor do salário bruto CLT

Deve possuir o arquivo 'calc.toml' configurado

[repository]
url = "https://github.com/lucaslimafernandes/clt-pj-calc"

`)
		return
	}

	if *generate_cfg {
		utilities.CreateToml()
		return
	}

	c := utilities.ReadToml()
	handler(*salBrt, c)

}

// Manipulador Handler
func handler(salBrt float64, cfg *utilities.Cfg) {

	inssContrib := utilities.CalcINSS(salBrt)
	irpfContrib := utilities.CalcIRPF(1, salBrt, inssContrib)

	ferias := utilities.Round2(salBrt * 0.09)
	umTercFerias := utilities.Round2(salBrt * 0.03)
	decimoTerc := utilities.Round2(salBrt * 0.09)
	fgts := utilities.Round2(salBrt * 0.08)
	planoSaude := cfg.CustosFixos["planoSaude"]
	valeRefeicao := cfg.CustosFixos["valeRefeicao"]
	contabilidade := cfg.CustosFixos["contabilidade"]

	salLiq := utilities.Round2(salBrt - inssContrib - irpfContrib)

	totalDespesas := ferias + umTercFerias + decimoTerc + fgts + planoSaude + valeRefeicao + contabilidade

	reservas := 0.0
	for _, v := range cfg.Reservas {
		if v <= 1.0 {
			reservas += salBrt * v
		} else {
			reservas += v
		}
	}

	fat := salBrt + totalDespesas
	liqPJ := fat - totalDespesas
	impPJ := fat * 0.093

	r := Relatorio{
		salarioCLT:    salBrt,
		contribINSS:   inssContrib,
		ImpostoIRPF:   irpfContrib,
		liquidoCLT:    salLiq,
		reservas:      reservas,
		faturamentoPJ: fat,
		liquidoPJ:     liqPJ,
		totalDespesas: totalDespesas,
		custosPJ:      contabilidade,
		impostosPJ:    impPJ,
		cfg:           cfg,
	}

	relatorio(&r)

}

// Imprimir o relatório
func relatorio(r *Relatorio) {

	fmt.Printf(
		`=== Comparativo CLT x PJ ===
CLT:
  Bruto:          R$ %.2f
  INSS:           R$ %.2f
  IRPF:           R$ %.2f
  Líquido:        R$ %.2f
  Reservas eq.:   R$ %.2f  (férias %.2f%%, 1/3 férias %.2f%%, 13º %.2f%%, FGTS %.2f%%)
  Vale-refeiçãp   R$ %.2f
  Plano de saúde  R$ %.2f

PJ (alvo: manter Líquido CLT + Reservas):
  Pro-labore :    R$ %.2f
  Reservas   :    R$ %.2f
  Simples (≈):    %.2f%%

  Faturamento:    R$ %.2f
  Impostos:       R$ %.2f
  Custos fixos:   R$ %.2f
  Líquido PJ:     R$ %.2f

Cálculo Valor/Hora => R$ %.2f por mês
  160hrs		  R$ %.2f por hora
  168hrs		  R$ %.2f por hora
`,
		r.salarioCLT,
		r.contribINSS*-1,
		r.ImpostoIRPF*-1,
		r.liquidoCLT,
		r.reservas, r.cfg.Reservas["ferias"]*100, r.cfg.Reservas["adicionalferias"]*100, r.cfg.Reservas["decimoterceiro"]*100, r.cfg.Reservas["fgts"]*100,
		r.cfg.CustosFixos["valeRefeicao"],
		r.cfg.CustosFixos["planoSaude"],
		r.salarioCLT,
		r.reservas,
		r.cfg.PJ["simples"]*100,
		r.faturamentoPJ,
		r.impostosPJ,
		r.custosPJ,
		r.liquidoPJ,
		r.faturamentoPJ,
		r.faturamentoPJ/160,
		r.faturamentoPJ/168,
	)
}
