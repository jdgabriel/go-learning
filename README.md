# Aprendendo GO
## Download & Instalação

### Download
```bash
https://go.dev/doc/install
```
### Instalando com ZSH
1. Instalando go (Substitua `X` pela versão atual)
```bash
 rm -rf /usr/local/go && tar -C /usr/local -xzf goX.XX.X.linux-amd64.tar.gz
```
2. Adicione os paths do GO no `.zshrc`
```bash
# Add Go lang paths
export GOPATH=$HOME/go
export GOROOT=/usr/local/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOPATH
export PATH=$PATH:$GOROOT/bin
```
3. Teste o funcionamento
```bash
go version
```
## Start
### Iniciando o módulo principal
Por convenção e padronização o go usa packages externos da linguagem com o link do repositório, seja Github ou de outro lugar.<br/>
Para iniciar uma aplicação em go, precisamos iniciar o `go mod` 
```bash
go mod init github.com/jdgabriel/go-learning
```
Ao iniciar o módulo será criado o arquivo: `go.mod` 
```go
// go.mod
module github.com/jdgabriel/go-learning

go 1.22.4
```
## Convenções
### Métodos públicos e privados do pacote
Todo novo módulo deve-se ter na primeira linha do arquivo o nome do pacote
```go
package NOME_DO_MODULO
// e.g. package main
// e.g. package router
// e.g. package user
```
| Norma  | Tipo | Exemplo |
|---|---|---|
|  Primeira letra maiúscula | PUBLICA  | ```func Initialize(){ //... }```|
|  Primeira letra minúscula | PRIVADA  | ```func initialize(){ //... }``` |
