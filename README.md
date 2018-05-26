# Prova de conceito: Calculador de compatibilidade VotaSP

Aqui será desenvolvida a prova de conceito da calculadora de compatibilidade do vota-sp, visando obter performance viável para execução em escala

## Prerequisitos para execução

* Docker

## Gerando dados de testes

os dados são gerados como um JSON, estruturado da seguinte maneira:
```json
{
  "candidato-1": ["CP","C","D","DP"],
  "candidato-2": ["C","DP","CP","DP"]
}
```
onde:
* CP: Concordo Plenamente
* C: Concordo
* D: Discordo
* DP: Discordo Plenamente

O json de respostas dos eleitores segue o mesmo padrão, mas com a opção adicional de resposta `I`, equivalente a Indiferente

para gerar os dados, basta executar:
```
docker run -it -v $PWD:/app -w /app node:10-alpine node ./gerador/gerador.js <modo> <quantidade> > dados.json
```
onde `<modo>` é ou `candidato` ou `eleitor`, e `<quantidade>` é o numero de registros a serem gerados

por exemplo:
```
docker run -it -v $PWD:/app -w /app node:10-alpine node ./gerador/gerador.js eleitor 10 > eleitores.json
```

## Algoritmo de pontuação

### Premissas

* Cada candidato deverá responder às 40 perguntas do questionário
  * as opções serão: `Concordo Plenamente`, `Concordo`, `Discordo` e `Discordo Plenamente`
* Cada eleitor deverá responder ao menos 20 perguntas para que sua compatibilidade seja calculada
   * as opções serão as mesmas do candidado, com uma opção adicional: `Indiferente`
   * opções `Indiferente` não são contabilizadas para o total de questões respondidas

### Cálculo da pontuação

Para cada pergunta, serão comparadas as respostas do candidato e do eleitor. A pontuação será acrescida conforme os seguintes critérios:

* Resposta do eleitor é `Indiferente`: **0 pontos**
* Respostas do candidato e do eleitor são distintas e discordantes: **0 pontos**
  * Um concorda e o outro discorda. Por exemplo: `Discordo Plenamente` e `Concordo`
* Respostas do candidato e do eleitor são distintas, porém concordantes: **1 pontos**
  * Por exemplo: `Concordo Plenamente` e `Concordo`
* Respostas do candidato e do eleitor são iguais: **2 pontos**

Com isso, a pontuação total de compatibilidade ficará entre 0 e 80

## Algoritmo de busca

Esse algoritmo deverá, dado o conjunto de respostas de um eleitor, e dados todos os conjuntos de respostas dos candidatos, encontrar o 10 candidatos com respostas mais compatíveis

## Avaliação dos algoritmos

para comparar os algoritmos criados para busca dos candidatos mais compatíveis com cada eleitor, vamos utilizar sempre as mesmas massas de dados nos testes
a contagem de tempo de execução deve ser feita no próprio algoritmo, e os parâmetros de comparação serão:

### Tempo para conversão dos dados

O tempo utilizado para converter a massa de dados de teste para o formato usado no seu algoritmo pode ser desconsiderado.
A massa de dados foi gerada em json visando facilidade de uso, mas na aplicação final o formato dos dados armazenados não necessariamente será esse.

### Tempo de preparação

Se o algoritmo precisa montar alguma estrutura de dados para tornar a busca mais rápida, esse tempo será contabilizado como `tempo de preparação`.

### Buscas por segundo

Esse valor será calculado com base no tempo de execução para que sejam encontrados todos os matches. A quantidade total de buscas (sempre 10 mil em nossa bateria de testes) dividida pelo tempo total de execução fornecerá a métrica de `buscas por segundo`

### Metodologia para testes

Para padronizar o benchmark, executaremos sempre com os seguintes conjuntos de dados, utilizando o gerador presente neste repositório:
* **10000 eleitores** / **1000 candidatos**
* **10000 eleitores** / **5000 candidatos**
* **10000 eleitores** / **10000 candidatos**
* **10000 eleitores** / **20000 candidatos**

Utilizar os dados disponíveis na pasta `massa-testes`

Para cada conjunto de dados, executar 10 vezes e tirar a média para considerar o resultado final.