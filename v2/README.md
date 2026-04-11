# go-resource v2

[English](./README.md) | [Português (Brasil)](./README.pt_br.md)

**ResourceBundle-style i18n for Go** — a modernized, safer, and more extensible rewrite of `go-resource`, inspired by the idea of Java ResourceBundles and tailored for Golang internationalization.

This is the **recommended line for new projects**.

## Install

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

## Why v2 exists

The original v1 worked well for small projects, but it mixed directory walking, parsing, loading, and lookup into a single flow, used `panic` in the main path, and had fallback behavior that could become inconsistent when a locale was missing.

v2 introduces:

- explicit error handling
- `LoadDir` and `LoadFS`
- support for `embed.FS`
- JSON, YAML, TOML, and Java-style `.properties` decoders
- flattening of nested structures with dot notation
- predictable fallback resolution
- configurable behavior for missing keys and duplicate keys
- thread-safe reads
- package-level examples for pkg.go.dev style documentation
- helper inspection methods such as `Has`, `HasFor`, and `Reset`

## Resource file conventions

### Folder-based namespaces

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

### Filename-based namespaces

```text
/resources
  en.errors.json
  pt_BR.errors.json
  en.messages.checkout.toml
```

### Java-style `.properties` files

```properties
title = Checkout
checkout.hello = Hello, %s
errors.validation.required: Required
checkout.description = Confirm your order \
  before leaving
welcome = Ol\u00E1
```

This format supports the usual Java-style separators (`=`, `:`, or whitespace), comments with `#` or `!`, line continuation with a trailing backslash, and escape sequences such as `\t`, `\n`, and `\uXXXX`.

### Nested objects

```json
{
  "checkout": {
    "button": {
      "confirm": "Confirm"
    }
  }
}
```

becomes:

```text
checkout.button.confirm
```

## Example

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

Runnable examples are available in [`./examples/basic`](./examples/basic) and [`./examples/properties`](./examples/properties).

## Docs

- [API reference](./docs/api-reference.md)
- [API reference (pt-BR)](./docs/api-reference.pt_br.md)
- [Architecture notes](./docs/architecture.md)
- [Architecture notes (pt-BR)](./docs/architecture.pt_br.md)
- [Migration guide](./docs/migration-v1-to-v2.md)
- [Migration guide (pt-BR)](./docs/migration-v1-to-v2.pt_br.md)

## Repository note

The repository root keeps the legacy v1 module for compatibility. This module lives in `/v2` to preserve Go major-version compatibility.
