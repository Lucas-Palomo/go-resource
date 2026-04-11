# go-resource

[English](./README.md) | [Português (Brasil)](./README.pt_br.md)

**Internacionalização para Go inspirada em ResourceBundle.**

`go-resource` é uma biblioteca de i18n baseada em arquivos para Go, inspirada no Java ResourceBundle e adaptada para um fluxo idiomático no ecossistema Go. Ela permite carregar labels, mensagens e erros localizados a partir de arquivos de recurso, em vez de espalhar textos por handlers, services e código de domínio.

Atualmente, este repositório carrega **duas linhas**:

- **v1** na raiz do repositório  
  Linha estável de manutenção para consumidores existentes. Release estável atual: **`v1.0.1`**.
- **v2** em [`/v2`](./v2)  
  Próxima major do projeto. Ela contém a reescrita arquitetural e a nova API.

> `v1.0.0` foi retraída. Use `v1.0.1+` para a linha legada.
>
> Se você consome apenas releases tagueadas, permaneça na **v1.0.1** até que **`v2.0.0`** seja publicada.

## Qual linha você deve usar?

### Use a v1 quando

- você já importa `github.com/Lucas-Palomo/go-resource/pkg/resource`
- você quer o menor custo de migração possível
- você só precisa da API original com as correções de estabilidade da `v1.0.1`

### Use a v2 quando

- você está começando trabalho novo ou uma migração estruturada
- precisa de tratamento explícito de erro
- quer suporte a `fs.FS` / `embed.FS`
- quer objetos aninhados achatados em notação por ponto
- precisa de políticas configuráveis para chaves ausentes e duplicadas
- quer suporte nativo a arquivos `.properties` no estilo Java

## Layout do repositório

```text
.
├── go.mod                    # módulo v1: github.com/Lucas-Palomo/go-resource
├── pkg/resource              # pacote v1
├── docs                      # documentação da raiz/v1
└── v2                        # módulo v2: github.com/Lucas-Palomo/go-resource/v2
```

## Instalação

### v1

```bash
go get github.com/Lucas-Palomo/go-resource@v1.0.1
```

Caminho de import:

```go
import "github.com/Lucas-Palomo/go-resource/pkg/resource"
```

### v2

Depois que a primeira tag da v2 for publicada:

```bash
go get github.com/Lucas-Palomo/go-resource/v2@latest
```

Caminho de import:

```go
import resource "github.com/Lucas-Palomo/go-resource/v2"
```

## Quick start da v1

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

## Quick start da v2

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

## Organização dos recursos

As duas linhas aceitam estruturas semânticas de diretórios como:

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

O projeto também aceita namespace por nome de arquivo:

```text
/resources
  en.errors.json
  pt_BR.errors.json
  en.messages.checkout.toml
```

Na v2, objetos aninhados são achatados automaticamente. Exemplo:

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

## Formatos suportados

### v1

- JSON
- YAML / YML
- TOML

### v2

- JSON
- YAML / YML
- TOML
- `.properties` no estilo Java

## Documentação

### v1

- [Estratégia de versionamento e release](./docs/versioning-strategy.pt_br.md)
- [Referência da API da v1](./docs/v1-reference.pt_br.md)

### v2

- [README da v2](./v2/README.pt_br.md)
- [Referência da API da v2](./v2/docs/api-reference.pt_br.md)
- [Notas de arquitetura da v2](./v2/docs/architecture.pt_br.md)
- [Migração da v1 para a v2](./v2/docs/migration-v1-to-v2.pt_br.md)

## Licença

Este repositório é distribuído sob **The Unlicense**.
