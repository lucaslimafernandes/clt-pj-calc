package readf

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type Calculo struct {
	Entrada struct {
		SalarioBruto float64
	}
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
