# Gerenciamento de pacotes em GoLang

## Contexto

A linguagem GoLang tem como padrão buscar seus pacotes importados em seu diretório raíz (*GOROOT*), normalmente localizado em: 
- C:\Program Files\Go\src\curso-go\ {pacote} (*windows*)
- /usr/local/go/src/ {pacote} (*Linux*)
---

## Solução

Para utilizarmos os pacotes que criamos em nosso projeto, de forma mais tranquila, sem precisarmos salvar nossos pacotes em outro diretório
usando o exemplo de importação contido no código main do diretório atual (*math/matematica*)

Iremos executar os seguintes comandos:

`go mod init math`

`go mod tidy`