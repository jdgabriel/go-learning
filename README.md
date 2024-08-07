# Aprendendo GO

# Sumário

- [1. Download & Instalação](#download_install)
  - [1.1. Download](#donwload)
  - [1.2. Instalando com ZSH](#install)
- [2. Start](#start)
  - [2.1. Iniciando o módulo principal](#start_init)
- [3. Convenções](#conventions)
  - [3.1. Métodos públicos e privados do pacote](#conventions_methods)
  - [3.2. Importando pacotes](#conventions_imports)
- [4. Tipos de estruturas e variáveis](#struct)
- [5. Makefile](#make)
- [6. Documentação com Swagger](#docs)

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

| Norma                    | Tipo    | Exemplo                      |
| ------------------------ | ------- | ---------------------------- |
| Primeira letra maiúscula | PUBLICA | `func Initialize(){ //... }` |
| Primeira letra minúscula | PRIVADA | `func initialize(){ //... }` |

### <a name="conventions_imports"></a> Importando pacotes

Para importar um pacote temos duas maneiras, a direta e a indireta

#### Método direto

```bash
go get -u github.com/gin-gonic/gin
```

#### Método indireto

Importar dentro do arquivo que vai usar o módulo

```bash
import "github.com/gin-gonic/gin"
```

Assim temos que executar o seguinte comando. Este comando faz o download de todos os pacotes relacionados diretamente e indiretamente com seu projeto, e pacotes existentes.

```bash
go mod tidy
```

Após o comando será criado um arquivo, o `go.sum` onde temos o lock de todas as versões dos pacotes instalados no projeto.<br>
Agora o o `go.mod` terá dois `require`

```go
module github.com/jdgabriel/go-learning

go 1.22.4

// direct
require github.com/gin-gonic/gin v1.10.0

require (
	... // indirect
)
```

### <a name="struct"></a> Tipos de estruturas e variáveis

| Tipo        | Padrão                            | ~                                                                                             | Exemplo                                        |
| ----------- | --------------------------------- | --------------------------------------------------------------------------------------------- | ---------------------------------------------- |
| Declarative | `:=`                              | Criar uma variável e atribuir o retorno                                                       | `v1 := router.Group("/api/v1")`                |
| Pointer     | `*`                               | Endereço do objeto em memória para ser usado em qualquer ponto do projeto                     | `*gin.Engine`                                  |
| Struct      | `struct`                          | Equivalente a um `objeto`                                                                     | `type Product struct {...}`                    |
| Address     | `&`                               | Indicação de endereço de uma struct criada                                                    | `return &Logger {...}`                         |
| Func Args   | `...interface{}` OU `interface{}` | Representa a interface ou um "array" de argumentos da função. Com interface fazia ou definida | `func (log *Logger) Info(v ...interface{}) {}` |

### <a name="make"></a> Makefile

#### <a name="make_script"></a> Scripts para rodar o projeto

Rode qualquer comando com `make <COMANDO>`

1. Rodar o projeto | `make run`

```bash
run:
	@go run main.go
```

2. Rodar o projeto com documentação `:8080/docs/index.html`

```bash
run-with-docs:
	@swag init
	@go run main.go
```

3. Build do projeto para a pasta `cmd/`

```bash
build:
	@go build -o cmd/$(APP_NAME) main.go
```

3. Build do projeto para a pasta `cmd/`

```bash
test:
	@go test ./ ...
```

4. Iniciar documentação

```bash
docs:
	@swag init
```

5. Limpar pastas de documentos e binário do projeto

```bash
clean:
	@rm -f cmd/$(APP_NAME)
	@rm -rf ./docs
```

### <a name="docs"></a> Documentação

#### <a name="docs_swagger"></a> Swagger

```bash
http://localhost:8080/docs/index.html
```
