package utilities

// Rendimentos previdenciários isentos para maiores de 65 anos: R$ 1.903,98
// Dedução mensal por dependente: R$ 189,59
// Limite mensal de desconto simplificado: R$ 564,80

// https://www27.receita.fazenda.gov.br/simulador-irpf/

// Exemplos com tabela defasada
// Base de Cálculo (R$)	Alíquota (%)	Dedução do IR (R$)
// Até R$ 2.259,20	zero	zero
// De R$ 2.259,21 até R$ 2.826,65	7,5%	R$ 169,44
// De R$ 2.826,66 até R$ 3.751,05	15%	R$ 381,44
// De R$ 3.751,06 até R$ 4.664,68	22,5%	R$ 662,77
// Acima de R$ 4.664,68	27,5%	R$ 896,00

// Salário Bruto: R$ 4.500,00.
// Base de Cálculo: R$ 4.500,00 (após dedução simplificada de R$ 564,80, restam R$ 3.935,20).
// Aplicando as alíquotas por faixa:
// Faixa 1 (até R$ 2.259,20): Isento.
// Faixa 2 (R$ 2.259,21 até R$ 2.826,65): 7,5% sobre R$ 567,44 = R$ 42,56.
// Faixa 3 (R$ 2.826,66 até R$ 3.751,05): 15% sobre R$ 924,39 = R$ 138,66.
// Faixa 4 (R$ 3.751,06 até R$ 4.664,68): 22,5% sobre R$ 184,15 = R$ 41,93.
// Com um salário bruto de R$ 4.500,00, o desconto de Imposto de Renda seria de R$ 223,15, considerando a tabela progressiva vigente.

const (
	dedDependent = 189.59
)

type FaixaIR struct {
	Limite   float64 // teto superior da faixa
	Aliquota float64 // ex.: 0.075
	Deducao  float64 // dedução fixa para a faixa
}

// Tabela mensal vigente maio/25
var tabelaIR = []FaixaIR{
	{2428.80, 0.000, 0.00},
	{2826.65, 0.075, 182.16},
	{3751.05, 0.150, 394.16},
	{4664.68, 0.225, 675.49},
	{1e15, 0.275, 908.73}, // um “infinito” prático
}

// Cálcula o valor do IRPF
func CalcIRPF(dependentes float64, salBrt float64, inss float64) float64 {

	if salBrt <= 0 {
		return 0.0
	}

	totDeducoes := inss + dependentes*dedDependent
	base := salBrt - totDeducoes

	for _, f := range tabelaIR {

		if base <= f.Limite {
			irDevido := base*f.Aliquota - f.Deducao
			if irDevido <= 0 {
				return 0
			}
			return Round2(irDevido)
		}

	}

	return 0

}
