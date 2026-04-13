# go-resource

[English](./README.md) | [Português (Brasil)](./README.pt_br.md)

Internacionalização baseada em arquivos para Go, inspirada no `ResourceBundle` do Java, mas adaptada para uso idiomático em Go.

Este repositório contém **duas linhas de módulo** e **quatro tags publicadas**:

| Tag | Linha | Path do módulo | Data       | Status |
|---|---|---|------------|---|
| `v1.0.0` | v1 | `github.com/Lucas-Palomo/go-resource` | 2024-08-20 | Release histórica, retraída |
| `v1.0.1` | v1 | `github.com/Lucas-Palomo/go-resource` | 2026-04-11 | Release estável de manutenção |
| `v2.0.0` | v2 | `github.com/Lucas-Palomo/go-resource/v2` | 2026-04-11 | Retraída por documentação semanticamente incorreta |
| `v2.0.1` | v2 | `github.com/Lucas-Palomo/go-resource/v2` | 2026-04-13 | Release estável atual da v2 |

## O que este projeto faz

`go-resource` carrega mensagens localizadas a partir de arquivos de recurso, em vez de espalhar textos por handlers, services e lógica de domínio.

Formatos suportados ao longo do repositório:

- JSON
- YAML / YML
- TOML
- `.properties` no estilo Java na v2

## Qual linha você deve usar?

### Use a v1 quando

- você já importa `github.com/Lucas-Palomo/go-resource/pkg/resource`
- quer o menor custo de migração possível
- precisa apenas da API original com as correções de estabilidade da `v1.0.1`

### Use a v2 quando

- está começando trabalho novo
- quer tratamento explícito de erro
- precisa de `fs.FS` ou `embed.FS`
- quer objetos aninhados achatados em notação por ponto
- quer composição de namespace por diretório e por segmentos do nome do arquivo
- precisa de políticas configuráveis para chaves ausentes ou duplicadas
- quer suporte nativo a `.properties`

## Layout do repositório

```text
.
├── go.mod                         # módulo v1: github.com/Lucas-Palomo/go-resource
├── pkg/resource                   # pacote v1
├── docs                           # documentação focada na v1
├── examples                       # recursos de exemplo da v1
└── v2                             # módulo v2: github.com/Lucas-Palomo/go-resource/v2
```

## Instalação

### v1

```bash
go get github.com/Lucas-Palomo/go-resource@v1.0.1
```

```go
import "github.com/Lucas-Palomo/go-resource/pkg/resource"
```

### v2

```bash
go get github.com/Lucas-Palomo/go-resource/v2@v2.0.1
```

```go
import resource "github.com/Lucas-Palomo/go-resource/v2"
```

## Modelo de recursos por linha

### v1

A v1 carrega catálogos **planos do tipo `map[string]string`** a partir de arquivos JSON, YAML e TOML.

Consequências importantes:

- os valores precisam ser strings
- objetos aninhados não são um modelo de recurso suportado
- diretórios servem apenas para organização física
- segmentos extras no nome do arquivo, como `en.errors.json`, **não** viram namespace de chave

Exemplo:

```text
resources/
  errors/
    en.json
  messages/
    pt_BR.yaml
```

Isso é válido na v1, mas as chaves continuam exatamente como foram declaradas dentro dos arquivos.

### v2

A v2 carrega arquivos de recurso para **catálogos planos em runtime**, com regras de entrada mais ricas:

- objetos aninhados são achatados em notação por ponto
- nomes de diretório viram segmentos de namespace
- segmentos extras do nome do arquivo viram segmentos de namespace
- `.properties` é suportado nativamente
- escalares como strings, números e booleanos são convertidos para string no catálogo final

Exemplo:

```text
resources/
  errors/
    en.json
  en.messages.checkout.toml
```

Com:

```json
{
  "validation": {
    "required": "Required field"
  }
}
```

Produz:

```text
errors.validation.required
```

## Quick start: v1

```go
package main

import (
    "fmt"
    "log"

    "github.com/Lucas-Palomo/go-resource/pkg/resource"
    "golang.org/x/text/language"
)

func main() {
    bundle := resource.NewBundle("./resources", language.English)

    if err := bundle.LoadWithError(); err != nil {
        log.Fatal(err)
    }

    bundle.SetLocale(language.BrazilianPortuguese)

    fmt.Println(bundle.Get("title"))
    fmt.Println(bundle.Get("hello", "Lucas"))
}
```

## Quick start: v2

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

## Documentação

### Repositório e v1

- [Estratégia de versionamento e release](./docs/versioning-strategy.pt_br.md)
- [Referência da v1](./docs/v1-reference.pt_br.md)
- [Changelog da raiz](./CHANGELOG.pt_br.md)

### v2

- [README da v2](./v2/README.pt_br.md)
- [Referência da API da v2](./v2/docs/api-reference.pt_br.md)
- [Notas de arquitetura da v2](./v2/docs/architecture.pt_br.md)
- [Migração da v1 para a v2](./v2/docs/migration-v1-to-v2.pt_br.md)
- [Changelog da v2](./v2/CHANGELOG.pt_br.md)

## Licença

Este repositório é distribuído sob **The Unlicense**.
