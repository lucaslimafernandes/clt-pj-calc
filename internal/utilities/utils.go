package utilities

import "math"

// Retorna valor arredondado 2 casas decimais
func Round2(v float64) float64 {
	return math.Round(v*100) / 100
}
