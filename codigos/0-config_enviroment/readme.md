# Configuração de ambiente

## Instalação

### Windows
[Download Link](https://www.go.dev/dl/)

### Linux
`sudo apt update`

`sudo apt install golang-go`

### Verificação de instalação foi bem sucedida
`go version`

Necessáriamente deverá aparecer a versão do Go

Caso resulte em "*comando inválido*" significa que a sua variável PATH não está apontando para a pasta do Go

- (*windows*) C:\Program Files\Go\
- (*Linux*) /usr/local/go/          


---

## Solução

Para utilizarmos os pacotes que criamos em nosso projeto, de forma mais tranquila, sem precisarmos salvar nossos pacotes em outro diretório
usando o exemplo de importação contido no código main do diretório atual (*math/matematica*)

Iremos executar os seguintes comandos:

`go mod init math`

`go mod tidy`