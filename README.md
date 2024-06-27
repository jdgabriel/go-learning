# Aprendendo GO

# Sumário
- [1. Download & Instalação](#download_install)
  - [1.1. Download](#donwload)
  - [1.2. Instalando com ZSH](#install)
- [2. Start](#start)
  - [2.1. Iniciando o módulo principal](#start_init)
- [3. Convenções](#conventions)
  - [3.1. Métodos públicos e privados do pacote](#conventions_methods)


## <a name="download_install"></a> Download & Instalação

### <a name="download"></a> Download
```bash
https://go.dev/doc/install
```
### <a name="install"></a> Instalando com ZSH
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
## <a name="start"></a> Start
### <a name="start_init"></a> Iniciando o módulo principal
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
## <a name="conventions"></a> Convenções
### <a name="conventions_methods"></a> Métodos públicos e privados do pacote
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
