# go-resource v2

[English](./README.md) | [Português (Brasil)](./README.pt_br.md)

**i18n em estilo ResourceBundle para Go** — uma reescrita mais moderna, segura e extensível do `go-resource`, inspirada na ideia dos Java ResourceBundles e adaptada para internacionalização em Golang.

Esta é a **linha recomendada para novos projetos**.

## Instalação

```bash
go get github.com/Lucas-Palomo/go-resource/v2@latest
```

## Import

```go
import (
	resource "github.com/Lucas-Palomo/go-resource/v2"
	"golang.org/x/text/language"
)
```

## Por que a v2 existe

A v1 original funcionava bem para projetos pequenos, mas misturava walking de diretórios, parsing, carregamento e lookup em um único fluxo, usava `panic` no caminho principal e tinha um fallback que podia ficar inconsistente quando uma locale estava ausente.

A v2 introduz:

- tratamento explícito de erros
- `LoadDir` e `LoadFS`
- suporte a `embed.FS`
- decoders para JSON, YAML, TOML e arquivos `.properties` no estilo Java
- flatten de estruturas aninhadas com notação por ponto
- resolução previsível de fallback
- comportamento configurável para chaves ausentes e duplicadas
- leituras thread-safe
- exemplos em nível de pacote no estilo pkg.go.dev
- helpers de inspeção como `Has`, `HasFor` e `Reset`

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
checkout.description = Confirm your order \
  before leaving
welcome = Ol\u00E1
```

Esse formato aceita os separadores usuais do Java (`=`, `:`, ou whitespace), comentários com `#` ou `!`, continuação de linha com barra invertida no final e escapes como `\t`, `\n` e `\uXXXX`.

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

	fmt.Println(bundle.Get("title"))
	fmt.Println(bundle.Get("checkout.hello", "Lucas"))
	fmt.Println(bundle.Get("errors.validation.required"))
}
```

Exemplos executáveis estão disponíveis em [`./examples/basic`](./examples/basic) e [`./examples/properties`](./examples/properties).

## Docs

- [Referência da API](./docs/api-reference.pt_br.md)
- [API reference](./docs/api-reference.md)
- [Notas de arquitetura](./docs/architecture.pt_br.md)
- [Architecture notes](./docs/architecture.md)
- [Guia de migração](./docs/migration-v1-to-v2.pt_br.md)
- [Migration guide](./docs/migration-v1-to-v2.md)

## Nota do repositório

A raiz do repositório mantém o módulo legado da v1 por compatibilidade. Este módulo vive em `/v2` para preservar o semantic import versioning de major version no Go.
