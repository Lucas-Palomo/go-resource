# go-resource

[English](./README.md) | [Português (Brasil)](./README.pt_br.md)

**Internacionalização para Go inspirada em ResourceBundle.**

`go-resource` é uma biblioteca leve em Go para gerenciar labels, mensagens, erros e textos localizados a partir de arquivos de recurso em JSON, YAML e TOML. O projeto nasceu com uma ideia parecida com Java ResourceBundles, adaptada ao ecossistema Go com uma API pequena e um fluxo simples baseado em arquivos.

Este repositório carrega duas linhas principais:

- **v1** na raiz do repositório: compatibilidade legada / modo de manutenção
- **v2** em [`/v2`](./v2): a linha recomendada para novos projetos

> A versão `v1.0.0` foi retraída. Use `v1.0.1+` para a linha legada ou migre para `v2` em novos trabalhos.

## Por que este projeto existe

Em muitas aplicações Go, i18n rapidamente degrada para:

- textos hardcoded espalhados em handlers e services
- troca de locale embutida na regra de negócio
- catálogos duplicados
- fallback frágil
- arquivos de tradução sem estrutura

`go-resource` oferece uma abordagem mais limpa com catálogos de recursos em estilo bundle, mais próxima da ergonomia que desenvolvedores Java conhecem com ResourceBundles, mas expressa de um jeito idiomático para Go.

## Layout do repositório

```text
.
├── go.mod                    # módulo v1: github.com/Lucas-Palomo/go-resource
├── pkg/resource              # pacote v1
└── v2                        # módulo v2: github.com/Lucas-Palomo/go-resource/v2
```

## Instalação

### v1

```bash
go get github.com/Lucas-Palomo/go-resource@latest
```

### v2

```bash
go get github.com/Lucas-Palomo/go-resource/v2@latest
```

## Quick start da v1

```go
package main

import (
	"github.com/Lucas-Palomo/go-resource/pkg/resource"
	"golang.org/x/text/language"
)

func main() {
	bundle := resource.NewBundle("./resources", language.English)
	bundle.Load()

	println(bundle.Get("title"))
	println(bundle.Get("hello", "Lucas"))
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

	fmt.Println(bundle.Get("title"))
	fmt.Println(bundle.Get("checkout.hello", "Lucas"))
}
```

## Formatos suportados

- JSON
- YAML / YML
- TOML

## Estruturas de recursos

### Baseada em pastas

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

### Baseada em nomes de arquivo

```text
/resources
  en.errors.json
  pt_BR.errors.json
  en.messages.checkout.toml
```

## Documentação

- [Estratégia de versionamento e release](./docs/versioning-strategy.pt_br.md)
- [Versioning and release strategy](./docs/versioning-strategy.md)
- [Migração da v1 para a v2](./v2/docs/migration-v1-to-v2.pt_br.md)
- [Migration from v1 to v2](./v2/docs/migration-v1-to-v2.md)
- [Notas de arquitetura da v2](./v2/docs/architecture.pt_br.md)
- [v2 architecture notes](./v2/docs/architecture.md)
- [Referência da API da v2](./v2/docs/api-reference.pt_br.md)
- [v2 API reference](./v2/docs/api-reference.md)

## SEO / descoberta

Palavras-chave relevantes para este projeto:

- biblioteca de i18n para Golang
- pacote de internacionalização para Go
- ResourceBundle para Go
- alternativa ao Java ResourceBundle em Go
- labels e mensagens localizadas em Go
- i18n baseado em arquivos para Golang

## Licença

Este repositório é distribuído sob **The Unlicense**.
