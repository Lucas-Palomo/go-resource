# go-resource v2

[English](./README.md) | [Português (Brasil)](./README.pt_br.md)

`go-resource/v2` é a linha major estável para desenvolvimento novo.

Ela preserva a ideia inspirada em `ResourceBundle`, mas melhora o contrato nos pontos que realmente importam para uma biblioteca Go reutilizável: tratamento de erro, modelo de recursos, pontos de extensão e abstração de filesystem.

## Status de release

- versão estável atual: `v2.0.1`
- tag estável anterior: `v2.0.0`, agora retraída porque a documentação publicada estava semanticamente incorreta
- path do módulo: `github.com/Lucas-Palomo/go-resource/v2`
- público recomendado: projetos novos e migrações estruturadas a partir da v1

## Instalação

```bash
go get github.com/Lucas-Palomo/go-resource/v2@v2.0.1
```

```go
import resource "github.com/Lucas-Palomo/go-resource/v2"
```

## Recursos principais

- erros explícitos de carregamento
- `LoadDir(root string)` e `LoadFS(fsys fs.FS, root string)`
- suporte a `fs.FS` e `embed.FS`
- JSON, YAML, TOML e `.properties` no estilo Java
- objetos aninhados achatados em notação por ponto
- composição de namespace a partir de pastas e segmentos do nome do arquivo
- cadeia de fallback de locale baseada em `golang.org/x/text/language`
- estratégias configuráveis para chave ausente e chave duplicada
- helpers de runtime como `Has`, `HasFor`, `Locales`, `Catalog` e `Reset`
- extensão de decoder via `RegisterDecoder`
- leituras thread-safe

## Convenções de recursos

### Namespace por pasta

```text
resources/
  errors/
    en.json
  messages/
    checkout/
      pt_BR.yaml
```

### Namespace por nome de arquivo

```text
resources/
  en.messages.checkout.toml
  pt_BR.messages.checkout.toml
```

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

Vira `checkout.button.confirm`.

Quando o arquivo já está escopado por namespace de pasta ou de nome de arquivo, os prefixos são combinados. Exemplo: `resources/errors/en.json` com o JSON acima produz `errors.checkout.button.confirm`.

### Normalização de escalares

A v2 normaliza o catálogo decodificado em `map[string]string`.

Isso significa:

- strings permanecem strings
- números são convertidos com `fmt.Sprint`
- booleanos são convertidos com `fmt.Sprint`
- valores `nil` são rejeitados
- valores folha compostos não suportados retornam `ErrInvalidResourceValue`

## `.properties` no estilo Java

Comportamentos suportados:

- separadores: `=`, `:`, ou whitespace
- comentários com `#` ou `!`
- continuação de linha com barra invertida no final
- escapes como `\t`, `\n`, `\r`, `\f` e `\uXXXX`

Exemplo:

```properties
title = Checkout
checkout.hello = Hello, %s
errors.validation.required: Required field
```

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

## Comportamento padrão

`resource.New()` inicia com estes defaults:

- locale de fallback: `language.English`
- locale atual: locale de fallback quando não for definida explicitamente
- estratégia para chave ausente: `ReturnKeyOnMissing`
- estratégia para chave duplicada: `ErrorOnDuplicate`

## Documentação

- [Referência da API](./docs/api-reference.pt_br.md)
- [Notas de arquitetura](./docs/architecture.pt_br.md)
- [Migração da v1](./docs/migration-v1-to-v2.pt_br.md)
- [Changelog da v2](./CHANGELOG.pt_br.md)
