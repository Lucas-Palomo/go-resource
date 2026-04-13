# go-resource v2

[English](./README.md) | [Português (Brasil)](./README.pt_br.md)

`go-resource/v2` é uma biblioteca de i18n orientada a arquivos para Go.

Ela mantém a ideia original inspirada em `ResourceBundle`: carregar textos localizados a partir de arquivos e resolver chaves por locale. Na v2, essa proposta ganha um contrato de runtime mais claro, carregamento explícito, melhores pontos de extensão e suporte às abstrações modernas de filesystem do Go, como `fs.FS` e `embed.FS`.

## Release atual

- release estável da v2: `v2.0.1`
- path do módulo: `github.com/Lucas-Palomo/go-resource/v2`
- recomendada para: projetos novos e migrações a partir da v1

## Instalação

```bash
go get github.com/Lucas-Palomo/go-resource/v2@v2.0.1
```

```go
import resource "github.com/Lucas-Palomo/go-resource/v2"
```

## O que a v2 entrega

- carregamento explícito com `LoadDir` e `LoadFS`
- suporte a `fs.FS` e `embed.FS`
- JSON, YAML, TOML e `.properties` no estilo Java
- achatamento de objetos aninhados em notação por ponto
- composição de namespace a partir de pastas e segmentos do nome do arquivo
- fallback de locale com base em `golang.org/x/text/language`
- estratégias configuráveis para chave ausente e chave duplicada
- helpers de runtime como `Has`, `Locales`, `Catalog` e `Reset`
- extensibilidade de decoder com `RegisterDecoder`
- leituras thread-safe após o carregamento

## Quick start

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

## Formatos de recurso suportados

Os decoders nativos cobrem:

- `.json`
- `.yaml`
- `.yml`
- `.toml`
- `.properties`

### `.properties` no estilo Java

Os comportamentos suportados incluem:

- separadores com `=`, `:` ou espaço em branco
- comentários com `#` ou `!`
- continuação de linha com barra invertida no final
- sequências de escape como `\t`, `\n`, `\r`, `\f` e `\uXXXX`

Exemplo:

```properties
title = Checkout
checkout.hello = Hello, %s
errors.validation.required: Required field
```

## Como os caminhos viram chaves

A v2 combina três fontes para formar a chave final de lookup:

1. nomes de pastas abaixo da raiz de recursos
2. segmentos do nome do arquivo depois do locale e antes da extensão
3. objetos aninhados dentro do arquivo decodificado

### Namespace por pasta

```text
resources/
  errors/
    en.json
  messages/
    checkout/
      pt_BR.yaml
```

Isso produz chaves sob `errors.*` e `messages.checkout.*`.

### Namespace por nome de arquivo

```text
resources/
  en.messages.checkout.toml
  pt_BR.messages.checkout.toml
```

Isso produz chaves sob `messages.checkout.*`.

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

Isso vira `checkout.button.confirm`.

Quando o namespace de pasta é combinado com objetos aninhados, os prefixos são unidos. Exemplo: `resources/errors/en.json` com o JSON acima vira `errors.checkout.button.confirm`.

## Comportamentos de runtime que importam

### Lookup estrito vs helper leniente

Use `Lookup` e `LookupFor` quando o erro importa.

Use `Get` e `GetFor` quando retornar a própria chave for um fallback aceitável.

Essa distinção importa: `Get` e `GetFor` retornam a chave original em erros de lookup, inclusive para chave ausente e `ErrBundleNotLoaded`.

### Carregamentos repetidos fazem merge em memória

`LoadDir` e `LoadFS` são incrementais. Um segundo carregamento bem-sucedido faz merge de novos catálogos no bundle já carregado em memória.

Quando você quiser semântica de substituição, chame `Reset()` antes de carregar novamente.

### Arquivos não suportados falham de forma explícita

Arquivos com extensão não suportada não são ignorados. O carregamento falha com `ErrUnsupportedFormat`.

Mantenha arquivos auxiliares fora da árvore de recursos, ou registre um decoder para a extensão desejada.

### Normalização de escalares

O catálogo de runtime é normalizado para `map[string]string`.

Isso significa:

- strings continuam strings
- números são convertidos com `fmt.Sprint`
- booleanos são convertidos com `fmt.Sprint`
- valores `nil` são rejeitados
- valores folha compostos não suportados retornam `ErrInvalidResourceValue`

## Comportamento padrão

`resource.New()` inicia com:

- fallback locale: `language.English`
- locale atual: fallback locale quando não definido explicitamente
- estratégia para chave ausente: `ReturnKeyOnMissing`
- estratégia para chave duplicada: `ErrorOnDuplicate`

## Documentação

- [Referência da API](./docs/api-reference.pt_br.md)
- [Notas de arquitetura](./docs/architecture.pt_br.md)
- [Migração da v1 para a v2](./docs/migration-v1-to-v2.pt_br.md)
- [Changelog da v2](./CHANGELOG.pt_br.md)
