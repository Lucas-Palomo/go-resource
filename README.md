# go-resource

**ResourceBundle-inspired internationalization for Go**.

`go-resource` is a lightweight Go library for managing **dynamic labels, messages, errors and localized texts** from **JSON, YAML and TOML** resource files. The project started with an idea similar to **Java ResourceBundles**, adapted to the Go ecosystem with a small API and a simple file-based workflow.

This repository now carries **two major lines**:

- **v1** at the repository root: legacy compatibility / maintenance mode
- **v2** in [`/v2`](./v2): the recommended line for new projects

> Version `v1.0.0` was retracted. Use `v1.0.1+` for the legacy line or migrate to `v2` for new work.

## Why this project exists

In many Go applications, i18n quickly degrades into:

- hardcoded texts spread across handlers and services
- locale switches embedded in business logic
- duplicated message catalogs
- fragile fallback behavior
- unstructured translation files

`go-resource` gives you a cleaner approach with **bundle-style resource catalogs**, closer to the ergonomics that Java developers know from **ResourceBundles**, but expressed in a Go-friendly way.

## Repository layout

```text
.
├── go.mod                    # v1 module: github.com/Lucas-Palomo/go-resource
├── pkg/resource              # v1 package
└── v2                        # v2 module: github.com/Lucas-Palomo/go-resource/v2
```

## Install

### v1

```bash
go get github.com/Lucas-Palomo/go-resource@latest
```

### v2

```bash
go get github.com/Lucas-Palomo/go-resource/v2@latest
```

## v1 quick start

```go
package main

import (
	"github.com/Lucas-Palomo/go-resource/pkg/resource"
	"golang.org/x/text/language"
)

func main() {
	bundle := resource.NewBundle("./resource", language.English)
	bundle.Load()

	println(bundle.Get("title"))
	println(bundle.Get("hello", "Lucas"))
}
```

## v2 quick start

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

## Supported resource formats

- JSON
- YAML / YML
- TOML

## Resource structures

### Folder-based

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

### Filename-based

```text
/resources
  en.errors.json
  pt_BR.errors.json
  en.messages.checkout.toml
```

## Documentation

- [Versioning and release strategy](./docs/versioning-strategy.md)
- [Migration from v1 to v2](./v2/docs/migration-v1-to-v2.md)
- [v2 architecture notes](./v2/docs/architecture.md)
- [v2 API reference](./v2/docs/api-reference.md)

## SEO / discoverability

Relevant keywords for this project:

- Golang i18n library
- Go internationalization package
- ResourceBundle for Go
- Java ResourceBundle alternative in Go
- localized labels and messages in Go
- file-based i18n for Golang

## License

This repository is distributed under **The Unlicense**.
