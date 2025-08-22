package utilities

import (
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type Cfg struct {
	PJ          map[string]float64
	PF          map[string]float64
	CustosFixos map[string]float64
	Reservas    map[string]float64
}

// type Calculo struct {
// 	Entrada      map[string]float64
// 	Impostos     map[string]float64
// 	PF           map[string]float64
// 	Obrigacoes   map[string]float64
// 	Recomendados map[string]float64
// 	Reservas     map[string]float64
// }

func ReadToml() *Cfg {

	var c Cfg

	f, err := os.ReadFile("calc.toml")
	if err != nil {
		log.Fatalln("Falha ao ler arquivo de cálculo")
	}

	_, err = toml.Decode(string(f), &c)
	if err != nil {
		log.Fatalln("Falha ao decodificar o arquivo de cálculo")
	}

	return &c

}

func CreateToml() error {

	archive := `
[Entrada]
salariobruto = 12000.00

[Impostos]
simples = 0.06

[PF]
dependentes = 1

[Obrigacoes]
prolabore = 8157.41
fgts = 0.08
ferias = 0.09
adicionalferias = 0.03
decimoterceiro = 0.09

[Reservas]
emergencia = 1000

[Beneficios] 
ValeRefeicao = 1000
PlanoSaude = 1500
	`

	f, err := os.Create("calc.toml")
	if err != nil {
		return fmt.Errorf("falha ao criar arquivo: %v", err)
	}
	defer f.Close()

	_, err = f.WriteString(archive)
	if err != nil {
		return fmt.Errorf("falha ao escrever no arquivo: %v", err)
	}

	return nil

}
