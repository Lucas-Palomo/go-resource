# go-resource v2

[English](./README.md) | [Português (Brasil)](./README.pt_br.md)

**i18n em estilo ResourceBundle para Go** — uma reescrita mais moderna, segura e extensível do `go-resource`.

Este módulo é a **próxima major** do projeto. Ele representa a direção de longo prazo da biblioteca e é o alvo recomendado de migração para quem excedeu os limites da v1.

> Até que `v2.0.0` seja tagueada, consumidores que usam apenas releases estáveis devem permanecer na `v1.0.1`.

## Path do módulo

```go
import resource "github.com/Lucas-Palomo/go-resource/v2"
```

## Instalação

Depois que a primeira release da v2 for publicada:

```bash
go get github.com/Lucas-Palomo/go-resource/v2@latest
```

## Por que a v2 existe

A v1 foi útil para projetos pequenos, mas tinha limites estruturais:

- falhas de carregamento eram historicamente expostas por `panic`
- walking de diretórios, parsing, carregamento, fallback e lookup estavam fortemente acoplados
- não havia suporte a `fs.FS` / `embed.FS`
- chaves ausentes e duplicadas não tinham política explícita
- objetos aninhados não eram um modelo de recurso de primeira classe

A v2 corrige o contrato em vez de apenas trocar nomes.

## O que a v2 adiciona

- tratamento explícito de erro
- `LoadDir(root string)` e `LoadFS(fsys fs.FS, root string)`
- suporte a `fs.FS` e `embed.FS`
- decoders para JSON, YAML, TOML e arquivos `.properties` no estilo Java
- flatten de estruturas aninhadas para notação por ponto
- namespacing por diretório e por segmentos do nome do arquivo
- resolução previsível de fallback
- políticas configuráveis para chaves ausentes e duplicadas
- leituras thread-safe
- helpers como `Has`, `HasFor` e `Reset`

## Convenções de arquivos de recurso

### Namespace por pastas

```text
/resources
  /errors
    en.json
    pt_BR.json
  /messages
    /checkout
      en.yaml
      pt_BR.yaml
```

### Namespace por nome de arquivo

```text
/resources
  en.errors.json
  pt_BR.errors.json
  en.messages.checkout.toml
```

### Arquivos `.properties` no estilo Java

```properties
title = Checkout
checkout.hello = Hello, %s
errors.validation.required: Required
checkout.description = Confirm your order  before leaving
welcome = Olá
```

Recursos suportados em `.properties`:

- separadores: `=`, `:`, ou whitespace
- comentários com `#` ou `!`
- continuação de linha com barra invertida no final
- escapes como `\t`, `\n` e `\uXXXX`

### Objetos aninhados

```json
{
  "checkout": {
    "button": {
      "confirm": "Confirm"
    }
  }
}
```

vira:

```text
checkout.button.confirm
```

## Exemplo

```go
package main

import (
	"fmt"
	"log"

	resource "github.com/Lucas-Palomo/go-resource/v2"
	"golang.org/x/text/language"
)

func main() {
	bundle := resource.New(
		resource.WithFallbackLocale(language.English),
		resource.WithLocale(language.BrazilianPortuguese),
	)

	if err := bundle.LoadDir("./resources"); err != nil {
		log.Fatal(err)
	}

	fmt.Println(bundle.Get("checkout.title"))
	fmt.Println(bundle.Get("checkout.hello", "Lucas"))
	fmt.Println(bundle.Get("errors.validation.required"))
}
```

Exemplos executáveis estão disponíveis em:

- [`./examples/basic`](./examples/basic)
- [`./examples/properties`](./examples/properties)

## Docs

- [Referência da API](./docs/api-reference.pt_br.md)
- [Notas de arquitetura](./docs/architecture.pt_br.md)
- [Guia de migração](./docs/migration-v1-to-v2.pt_br.md)

## Relação com a raiz do repositório

A raiz do repositório mantém a linha de manutenção da v1 por compatibilidade. A nova major vive em `/v2` para preservar o semantic import versioning do Go.
