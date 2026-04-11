# go-resource

**Internacionalização para Go inspirada no ResourceBundle.**

[English](./README.md) | [Português (Brasil)](./README.pt_br.md)

`go-resource` é uma biblioteca leve em Go para gerenciar labels, mensagens, erros e textos localizados a partir de arquivos de recurso JSON, YAML e TOML. O projeto surgiu com uma ideia semelhante ao Java ResourceBundles, adaptada ao ecossistema Go com uma API pequena e um fluxo simples baseado em arquivos.

> Este pacote é a **linha de manutenção da v1**.
> A versão `v1.0.0` foi retraída. Use `v1.0.1+` para compatibilidade legada.

## Por que esta release existe

Esta linha `v1.0.1` existe para estabilizar o pacote original antes da publicação separada da nova major.

Objetivos principais desta patch release:
- parar de usar `panic` como modo padrão de falha em uma biblioteca
- preservar o import path existente: `github.com/Lucas-Palomo/go-resource/pkg/resource`
- manter pequeno o custo de migração para consumidores existentes
- documentar com clareza o comportamento suportado da v1

## Instalação

```bash
go get github.com/Lucas-Palomo/go-resource@v1.0.1
```

## Início rápido

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

## Carregamento compatível com legado

O método antigo `Load()` ainda existe por compatibilidade, mas não usa mais `panic`.
Em vez disso, ele armazena internamente o último erro de carregamento.

```go
bundle.Load()
if err := bundle.Err(); err != nil {
	log.Fatal(err)
}
```

Para código novo, prefira `LoadWithError()`.

## Formatos de recurso suportados

- JSON
- YAML / YML
- TOML

## Contrato dos arquivos de recurso

A v1 espera nomes de arquivo no formato:

```text
<locale>.<ext>
```

Exemplos:

```text
resources/
  en.json
  pt_BR.yaml
  es.toml
```

Diretórios aninhados são permitidos. Todo arquivo de recurso válido encontrado sob a raiz configurada é mesclado no catálogo da locale correspondente.

## Comportamento de fallback

`Get()` agora consulta primeiro a locale atual e depois a locale default configurada.
Se a chave ainda não for encontrada, a própria chave é retornada.

## Documentação

- EN
	- [Versioning and release strategy](./docs/versioning-strategy.md)
	- [v1 reference](./docs/v1-reference.md)
- PT-BR
	- [Estratégia de versionamento e release](./docs/versioning-strategy.pt_br.md)
	- [Referência da v1](./docs/v1-reference.pt_br.md)

## Package path

```go
import "github.com/Lucas-Palomo/go-resource/pkg/resource"
```

## Licença

Este repositório é distribuído sob a **The Unlicense**.
