# clt-pj-calc

Autor: Lucas Lima Fernandes
Github: https://github.com/lucaslimafernandes
Linkedin: https://www.linkedin.com/in/lucaslimafernandes/

Calculadora de comparação entre contratação CLT e PJ.  
O objetivo é estimar o custo total e o valor líquido em cada regime, considerando impostos, reservas e custos fixos.

## Instalação

Clone o repositório:

```bash
git clone https://github.com/lucaslimafernandes/clt-pj-calc.git
cd clt-pj-calc
```

Baixe as dependências:

```bash
go mod tidy
```

## Configuração

É necessário configurar o arquivo calc.toml com os parâmetros de custos, reservas e impostos.
Um exemplo de configuração já está disponível no repositório.

Para gerar um arquivo de configuração padrão:

```bash
go run main.go -gen-config
```

### Arquivo calc.toml

O arquivo calc.toml define os parâmetros usados nos cálculos.
Exemplo:

```toml
[PJ]
simples = 0.06
inss = 0.033

[PF]
dependentes = 1.00

[CustosFixos]
valeRefeicao = 1000.00
planoSaude = 700.00
contabilidade = 300.00

[Reservas]
fgts = 0.08
ferias = 0.09
adicionalferias = 0.03
decimoterceiro = 0.09
emergencia = 0.0
```

Parâmetros:
[PJ]
- simples: alíquota do Simples Nacional aplicada sobre o faturamento.
- inss: contribuição do INSS sobre o pró-labore.

[PF]
- dependentes: número de dependentes para dedução do IRPF.

[CustosFixos]
- valeRefeicao, planoSaude, contabilidade: custos mensais adicionados ao cálculo do PJ.

[Reservas]
- Percentuais do salário bruto destinados a FGTS, férias, adicional de férias e 13º salário
- reserva de emergência.

## Uso

Rodar a aplicação passando o salário bruto CLT:

```bash
go run main.go -salario 7500
```

Também é possível exibir a ajuda:

```bash
go run main.go -help
```

## Estrutura do Projeto

main.go: ponto de entrada da aplicação.

internal/utilities/: funções auxiliares de cálculo (INSS, IRPF, arredondamentos, leitura de config).

calc.toml: arquivo de configuração de custos e impostos.

## Contribua

Contribuições são bem-vindas.
Para colaborar:

1. Faça um fork do repositório.

2. Crie uma branch para sua feature ou correção:

```bash
git checkout -b minha-feature
```

3. Commit suas alterações:

```bash
git commit -m "Descrição da mudança"
```

4. Envie a branch para o seu fork:

```bash
git push origin minha-feature
```

Abra um Pull Request.

## Licença

Este projeto está licenciado sob os termos da licença MIT. Veja o arquivo LICENSE
 para mais informações.

