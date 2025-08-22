package utilities

import "math"

// Exemplo:
// Para um salário de R$ 3.500,00, o cálculo seria:
// Faixa 1 (até R$ 1.518,00): R$ 1.518,00 * 7,5% = R$ 113,85.
// Faixa 2 (de R$ 1.518,01 até R$ 2.793,88): (R$ 2.793,88 - R$ 1.518,00) * 9% = R$ 114,83.
// Faixa 3 (de R$ 2.793,89 até R$ 3.500,00): (R$ 3.500,00 - R$ 2.793,89) * 12% = R$ 84,73.
// Contribuição total: R$ 113,85 + R$ 114,83 + R$ 84,73 = R$ 313,41.

type Faixa struct {
	Teto     float64
	Aliquota float64
}

// Valores de contribuição do INSS 2025
var faixas = []Faixa{
	{1518.00, 0.075},
	{2793.88, 0.090},
	{4190.83, 0.120},
	{8157.41, 0.140}, // teto máximo de contribuição; acima disso não aumenta
}

// Cálcula o valor de contribuição do INSS
func CalcINSS(val float64) float64 {

	if val <= 0 {
		return 0
	}

	contrib := 0.0
	limInf := 0.0

	for _, f := range faixas {
		base := math.Min(val, f.Teto) - limInf
		if base <= 0 {
			break
		}
		contrib += base * f.Aliquota
		limInf = f.Teto
	}

	return Round2(contrib)
}
