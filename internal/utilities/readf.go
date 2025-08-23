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

// Função para salvar arquivo base calc.toml
func CreateToml() error {

	archive := `
[PJ]
simples = 0.06
inss = 0.033
prolabore_min = 1412.00

[PF]
dependentes = 1.00

[CustosFixos]
valeRefeicao = 1000.00
planoSaude = 1500.00
contabilidade = 300.00

[Reservas]
fgts = 0.08
ferias = 0.09
adicionalferias = 0.03
decimoterceiro = 0.09
emergencia = 0.0

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
