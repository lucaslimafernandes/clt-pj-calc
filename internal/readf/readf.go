package readf

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type Calculo struct {
	Entrada      map[string]float64
	Impostos     map[string]float64
	Obrigacoes   map[string]float64
	Recomendados map[string]float64
	Reservas     map[string]float64
}

func ReadToml() *Calculo {

	var calc Calculo

	f, err := os.ReadFile("calc.toml")
	if err != nil {
		log.Fatalln("Falha ao ler arquivo de cálculo")
	}

	_, err = toml.Decode(string(f), &calc)
	if err != nil {
		log.Fatalln("Falha ao decodificar o arquivo de cálculo")
	}

	return &calc

}
